package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	OpenID    string         `gorm:"uniqueIndex;size:64" json:"openid"`
	Nickname  string         `gorm:"size:64" json:"nickname"`
	Avatar    string         `gorm:"size:512" json:"avatar"`
	School    string         `gorm:"size:128" json:"school"`
	Phone     string         `gorm:"size:20" json:"phone"`
	Rating    float64        `gorm:"default:5.0" json:"rating"`
	SellCount int            `gorm:"default:0" json:"sell_count"`
	BuyCount  int            `gorm:"default:0" json:"buy_count"`
	Role      string         `gorm:"size:16;default:user" json:"role"` // user | admin
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Service 服务模型（卖板块 - 服务）
type Service struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      uint           `gorm:"index" json:"user_id"`
	Title       string         `gorm:"size:256" json:"title"`
	Category    string         `gorm:"size:64" json:"category"`
	Price       float64        `json:"price"`
	Description string         `gorm:"type:text" json:"description"`
	Images      string         `gorm:"type:text" json:"images"` // JSON 数组
	Status      string         `gorm:"size:16;default:online" json:"status"` // online | offline | sold
	ViewCount   int            `gorm:"default:0" json:"view_count"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	User        User           `gorm:"foreignKey:UserID" json:"seller,omitempty"`
}

// Goods 物品模型（卖板块 - 二手物品）
type Goods struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	UserID        uint           `gorm:"index" json:"user_id"`
	Title         string         `gorm:"size:256" json:"title"`
	Category      string         `gorm:"size:64" json:"category"`
	Price         float64        `json:"price"`
	OriginalPrice float64        `json:"original_price"`
	Condition     string         `gorm:"size:32" json:"condition"` // 全新 | 九成新 | ...
	Description   string         `gorm:"type:text" json:"description"`
	Images        string         `gorm:"type:text" json:"images"`
	Status        string         `gorm:"size:16;default:online" json:"status"`
	ViewCount     int            `gorm:"default:0" json:"view_count"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	User          User           `gorm:"foreignKey:UserID" json:"seller,omitempty"`
}

// Demand 需求模型（买板块 - 用户发布需求）
type Demand struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      uint           `gorm:"index" json:"user_id"`
	Type        string         `gorm:"size:16" json:"type"` // service | goods
	Title       string         `gorm:"size:256" json:"title"`
	Budget      float64        `json:"budget"`
	Description string         `gorm:"type:text" json:"description"`
	IsUrgent    bool           `gorm:"default:false" json:"is_urgent"`
	Status      string         `gorm:"size:16;default:open" json:"status"` // open | closed
	OfferCount  int            `gorm:"default:0" json:"offer_count"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	User        User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Order 订单模型
type Order struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	OrderNo        string         `gorm:"uniqueIndex;size:32" json:"order_no"`
	BuyerID        uint           `gorm:"index" json:"buyer_id"`
	SellerID       uint           `gorm:"index" json:"seller_id"`
	ItemType       string         `gorm:"size:16" json:"item_type"` // service | goods | demand
	ItemID         uint           `json:"item_id"`
	Amount         float64        `json:"amount"`
	Status         string         `gorm:"size:16;default:pending" json:"status"` // pending | paid | ongoing | done | cancelled
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	Buyer          User           `gorm:"foreignKey:BuyerID" json:"buyer,omitempty"`
	Seller         User           `gorm:"foreignKey:SellerID" json:"seller,omitempty"`
}

// Message 消息模型
type Message struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	FromUserID uint      `gorm:"index" json:"from_user_id"`
	ToUserID   uint      `gorm:"index" json:"to_user_id"`
	Content    string    `gorm:"type:text" json:"content"`
	IsRead     bool      `gorm:"default:false" json:"is_read"`
	CreatedAt  time.Time `json:"created_at"`
}