package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Dusheh/campus-market/internal/model"
)

// ListServices 获取服务列表
func (h *Handler) ListServices(c *gin.Context) {
	var p Pagination
	p.Default()

	var services []model.Service
	query := h.DB.Model(&model.Service{}).Where("status = ?", "online").Preload("User")

	category := c.Query("category")
	if category != "" {
		query = query.Where("category = ?", category)
	}

	keyword := c.Query("keyword")
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}

	var total int64
	query.Count(&total)
	query.Offset(p.Offset()).Limit(p.PageSize).Order("created_at DESC").Find(&services)

	Success(c, gin.H{
		"total":    total,
		"page":     p.Page,
		"page_size": p.PageSize,
		"items":    services,
	})
}

// GetService 获取服务详情
func (h *Handler) GetService(c *gin.Context) {
	id := c.Param("id")

	var service model.Service
	if err := h.DB.Preload("User").First(&service, id).Error; err != nil {
		Error(c, 404, "服务不存在")
		return
	}

	// 增加浏览量
	h.DB.Model(&service).UpdateColumn("view_count", service.ViewCount+1)

	Success(c, service)
}

// CreateService 创建服务
func (h *Handler) CreateService(c *gin.Context) {
	userID := c.GetUint("user_id")

	var service model.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		Error(c, 400, "参数错误")
		return
	}

	service.UserID = userID
	service.Status = "online"

	if err := h.DB.Create(&service).Error; err != nil {
		Error(c, 500, "创建失败")
		return
	}

	Success(c, service)
}

// ListGoods 获取物品列表
func (h *Handler) ListGoods(c *gin.Context) {
	var p Pagination
	p.Default()

	var goods []model.Goods
	query := h.DB.Model(&model.Goods{}).Where("status = ?", "online").Preload("User")

	category := c.Query("category")
	if category != "" {
		query = query.Where("category = ?", category)
	}

	keyword := c.Query("keyword")
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}

	var total int64
	query.Count(&total)
	query.Offset(p.Offset()).Limit(p.PageSize).Order("created_at DESC").Find(&goods)

	Success(c, gin.H{
		"total":     total,
		"page":      p.Page,
		"page_size": p.PageSize,
		"items":     goods,
	})
}

// GetGoods 获取物品详情
func (h *Handler) GetGoods(c *gin.Context) {
	id := c.Param("id")

	var goods model.Goods
	if err := h.DB.Preload("User").First(&goods, id).Error; err != nil {
		Error(c, 404, "物品不存在")
		return
	}

	h.DB.Model(&goods).UpdateColumn("view_count", goods.ViewCount+1)

	Success(c, goods)
}

// CreateGoods 创建物品
func (h *Handler) CreateGoods(c *gin.Context) {
	userID := c.GetUint("user_id")

	var goods model.Goods
	if err := c.ShouldBindJSON(&goods); err != nil {
		Error(c, 400, "参数错误")
		return
	}

	goods.UserID = userID
	goods.Status = "online"

	if err := h.DB.Create(&goods).Error; err != nil {
		Error(c, 500, "创建失败")
		return
	}

	Success(c, goods)
}