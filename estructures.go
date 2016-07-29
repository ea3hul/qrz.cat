package main

type SActivitatSolar struct {
	Solar struct {
		Solardata struct {
			Aindex string `json:"aindex,omitempty"`
			Aurora string `json:"aurora,omitempty"`
			Calculatedconditions struct {
				Band []struct {
					Text string `json:"#text,omitempty"`
					Name string `json:"-name,omitempty"`
					Time string `json:"-time,omitempty"`
				} `json:"band,omitempty"`
			} `json:"calculatedconditions,omitempty"`
			Calculatedvhfconditions struct {
				Phenomenon []struct {
					Text string `json:"#text,omitempty"`
					Location string `json:"-location,omitempty"`
					Name string `json:"-name,omitempty"`
				} `json:"phenomenon,omitempty"`
			} `json:"calculatedvhfconditions,omitempty"`
			Electonflux string `json:"electonflux,omitempty"`
			Fof2 string `json:"fof2,omitempty"`
			Geomagfield string `json:"geomagfield,omitempty"`
			Heliumline string `json:"heliumline,omitempty"`
			Kindex string `json:"kindex,omitempty"`
			Kindexnt string `json:"kindexnt,omitempty"`
			Latdegree string `json:"latdegree,omitempty"`
			Magneticfield string `json:"magneticfield,omitempty"`
			Muf string `json:"muf,omitempty"`
			Muffactor string `json:"muffactor,omitempty"`
			Normalization string `json:"normalization,omitempty"`
			Protonflux string `json:"protonflux,omitempty"`
			Signalnoise string `json:"signalnoise,omitempty"`
			Solarflux string `json:"solarflux,omitempty"`
			Solarwind string `json:"solarwind,omitempty"`
			Source struct {
				Text string `json:"#text,omitempty"`
				URL string `json:"-url,omitempty"`
			} `json:"source,omitempty"`
			Sunspots string `json:"sunspots,omitempty"`
			Updated string `json:"updated,omitempty"`
			Xray string `json:"xray,omitempty"`
		} `json:"solardata,omitempty"`
	} `json:"solar,omitempty"`
}
type SQrzAuth struct {
	Usuari      string `json:"username"`
	Contrasenya string `json:"password"`
}

type SQrzRespostaLogin struct {
	QRZDatabase struct {
		Version string `json:"-version"`
		Xmlns   string `json:"-xmlns"`
		Session struct {
			Count  string `json:"Count,omitempty"`
			GMTime string `json:"GMTime,omitempty"`
			Key    string `json:"Key,omitempty"`
			Remark string `json:"Remark,omitempty"`
			SubExp string `json:"SubExp,omitempty"`
			Error  string `json:"Error,omitempty"`
		} `json:"Session"`
	} `json:"QRZDatabase"`
}

type SQrzRespostaIndicatiu struct {
	QRZDatabase struct {
		Version  string `json:"-version"`
		Xmlns    string `json:"-xmlns"`
		Callsign struct {
			AreaCode  string `json:"AreaCode,omitempty"`
			DST       string `json:"DST,omitempty"`
			GMTOffset string `json:"GMTOffset,omitempty"`
			MSA       string `json:"MSA,omitempty"`
			TimeZone  string `json:"TimeZone,omitempty"`
			Addr1     string `json:"addr1,omitempty"`
			Addr2     string `json:"addr2,omitempty"`
			Aliases   string `json:"aliases,omitempty"`
			Bio       string `json:"bio,omitempty"`
			Biodate   string `json:"biodate,omitempty"`
			Born      string `json:"born,omitempty"`
			Call      string `json:"call,omitempty"`
			Ccode     string `json:"ccode,omitempty"`
			Class     string `json:"class,omitempty"`
			Codes     string `json:"codes,omitempty"`
			Country   string `json:"country,omitempty"`
			County    string `json:"county,omitempty"`
			Cqzone    string `json:"cqzone,omitempty"`
			Dxcc      string `json:"dxcc,omitempty"`
			Efdate    string `json:"efdate,omitempty"`
			Email     string `json:"email,omitempty"`
			Eqsl      string `json:"eqsl,omitempty"`
			Expdate   string `json:"expdate,omitempty"`
			Fips      string `json:"fips,omitempty"`
			Fname     string `json:"fname,omitempty"`
			Geoloc    string `json:"geoloc,omitempty"`
			Grid      string `json:"grid,omitempty"`
			Image     string `json:"image,omitempty"`
			Imageinfo string `json:"imageinfo,omitempty"`
			Iota      string `json:"iota,omitempty"`
			Ituzone   string `json:"ituzone,omitempty"`
			Land      string `json:"land,omitempty"`
			Lat       string `json:"lat,omitempty"`
			Lon       string `json:"lon,omitempty"`
			Lotw      string `json:"lotw,omitempty"`
			Moddate   string `json:"moddate,omitempty"`
			Mqsl      string `json:"mqsl,omitempty"`
			Name      string `json:"name,omitempty"`
			PCall     string `json:"p_call,omitempty"`
			Qslmgr    string `json:"qslmgr,omitempty"`
			Serial    string `json:"serial,omitempty"`
			State     string `json:"state,omitempty"`
			UViews    string `json:"u_views,omitempty"`
			Url       string `json:"url,omitempty"`
			User      string `json:"user,omitempty"`
			Xref      string `json:"xref,omitempty`
			Zip       string `json:"zip,omitempty"`
		} `json:"Callsign,omitempty"`
		Session struct {
			Count  string `json:"Count,omitempty"`
			Error  string `json:"Error,omitempty"`
			GMTime string `json:"GMTime,omitempty"`
			Key    string `json:"Key,omitempty"`
			Remark string `json:"Remark,omitempty"`
			SubExp string `json:"SubExp,omitempty"`
		} `json:"Session"`
	} `json:"QRZDatabase"`
}

type SQrzEntitat struct {
	QRZDatabase struct {
		Version string `json:"-version"`
		Xmlns   string `json:"-xmlns"`
		DXCC    []struct {
			Cqzone    string `json:"cqzone"`
			Dxcc      string `json:"dxcc"`
			Ituzone   string `json:"ituzone"`
			Lat       string `json:"lat"`
			Lon       string `json:"lon"`
			Name      string `json:"name"`
			Timezone  string `json:"timezone"`
			Cc        string `json:"cc,omitempty"`
			Ccc       string `json:"ccc,omitempty"`
			Continent string `json:"continent,omitempty"`
			Notes     string `json:"notes,omitempty"`
		} `json:"DXCC,omitempty"`
		Session struct {
			Count  string `json:"Count,omitempty"`
			Error  string `json:"Error,omitempty"`
			GMTime string `json:"GMTime,omitempty"`
			Key    string `json:"Key,omitempty"`
			Remark string `json:"Remark,omitempty"`
			SubExp string `json:"SubExp,omitempty"`
		} `json:"Session"`
	} `json:"QRZDatabase"`
}

type SPoblacio struct {
	Comunitat string  `json:"comunitat,omitempty"`
	Provincia string  `json:"provincia,omitempty"`
	Poblacio  string  `json:"poblacio,omitempty"`
	Lat       float64 `json:"lat,omitempty"`
	Lon       float64 `json:"lon,omitempty"`
	Alt       float64 `json:"alt,omitempty"`
}

type SPoblacions []struct {
	Comunitat string  `json:"comunitat,omitempty"`
	Provincia string  `json:"provincia,omitempty"`
	Poblacio  string  `json:"poblacio,omitempty"`
	Lat       float64 `json:"lat,omitempty"`
	Lon       float64 `json:"lon,omitempty"`
	Alt       float64 `json:"alt,omitempty"`
}
