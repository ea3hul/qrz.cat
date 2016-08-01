package main

import (
	. "github.com/clbanning/mxj"
	"io/ioutil"
	"log"
	"net/http"
)


func toUtf8(iso8859_1_buf []byte) []byte {
    buf := make([]rune, len(iso8859_1_buf))
    for i, b := range iso8859_1_buf {
        buf[i] = rune(b)
    }
    return (buf)
}


func RestGetPoblacio(φ1, λ1 float64) SPoblacio {
	return GetPoblacio(φ1, λ1)
}

func RestGetActivitatSolar() ([]byte,error){
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

func RestGetQrzEntitat(entitat string) ([]byte, error) {

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

func RestGetQrzIndicatiuBio(indicatiu string) ([]byte, error) {

	url := "http://xmldata.qrz.com/xml/current/?s=" + token + ";html=" + indicatiu
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		if b, err := ioutil.ReadAll(response.Body); err == nil {

			bio := toUtf8(b)

			return bio, err
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil, err

}

func RestGetQrzIndicatiu(indicatiu string) ([]byte, error) {

	url := "http://xmldata.qrz.com/xml/current/?s=" + token + ";callsign=" + indicatiu
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

func XmlToJson(xmlVal []byte, safeEncoding ...bool) ([]byte, error) {
	m, err := NewMapXml(xmlVal)
	if err != nil {
		return nil, err
	}
	return m.Json(safeEncoding...)
}

func RestGetQrzToken(usuari string, contrasenya string) ([]byte, error) {

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
