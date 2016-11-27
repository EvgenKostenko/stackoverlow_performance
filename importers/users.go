package importers

import (
	"encoding/xml"
	"fmt"
	"github.com/EvgenKostenko/stackoverlow_performance/database"
	"github.com/EvgenKostenko/stackoverlow_performance/models"
)

type Users struct {
	Data []models.User `xml:"row"`
	BaseEntity
}

func (u *Users) LoadDataToDB() {

	file := u.Parse()
	if file == nil {
		panic("File is nil")
	}
	err := xml.Unmarshal(*file, &u)

	if err != nil {
		fmt.Println(err)
	}
	// Init db
	db := database.MgoDb{}
	db.Init()
	defer db.Close()

	for _, row := range u.Data {
		err := db.C(u.Collection).Insert(&row)
		if err != nil {
			fmt.Println(err)
		}
	}
}
