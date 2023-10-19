package middleware

import (
	"errors"
	usermodel "go01-airbnb/internal/user/model"
	"go01-airbnb/pkg/common"

	"github.com/gin-gonic/gin"
)

func (m *middlewareManager) RequiredRoles(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("User").(*usermodel.User)
		for i := range roles {
			if user.Role == roles[i] {
				ctx.Next()
				return
			}
		}
		//ctx.JSON(http.StatusForbidden, gin.H{"error": "you have no permission"})
		panic(common.ErrForbidden(errors.New("have no permission")))
	}
}
