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

CREATE TABLE `t_user_friend`  (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL,
    `friend_id` bigint NOT NULL,
    `closeness` decimal(5, 4) ZEROFILL NOT NULL COMMENT '亲密度，0-1',
    PRIMARY KEY (`id`)
);

-- 开始执行get_sql.py
-- 插入测试数据示例


-- 规范化数据
UPDATE t_ad_creative c
JOIN t_ad_plan p ON c.plan_id = p.plan_id
SET c.description = CONCAT(
    '广告描述_',
    JSON_UNQUOTE(JSON_EXTRACT(p.targeting_rule, '$.interest')),
    '_',
    JSON_UNQUOTE(JSON_EXTRACT(p.targeting_rule, '$.region'))
);


-- 协同过滤测试数据
INSERT INTO user_profile_base(user_id, gender, age, region, device_type, create_time, update_time)
VALUES (1101, 1, 25, '上海', 'iPhone', NOW(), NOW());

INSERT INTO t_user_friend(user_id, friend_id, closeness)
VALUES (1101, 1001, 0.9), (1101, 1002, 0.1);
