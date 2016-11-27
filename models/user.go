package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	IdBson          bson.ObjectId `bson:"_id,omitempty" json:"-"`
	WebsiteUrl      string        `xml:"WebsiteUrl,attr" json:"WebsiteUrl" bson:"WebsiteUrl"`
	UpVotes         int           `xml:"UpVotes,attr" json:"UpVotes" bson:"UpVotes"`
	Id              int           `xml:"Id,attr" json:"Id" bson:"Id"`
	Age             int           `xml:"Age,attr" json:"Age" bson:"Age"`
	DisplayName     string        `xml:"DisplayName,attr" json:"DisplayName" bson:"DisplayName"`
	AboutMe         string        `xml:"AboutMe,attr" json:"AboutMe" bson:"AboutMe"`
	Reputation      int           `xml:"Reputation,attr" bson:"Reputation" bson:"Reputation"`
	LastAccessDate  string        `xml:"LastAccessDate,attr" json:"LastAccessDate" bson:"LastAccessDate"`
	DownVotes       int           `xml:"DownVotes,attr" json:"DownVotes" bson:"DownVotes"`
	AccountId       int           `xml:"AccountId,attr" json:"AccountId" bson:"AccountId"`
	Location        string        `xml:"Location,attr" xml:"Location" xml:"Location"`
	Views           int           `xml:"Views,attr" json:"Views" bson:"Views"`
	CreationDate    string        `xml:"CreationDate,attr" json:"CreationDate" bson:"CreationDate"`
	ProfileImageUrl string        `xml:"ProfileImageUrl,attr" bson:"ProfileImageUrl" json:"ProfileImageUrl"`
	PostCount       int           `json:"PostCount,omitempty"`
}
