package routes

import (
	"mginmall/api/v1"
	middleware "mginmall/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)

		v1.GET("/carousels", api.ListCarousel)

		v1.GET("/products", api.ListProduct)
		v1.GET("/products/:id", api.ShowProduct)
		v1.GET("/imgs/:id", api.ListProductImg)
		v1.GET("/categorys", api.ListCategory)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			authed.POST("money", api.ShowMoney)

			authed.POST("product", api.CreateProduct)
			authed.POST("products", api.SearchProduct)

			authed.GET("/favorite", api.ListFavorite)
			authed.POST("/favorite", api.CreateFavorite)
			authed.DELETE("/favorite/:id", api.DeleteFavorite)
		}
	}
	return r
}
