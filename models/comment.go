package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
	IdBson          bson.ObjectId `bson:"_id,omitempty"`
	UserId          int           `xml:"UserId,attr" json:"UserId" bson:"UserId"`
	Id              int           `xml:"Id,attr" json:"Id" bson:"Id"`
	PostId          int           `xml:"PostId,attr" json:"PostId" bson:"PostId"`
	Score           int           `xml:"Score,attr" json:"Score" bson:"Score"`
	Text            string        `xml:"Text,attr" json:"Text" bson:"Text"`
	CreationDate    string        `xml:"CreationDate,attr" bson:"CreationDate" json:"CreationDate"`
	UserDisplayName string        `xml:"UserDisplayName,attr" bson:"UserDisplayName" json:"UserDisplayName"`
}
