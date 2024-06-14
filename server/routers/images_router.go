package routers

import "server/api"

func (RG RouterGroup) ImagesRouter() {
	imagesApi := api.Groups.ImageApi
	RG.Router.GET("/images", imagesApi.ImageList)
	RG.Router.GET("/imagesName", imagesApi.ImageNameList)
	RG.Router.POST("/images", imagesApi.ImageUpload)
	RG.Router.PUT("/images", imagesApi.ImageUpdate)
	RG.Router.DELETE("/images", imagesApi.ImageRemove)
}
