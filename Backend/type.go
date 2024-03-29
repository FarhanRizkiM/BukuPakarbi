package pakarbibackend

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" `
	UsernameId   string             `json:"usernameid" bson:"usernameid"`
	Username     string             `json:"username" bson:"username"`
	NPM          string             `json:"npm" bson:"npm"`
	Password     string             `json:"password" bson:"password"`
	PasswordHash string             `json:"passwordhash" bson:"passwordhash"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Role         string             `json:"role,omitempty" bson:"role,omitempty"`
	Token        string             `json:"token,omitempty" bson:"token,omitempty"`
	Private      string             `json:"private,omitempty" bson:"private,omitempty"`
	Public       string             `json:"public,omitempty" bson:"public,omitempty"`
}

type Admin struct {
	UsernameId   string `json:"usernameid" bson:"usernameid"`
	Username     string `json:"username" bson:"username"`
	Password     string `json:"password" bson:"password"`
	PasswordHash string `json:"passwordhash" bson:"passwordhash"`
	Email        string `bson:"email,omitempty" json:"email,omitempty"`
	Role         string `json:"role,omitempty" bson:"role,omitempty"`
	Token        string `json:"token,omitempty" bson:"token,omitempty"`
	Private      string `json:"private,omitempty" bson:"private,omitempty"`
	Public       string `json:"public,omitempty" bson:"public,omitempty"`
}

type Parkiran struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" `
	Parkiranid     string             `json:"parkiranid,omitempty" bson:"parkiranid,omitempty"`
	Nama           string             `json:"nama,omitempty" bson:"nama,omitempty"`
	NPM            string             `json:"npm,omitempty" bson:"npm,omitempty"`
	Prodi          string             `json:"prodi,omitempty" bson:"prodi,omitempty"`
	NamaKendaraan  string             `json:"namakendaraan,omitempty" bson:"namakendaraan,omitempty"`
	NomorKendaraan string             `json:"nomorkendaraan,omitempty" bson:"nomorkendaraan,omitempty"`
	JenisKendaraan string             `json:"jeniskendaraan,omitempty" bson:"jeniskendaraan,omitempty"`
	Status         Status             `json:"status" bson:"status"`
	Base64Image    string             `json:"base64image,omitempty" bson:"base64image,omitempty"`
	LogoBase64     string             `json:"logobase64,omitempty" bson:"logobase64,omitempty"`
}

type Parkiran2 struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" `
	Parkiranid     string             `json:"parkiranid,omitempty" bson:"parkiranid,omitempty"`
	Nama           string             `json:"nama,omitempty" bson:"nama,omitempty"`
	NPM            string             `json:"npm,omitempty" bson:"npm,omitempty"`
	Prodi          string             `json:"prodi,omitempty" bson:"prodi,omitempty"`
	NamaKendaraan  string             `json:"namakendaraan,omitempty" bson:"namakendaraan,omitempty"`
	NomorKendaraan string             `json:"nomorkendaraan,omitempty" bson:"nomorkendaraan,omitempty"`
	JenisKendaraan string             `json:"jeniskendaraan,omitempty" bson:"jeniskendaraan,omitempty"`
	Status         Status             `json:"status" bson:"status"`
	QRCodePath     string             `json:"qrcodepath" bson:"qrcodepath"`
}

type QRData struct {
	ParkiranID string `bson:"parkiranid"`
	Image      string `bson:"image"`
}

type Time struct {
	Message     string `json:"message,omitempty" bson:"message,omitempty"`
	WaktuMasuk  string `json:"waktumasuk,omitempty" bson:"waktumasuk,omitempty"`
	WaktuKeluar string `json:"waktukeluar,omitempty" bson:"waktukeluar,omitempty"`
}

type Status struct {
	Message     string `json:"message,omitempty" bson:"message,omitempty"`
	WaktuMasuk  string `json:"waktumasuk,omitempty" bson:"waktumasuk,omitempty"`
	WaktuKeluar string `json:"waktukeluar,omitempty" bson:"waktukeluar,omitempty"`
}

type Status2 struct {
	Status          string          `json:"status,omitempty" bson:"status,omitempty"`
	Message         string          `json:"message,omitempty" bson:"message,omitempty"`
	DataParkir      interface{}     `json:"dataparkir,omitempty" bson:"dataparkir,omitempty"`
	RequestParkiran RequestParkiran `json:"requestparkiran,omitempty" bson:"requestparkiran,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Data    string `json:"data,omitempty" bson:"data,omitempty"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type ResponseParkiran struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Parkiran `json:"data"`
}

type RequestParkiran struct {
	Parkiranid string `json:"parkiranid"`
}

type Payload struct {
	User     string    `json:"user"`
	Parkiran string    `json:"parkiran"`
	Role     string    `json:"role"`
	Exp      time.Time `json:"exp"`
	Iat      time.Time `json:"iat"`
	Nbf      time.Time `json:"nbf"`
}

type EmailValidator struct {
	regexPattern string
}
