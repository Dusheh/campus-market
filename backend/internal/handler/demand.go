package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Dusheh/campus-market/internal/model"
)

// ListDemands 获取需求列表
func (h *Handler) ListDemands(c *gin.Context) {
	var p Pagination
	p.Default()

	var demands []model.Demand
	query := h.DB.Model(&model.Demand{}).Where("status = ?", "open").Preload("User")

	demandType := c.Query("type")
	if demandType != "" {
		query = query.Where("type = ?", demandType)
	}

	keyword := c.Query("keyword")
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}

	isUrgent := c.Query("is_urgent")
	if isUrgent == "true" {
		query = query.Where("is_urgent = ?", true)
	}

	var total int64
	query.Count(&total)
	query.Offset(p.Offset()).Limit(p.PageSize).Order("is_urgent DESC, created_at DESC").Find(&demands)

	Success(c, gin.H{
		"total":     total,
		"page":      p.Page,
		"page_size": p.PageSize,
		"items":     demands,
	})
}

// GetDemand 获取需求详情
func (h *Handler) GetDemand(c *gin.Context) {
	id := c.Param("id")

	var demand model.Demand
	if err := h.DB.Preload("User").First(&demand, id).Error; err != nil {
		Error(c, 404, "需求不存在")
		return
	}

	Success(c, demand)
}

// CreateDemand 创建需求
func (h *Handler) CreateDemand(c *gin.Context) {
	userID := c.GetUint("user_id")

	var demand model.Demand
	if err := c.ShouldBindJSON(&demand); err != nil {
		Error(c, 400, "参数错误")
		return
	}

	demand.UserID = userID
	demand.Status = "open"

	if err := h.DB.Create(&demand).Error; err != nil {
		Error(c, 500, "创建失败")
		return
	}

	Success(c, demand)
}