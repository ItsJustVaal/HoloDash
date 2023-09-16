package controllers

import (
	"log"

	"github.com/ItsJustVaal/Holodash/database"
)

type Channel struct {
	UniqueId int
	Name     string
	Id       string
	Region   string
}

func HandleGetAllChannels() []Channel {
	db, err := database.ConnectDatase()

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Db.Close()

	channels, err := db.Db.Query("SELECT * FROM channels")

	if err != nil {
		log.Fatalln(err)
	}

	var channel Channel
	var channelMap []Channel

	for channels.Next() {
		e := channels.Scan(&channel.UniqueId, &channel.Id, &channel.Name, &channel.Region)
		if e != nil {
			log.Fatalln(e)
		}

		channelMap = append(channelMap, channel)
	}

	return channelMap

}

// ctx *gin.Context
