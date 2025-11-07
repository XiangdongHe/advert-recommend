import random
import json
from datetime import datetime, timedelta

# 配置
NUM_USERS = 100
NUM_AD_PLANS = 5000
CREATIVE_PER_PLAN = (1, 2)
INTEREST_PER_USER = 8

# 随机池
regions = ['北京', '上海', '广州', '深圳', '杭州', '成都', '武汉', '南京', '重庆', '天津', '青岛', '苏州', '厦门', '郑州']
devices = ['Android', 'iOS']
objectives = ['click', 'download', 'conversion']
bid_modes = ['CPC 0.5元', 'CPM 0.01元', 'CPA 2元']
interests = [
    # --- 生活与家庭 ---
    "居家收纳", "家装设计", "厨房小技巧", "家电测评", "家庭园艺", "阳台花园", "手工改造", "整理收纳", "节日布置", "家庭理财",
    "健康饮食", "早晨习惯", "时间管理", "生活妙招", "断舍离", "极简生活", "二手交易", "家庭教育", "亲子阅读", "母婴护理",
    "婴儿用品推荐", "育儿心得", "家庭摄影", "家居DIY", "宠物养护", "宠物训练", "萌宠日常", "宠物摄影", "宠物营养",

    # --- 时尚美妆 ---
    "穿搭分享", "季节穿搭", "极简穿搭", "通勤穿搭", "彩妆教程", "护肤步骤", "防晒测评", "香水推荐", "口红试色", "化妆品评测",
    "发型设计", "美甲艺术", "发色趋势", "皮肤护理", "化妆工具", "卸妆产品", "护发技巧", "时尚趋势", "配饰搭配", "鞋包测评",
    "风格搭配", "节日穿搭", "约会妆容", "职场妆容", "日韩妆容", "欧美妆容", "学生党穿搭", "明星同款", "轻奢品牌", "小众品牌",

    # --- 教育成长 ---
    "学习方法", "英语学习", "考研经验", "四六级备考", "留学申请", "留学生活", "学术科研", "论文写作", "在线教育", "课程推荐",
    "编程学习", "人工智能基础", "数据结构", "机器学习入门", "数学思维", "逻辑推理", "阅读分享", "读书笔记", "历史解读", "哲学思考",
    "心理学知识", "时间管理术", "笔记技巧", "学习规划", "个人成长", "高效记忆法", "学习软件推荐", "学习工具测评", "成长日记", "教育观察",

    # --- 运动健康 ---
    "健身计划", "跑步打卡", "健身打卡", "减脂餐", "力量训练", "瑜伽教学", "普拉提", "骑行记录", "游泳技巧", "篮球技巧",
    "羽毛球", "乒乓球", "滑雪", "登山", "户外徒步", "露营装备", "健康饮食", "营养搭配", "体态矫正", "冥想练习",
    "作息规律", "中医养生", "心理健康", "健康检测", "健身装备", "运动鞋测评", "手环评测", "运动饮食", "体重管理", "健康打卡",

    # --- 游戏娱乐 ---
    "王者荣耀", "原神", "和平精英", "LOL英雄联盟", "DOTA2", "CSGO", "Steam新游", "Switch游戏", "手游推荐", "游戏评测",
    "独立游戏", "怀旧游戏", "二次元文化", "漫展探店", "手办收藏", "漫画推荐", "动画解析", "游戏攻略", "电竞赛事", "直播技巧",
    "B站文化", "短视频创作", "剧情向游戏", "视觉小说", "游戏美术", "游戏音乐", "游戏开发", "虚拟主播", "数码娱乐", "直播带货",

    # --- 财经与商业 ---
    "投资理财", "基金分析", "股票策略", "期货交易", "保险知识", "经济观察", "宏观分析", "商业新闻", "创业经验", "企业管理",
    "职场成长", "面试技巧", "简历优化", "人际沟通", "谈判技巧", "领导力培养", "项目管理", "时间效率", "远程办公", "副业项目",
    "自由职业", "内容创业", "品牌建设", "市场营销", "广告投放", "数据分析", "AI商业化", "科技创业", "商业洞察", "新媒体营销",

    # --- 科技与汽车 ---
    "数码测评", "手机评测", "电脑装机", "笔记本推荐", "摄影器材", "智能手表", "AI应用", "机器人", "无人机航拍", "智能家居",
    "汽车测评", "新能源车", "电动车体验", "汽车改装", "驾驶技巧", "汽车保养", "赛车运动", "汽车文化", "特斯拉动态", "智能驾驶",
    "工业科技", "芯片发展", "半导体趋势", "ARVR体验", "未来科技", "互联网趋势", "编程技术", "区块链知识", "数字货币", "云计算",

    # --- 旅行与摄影 ---
    "城市旅行", "露营生活", "自然风光", "旅行攻略", "打卡圣地", "网红景点", "节日旅行", "城市探店", "国外旅行", "民宿体验",
    "摄影技巧", "航拍摄影", "人像摄影", "夜景拍摄", "光影摄影", "手机摄影", "照片修图", "后期调色", "旅行日记", "旅拍Vlog",
    "文化探访", "老街漫游", "建筑摄影", "旅行穿搭", "旅途故事", "旅拍摄影", "地理奇观", "边境探索", "海岛旅行", "雪山之旅",

    # --- 艺术与手工 ---
    "绘画创作", "水彩画", "油画", "素描", "插画", "数字绘画", "陶艺制作", "手工皮具", "木工创作", "刺绣",
    "十字绣", "DIY饰品", "蜡烛制作", "布艺", "手帐设计", "艺术鉴赏", "艺术史", "艺术展览", "画展探店", "文艺生活",
    "摄影艺术", "字体设计", "平面设计", "UI设计", "室内设计", "建筑艺术", "雕塑", "书法", "篆刻", "舞蹈艺术",

    # --- 心理与情感 ---
    "恋爱技巧", "婚姻关系", "分手心理", "情感修复", "社交心理学", "孤独研究", "幸福感提升", "自我认知", "人格类型", "心理疗愈",
    "焦虑管理", "情绪稳定", "冥想放松", "睡眠改善", "生活哲学", "人生思考", "励志成长", "成功学", "幸福哲学", "人生规划",
    "社交礼仪", "人际关系", "职场沟通", "共情力培养", "心理测试", "心灵鸡汤", "自我探索", "习惯养成", "时间管理", "感恩日记"
]

# ----------------------------
# 1️⃣ 生成广告计划
# ----------------------------
ad_plan_sql = open("ad_plan.sql", "w", encoding="utf-8")
ad_plan_sql.write("INSERT INTO t_ad_plan (name, objective, budget, bid_price, targeting_rule, start_time, end_time, status) VALUES\n")

for i in range(1, NUM_AD_PLANS + 1):
    name = f"广告计划_{i}"
    obj = random.choice(objectives)
    budget = round(random.uniform(1000, 50000), 2)
    bid = random.choice(bid_modes)
    region = random.choice(regions)
    age_range = f"{random.randint(18, 25)}-{random.randint(30, 50)}"
    interest = random.choice(interests)
    device = random.choice(devices)
    rule = json.dumps({'region': region, 'age': age_range, 'interest': interest, 'device': device}, ensure_ascii=False)
    start_time = datetime(2025, 11, random.randint(1, 28))
    end_time = start_time + timedelta(days=random.randint(30, 120))
    status = random.choice([0, 1, 2])
    ad_plan_sql.write(f"('{name}', '{obj}', {budget}, '{bid}', '{rule}', '{start_time}', '{end_time}', {status})")
    ad_plan_sql.write(",\n" if i < NUM_AD_PLANS else ";\n")
ad_plan_sql.close()

# ----------------------------
# 2️⃣ 生成广告创意
# ----------------------------
ad_creative_sql = open("ad_creative.sql", "w", encoding="utf-8")
ad_creative_sql.write("INSERT INTO t_ad_creative (plan_id, creative_type, media_url, title, description) VALUES\n")

creative_id = 1
for plan_id in range(1, NUM_AD_PLANS + 1):
    for _ in range(random.randint(*CREATIVE_PER_PLAN)):
        ctype = random.choice([1, 2, 3])
        if ctype == 1:
            media_url = f"https://cdn.ad.com/images/{plan_id}_{creative_id}.jpg"
        elif ctype == 2:
            media_url = f"https://cdn.ad.com/videos/{plan_id}_{creative_id}.mp4"
        else:
            media_url = ''
        title = f"创意标题_{plan_id}_{creative_id}"
        desc = f"广告描述_{random.choice(interests)}_{region}"
        ad_creative_sql.write(f"({plan_id}, {ctype}, '{media_url}', '{title}', '{desc}')")
        creative_id += 1
        ad_creative_sql.write(",\n" if plan_id < NUM_AD_PLANS or _ < CREATIVE_PER_PLAN[1] else ";\n")
ad_creative_sql.close()

# ----------------------------
# 3️⃣ 生成用户基础画像
# ----------------------------
user_sql = open("user_profile.sql", "w", encoding="utf-8")
user_sql.write("INSERT INTO user_profile_base (user_id, gender, age, region, device_type) VALUES\n")

for uid in range(1001, 1001 + NUM_USERS):
    gender = random.choice([1, 2])
    age = random.randint(18, 45)
    region = random.choice(regions)
    device = random.choice(devices)
    user_sql.write(f"({uid}, {gender}, {age}, '{region}', '{device}')")
    user_sql.write(",\n" if uid < 1000 + NUM_USERS else ";\n")
user_sql.close()

# ----------------------------
# 4️⃣ 生成用户兴趣画像
# ----------------------------
interest_sql = open("user_interest.sql", "w", encoding="utf-8")
interest_sql.write("INSERT INTO user_profile_interest (user_id, tag, weight) VALUES\n")

cnt = 0
for uid in range(1001, 1001 + NUM_USERS):
    tags = random.sample(interests, INTEREST_PER_USER)
    for tag in tags:
        weight = round(random.uniform(0.5, 1.0), 4)
        cnt += 1
        interest_sql.write(f"({uid}, '{tag}', {weight})")
        interest_sql.write(",\n" if cnt < NUM_USERS * INTEREST_PER_USER else ";\n")
interest_sql.close()

# ----------------------------
# 5 生成用户行为日志
# ----------------------------
interest_sql = open("user_event.sql", "w", encoding="utf-8")
interest_sql.write("INSERT INTO user_ad_event_log (user_id, creative_id, event_type, ts) VALUES\n")

cnt = 0
for i in range(10000):
    uid = random.randint(1001, 1100)
    creative_id = random.randint(1, 7516)
    event_type = random.randint(1, 3)

    interest_sql.write(f"({uid}, {creative_id}, {event_type}, NOW()),")
interest_sql.close()

print("✅ SQL 数据文件生成完毕：")
