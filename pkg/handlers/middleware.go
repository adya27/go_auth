package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	fmt.Println("headerParts", headerParts)

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponce(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}
