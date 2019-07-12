package main

import (
	"videoSecurity/logwriter"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/syndtr/goleveldb/leveldb"
)

//Router support version
var Router *gin.Engine

//Init initialization routers
func initRouters(serviceContainer IServiceContainer, log *logwriter.Logger, db *leveldb.DB) (err error) {
	Router = gin.New()
	Router.Use(middlewareCORS())
	Router.Use(middlewareResponse())

	Router.Use(static.Serve("/", static.LocalFile("www", true)))
	Router.LoadHTMLGlob("www/index.html")

	{
		searchController := serviceContainer.InjectSearchController()
		videoController := serviceContainer.InjectVideoController()

		deviceController := serviceContainer.InjectDeviceController()
		deviceAuthController := serviceContainer.InjectDeviceAuthController()

		api := Router.Group("/v1")
		//	api.Handle("/get-token", auth.GetTokenHandler)
		api.GET("/autosearch", searchController.Auto)
		api.GET("/manualsearch", searchController.Manual)

		api.GET("/videostream", videoController.Run)

		api.GET("/device/all", deviceController.All)
		api.GET("/device/add/:ip/:port", deviceController.AddOrUpdate)
		api.DELETE("/device/remove/:ip", deviceController.Remove)

		api.GET("/device/auth/all", deviceAuthController.All)
		api.GET("/device/auth/add/:login/:password", deviceAuthController.AddOrUpdate)
		api.DELETE("/device/auth/remove/:login", deviceAuthController.Remove)
	}

	return
}
