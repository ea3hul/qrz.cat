package main

import (
_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"fmt"
)

type Database struct{
	Db *gorm.DB
	Connectat bool
}

func (database *Database) Connecta (usuari,contrasenya,basedades string) error{

	var err error

	conn := (usuari+":"+contrasenya+"@/"+basedades+"?charset=utf8&parseTime=True&loc=Local")

	database.Db,err = gorm.Open("mysql",conn)

	if err != nil {
		database.Connectat = false
		
	}else{
		database.Connectat = true
		fmt.Println("Connectat a la base de dades: ",basedades)
	}

	return err
}

func (database *Database) InsQrz (indicatiu Callsign) error {

	database.Db.AutoMigrate(&Callsign{})

	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	database.Db.NewRecord(indicatiu) // => returns `true` as primary key is blank

	database.Db.Create(&indicatiu)

	return nil
}

func (database *Database) GetQrz (indicatiu string, qrz Callsign) error {
	database.Db.Where("call = ?", indicatiu).First(&qrz)
	return nil
}


