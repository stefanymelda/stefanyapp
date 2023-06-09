package url

import (
	"github.com/stefanymelda/stefanyapp/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/gofiber/swagger" // swagger handler
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage) //ujicoba panggil package musik
	page.Get("/presensiold", controller.GetPresensi)
	page.Get("/kuisioner", controller.GetKuesioner)
	page.Get("/responden", controller.GetResponden)
	// page.Get("/jampengisian", controller.GetJamPengisian)
	page.Get("/lokasi", controller.GetLokasi)
	page.Get("/kuis", controller.GetAllKuesionerOld)
	page.Get("/presensi", controller.GetAllPresensi) //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	page.Get("/kuesioner", controller.GetAllKuesioner) //menampilkan seluruh data kuesioner
	page.Post("/ins", controller.InsertData)
	page.Put("/upd/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePresensiByID)
	page.Get("/docs/*", swagger.HandlerDefault)
	page.Post("/ins1", controller.InsertData)
}
