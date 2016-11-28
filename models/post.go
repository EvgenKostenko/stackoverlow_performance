package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	IdBson                bson.ObjectId `bson:"_id,omitempty"`
	OwnerUserId           int           `xml:"OwnerUserId,attr"  json:"OwnerUserId"  bson:"OwnerUserId"`
	Body                  string        `xml:"Body,attr"  json:"Body"  bson:"Body"`
	LastEditDate          string        `xml:"LastEditDate,attr"  json:"LastEditDate"  bson:"LastEditDate"`
	Title                 string        `xml:"Title,attr"  json:"Title"  bson:"Title"`
	CommentCount          int           `xml:"CommentCount,attr"  json:"CommentCount"  bson:"CommentCount"`
	Id                    int           `xml:"Id,attr"  json:"Id"  bson:"Id"`
	ParentId              int           `xml:"ParentId,attr"  json:"ParentId"  bson:"ParentId"`
	AnswerCount           int           `xml:"AnswerCount,attr"  json:"AnswerCount"  bson:"AnswerCount"`
	CommunityOwnedDate    string        `xml:"CommunityOwnedDate,attr"  json:"CommunityOwnedDate"  bson:"CommunityOwnedDate"`
	PostTypeId            int           `xml:"PostTypeId,attr"  json:"PostTypeId"  bson:"PostTypeId"`
	LastActivityDate      string        `xml:"LastActivityDate,attr"  json:"LastActivityDate"  bson:"LastActivityDate"`
	OwnerDisplayName      string        `xml:"OwnerDisplayName,attr"  json:"OwnerDisplayName"  bson:"OwnerDisplayName"`
	LastEditorUserId      int           `xml:"LastEditorUserId,attr"  json:"LastEditorUserId"  bson:"LastEditorUserId"`
	CreationDate          string        `xml:"CreationDate,attr"  json:"CreationDate"  bson:"CreationDate"`
	AcceptedAnswerId      int           `xml:"AcceptedAnswerId,attr"  json:"AcceptedAnswerId"  bson:"AcceptedAnswerId"`
	LastEditorDisplayName string        `xml:"LastEditorDisplayName,attr"  json:"LastEditorDisplayName"  bson:"LastEditorDisplayName"`
	Tags                  string        `xml:"Tags,attr"  json:"Tags"  bson:"Tags"`
	Score                 int           `xml:"Score,attr"  json:"Score"  bson:"Score"`
	FavoriteCount         int           `xml:"FavoriteCount,attr"  json:"FavoriteCount"  bson:"FavoriteCount"`
	ClosedDate            string        `xml:"ClosedDate,attr"  json:"ClosedDate"  bson:"ClosedDate"`
	ViewCount             int           `xml:"ViewCount,attr"  json:"ViewCount"  bson:"ViewCount"`
}
