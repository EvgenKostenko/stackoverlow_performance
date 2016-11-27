package importers

import (
	"fmt"
	"io/ioutil"
)

type BaseEntity struct {
	FilePath   string
	Collection string
}

func (b *BaseEntity) Parse() *[]byte {
	xmlFile, err := ioutil.ReadFile(b.FilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	return &xmlFile
}
