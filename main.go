package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"io/ioutil"
	"os"
	"time"
	"github.com/olahol/melody"
	"net/http"
    
)

var (
	qau   SQrzAuth
	qrl   SQrzRespostaLogin
	token string
	pb    SPoblacions
	caducitat time.Time
	ent SQrzEntitat
)

const (
	Port = "8088"
)

func main() {
	Start()
}

// Obté el token del QRZ
func Auth() {

	var err error

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

func Entitats() ([]byte,error) {

	// Carrega del fitxer de les poblacions
	file, err := ioutil.ReadFile("./lib/entitats.json")

	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	json.Unmarshal(file, &ent)

	return file,err
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

	// Carrega entitats
	Entitats()

	router := gin.Default()
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	fmt.Println(m)

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
	
	router.GET("/ws",func(c *gin.Context) {
  		m.HandleRequest(c.Writer, c.Request)
  	})
  	
  	m.HandleMessage(func(s *melody.Session, msg []byte) {
  		m.Broadcast(msg)
  	})

	api := router.Group("/api")

	// Consultes a la api
	api.GET("/qrz/:indicatiu", ApiQrzIndicatiu)
	api.GET("/qrz/:indicatiu/bio", ApiQrzIndicatiuBio)
	api.GET("/dxcc/:entitat", ApiQrzEntitat)
	api.GET("/dxcc", ApiQrzEntitats)
	api.GET("/geo/poblacio/:lat/:lon", ApiPoblacio)
	api.GET("/solar/activitat", ApiActivitatSolar)

	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)

}
