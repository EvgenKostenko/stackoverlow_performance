package api

import (
	"github.com/EvgenKostenko/stackoverlow_performance/database"
	"github.com/EvgenKostenko/stackoverlow_performance/lib"
	"github.com/EvgenKostenko/stackoverlow_performance/models"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

// GET /topusers/:word
func GetTopUsersByWord(ctx *iris.Context) {

	db := database.MgoDb{}
	// Init DB
	db.Init()

	word := ctx.Param("word")

	var topUsersIds = make(map[int]int)

	var posts []struct {
		OwnerUserId int `bson:"OwnerUserId"`
	}
	if err := db.C("posts").Pipe([]bson.M{{"$match": bson.M{"$text": bson.M{"$search": word}}}}).All(&posts); err == nil {
		for _, row := range posts {
			if row.OwnerUserId > 0 {
				if _, ok := topUsersIds[row.OwnerUserId]; ok {
					topUsersIds[row.OwnerUserId]++
				} else {
					topUsersIds[row.OwnerUserId] = 1
				}
			}
		}
	}

	var comments []struct {
		UserId int `bson:"UserId"`
	}
	if err := db.C("comments").Pipe([]bson.M{{"$match": bson.M{"$text": bson.M{"$search": word}}}}).All(&comments); err == nil {
		for _, row := range comments {
			if row.UserId > 0 {
				if _, ok := topUsersIds[row.UserId]; ok {
					topUsersIds[row.UserId]++
				} else {
					topUsersIds[row.UserId] = 1
				}
			}
		}
	}

	result := []models.User{}

	if len(topUsersIds) < 10 {
		result = getUsersList(lib.SortedKeysInt(topUsersIds), len(topUsersIds), &topUsersIds)
	} else {
		result = getUsersList(lib.SortedKeysInt(topUsersIds), 10, &topUsersIds)
	}

	if len(result) < 1 {
		ctx.JSON(iris.StatusOK, models.Err("NoResults"))
		return
	} else {
		ctx.JSON(iris.StatusOK, &result)
	}

	//Close db
	db.Close()
}

func getUser(id int) models.User {

	mdb := database.MgoDb{}
	mdb.Init()

	defer mdb.Close()

	thisUser := models.User{}
	if err := mdb.C("users").Find(bson.M{"Id": id}).One(&thisUser); err == nil {
		return thisUser
	}
	return models.User{}
}

func getUsersList(listKeys lib.PairIntList, len int, topUsersIds *map[int]int) []models.User {
	resultUsersList := []models.User{}
	for _, res := range listKeys[:len] {
		resultUsersList = append(resultUsersList, getUser(res.Key))
	}
	return resultUsersList

}
