package middleware

import (
	"github.com/gin-gonic/gin"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
	"strings"
)

func extractTokenFromHeader(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")

	parts := strings.Split(bearerToken, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", utils.ErrInvalidToken
	}
	return parts[1], nil
}

// check if token is valid or not
// B1: Get token from header
// B2: Validate token and get payload
// B3: From payload, use email to find user in db
func (m *middlewareManager) RequiredAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeader(c.Request)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			//return
			panic(common.ErrAuthorized(err))
		}

		payload, err := utils.ValidateJWT(token, m.cfg)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			//return
			panic(common.ErrAuthorized(err))
		}
		user, err := m.userRepo.FindDataWithCondition(c.Request.Context(), map[string]any{"email": payload.Email})
		if err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"error": err})
			//return
			panic(common.ErrBadRequest(err))
		}

		c.Set("User", user)

		c.Next()
	}
}
