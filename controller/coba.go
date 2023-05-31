package controller

import (
	"errors"
	"fmt"
	"github.com/aiteung/musik"
	inimodel "github.com/stefanymelda/be_kuesioner/model"
	inimodule "github.com/stefanymelda/be_kuesioner/module"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
	tuhmodelloh "github.com/indrariksa/be_presensi/model"
	tuhmoduleloh "github.com/indrariksa/be_presensi/module"
	// inimodultugas "github.com/stefanymelda/be_kuesioner/module"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	"github.com/stefanymelda/stefanyapp/config"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Home(c *fiber.Ctx) error {
 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
 		"github_repo": "https://github.com/stefanymelda/stefanyapp",
 		"message":     "You are at the root endpoint 😉",
 		"success":     true,
 	})
 }

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

func GetKuesioner(c *fiber.Ctx) error {
	ps := inimodule.GetKuesionerFromStatus("done", config.Ulbimongoconn, "kuesioner")
	return c.JSON(ps)
}

func GetResponden(c *fiber.Ctx) error {
	ps := inimodule.GetRespondenFromUsia(19, config.Ulbimongoconn, "responden")
	return c.JSON(ps)
}

// func GetJamPengisian(c *fiber.Ctx) error {
// 	ps := inimodule.GetJamPengisianFromDurasi(2, config.Ulbimongoconn, "jampengisian")
// 	return c.JSON(ps)
// }

func GetLokasi(c *fiber.Ctx) error {
	ps := inimodule.GetLokasiFromNama("ULBI", config.Ulbimongoconn, "lokasi")
	return c.JSON(ps)
}

func GetAllKuesionerOld(c *fiber.Ctx) error {
	ps := inimodule.GetAllKuesionerFromEmail("arykann20@gmail.com", config.Ulbimongoconn, "kuesioner")
	return c.JSON(ps)
}

func InsertDataKuesioner(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var kuesioner inimodel.Kuesioner
	if err := c.BodyParser(&kuesioner); err != nil {
		return err
	}
	insertedID := inimodule.InsertKuesioner(db, "kuesioner",
	kuesioner.Latitude,
	kuesioner.Longitude, 
	kuesioner.Location, 
	kuesioner.Email, 
	kuesioner.Status,
	kuesioner.Biodata)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataResponden(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var responden inimodel.Responden
	if err := c.BodyParser(&responden); err != nil {
		return err
	}
	insertedID := inimodule.InsertResponden(db, "responden",
	responden.Nama,
    responden.Jenis_kelamin,
    responden.Usia,
    responden.Email,
	responden.Phone_number)

	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// func InsertDataJamPengisian(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn
// 	var jampengisian inimodel.JamPengisian
// 	if err := c.BodyParser(&jampengisian); err != nil {
// 		return err
// 	}
// 	insertedID := inimodule.InsertJamPengisian(db, "jampengisian",
// 	jampengisian.Durasi,
// 	jampengisian.Jam_mulai,
// 	jampengisian.Jam_selesai,
// 	jampengisian.Deskripsi)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "Data berhasil disimpan.",
// 		"inserted_id": insertedID,
// 	})
// }

func InsertDataLokasi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var lokasi inimodel.Lokasi
	if err := c.BodyParser(&lokasi); err != nil {
		return err
	}
	insertedID := inimodule.InsertLokasi(db, "lokasi",
	lokasi.Nama,
	lokasi.Kategori)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataSurvey(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var survey inimodel.Survey
	if err := c.BodyParser(&survey); err != nil {
		return err
	}
	insertedID := inimodule.InsertSurvey(db, "survey",
	survey.Kode,
	survey.Title,
	survey.Soal)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetAllKuesioner(c *fiber.Ctx) error {
	ps := inimodule.GetAllKuesioner(config.Ulbimongoconn, "kuesioner")
	return c.JSON(ps)
}

// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi tuhmodelloh.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := tuhmoduleloh.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi tuhmodelloh.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = tuhmoduleloh.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePresensiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = tuhmoduleloh.DeletePresensiByID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}