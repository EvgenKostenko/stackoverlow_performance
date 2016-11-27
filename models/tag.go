package models

import "gopkg.in/mgo.v2/bson"

type Tag struct {
	IdBson        bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Count         int           `xml:"Count,attr" json:"Count" bson:"Count"`
	ExcerptPostId int           `xml:"ExcerptPostId,attr" json:"ExcerptPostId" bson:"ExcerptPostId"`
	WikiPostId    int           `xml:"WikiPostId,attr" json:"WikiPostId" bson:"WikiPostId"`
	Id            int           `xml:"Id,attr" json:"Id" bson:"Id"`
	TagName       string        `xml:"TagName,attr" json:"TagName" bson:"TagName"`
}
