package main

import (
	"github.com/labstack/echo/v4"
	"echolabstack/routes"
	

)


	

func main() {
	e := echo.New()


     routes.Echoroutes(e)
	
   
	e.Logger.Fatal(e.Start(":8080"))
}
