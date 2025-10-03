package handler

import (
	"net/http"

	"github.com/amenshenin/go_auth/internal/schemas"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input schemas.UserInput
	if err := c.BindJSON(&input); err != nil {
		h.logger.Error("Error in signUp", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"error": "wrong input structure",
		})
		return
	}
	id, err := h.service.Autorization.CreateUser(&input)
	if err != nil {
		h.logger.Error("Error in signUp", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"error": "cannot create user",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input schemas.UserInput
	if err := c.BindJSON(&input); err != nil {
		h.logger.Error("Error in signIn", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"error": "wrong input structure",
		})
		return
	}
	token, err := h.service.Autorization.GenerateToken(&input)
	if err != nil {
		h.logger.Error("Error in signIn", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"error": "cannot generate tocken",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
