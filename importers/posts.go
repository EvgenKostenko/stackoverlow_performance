package importers

import (
	"encoding/xml"
	"fmt"
	"github.com/EvgenKostenko/stackoverlow_performance/database"
	"github.com/EvgenKostenko/stackoverlow_performance/models"
	"gopkg.in/mgo.v2"
)

type Posts struct {
	Data []models.Post `xml:"row"`
	BaseEntity
}

func (p *Posts) LoadDataToDB() {

	file := p.Parse()
	err := xml.Unmarshal(*file, &p)

	if err != nil {
		fmt.Println(err)
	}

	// Init db
	db := database.MgoDb{}
	db.Init()
	defer db.Close()

	for _, row := range p.Data {
		err := db.C(p.Collection).Insert(&row)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Creating index for posts")
	index := mgo.Index{
		Key: []string{"$text:Body"},
	}

	fmt.Println("Ensure index for posts")
	err = db.C(p.Collection).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
