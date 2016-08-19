package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

type Rutes struct {
	rest *RestClient
}

func (r *Rutes) Init (){
	r.rest = new(RestClient)
}

func (r *Rutes) ApiActivitatSolar(c *gin.Context){

	resp,_ := r.rest.RestGetActivitatSolar()

	c.Data(200, "application/json", resp)
}

func (r *Rutes) ApiPoblacio(c *gin.Context) {
	lat := c.Params.ByName("lat")
	lon := c.Params.ByName("lon")

	flat, _ := strconv.ParseFloat(lat, 64)
	flon, _ := strconv.ParseFloat(lon, 64)

	resp := r.rest.RestGetPoblacio(flat, flon)

	b, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("error:", err)
	}

	c.Data(200, "application/json", b)

}

// Retorna el token de validacio
func (r *Rutes) ApiQrzIndicatiu(c *gin.Context) {

	indicatiu := c.Params.ByName("indicatiu")

	resp, err := r.rest.RestGetQrzIndicatiu(indicatiu)

	if err == nil {
		c.Data(200, "application/json", resp)
	} else {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

}

func (r *Rutes) ApiQrzEntitat(c *gin.Context) {

	entitat := c.Params.ByName("entitat")

	resp, err := r.rest.RestGetQrzEntitat(entitat)

	if err == nil {
		c.Data(200, "application/json", resp)
	} else {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

}

func (r *Rutes) ApiQrzEntitats(c *gin.Context) {

	entitats,err := r.rest.RestGetQrzEntitats()

	if err == nil {
		c.Data(200, "application/json", entitats)
	} else {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

}


func (r *Rutes) ApiQrzIndicatiuBio(c *gin.Context) {

	indicatiu := c.Params.ByName("indicatiu")

	resp, err := r.rest.RestGetQrzIndicatiuBio(indicatiu)

	if err == nil {
		c.Data(200, "text/html", resp)
	} else {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

}
