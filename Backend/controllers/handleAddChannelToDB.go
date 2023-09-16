package controllers

import (
	"fmt"
	"log"

	"github.com/ItsJustVaal/Holodash/database"
)

func HandleAddChannelToDB(c Channel) {
	data, err := database.ConnectDatase()

	if err != nil {
		log.Fatalln(err)
	}
	executeString := fmt.Sprintf("INSERT INTO channels(channel_id, channel_name, region) VALUES (%v, %v, %v)", c.Id, c.Name, c.Region)
	res, err := data.Db.Exec(executeString)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.RowsAffected())
}
