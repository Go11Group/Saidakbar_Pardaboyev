package api

import (
	"client/api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(handler *handler.Handler) *http.Server {
	r := gin.Default()
	mainRouter := r.Group("/exam")

	// Create subRouter for Users
	userRouter := mainRouter.Group("/users")

	userRouter.POST("/create", handler.Universal)
	userRouter.GET("/getall", handler.Universal)
	userRouter.GET("/:id", handler.Universal)
	userRouter.GET("/:id/courses", handler.Universal)
	userRouter.PUT("/update/:id", handler.Universal)
	userRouter.DELETE("/delete/:id", handler.Universal)

	// Create subRouter for Courses
	courseRouter := mainRouter.Group("/courses")

	courseRouter.POST("/create", handler.Create)
	courseRouter.GET("/getall", handler.Get)
	courseRouter.GET("/:id", handler.Get)
	courseRouter.GET("/:id/lessons", handler.Get)
	courseRouter.GET("/:id/enrollments", handler.Get)
	courseRouter.GET("/popular", handler.Get)
	courseRouter.PUT("/update/:id", handler.Update)
	courseRouter.DELETE("/delete/:id", handler.Delete)

	// Create subRouter for Lessons
	lessonRouter := mainRouter.Group("/lessons")

	lessonRouter.POST("/create", handler.Universal)
	lessonRouter.GET("/getall", handler.Universal)
	lessonRouter.GET("/:id", handler.Universal)
	lessonRouter.PUT("/update/:id", handler.Universal)
	lessonRouter.DELETE("/delete/:id", handler.Universal)

	// Create subRouter for Enrollments
	enrollmentRouter := mainRouter.Group("/enrollments")

	enrollmentRouter.POST("/create", handler.Universal)
	enrollmentRouter.GET("/getall", handler.Universal)
	enrollmentRouter.GET("/:id", handler.Universal)
	enrollmentRouter.PUT("/update/:id", handler.Universal)
	enrollmentRouter.DELETE("/delete/:id", handler.Universal)

	return &http.Server{
		Addr:    ":8088",
		Handler: r,
	}
}
