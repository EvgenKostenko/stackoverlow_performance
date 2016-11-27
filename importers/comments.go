package importers

import (
	"encoding/xml"
	"fmt"
	"github.com/EvgenKostenko/stackoverlow_performance/database"
	"github.com/EvgenKostenko/stackoverlow_performance/models"
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
}
