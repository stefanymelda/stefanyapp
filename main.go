package main

import (
	"log"

	"github.com/stefanymelda/stefanyapp/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/stefanymelda/stefanyapp/url"

	"github.com/gofiber/fiber/v2"
	"github.com/stefanymelda/stefanyapp/docs"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
