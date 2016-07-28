package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ApiPoblacio(c *gin.Context) {
	lat := c.Params.ByName("lat")
	lon := c.Params.ByName("lon")

	flat, _ := strconv.ParseFloat(lat, 64)
	flon, _ := strconv.ParseFloat(lon, 64)

	resp := GetPoblacio(flat, flon)

	b, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("error:", err)
	}

	c.Data(200, "application/json", b)

	/*} else {
		c.JSON(404, gin.H{"comunitat": "no s'ha trobat"})
	}*/

}

// Retorna el token de validacio
func ApiQrzIndicatiu(c *gin.Context) {

	indicatiu := c.Params.ByName("indicatiu")

	resp, err := RestGetQrzIndicatiu(indicatiu)

	if err == nil {
		c.Data(200, "application/json", resp)
	} else {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

}

func ApiQrzEntitat(c *gin.Context) {

	entitat := c.Params.ByName("entitat")

	resp, err := RestGetQrzEntitat(entitat)

	if err == nil {
		c.Data(200, "application/json", resp)
	} else {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

}

func ApiQrzIndicatiuBio(c *gin.Context) {

	indicatiu := c.Params.ByName("indicatiu")

	resp, err := RestGetQrzIndicatiuBio(indicatiu)

	if err == nil {
		c.Data(200, "text/html", resp)
	} else {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

}
