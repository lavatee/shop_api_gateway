package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/lavatee/shop_api_gateway/internal/service"
	pb "github.com/lavatee/shop_protos/gen"
	"github.com/sirupsen/logrus"
)

func Err(c *gin.Context, code int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(code, map[string]string{
		"message": message,
	})

}

type Endpoint struct {
	Services     *service.Service
	ProductsConn pb.ProductsClient
	ReviewsConn  pb.ReviewsClient
	OrdersConn   pb.OrdersClient
	SavedConn    pb.SavedClient
}

func (e *Endpoint) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
		ctx.Writer.Header().Set("Access-Control-Allow-Creditionals", "true")
	})

	auth := router.Group("/auth")
	{
		auth.POST("/signin", e.SignIn)
		auth.POST("/signup", e.SignUp)
		auth.POST("/refresh", e.Refresh)
	}
	api := router.Group("/api")
	{
		api.GET("/products/:category", e.GetProducts)
		api.GET("/products/:id", e.GetOneProduct)
		api.POST("/products", e.PostProduct)
		api.DELETE("/products/:id", e.DeleteProduct)
		api.POST("/orders", e.PostOrder)
		api.GET("/orders/:user_id", e.GetOrders)
		api.GET("/liked/:user_id", e.GetLiked)
		api.POST("/liked", e.PostLikedProduct)
		api.DELETE("/liked/:user_id", e.DeleteLikedProduct)
		api.GET("/cart/:user_id", e.GetCart)
		api.POST("/cart", e.PostCartProduct)
		api.DELETE("/cart/:user_id", e.DeleteCartProduct)
		api.POST("/reviews", e.PostReview)
		api.DELETE("/reviews", e.DeleteReview)
	}

	return router
}
