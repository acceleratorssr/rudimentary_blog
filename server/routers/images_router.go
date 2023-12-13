package routers

import "server/api"

func (RG RouterGroup) ImagesRouter() {
	imagesApi := api.Groups.ImageApi
	RG.Router.GET("/images", imagesApi.ImageListView)
	RG.Router.POST("/images", imagesApi.ImageUploadView)
	RG.Router.DELETE("/images", imagesApi.ImageDeleteView)
	RG.Router.PUT("/images", imagesApi.ImageUpdateView)
}
