# 校园服务买卖平台

一个基于微信小程序的校园 C2C 服务与二手物品交易平台。

## 功能概览

### 卖板块
- 接单者可发布**服务**（代取快递、家教辅导、摄影设计等）
- 所有用户可挂售**二手物品**（教材、电子产品、生活用品等）
- 客户可浏览并购买服务或物品

### 买板块
- 客户可发布**求购需求**（急需物品、寻求服务等）
- 接单者浏览需求并选择接单
- 支持"急需"标记，高优先级展示

### 通用功能
- 微信一键登录
- 订单管理与状态追踪
- 实时消息通讯
- 用户评价体系

## 技术栈

| 层级 | 技术 |
|------|------|
| 前端 | 微信小程序原生框架 |
| 后端 | Go 1.22 + Gin |
| 数据库 | MySQL 8.0 + GORM |
| 缓存 | Redis 7 |
| 认证 | JWT |
| 部署 | Docker + Nginx |

## 项目结构

```
campus-market/
├── miniprogram/          # 微信小程序前端
│   ├── app.js/json/wxss  # 应用入口
│   ├── pages/            # 页面
│   │   ├── sell/         # 卖板块（服务+物品列表）
│   │   ├── buy/          # 买板块（需求列表）
│   │   ├── publish-sell/ # 发布卖（服务/物品）
│   │   ├── publish-buy/  # 发布买（需求）
│   │   ├── chat/         # 消息列表
│   │   ├── order/        # 订单列表
│   │   └── profile/      # 个人中心
│   ├── components/       # 公共组件
│   ├── utils/            # 工具函数
│   └── images/           # 图片资源
├── backend/              # Go 后端
│   ├── cmd/server/       # 入口
│   ├── internal/
│   │   ├── config/       # 配置
│   │   ├── handler/      # 请求处理
│   │   ├── middleware/   # 中间件
│   │   ├── model/        # 数据模型
│   │   ├── repository/   # 数据库操作
│   │   ├── router/       # 路由
│   │   └── service/      # 业务逻辑
│   ├── config/           # 配置文件
│   └── Dockerfile
├── scripts/              # 部署脚本
├── docs/                 # 文档
└── README.md
```

## 快速开始

### 前置要求
- 微信开发者工具
- Go 1.22+
- MySQL 8.0+
- Redis 7+

### 后端启动
```bash
cd backend
# 安装依赖
go mod tidy
# 修改 config/config.yaml 中的数据库配置
# 运行
go run ./cmd/server/
```

### 小程序启动
1. 打开微信开发者工具
2. 导入项目，选择 `miniprogram/` 目录
3. 修改 `project.config.json` 中的 `appid` 为你的小程序 AppID
4. 修改 `miniprogram/app.js` 中的 `baseUrl` 为你的后端地址

### Docker 部署
```bash
# 启动 MySQL + Redis
cd scripts && docker-compose up -d

# 构建并运行后端
cd backend && docker build -t campus-market . && docker run -p 8080:8080 campus-market
```

## API 文档

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/auth/login | 微信登录 |
| GET | /api/user/profile | 获取用户信息 |
| GET | /api/services | 服务列表 |
| GET | /api/services/:id | 服务详情 |
| POST | /api/services | 创建服务 |
| GET | /api/goods | 物品列表 |
| GET | /api/goods/:id | 物品详情 |
| POST | /api/goods | 创建物品 |
| GET | /api/demands | 需求列表 |
| GET | /api/demands/:id | 需求详情 |
| POST | /api/demands | 创建需求 |
| GET | /api/orders | 订单列表 |
| POST | /api/orders | 创建订单 |
| PUT | /api/orders/:id/status | 更新订单状态 |

## 开发计划

- [ ] 用户认证与微信登录
- [ ] 服务发布与浏览
- [ ] 物品挂售与购买
- [ ] 需求发布与接单
- [ ] 订单管理
- [ ] 实时消息
- [ ] 用户评价
- [ ] 管理后台

## License

MIT