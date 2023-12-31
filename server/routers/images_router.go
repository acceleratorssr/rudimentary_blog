package routers

import "server/api"

func (RG RouterGroup) ImagesRouter() {
	imagesApi := api.Groups.ImageApi
	RG.Router.GET("/images", imagesApi.ImageListView)
	RG.Router.GET("/imagesName", imagesApi.ImageNameListView)
	RG.Router.POST("/images", imagesApi.ImageUploadView)
	RG.Router.PUT("/images", imagesApi.ImageUpdateView)
	RG.Router.DELETE("/images", imagesApi.ImageRemoveView)
}
