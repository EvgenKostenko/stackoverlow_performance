package main

import (
	"fmt"
	"github.com/EvgenKostenko/stackoverlow_performance/config"
	"github.com/EvgenKostenko/stackoverlow_performance/importers"
)

func main() {
	//loaded user to DB
	fmt.Println("Load users")
	u := importers.Users{BaseEntity: importers.BaseEntity{Collection: "users", FilePath: config.Config.DUMP.Path + "Users.xml"}}
	u.LoadDataToDB()
	fmt.Println("users loaded to db")
	//loaded tags to DB
	fmt.Println("Load tags")
	t := importers.Tags{BaseEntity: importers.BaseEntity{Collection: "tags", FilePath: config.Config.DUMP.Path + "Tags.xml"}}
	t.LoadDataToDB()
	fmt.Println("tags loaded to db")
	//loaded comments to DB
	fmt.Println("Load comments")
	c := importers.Comments{BaseEntity: importers.BaseEntity{Collection: "comments", FilePath: config.Config.DUMP.Path + "Comments.xml"}}
	c.LoadDataToDB()
	fmt.Println("comments loaded to db")
	// Load posts to DB
	fmt.Println("Load posts")
	p := importers.Posts{BaseEntity: importers.BaseEntity{Collection: "posts", FilePath: config.Config.DUMP.Path + "Posts.xml"}}
	p.LoadDataToDB()
	fmt.Println("posts loaded to db")
}
