-- 广告推荐系统数据库初始化脚本

-- 广告计划表
CREATE TABLE t_ad_plan (
    plan_id        BIGINT PRIMARY KEY AUTO_INCREMENT,
    name           VARCHAR(128) NOT NULL,
    objective      VARCHAR(64)  NOT NULL COMMENT '投放目标，click/download/conversion',
    budget         DECIMAL(16,2) NOT NULL COMMENT '总预算',
    bid_price      VARCHAR(64) NOT NULL COMMENT '出价模式，CPC/CPM/CPA 0.5元/0.01元/5元',
    targeting_rule JSON NOT NULL COMMENT '地域/年龄/性别/兴趣/设备等定向条件',
    start_time     DATETIME NOT NULL,
    end_time       DATETIME NOT NULL,
    status         TINYINT NOT NULL DEFAULT 1 COMMENT '1=active,0=paused,2=ended',
    create_time    DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_time    DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_status (status),
    INDEX idx_time_range (start_time, end_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告计划表';

-- 广告创意（具体的图片/视频）
CREATE TABLE t_ad_creative (
    creative_id      BIGINT PRIMARY KEY AUTO_INCREMENT,
    plan_id         BIGINT NOT NULL,
    creative_type    TINYINT NOT NULL COMMENT '1=image,2=video,3=text',
    media_url        VARCHAR(512) NOT NULL,
    title            VARCHAR(256),
    description      VARCHAR(512),
    status           TINYINT NOT NULL DEFAULT 1,
    create_time      DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_time      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_plan_id (plan_id),
    INDEX idx_status (status),
    FOREIGN KEY (plan_id) REFERENCES t_ad_plan(plan_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告创意表';

-- 用户静态画像（基础属性）
CREATE TABLE user_profile_base (
    user_id      BIGINT PRIMARY KEY,
    gender       TINYINT COMMENT '0=unknown,1=male,2=female',
    age          INT,
    region       VARCHAR(64),
    device_type  VARCHAR(64),
    create_time  DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_time  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_region (region),
    INDEX idx_age (age)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户基础画像表';

-- 用户兴趣画像（标签 + 权重）
CREATE TABLE user_profile_interest (
    id           BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id      BIGINT NOT NULL,
    tag          VARCHAR(128) NOT NULL,
    weight       DECIMAL(5,4) NOT NULL COMMENT '0~1',
    update_time  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_user_tag(user_id, tag),
    INDEX idx_tag (tag)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户兴趣画像表';

-- 用户广告行为日志
CREATE TABLE user_ad_event_log (
    event_id    BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id     BIGINT NOT NULL,
    creative_id BIGINT NOT NULL,
    event_type  TINYINT NOT NULL COMMENT '1=exposure,2=click,3=conversion',
    ts          DATETIME NOT NULL,
    extra       JSON NULL,
    INDEX idx_user_ts(user_id, ts),
    INDEX idx_creative_ts(creative_id, ts),
    INDEX idx_event_type(event_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户广告行为日志表';

-- 插入测试数据示例

-- 广告计划测试数据
INSERT INTO t_ad_plan (name, objective, budget, bid_price, targeting_rule, start_time, end_time, status) VALUES
('新品推广计划', 'click', 10000.00, 'CPC:0.5', '{"region":["北京","上海"],"age":[18,35],"gender":[1,2]}', '2024-01-01 00:00:00', '2024-03-31 23:59:59', 1),
('品牌曝光计划', 'conversion', 50000.00, 'CPM:0.01', '{"region":["全国"],"age":[25,45],"interests":["科技","数码"]}', '2024-02-01 00:00:00', '2024-06-30 23:59:59', 1);

-- 广告创意测试数据
INSERT INTO t_ad_creative (plan_id, creative_type, media_url, title, description, status) VALUES
(1, 1, 'https://example.com/ad1.jpg', '新品上市限时优惠', '全新产品，限时5折优惠', 1),
(1, 2, 'https://example.com/ad1.mp4', '产品视频介绍', '精彩视频展示产品特性', 1),
(2, 1, 'https://example.com/ad2.jpg', '品牌形象广告', '展示品牌价值和理念', 1);

-- 用户画像测试数据
INSERT INTO user_profile_base (user_id, gender, age, region, device_type) VALUES
(1001, 1, 28, '北京', 'iPhone 14'),
(1002, 2, 32, '上海', 'Huawei P60'),
(1003, 1, 25, '深圳', 'Xiaomi 13');

-- 用户兴趣测试数据
INSERT INTO user_profile_interest (user_id, tag, weight) VALUES
(1001, '科技', 0.85),
(1001, '数码', 0.75),
(1001, '游戏', 0.60),
(1002, '时尚', 0.90),
(1002, '美妆', 0.80),
(1003, '运动', 0.70),
(1003, '科技', 0.65);

-- 用户行为日志测试数据
INSERT INTO user_ad_event_log (user_id, creative_id, event_type, ts, extra) VALUES
(1001, 1, 1, '2024-01-15 10:30:00', '{"source":"feed","position":1}'),
(1001, 1, 2, '2024-01-15 10:30:05', '{"source":"feed","position":1}'),
(1002, 2, 1, '2024-01-15 11:20:00', '{"source":"search","keyword":"新品"}'),
(1003, 3, 1, '2024-02-01 14:00:00', '{"source":"recommend","score":0.92}'),
(1003, 3, 2, '2024-02-01 14:00:10', '{"source":"recommend","score":0.92}'),
(1003, 3, 3, '2024-02-01 14:05:00', '{"source":"recommend","order_id":"ORD123456"}');
