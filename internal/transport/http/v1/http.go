package v1

import (
	"fr/internal/service"
	"fr/pkg/logger"
	"github.com/gin-gonic/gin"
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
		h1.POST("/", r.addClient)
		h1.PUT("/", r.updateClient)
		h1.DELETE("/", r.deleteClient)

	}
	h2 := handler.Group("/newsletter")
	{
		h2.POST("/create", r.createNewsletter)
		h2.DELETE("/", r.deleteNewsletter)
		h2.PATCH("/", r.updateNewsletter)

	}
}
func NewRouter(handler *gin.Engine, l logger.Interface, t service.Service) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	//// Swagger
	//swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	//handler.GET("/swagger/*any", swaggerHandler)

	//// K8s probe
	//handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	//
	//// Prometheus metrics
	//handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newTransport(h, t, l)
	}
}
