package main

import (
	"github.com/alouche/go-geolib"
)

func GetDistancia(φ1, λ1, φ2, λ2 float64) float64 {

	//return geolib.EquirectangularDistance(φ1, λ1, φ2, λ2)
	return geolib.HaversineDistance(φ1, λ1, φ2, λ2)
}

func GetPoblacio(φ1, λ1 float64) SPoblacio {

	var distant float64
	var poblacio SPoblacio

	distant = 300000.0

	for _, elem := range pb {
		distancia := GetDistancia(φ1, λ1, elem.Lat, elem.Lon)

		if distancia <= distant {
			distant = distancia
			poblacio = elem
		}

	}

	return poblacio

}
