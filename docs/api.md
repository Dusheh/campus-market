# API 接口文档

Base URL: `https://api.campus-market.example.com/api`

## 认证

所有需认证的接口需在 Header 中携带：
```
Authorization: Bearer <token>
```

---

## 1. 认证模块

### POST /auth/login — 微信登录
```
Request:
{
  "code": "wx.login() 返回的 code"
}

Response:
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "jwt_token_string",
    "user_id": 1,
    "nickname": "用户昵称"
  }
}
```

---

## 2. 用户模块

### GET /user/profile — 获取用户信息
```
Response:
{
  "code": 200,
  "data": {
    "id": 1,
    "nickname": "张三",
    "avatar": "https://...",
    "school": "XX大学",
    "rating": 4.8,
    "sell_count": 12,
    "buy_count": 5
  }
}
```

### PUT /user/profile — 更新用户信息
```
Request:
{
  "nickname": "新昵称",
  "school": "XX大学",
  "phone": "13800138000"
}
```

---

## 3. 服务模块

### GET /services — 服务列表
```
Query:
  page       int    页码 (默认 1)
  page_size  int    每页数量 (默认 20, 最大 50)
  category   string 分类筛选
  keyword    string 关键词搜索

Response:
{
  "code": 200,
  "data": {
    "total": 100,
    "page": 1,
    "page_size": 20,
    "items": [...]
  }
}
```

### GET /services/:id — 服务详情
### POST /services — 创建服务
```
Request:
{
  "title": "代拿快递",
  "category": "代取快递",
  "price": 5.00,
  "description": "校门口快递点代取",
  "images": "[\"url1\",\"url2\"]"
}
```

---

## 4. 物品模块

### GET /goods — 物品列表
### GET /goods/:id — 物品详情
### POST /goods — 创建物品
```
Request:
{
  "title": "高等数学教材",
  "category": "教材书籍",
  "price": 15.00,
  "original_price": 42.00,
  "condition": "九成新",
  "description": "几乎没用过",
  "images": "[\"url1\"]"
}
```

---

## 5. 需求模块

### GET /demands — 需求列表
```
Query:
  page       int    页码
  page_size  int    每页数量
  type       string service|goods
  keyword    string 关键词
  is_urgent  bool   是否急需

Response:
{
  "code": 200,
  "data": {
    "total": 50,
    "items": [
      {
        "id": 1,
        "type": "service",
        "title": "急需一名代课同学",
        "budget": 50.00,
        "is_urgent": true,
        "offer_count": 3,
        "user": { "id": 1, "nickname": "李四" }
      }
    ]
  }
}
```

### GET /demands/:id — 需求详情
### POST /demands — 创建需求
```
Request:
{
  "type": "service",
  "title": "急需一名代课同学",
  "budget": 50.00,
  "description": "周三下午2点，需要代一节课",
  "is_urgent": true
}
```

---

## 6. 订单模块

### GET /orders — 订单列表
```
Query:
  role    string buyer|seller  角色
  status  string pending|ongoing|done  状态筛选

Response:
{
  "code": 200,
  "data": {
    "items": [
      {
        "id": 1,
        "order_no": "CM17170000001",
        "buyer": { "id": 1, "nickname": "张三" },
        "seller": { "id": 2, "nickname": "李四" },
        "item_type": "service",
        "amount": 5.00,
        "status": "pending"
      }
    ]
  }
}
```

### POST /orders — 创建订单
```
Request:
{
  "item_type": "service",
  "item_id": 1,
  "amount": 5.00,
  "seller_id": 2
}
```

### PUT /orders/:id/status — 更新订单状态
```
Request:
{
  "status": "paid"  // pending | paid | ongoing | done | cancelled
}
```

---

## 错误码

| 状态码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 参数错误 |
| 401 | 未登录 / token 无效 |
| 404 | 资源不存在 |
| 500 | 服务器错误 |