package main

import (
	"skeleton/app/global/variable"
	_ "skeleton/bootstrap"
	"skeleton/routes"
)

func main() {
	router := routes.InitWebRouters()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Web.Port"))
}
