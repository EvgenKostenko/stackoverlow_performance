package importers

import (
	"encoding/xml"
	"fmt"
	"github.com/EvgenKostenko/stackoverlow_performance/database"
	"github.com/EvgenKostenko/stackoverlow_performance/models"
)

type Tags struct {
	Data []models.Tag `xml:"row"`
	BaseEntity
}

func (t *Tags) LoadDataToDB() {

	file := t.Parse()
	err := xml.Unmarshal(*file, &t)

	if err != nil {
		fmt.Println(err)
	}

	// Init db
	db := database.MgoDb{}
	db.Init()
	defer db.Close()

	for _, row := range t.Data {
		err := db.C(t.Collection).Insert(&row)
		if err != nil {
			fmt.Println(err)
		}
	}
}
