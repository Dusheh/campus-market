package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Dusheh/campus-market/internal/model"
)

// CreateOrder 创建订单
func (h *Handler) CreateOrder(c *gin.Context) {
	buyerID := c.GetUint("user_id")

	var req struct {
		ItemType string  `json:"item_type" binding:"required"` // service | goods | demand
		ItemID   uint    `json:"item_id" binding:"required"`
		Amount   float64 `json:"amount" binding:"required"`
		SellerID uint    `json:"seller_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, 400, "参数错误")
		return
	}

	order := model.Order{
		OrderNo:  fmt.Sprintf("CM%d%d", time.Now().Unix(), buyerID),
		BuyerID:  buyerID,
		SellerID: req.SellerID,
		ItemType: req.ItemType,
		ItemID:   req.ItemID,
		Amount:   req.Amount,
		Status:   "pending",
	}

	if err := h.DB.Create(&order).Error; err != nil {
		Error(c, 500, "创建订单失败")
		return
	}

	Success(c, order)
}

// ListOrders 获取订单列表
func (h *Handler) ListOrders(c *gin.Context) {
	userID := c.GetUint("user_id")
	role := c.Query("role") // buyer | seller
	status := c.Query("status")

	var p Pagination
	p.Default()

	var orders []model.Order
	query := h.DB.Model(&model.Order{}).Preload("Buyer").Preload("Seller")

	if role == "seller" {
		query = query.Where("seller_id = ?", userID)
	} else {
		query = query.Where("buyer_id = ?", userID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)
	query.Offset(p.Offset()).Limit(p.PageSize).Order("created_at DESC").Find(&orders)

	Success(c, gin.H{
		"total":     total,
		"page":      p.Page,
		"page_size": p.PageSize,
		"items":     orders,
	})
}

// UpdateOrderStatus 更新订单状态
func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, 400, "参数错误")
		return
	}

	if err := h.DB.Model(&model.Order{}).Where("id = ?", orderID).Update("status", req.Status).Error; err != nil {
		Error(c, 500, "更新失败")
		return
	}

	Success(c, nil)
}