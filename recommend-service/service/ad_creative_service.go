package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/config"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/database"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/models"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/utils"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// AdCreativeService 广告创意服务
type AdCreativeService struct{}

// NewAdCreativeService 创建广告创意服务实例
func NewAdCreativeService() *AdCreativeService {
	return &AdCreativeService{}
}

// CreateAdCreative 创建广告创意
func (s *AdCreativeService) CreateAdCreative(planID int64, creativeType int32, mediaURL, title, description string) (int64, error) {
	creative := &models.AdCreative{
		PlanID:       planID,
		CreativeType: creativeType,
		MediaURL:     mediaURL,
		Title:        title,
		Description:  description,
		Status:       1, // 默认激活
	}

	if err := database.DB.Create(creative).Error; err != nil {
		return 0, err
	}

	return creative.CreativeID, nil
}

// UpdateAdCreative 更新广告创意
func (s *AdCreativeService) UpdateAdCreative(creativeID int64, updates map[string]interface{}) error {
	result := database.DB.Model(&models.AdCreative{}).Where("creative_id = ?", creativeID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("creative not found")
	}

	return nil
}

// GetAdCreative 获取广告创意
func (s *AdCreativeService) GetAdCreative(creativeID int64) (*models.AdCreative, error) {
	var creative models.AdCreative
	if err := database.DB.Where("creative_id = ?", creativeID).First(&creative).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("creative not found")
		}
		return nil, err
	}
	return &creative, nil
}

// ListAdCreatives 获取广告创意列表
func (s *AdCreativeService) ListAdCreatives(page, pageSize int, planID *int64) ([]*models.AdCreative, int64, error) {
	var creatives []*models.AdCreative
	var total int64

	query := database.DB.Model(&models.AdCreative{})
	if planID != nil {
		query = query.Where("plan_id = ?", *planID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&creatives).Error; err != nil {
		return nil, 0, err
	}

	return creatives, total, nil
}

// DeleteAdCreative 删除广告创意
func (s *AdCreativeService) DeleteAdCreative(creativeID int64) error {
	result := database.DB.Where("creative_id = ?", creativeID).Delete(&models.AdCreative{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("creative not found")
	}

	return nil
}

// 获取推荐广告列表
func (s *AdCreativeService) GetAdvertRecommend(userId int64) ([]*models.AdCreative, int64, error) {
	var ansCreatives []*models.AdCreative
	var ansAdPlans []*models.AdPlan
	var total int64
	// TODO 完成广告推荐的逻辑
	// 提取用户的基本信息
	t1 := utils.StartTiming("GetUserProfileBaseCached")
	userInfo, _ := GetUserProfileBaseCached(userId)
	t1.End()
	t2 := utils.StartTiming("GetUserInterestsCached")
	userInterests, _ := GetUserInterestsCached(userId)
	t2.End()
	// 1. 基于协同过滤匹配扩充用户兴趣爱好
	t3 := utils.StartTiming("GetCollaborativeInterests")
	collaborativeInterests, _ := GetCollaborativeInterests(userId)
	t3.End()
	t4 := utils.StartTiming("MergeInterests")
	allInterests := MergeInterests(userInterests, collaborativeInterests)
	t4.End()
	// 2. 基于规则匹配到的广告集合
	t5 := utils.StartTiming("GetInterestAdPlansCachedV2")
	interestAdPlans, _ := GetInterestAdPlansCachedV2(allInterests)
	t5.End()
	ansAdPlans = append(ansAdPlans, interestAdPlans...)
	// TODO 3.基于内容匹配到的广告集合-感觉意义不大，暂不考虑
	// TODO 4.基于向量召回匹配到的广告集合
	// 筛选
	t6 := utils.StartTiming("FilterAdPlansByUser")
	ansAdPlans = FilterAdPlansByUser(ansAdPlans, userInfo.Region, int(userInfo.Age))
	t6.End()
	// 根据兴趣权重进行粗排
	t7 := utils.StartTiming("SortAdPlansByInterest")
	ansCreatives = SortAdPlansByInterest(ansAdPlans, allInterests)
	t7.End()
	// TODO 根据CTR模型预测，进行粗排
	return ansCreatives, total, nil
}

func GetUserProfileBaseCached(userId int64) (*models.UserProfileBase, error) {
	key := fmt.Sprintf("user:profile:%d", userId)
	var user models.UserProfileBase
	ctx := context.Background()
	// 尝试从 Redis 取
	val, err := database.RDB.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &user); err == nil {
			return &user, nil
		}
	}
	// 缓存未命中 → 访问数据库
	if err := database.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	// 写入缓存
	data, _ := json.Marshal(user)
	database.RDB.Set(ctx, key, data, time.Hour)
	return &user, nil
}

func GetUserInterestsCached(userId int64) ([]*models.UserProfileInterest, error) {
	key := fmt.Sprintf("user:interests:%d", userId)
	ctx := context.Background()
	var interests []*models.UserProfileInterest

	// Redis
	val, err := database.RDB.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &interests); err == nil {
			return interests, nil
		}
	}

	// DB
	if err := database.DB.Where("user_id = ?", userId).Find(&interests).Error; err != nil {
		return nil, err
	}

	// 缓存
	data, _ := json.Marshal(interests)
	database.RDB.Set(ctx, key, data, 30*time.Minute)

	return interests, nil
}

func GetInterestAdPlansCachedV1(userInterests []*models.UserProfileInterest) ([]*models.AdPlan, error) {
	ctx := context.Background()

	// 批量取所有 interest -> plan 集合
	pipe := database.RDB.Pipeline()
	cmds := make([]*redis.StringSliceCmd, len(userInterests))
	for i, it := range userInterests {
		cmds[i] = pipe.SMembers(ctx, fmt.Sprintf("interest:%s", it.Tag))
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	// 合并所有 PlanID
	planIDSet := make(map[int64]struct{})
	for _, cmd := range cmds {
		ids, _ := cmd.Result()
		for _, idStr := range ids {
			id, _ := strconv.ParseInt(idStr, 10, 64)
			planIDSet[id] = struct{}{}
		}
	}
	if len(planIDSet) == 0 {
		return nil, nil
	}

	// Pipeline 批量 MGET 广告计划
	var planKeys []string
	for id := range planIDSet {
		planKeys = append(planKeys, fmt.Sprintf("ad:plan:%d", id))
	}

	pipe = database.RDB.Pipeline()
	mgetCmds := make([]*redis.StringCmd, len(planKeys))
	for i, k := range planKeys {
		mgetCmds[i] = pipe.Get(ctx, k)
	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	var adPlans []*models.AdPlan
	for _, cmd := range mgetCmds {
		data, _ := cmd.Result()
		if data == "" {
			continue
		}
		var plan models.AdPlan
		if err := json.Unmarshal([]byte(data), &plan); err == nil && plan.Status == 1 {
			adPlans = append(adPlans, &plan)
		}
	}

	// Pipeline 批量取创意
	pipe = database.RDB.Pipeline()
	creativeCmds := make([]*redis.StringSliceCmd, len(adPlans))
	for i, plan := range adPlans {
		creativeCmds[i] = pipe.SMembers(ctx, fmt.Sprintf("ad:plan:%d:creatives", plan.PlanID))
	}
	_, _ = pipe.Exec(ctx)

	for i, plan := range adPlans {
		cids, _ := creativeCmds[i].Result()
		if len(cids) == 0 {
			continue
		}

		// 再次使用 pipeline 批量取创意详情
		subPipe := database.RDB.Pipeline()
		subCmds := make([]*redis.StringCmd, len(cids))
		for j, cid := range cids {
			subCmds[j] = subPipe.Get(ctx, fmt.Sprintf("ad:creative:%s", cid))
		}
		_, _ = subPipe.Exec(ctx)
		for _, cmd := range subCmds {
			val, _ := cmd.Result()
			if val == "" {
				continue
			}
			var c models.AdCreative
			if err := json.Unmarshal([]byte(val), &c); err == nil && c.Status == 1 {
				plan.Creatives = append(plan.Creatives, &c)
			}
		}
	}

	return adPlans, nil
}

const getInterestAdPlansLua = `
local result = {}
local planIDSet = {}
local limit = 50

for i, tag in ipairs(ARGV) do
    local key = "interest:" .. tag
    local pids = redis.call("SRANDMEMBER", key, limit)
    for _, pid in ipairs(pids) do
        planIDSet[pid] = true
    end
end

local planIDs = {}
for pid, _ in pairs(planIDSet) do
    table.insert(planIDs, pid)
end

math.randomseed(redis.call("TIME")[2])
for i = #planIDs, 2, -1 do
    local j = math.random(i)
    planIDs[i], planIDs[j] = planIDs[j], planIDs[i]
end

if #planIDs > limit then
    for i = limit + 1, #planIDs do
        planIDs[i] = nil
    end
end

for _, pid in ipairs(planIDs) do
    local planJSON = redis.call("GET", "ad:plan:" .. pid)
    if planJSON then
        table.insert(result, planJSON)
    end
end
return result
`

func GetInterestAdPlansCachedV2(userInterests []*models.UserProfileInterest) ([]*models.AdPlan, error) {
	ctx := context.Background()

	args := make([]interface{}, len(userInterests))
	for i, t := range userInterests {
		args[i] = t.Tag
	}
	res, err := database.RDB.Eval(ctx, getInterestAdPlansLua, []string{}, args...).Result()
	if err != nil {
		log.Fatalf("Redis Lua eval failed: %v", err)
	}
	adPlans := []*models.AdPlan{}
	items, ok := res.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", res)
	}

	for _, item := range items {
		str, ok := item.(string)
		if !ok {
			continue
		}

		var plan models.AdPlan
		if err := json.Unmarshal([]byte(str), &plan); err == nil {
			adPlans = append(adPlans, &plan)
		}
	}
	return adPlans, nil
}

// 根据协同过滤查询筛选兴趣爱好
func GetCollaborativeInterests(userId int64) ([]*models.UserProfileInterest, error) {
	var friendInterests []*models.UserProfileInterest

	sql := `
		SELECT 
			i.id,
			i.user_id,
			i.tag,
			(f.closeness * i.weight) AS weight,
			i.update_time
		FROM t_user_friend AS f
		JOIN user_profile_interest AS i 
			ON f.friend_id = i.user_id
		WHERE f.user_id = ?
		ORDER BY weight DESC
		LIMIT ?
	`
	topN := config.Global.RecommendConfig.CollaborativeCount
	if err := database.DB.Raw(sql, userId, topN).Scan(&friendInterests).Error; err != nil {
		return nil, err
	}
	return friendInterests, nil
}

// MergeInterests 合并两个兴趣切片，并根据 ID 去重
func MergeInterests(a, b []*models.UserProfileInterest) []*models.UserProfileInterest {
	idMap := make(map[int64]bool)
	var merged []*models.UserProfileInterest

	// 遍历第一个列表
	for _, item := range a {
		if !idMap[item.ID] {
			merged = append(merged, item)
			idMap[item.ID] = true
		}
	}

	// 遍历第二个列表
	for _, item := range b {
		if !idMap[item.ID] {
			merged = append(merged, item)
			idMap[item.ID] = true
		}
	}

	return merged
}

type TargetingRule struct {
	Age      string `json:"age"`
	Device   string `json:"device"`
	Region   string `json:"region"`
	Interest string `json:"interest"`
}

// 筛选广告计划
func FilterAdPlansByUser(plans []*models.AdPlan, userRegion string, userAge int) []*models.AdPlan {
	var result []*models.AdPlan

	for _, plan := range plans {
		var rule TargetingRule
		if err := json.Unmarshal([]byte(plan.TargetingRule), &rule); err != nil {
			continue
		}

		matchRegion := rule.Region == "" || strings.Contains(userRegion, rule.Region) || strings.Contains(rule.Region, userRegion)

		// 年龄匹配
		matchAge := false
		if rule.Age != "" {
			var minAge, maxAge int
			fmt.Sscanf(rule.Age, "%d-%d", &minAge, &maxAge)
			if userAge >= minAge && userAge <= maxAge {
				matchAge = true
			}
		} else {
			matchAge = true
		}

		if matchRegion && matchAge {
			result = append(result, plan)
		}
	}

	return result
}

func SortAdPlansByInterest(plans []*models.AdPlan, userInterests []*models.UserProfileInterest) []*models.AdCreative {
	interestWeight := make(map[string]float64)
	for _, ui := range userInterests {
		interestWeight[ui.Tag] = ui.Weight
	}
	var ansCreatives []*models.AdCreative
	for _, adPlan := range plans {
		var targeting map[string]string
		if err := json.Unmarshal([]byte(adPlan.TargetingRule), &targeting); err != nil {
			continue
		}
		interest, ok := targeting["interest"]
		if !ok {
			continue
		}
		w := interestWeight[interest]
		for _, ad := range adPlan.Creatives {
			ad.Weight = w
			ansCreatives = append(ansCreatives, ad)
		}
	}

	sort.SliceStable(ansCreatives, func(i, j int) bool {
		return ansCreatives[i].Weight > ansCreatives[j].Weight
	})
	return ansCreatives
}
