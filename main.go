package main

import (
	"github.com/gin-gonic/gin"
	"go01-airbnb/cache"
	"go01-airbnb/config"
	amenityhttp "go01-airbnb/internal/amenity/delivery/http"
	amenityrepository "go01-airbnb/internal/amenity/repository"
	amenityusecase "go01-airbnb/internal/amenity/usecase"
	bookinghttp "go01-airbnb/internal/booking/delivery/http"
	bookingrepository "go01-airbnb/internal/booking/repository"
	bookingusecase "go01-airbnb/internal/booking/usecase"
	locationhttp "go01-airbnb/internal/location/delivery/http"
	locationrepository "go01-airbnb/internal/location/repository"
	locationusecase "go01-airbnb/internal/location/usecase"
	"go01-airbnb/internal/middleware"
	placehttp "go01-airbnb/internal/place/delivery/http"
	placerepository "go01-airbnb/internal/place/repository"
	placeusecase "go01-airbnb/internal/place/usecase"
	placeamenitieshttp "go01-airbnb/internal/placeamenities/delivery/http"
	placeamenitiesrepository "go01-airbnb/internal/placeamenities/repository"
	placeamenitiesusecase "go01-airbnb/internal/placeamenities/usecase"
	placelikehttp "go01-airbnb/internal/placelike/delivery/http"
	placelikerepository "go01-airbnb/internal/placelike/repository"
	placelikeusecase "go01-airbnb/internal/placelike/usecase"
	reviewhttp "go01-airbnb/internal/review/delivery/http"
	reviewrepository "go01-airbnb/internal/review/repository"
	reviewusecase "go01-airbnb/internal/review/usecase"
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
	placeLikeHdl := placelikehttp.NewUserLikePlaceHandler(placeLikeUC, hasher)

	amenityRepo := amenityrepository.NewAmenityRepository(db, sugarLogger)
	amenityUC := amenityusecase.NewAmenityUseCase(amenityRepo)
	amenityHdl := amenityhttp.NewAmenityHandler(amenityUC, hasher)

	checkPlaceOwner := placerepository.NewPlaceRepository(db, sugarLogger)
	placeAmenityRepo := placeamenitiesrepository.NewPlaceAmenitiesRepo(db, sugarLogger)
	placeAmenityUC := placeamenitiesusecase.NewPlaceAmenitiesUseCase(placeAmenityRepo, checkPlaceOwner)
	placeAmenityHdl := placeamenitieshttp.NewPlaceAmenitiesHandler(placeAmenityUC, hasher)

	bookingRepo := bookingrepository.NewBookingRepository(db, sugarLogger)
	bookingUC := bookingusecase.NewBookingUseCase(bookingRepo)
	bookingHdl := bookinghttp.NewBookingHandler(bookingUC, hasher)

	reviewRepo := reviewrepository.NewReviewRepository(db, sugarLogger)
	reviewUC := reviewusecase.NewReviewUseCase(reviewRepo)
	reviewHdl := reviewhttp.NewReviewHandler(reviewUC, hasher)

	locationRepo := locationrepository.NewLocationRepository(db, sugarLogger)
	locationUC := locationusecase.NewLocationUseCase(locationRepo)
	locationHdl := locationhttp.NewLocationHandler(locationUC, hasher)

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
	v1.PUT("/user", middlewares.RequiredAuth(), userHdl.UpdateProfile())

	// Place Like
	v1.POST("/like/:id", middlewares.RequiredAuth(), placeLikeHdl.UserLikePlace())
	v1.DELETE("/unlike/:id", middlewares.RequiredAuth(), placeLikeHdl.UserUnLikePlace())
	v1.GET("/like", middlewares.RequiredAuth(), placeLikeHdl.GetPlacesLikedByUser())

	// Amenity
	v1.POST("/amenities", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), amenityHdl.CreateAmenity())
	v1.GET("/amenities", amenityHdl.GetAmenities())
	v1.GET("/amenities/:id", amenityHdl.GetAmenityById())
	v1.DELETE("/amenities/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), amenityHdl.DeleteAmenity())
	v1.PUT("/amenities/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), amenityHdl.UpdateAmenity())

	// PlaceAmenities
	v1.POST("/place_amenities/:pid/:aid", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), placeAmenityHdl.CreatePlaceAmenities())
	v1.DELETE("/place_amenities/:pid/:aid", middlewares.RequiredAuth(), placeAmenityHdl.DeletePlaceAmenities())
	v1.GET("/place_amenities", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), placeAmenityHdl.GetPlaceAmenities())
	v1.GET("/place_amenities/:place_id", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), placeAmenityHdl.GetAmenitiesByPlaceId())

	// Bookings
	v1.POST("/bookings/:place_id", middlewares.RequiredAuth(), bookingHdl.CreateBooking())
	v1.GET("/bookings", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), bookingHdl.GetAllBooking())
	v1.GET("/bookings/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), bookingHdl.GetBookingById())
	v1.PUT("/bookings/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), bookingHdl.UpdateBooking())
	v1.DELETE("/bookings/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("admin", "host"), bookingHdl.DeleteBooking())

	//Reviews
	v1.POST("/reviews/:booking_id", middlewares.RequiredAuth(), reviewHdl.CreateReview())
	v1.DELETE("/reviews/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("host", "admin"), reviewHdl.DeleteReview())
	v1.GET("/reviews", middlewares.RequiredAuth(), middlewares.RequiredRoles("host", "admin"), reviewHdl.GetAllReview())
	v1.GET("/reviews/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("host", "admin"), reviewHdl.GetReviewById())
	v1.GET("/place_reviews/:place_id", middlewares.RequiredAuth(), reviewHdl.GetAllReviewByPlaceId())

	//Locations
	v1.POST("/locations", middlewares.RequiredAuth(), middlewares.RequiredRoles("host", "admin"), locationHdl.CreateLocation())
	v1.DELETE("/locations/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("host", "admin"), locationHdl.DeleteLocation())
	v1.GET("/locations", middlewares.RequiredAuth(), middlewares.RequiredRoles("host", "admin"), locationHdl.GetAllLocation())
	v1.PUT("/locations/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("host", "admin"), locationHdl.UpdateLocation())
	v1.GET("/locations/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles("host", "admin"), locationHdl.GetLocationById())
	//router.Run()
	router.Run(":" + cfg.App.Port)
}
