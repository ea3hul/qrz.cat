package main

import (
	. "github.com/clbanning/mxj"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

type RestClient struct {}


func (rest *RestClient) RestGetPoblacio(φ1, λ1 float64) SPoblacio {
	return GetPoblacio(φ1, λ1)
}

func (rest *RestClient) RestGetActivitatSolar() ([]byte,error){
	url := "http://www.hamqsl.com/solarxml.php"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		if b, err := ioutil.ReadAll(response.Body); err == nil {

			data, err := XmlToJson(b)

			return data, err
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil, err
}

func (rest *RestClient) RestGetQrzEntitat(entitat string) ([]byte, error) {

	url := "http://xmldata.qrz.com/xml/current/?s=" + token + ";dxcc=" + entitat
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		if b, err := ioutil.ReadAll(response.Body); err == nil {

			data, err := XmlToJson(b)

			return data, err
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil, err
}

func (rest *RestClient) RestGetQrzEntitats() ([]byte, error) {

	entitats,err := Entitats()

	return entitats,err
}

func (rest *RestClient) RestGetQrzIndicatiuBio(indicatiu string) ([]byte, error) {

	url := "http://xmldata.qrz.com/xml/current/?s=" + token + ";html=" + indicatiu
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		if bio, err := ioutil.ReadAll(response.Body); err == nil {

			return bio, err
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil, err

}

func (rest *RestClient) RestGetQrzIndicatiu(indicatiu string) ([]byte, error) {

	url := "http://xmldata.qrz.com/xml/current/?s=" + token + ";callsign=" + indicatiu
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		if b, err := ioutil.ReadAll(response.Body); err == nil {

			var la SQrzRespostaIndicatiu

			data, err := XmlToJson(b)

			json.Unmarshal(data,&la)

			//rest.database.InsQrz(la.QRZDatabase.Callsign)

			return data, err
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil, err

}

func XmlToJson(xmlVal []byte, safeEncoding ...bool) ([]byte, error) {
	m, err := NewMapXml(xmlVal)
	if err != nil {
		return nil, err
	}
	return m.Json(safeEncoding...)
}

func (rest *RestClient) RestGetQrzToken(usuari string, contrasenya string) ([]byte, error) {

	url := "http://xmldata.qrz.com/xml/current/?username=" + usuari + ";password=" + contrasenya
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		if b, err := ioutil.ReadAll(response.Body); err == nil {

			json, err := XmlToJson(b)

			return json, err
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil, err
}
