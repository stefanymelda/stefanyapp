package main

import (
	"log"

	"github.com/stefanymelda/stefanyapp/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/stefanymelda/stefanyapp/url"

	"github.com/gofiber/fiber/v2"
	_ "github.com/stefanymelda/stefanyapp/docs"
)

// @title TES SWAG
// @version 1.0
// @description This is a sample server.


// @contact.name API Support
// @contact.url https://github.com/stefanymelda
// @contact.email 1214026@std.ulbi.ac.id

// @host jaehyun.herokuapp.com
// @BasePath /
// @schemes https http
func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
