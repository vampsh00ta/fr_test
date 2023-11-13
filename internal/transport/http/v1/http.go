package v1

import (
	"fr/internal/service"
	"fr/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type transport struct {
	service service.Service
	logger  logger.Interface
}

func newTransport(handler *gin.RouterGroup, t service.Service, l logger.Interface) {
	r := &transport{t, l}
	//r = r
	h1 := handler.Group("/client")
	{
		h1.POST("/add", r.addClient)
		h1.PUT("/", r.updateClient)
		h1.DELETE("/", r.deleteClient)

	}
	h2 := handler.Group("/newsletter")
	{
		h2.POST("/create", r.createNewsletter)
		h2.DELETE("/", r.deleteNewsletter)
		h2.PATCH("/", r.updateNewsletter)
		h2.GET("/:id/messages", r.getLastMessageStatuses)
		h2.GET("/:id", r.getNewsletter)

	}
}
func NewRouter(handler *gin.Engine, l logger.Interface, t service.Service) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"localhost:8000"}
	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/docs/*any", swaggerHandler)

	// Routers
	h := handler.Group("/v1")
	{
		newTransport(h, t, l)
	}
}
