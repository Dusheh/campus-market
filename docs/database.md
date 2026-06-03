# 数据库设计

## ER 图

```
┌──────────┐       ┌──────────┐       ┌──────────┐
│   User   │       │ Service  │       │  Goods   │
├──────────┤       ├──────────┤       ├──────────┤
│ id       │──┐    │ id       │       │ id       │
│ openid   │  │    │ user_id  │──┐    │ user_id  │──┐
│ nickname │  │    │ title    │  │    │ title    │  │
│ avatar   │  │    │ category │  │    │ category │  │
│ school   │  │    │ price    │  │    │ price    │  │
│ phone    │  │    │ desc     │  │    │ cond     │  │
│ rating   │  ├───>│ status   │  ├───>│ status   │  │
└──────────┘  │    └──────────┘  │    └──────────┘  │
              │                  │                  │
              │    ┌──────────┐  │                  │
              │    │ Demand   │  │                  │
              │    ├──────────┤  │                  │
              │    │ id       │  │                  │
              ├───>│ user_id  │──┘                  │
              │    │ type     │                     │
              │    │ title    │                     │
              │    │ budget   │                     │
              │    │ is_urgent│                     │
              │    └──────────┘                     │
              │                                     │
              │    ┌──────────────────┐             │
              │    │      Order       │             │
              │    ├──────────────────┤             │
              ├───>│ buyer_id         │             │
              ├───>│ seller_id        │<────────────┘
              │    │ item_type        │
              │    │ item_id          │
              │    │ amount           │
              │    │ status           │
              │    └──────────────────┘
              │
              │    ┌──────────┐
              │    │ Message  │
              │    ├──────────┤
              ├───>│from_user │
              ├───>│ to_user  │
              │    │ content  │
              │    └──────────┘
```

## 表结构

### users 用户表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键 |
| open_id | varchar(64) | 微信 OpenID |
| nickname | varchar(64) | 昵称 |
| avatar | varchar(512) | 头像 URL |
| school | varchar(128) | 学校 |
| phone | varchar(20) | 手机号 |
| rating | decimal(3,2) | 评分 (默认 5.0) |
| sell_count | int | 卖出次数 |
| buy_count | int | 买入次数 |
| role | varchar(16) | 角色 (user/admin) |

### services 服务表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键 |
| user_id | bigint | 发布者 ID |
| title | varchar(256) | 标题 |
| category | varchar(64) | 分类 |
| price | decimal(10,2) | 价格 |
| description | text | 描述 |
| images | text | 图片 JSON 数组 |
| status | varchar(16) | online/offline/sold |
| view_count | int | 浏览次数 |

### goods 物品表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键 |
| user_id | bigint | 发布者 ID |
| title | varchar(256) | 标题 |
| category | varchar(64) | 分类 |
| price | decimal(10,2) | 售价 |
| original_price | decimal(10,2) | 原价 |
| condition | varchar(32) | 成色 |
| description | text | 描述 |
| images | text | 图片 JSON 数组 |
| status | varchar(16) | online/offline/sold |

### demands 需求表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键 |
| user_id | bigint | 发布者 ID |
| type | varchar(16) | service/goods |
| title | varchar(256) | 标题 |
| budget | decimal(10,2) | 预算 |
| description | text | 描述 |
| is_urgent | tinyint | 是否急需 |
| status | varchar(16) | open/closed |
| offer_count | int | 响应人数 |

### orders 订单表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键 |
| order_no | varchar(32) | 订单编号 |
| buyer_id | bigint | 买家 ID |
| seller_id | bigint | 卖家 ID |
| item_type | varchar(16) | service/goods/demand |
| item_id | bigint | 关联项目 ID |
| amount | decimal(10,2) | 金额 |
| status | varchar(16) | pending/paid/ongoing/done/cancelled |

### messages 消息表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键 |
| from_user_id | bigint | 发送者 ID |
| to_user_id | bigint | 接收者 ID |
| content | text | 消息内容 |
| is_read | tinyint | 是否已读 |

## 索引建议

```sql
-- 服务查询优化
CREATE INDEX idx_services_status_category ON services(status, category);
CREATE INDEX idx_services_user_id ON services(user_id);

-- 物品查询优化
CREATE INDEX idx_goods_status_category ON goods(status, category);
CREATE INDEX idx_goods_user_id ON goods(user_id);

-- 需求查询优化
CREATE INDEX idx_demands_status_type ON demands(status, type);
CREATE INDEX idx_demands_is_urgent ON demands(is_urgent);

-- 订单查询优化
CREATE INDEX idx_orders_buyer_status ON orders(buyer_id, status);
CREATE INDEX idx_orders_seller_status ON orders(seller_id, status);
```