package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"io/ioutil"
	"os"
	"time"
)

var (
	qau   SQrzAuth
	qrl   SQrzRespostaLogin
	token string
	pb    SPoblacions
)

const (
	Port = "8088"
)

func main() {
	Start()
}

// Obté el token del QRZ
func Auth() {

	// Carrega del fitxer auth per qrz.com
	file, err := ioutil.ReadFile("qrzauth.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(file, &qau)

	res, err := RestGetQrzToken(qau.Usuari, qau.Contrasenya)

	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(res, &qrl)

	token = qrl.QRZDatabase.Session.Key

	fmt.Println("Token actual >", token)

}

func Entitats() {

}

func Poblacions() {

	// Carrega del fitxer de les poblacions
	file, err := ioutil.ReadFile("./lib/poblacions.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(file, &pb)

}

// Funció principal
func Start() {

	// Autorització QRZ
	Auth()

	// Carrega poblacions
	Poblacions()

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.Static("/public", "./public")
	router.StaticFile("/favicon.ico", "./public/images/favicon.png")

	api := router.Group("/api")

	// Consultes a la api
	api.GET("/qrz/:indicatiu", ApiQrzIndicatiu)
	api.GET("/qrz/:indicatiu/bio", ApiQrzIndicatiuBio)
	api.GET("/dxcc/:entitat", ApiQrzEntitat)
	api.GET("/geo/poblacio/:lat/:lon", ApiPoblacio)

	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)

}
