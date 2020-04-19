package router

import (
	"beer_dmeo/filter"
	"beer_dmeo/service"
	"github.com/houzhongjian/beer"
	"log"
)
func init (){
	log.Println("router init")
}

func Init() beer.Engine {
	beer.Loadini("./public/conf/app.ini")
	srv := beer.New()

	srv.Use(filter.FilterLogin)
	srv.Static("/img/", "./public/image")
	srv.Static("/conf/", "./public/conf/")
	srv.Static("/css/", "./public/css/")
	srv.Static("/js/", "./public/js/")

	srv.SetTemplateDir("./views/")

	srv.GET("/", service.Default)
	srv.GET("/detail/:id", service.Detail)
	srv.GET("/login", service.Login)
	srv.GET("/rem", service.Rem)
	return srv
}
