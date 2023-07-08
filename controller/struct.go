package controller

// import (
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

type Responden struct {
	// ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama             string             `bson:"nama,omitempty" json:"nama,omitempty" example:"Jasmine"`
	Jenis_kelamin    string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty" example:"Prempuan"`
	Usia             int      			 `bson:"usia,omitempty" json:"usia,omitempty" example:"18"`
	Email      	     string             `bson:"email,omitempty" json:"email,omitempty" example:"jasmine@gmail.com"`
	Phone_number     string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"081243284212"`
	// Jam_pengisian    JamPengisian     `bson:"jam_pengisian,omitempty" json:"jam_pengisian,omitempty"`
	// Hari_pengisian   string           `bson:"hari_pengisian,omitempty" json:"hari_pengisian,omitempty"`
}

// type JamPengisian struct {
// 	Durasi      int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
// 	Jam_mulai   string   `bson:"jam_mulai,omitempty" json:"jam_mulai,omitempty"`
// 	Jam_selesai string   `bson:"jam_selesai,omitempty" json:"jam_selesai,omitempty"`
// 	Deskripsi 	string	 `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
// 	// Gmt         int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
// 	// Hari        []string `bson:"hari,omitempty" json:"hari,omitempty"`
// }

type Kuesioner struct {
	// ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty" example:"123.11"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty" example:"123.11"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty" example:"Bandung"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty" example:"jasmine@gmail.com"`
	// Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty" example:"Done"`
	Biodata      Responden           `bson:"biodata,omitempty" json:"biodata,omitempty"`
	// Lokasi_kuis  Lokasi              `bson:"lokasi_kuis,omitempty" json:"lokasi_kuis,omitempty"`
}

type Lokasi struct {
	// ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty" example:"Bandung"`
	// Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty" example:"Kota"`
}

type Question struct {
	Nomor    int				`bson:"nomor,omitempty" json:"nomor,omitempty" example:"1"`
	Text     string             `bson:"text,omitempty" json:"text,omitempty" example:"Siapa Presiden pertama Indonesia?"`
	Options  string             `bson:"options,omitempty" json:"options,omitempty" example:"a.Sukarno b.Jokowi"`
}

// type Answer struct {
// 	Question_Nomor     int		`bson:"question_nomor,omitempty" json:"question_nomor,omitempty"`
// 	Text2     string             `bson:"text2,omitempty" json:"text2,omitempty"`
// }

type Survey struct {
	Kode     int		`bson:"kode,omitempty" json:"kode,omitempty" example:"101"`
	Title     string             `bson:"title,omitempty" json:"title,omitempty" example:"Sejarah"`
	Soal	Question	`bson:"soal,omitempty" json:"soal,omitempty"`
}

type Karyawan struct {
	// ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty" example:"Tes Swagger"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	Jabatan     string             `bson:"jabatan,omitempty" json:"jabatan,omitempty" example:"Anonymous"`
	Jam_kerja   []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja  []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty" example:"8"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty" example:"08:00"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty" example:"16:00"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty" example:"7"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty" example:"2"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty" example:"Piket Z"`
}

type Presensi struct {
	// ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty" example:"123.11"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty" example:"123.11"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty" example:"Bandung"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	//Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin string   `bson:"checkin,omitempty" json:"checkin,omitempty" example:"MASUK"`
	Biodata Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

// type Lokasi struct {
// 	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
// 	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
// 	Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
// 	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
// }

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}