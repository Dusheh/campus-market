package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/Dusheh/campus-market/internal/model"
)

// LoginRequest 登录请求
type LoginRequest struct {
	Code string `json:"code" binding:"required"`
}

// Login 微信登录
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, 400, "参数错误")
		return
	}

	// TODO: 调用微信 API 换取 openid
	// 这里先做模拟，实际需要用 req.Code 请求微信接口
	openID := "mock_openid_" + req.Code

	// 查找或创建用户
	var user model.User
	result := h.DB.Where("open_id = ?", openID).First(&user)
	if result.Error != nil {
		user = model.User{
			OpenID:   openID,
			Nickname: "新用户",
			Rating:   5.0,
		}
		h.DB.Create(&user)
	}

	// 生成 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your-jwt-secret"))
	if err != nil {
		Error(c, 500, "token 生成失败")
		return
	}

	Success(c, gin.H{
		"token":    tokenString,
		"user_id":  user.ID,
		"nickname": user.Nickname,
	})
}

// GetProfile 获取用户信息
func (h *Handler) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var user model.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		Error(c, 404, "用户不存在")
		return
	}

	Success(c, user)
}