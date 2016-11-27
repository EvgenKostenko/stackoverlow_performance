package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Vote struct {
	IdBson       bson.ObjectId `bson:"_id,omitempty"`
	Id           int           `xml:"Id,attr" json:"Id" bson:"Id"`
	PostId       int           `xml:"PostId,attr" json:"PostId" bson:"PostId"`
	VoteTypeId   int           `xml:"VoteTypeId,attr" json:"VoteTypeId" bson:"VoteTypeId"`
	CreationDate string        `xml:"CreationDate,attr" json:"CreationDate" bson:"CreationDate"`
	UserId       int           `xml:"UserId,attr" json:"UserId" bson:"UserId"`
}
