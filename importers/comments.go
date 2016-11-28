package importers

import (
	"encoding/xml"
	"fmt"
	"github.com/EvgenKostenko/stackoverlow_performance/database"
	"github.com/EvgenKostenko/stackoverlow_performance/models"
	"gopkg.in/mgo.v2"
)

type Comments struct {
	Data []models.Comment `xml:"row"`
	BaseEntity
}

func (c *Comments) LoadDataToDB() {

	file := c.Parse()
	err := xml.Unmarshal(*file, &c)

	if err != nil {
		fmt.Println(err)
	}

	// Init db
	db := database.MgoDb{}
	db.Init()
	defer db.Close()

	for _, row := range c.Data {
		err := db.C(c.Collection).Insert(&row)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Creating index for comments")
	index := mgo.Index{
		Key: []string{"$text:Text"},
	}

	fmt.Println("Ensure index for comments")
	err = db.C(c.Collection).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
