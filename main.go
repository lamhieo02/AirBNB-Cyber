package main

import (
	"github.com/gin-gonic/gin"
	"go01-airbnb/config"
	"go01-airbnb/internal/middleware"
	placehttp "go01-airbnb/internal/place/delivery/http"
	placerepository "go01-airbnb/internal/place/repository"
	placeusecase "go01-airbnb/internal/place/usecase"
	uploadhttp "go01-airbnb/internal/upload/delivery/http"
	userhttp "go01-airbnb/internal/user/delivery/http"
	userrepository "go01-airbnb/internal/user/repository"
	userusecase "go01-airbnb/internal/user/usecase"
	"go01-airbnb/pkg/db/mysql"
	"go01-airbnb/pkg/logger"
	"go01-airbnb/pkg/upload"
	"go01-airbnb/pkg/utils"
	"log"
)

func main() {
	// dsn := "root:109339Lam@@tcp(127.0.0.1:3306)/airbnb?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalln("Failed connect to MySQL")
	// }

	// log.Println("MySQL connected", db)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Get config error", err)
		return
	}
	// Declare DB
	db, err := mysql.NewMySQL(cfg)

	if err != nil {
		log.Fatalln("Can not connect Mysql: ", err)
	}

	// Declare S3 AWS
	s3Provider := upload.NewS3Provider(cfg)

	// Declare logger
	sugarLogger := logger.NewZapLogger()

	// Declare hashids
	hasher := utils.NewHashIds(cfg.App.Secret, 10)
	//db.AutoMigrate(placemodel.Place{})

	// declare dependencies
	placeRepo := placerepository.NewPlaceRepository(db, sugarLogger)
	placeUC := placeusecase.NewPlaceUseCase(placeRepo)
	placeHdl := placehttp.NewPlaceHandler(placeUC, hasher)

	userRepo := userrepository.NewUserRepository(db)
	userUC := userusecase.NewUserUseCase(cfg, userRepo)
	userHdl := userhttp.NewUserHandler(userUC)

	uploadHdl := uploadhttp.NewUploadHandler(s3Provider)

	middlewares := middleware.NewMiddlewareManager(cfg, userRepo)
	router := gin.Default()

	// Global middleware, nghĩa là tất cả các routers đều phải đi qua middleware này
	router.Use(middlewares.Recover())
	router.Static("/static", "./static")

	v1 := router.Group("/api/v1")

	v1.POST("upload", uploadHdl.Upload())

	v1.POST("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), placeHdl.CreatePlace())
	v1.GET("/places", placeHdl.GetPlaces())
	v1.GET("/places/:id", placeHdl.GetPlaceById())
	v1.PUT("/places/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), placeHdl.UpdatePlace())
	v1.DELETE("places/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), placeHdl.DeletePlace())

	//User
	v1.GET("/profiles", middlewares.RequiredAuth(), userHdl.Profiles())
	v1.GET("/admin", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), userHdl.Profiles())
	v1.POST("/register", userHdl.Register())
	v1.POST("/login", userHdl.Login())
	//router.Run()
	router.Run(":" + cfg.App.Port)
	//hd := hashids.NewData()
	//hd.Salt = "this is my salt"
	//hd.MinLength = 30
	//h, _ := hashids.NewWithData(hd)
	//e, _ := h.Encode([]int{45, 434, 1313, 99})
	//fmt.Println(e)
	//d, _ := h.DecodeWithError(e)
	//fmt.Println(d)
	//

}
