package userhttp

import (
	"context"
	"github.com/gin-gonic/gin"
	usermodel "go01-airbnb/internal/user/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
)

type userUseCase interface {
	Register(context.Context, *usermodel.UserRegister) error
	Login(context.Context, *usermodel.UserLogin) (*utils.Token, error)
	UpdateProfile(context.Context, *usermodel.UserUpdate, string) error
	//GetProfileOfCurrentUser
}

type userHandler struct {
	userUC userUseCase
}

func NewUserHandler(userUseCase userUseCase) *userHandler {
	return &userHandler{userUC: userUseCase}
}

func (hdl *userHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserRegister

		if err := c.ShouldBind(&data); err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"error": err})
			//return
			panic(err)
		}

		if err := hdl.userUC.Register(c.Request.Context(), &data); err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			//return
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}
func (hdl *userHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials usermodel.UserLogin

		if err := c.ShouldBind(&credentials); err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"error": err})
			//return
			panic(err)
		}

		token, err := hdl.userUC.Login(c.Request.Context(), &credentials)
		if err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			//return
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": token})
	}
}

func (hdl *userHandler) Profiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("User")
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (hdl *userHandler) UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("User").(common.Requester)
		var userUpdate usermodel.UserUpdate
		if err := c.ShouldBind(&userUpdate); err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := hdl.userUC.UpdateProfile(c, &userUpdate, user.GetUserEmail()); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.Response(true))
	}
}
