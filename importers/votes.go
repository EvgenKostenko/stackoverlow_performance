package importers

import (
	"encoding/xml"
	"fmt"
	"github.com/EvgenKostenko/stackoverlow_performance/database"
	"github.com/EvgenKostenko/stackoverlow_performance/models"
)

type Votes struct {
	Data []models.Vote `xml:"row"`
	BaseEntity
}

func (v *Votes) LoadDataToDB() {

	file := v.Parse()
	err := xml.Unmarshal(*file, &v)

	if err != nil {
		fmt.Println(err)
	}
	// Init db
	db := database.MgoDb{}
	db.Init()
	defer db.Close()

	for _, row := range v.Data {
		err := db.C(v.Collection).Insert(&row)
		if err != nil {
			fmt.Println(err)
		}
	}
}
