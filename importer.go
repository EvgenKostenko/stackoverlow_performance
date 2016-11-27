package main

import (
	"fmt"
	"github.com/EvgenKostenko/stackoverlow_performance/importers"
)

func main() {
	//loaded user to DB
	fmt.Println("Load users")
	u := importers.Users{BaseEntity: importers.BaseEntity{Collection: "users", FilePath: "scifi.stackexchange.com/Users.xml"}}
	u.LoadDataToDB()
	fmt.Println("users loaded to db")
	//loaded tags to DB
	t := importers.Tags{BaseEntity: importers.BaseEntity{Collection: "tags", FilePath: "scifi.stackexchange.com/Tags.xml"}}
	t.LoadDataToDB()
	fmt.Println("tags loaded to db")
	//loaded comments to DB
	c := importers.Comments{BaseEntity: importers.BaseEntity{Collection: "comments", FilePath: "scifi.stackexchange.com/Comments.xml"}}
	c.LoadDataToDB()
	fmt.Println("comments loaded to db")
	// Load posts to DB
	p := importers.Posts{BaseEntity: importers.BaseEntity{Collection: "posts", FilePath: "./scifi.stackexchange.com/Posts.xml"}}
	p.LoadDataToDB()
	fmt.Println("posts loaded to db")
	//Load votes to DB
	v := importers.Votes{BaseEntity: importers.BaseEntity{Collection: "votes", FilePath: "scifi.stackexchange.com/Votes.xml"}}
	v.LoadDataToDB()
	fmt.Println("votes loaded to db")
}
