package url

import (
	"github.com/stefanymelda/stefanyapp/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage) //ujicoba panggil package musik
	// page.Get("/presensi", controller.GetPresensi)
	page.Get("/kuesioner", controller.GetKuesioner)
	page.Get("/responden", controller.GetResponden)
	page.Get("/jampengisian", controller.GetJamPengisian)
	page.Get("/lokasi", controller.GetLokasi)
	page.Get("/kuis", controller.GetAllKuesioner)

}
