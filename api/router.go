// Sets up and configures Gin HTTP routes
package api

import (
	"time"

	"appnest.co/aisland-go/api/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/v1")
	{
		deviceGroup := v1.Group("/device")
		{
			//Startup
			deviceGroup.POST("/startup", handlers.AccountHandler)
		}
		accountGroup := v1.Group("/account")
		{
			//FastRegister
			accountGroup.POST("/fast_register", handlers.DeviceHandler)
			//RefreshToken
			accountGroup.POST("/refresh_token", handlers.DeviceHandler)
		}
		taskGroup := v1.Group("/tasks")
		{
			taskGroup.POST("/ask", handlers.AskHandler)
			taskGroup.POST("/reply", handlers.ReplyHandler)
			taskGroup.POST("/improve", handlers.ImproveHandler)
			// taskGroup.POST("/errors", handlers.Errors)
		}
	}

	return r
}
