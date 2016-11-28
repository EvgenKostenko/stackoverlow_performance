package api

import (
	"errors"
	"github.com/EvgenKostenko/stackoverlow_performance/database"
	"github.com/EvgenKostenko/stackoverlow_performance/lib"
	"github.com/EvgenKostenko/stackoverlow_performance/models"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

// GET /toptags/:userid
func GetTopUsersTags(ctx *iris.Context) {

	db := database.MgoDb{}
	// Init DB
	db.Init()
	defer db.Close()

	userIDparam := ctx.Param("userid")

	userID, err := strconv.ParseInt(userIDparam, 10, 32)

	if err != nil {
		ctx.JSON(iris.StatusNotAcceptable, models.Err("Request Error."))
		return
	}

	var topTags = make(map[string]int)

	var posts []struct {
		OwnerUserId int    `bson:"OwnerUserId"`
		Tags        string `bson:"Tags"`
	}
	if err := db.C("posts").Find(bson.M{"OwnerUserId": userID}).All(&posts); err == nil {
		for _, row := range posts {
			tagsArray := strings.Split(row.Tags, "><")
			for _, tag := range tagsArray {
				cleanTag := strings.Trim(tag, "<>")
				if _, ok := topTags[cleanTag]; ok {
					topTags[cleanTag]++
				} else {
					topTags[cleanTag] = 1
				}

			}
		}
	}

	result := []models.Tag{}

	var fringe int

	if len(topTags) < 11 {
		fringe = len(topTags)
	} else {
		fringe = 11
	}

	for _, tag := range lib.SortedKeysStr(topTags)[:fringe] {
		if fullTag, err := getTag(tag.Key); err == nil {
			result = append(result, fullTag)
		}
	}

	if len(result) < 1 {
		ctx.JSON(iris.StatusOK, models.Err("NoResults"))
		return
	} else {
		ctx.JSON(iris.StatusOK, &result)
	}
}

//
func getTag(tag string) (models.Tag, error) {

	mdb := database.MgoDb{}
	mdb.Init()

	defer mdb.Close()

	thisTag := models.Tag{}
	if err := mdb.C("tags").Find(bson.M{"TagName": tag}).One(&thisTag); err == nil {
		return thisTag, nil
	}
	return models.Tag{}, errors.New("Cant get tag from DB")
}
