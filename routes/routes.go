package routes

import (
	"jwt-go-rbac/controller"
	"jwt-go-rbac/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	authRoutes := router.Group("/auth/user")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)
	authRoutes.GET("/verify/:token", controller.VerifyEmail)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.JWTAuth())
	adminRoutes.GET("/users", controller.GetUsers)
	adminRoutes.GET("/user/:id", controller.GetUser)
	adminRoutes.PUT("/user/:id", controller.UpdateUser)
	adminRoutes.POST("/user/role", controller.CreateRole)
	adminRoutes.GET("/user/roles", controller.GetRoles)
	adminRoutes.PUT("/user/role/:id", controller.UpdateRole)
	adminRoutes.POST("/room/add", controller.CreateRoom)
	adminRoutes.PUT("/room/:id", controller.UpdateRoom)
	adminRoutes.GET("/room/bookings", controller.GetBookings)

	publicRoutes := router.Group("/api/view")
	publicRoutes.GET("/rooms", controller.GetRooms)
	publicRoutes.GET("/room/:id", controller.GetRoom)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthCustomer())
	protectedRoutes.GET("/rooms/booked", controller.GetUserBookings)
	protectedRoutes.POST("/room/book", controller.CreateBooking)

}
