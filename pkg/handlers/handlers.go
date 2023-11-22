package handlers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/novriantama/question1/pkg/models"
	"github.com/novriantama/question1/pkg/services"
)

type Handlers struct {
	service services.Service
}

func NewHandlers(service services.Service) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) GetUserByID(c *gin.Context) {
	userID := parseUserID(c.Param("id"))
	user, err := h.service.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		// Handle error
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func parseUserID(param string) int {
	userID, _ := strconv.Atoi(param)
	return userID
}

// handler for create user
func (h *Handlers) CreateUser(c *gin.Context) {
	var userPayload models.UserPayload

	// Bind JSON payload to the struct
	if err := c.ShouldBindJSON(&userPayload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create user
	if err := h.service.CreateUser(c.Request.Context(), userPayload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "success")
}

// handler for generate otp
func (h *Handlers) GenerateOtp(c *gin.Context) {
	var OtpPayload models.SetOtpPayload

	// Bind JSON payload to the struct
	if err := c.ShouldBindJSON(&OtpPayload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// generate otp
	if err := h.service.GenerateOtp(c.Request.Context(), OtpPayload); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "success")
}

// handler for verify otp
func (h *Handlers) VerifyOtp(c *gin.Context) {
	var OtpPayload models.GetOtpPayload

	// Bind JSON payload to the struct
	if err := c.ShouldBindJSON(&OtpPayload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// verify otp
	if err := h.service.VerifyOtp(c.Request.Context(), OtpPayload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "success")
}
