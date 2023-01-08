package main

import (
	"github.com/gin-gonic/gin"
	"go01-airbnb/cache"
	"go01-airbnb/config"
	amenityhttp "go01-airbnb/internal/amenity/delivery/http"
	amenityrepository "go01-airbnb/internal/amenity/repository"
	amenityusecase "go01-airbnb/internal/amenity/usecase"
	"go01-airbnb/internal/middleware"
	placehttp "go01-airbnb/internal/place/delivery/http"
	placerepository "go01-airbnb/internal/place/repository"
	placeusecase "go01-airbnb/internal/place/usecase"
	placelikehttp "go01-airbnb/internal/placelike/delivery/http"
	placelikerepository "go01-airbnb/internal/placelike/repository"
	placelikeusecase "go01-airbnb/internal/placelike/usecase"
	uploadhttp "go01-airbnb/internal/upload/delivery/http"
	userhttp "go01-airbnb/internal/user/delivery/http"
	userrepository "go01-airbnb/internal/user/repository"
	userusecase "go01-airbnb/internal/user/usecase"
	"go01-airbnb/pkg/db/mysql"
	dbredis "go01-airbnb/pkg/db/redis"
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
		log.Fatalln("Can not connect mysql: ", err)
	}
	db = db.Debug()

	utils.RunDBMigration(cfg)
	// Declare redis
	redis, err := dbredis.NewRedisClient(cfg)
	if err != nil {
		log.Fatalln("Can not connect redis: ", err)
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
	userCache := cache.NewAuthUserCache(userRepo, cache.NewRedisCache(redis))
	userUC := userusecase.NewUserUseCase(cfg, userRepo)
	userHdl := userhttp.NewUserHandler(userUC)

	uploadHdl := uploadhttp.NewUploadHandler(s3Provider)

	placeLikeRepo := placelikerepository.NewPlaceLikeRepository(db, sugarLogger)
	placeLikeUC := placelikeusecase.NewUserLikePlaceUseCase(placeLikeRepo)
	placeLikeHdl := placelikehttp.NewUserLikePlaceUseCase(placeLikeUC, hasher)

	amenityRepo := amenityrepository.NewAmenityRepository(db, sugarLogger)
	amenityUC := amenityusecase.NewAmenityUseCase(amenityRepo)
	amenityHdl := amenityhttp.NewAmenityHandler(amenityUC, hasher)

	middlewares := middleware.NewMiddlewareManager(cfg, userCache)
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

	// Place Like
	v1.POST("/:id/like", middlewares.RequiredAuth(), placeLikeHdl.UserLikePlace())
	v1.DELETE("/:id/unlike", middlewares.RequiredAuth(), placeLikeHdl.UserUnLikePlace())
	v1.GET("/like", middlewares.RequiredAuth(), placeLikeHdl.GetPlacesLikedByUser())

	// Amenity
	v1.POST("/amenities", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), amenityHdl.CreateAmenity())
	v1.GET("/amenities", amenityHdl.GetAmenities())
	v1.DELETE("/:id/amenities", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), amenityHdl.DeleteAmenity())
	v1.PUT("/:id/amenities", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), amenityHdl.UpdateAmenity())

	//router.Run()
	router.Run(":" + cfg.App.Port)
}
