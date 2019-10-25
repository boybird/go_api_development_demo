package routers

import (
	"net/http"

	// 注册 swager 文档
	_ "github.com/boybird/hello/docs"
	v1 "github.com/boybird/hello/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/boybird/hello/pkg/qrcode"
	"github.com/boybird/hello/pkg/upload"
	"github.com/boybird/hello/middleware/jwt"
	"github.com/boybird/hello/pkg/export"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiv1.GET("TAGS", v1.GetTags)
		// 新建标签
		apiv1.POST("TAGS", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
	}

	return r

}
