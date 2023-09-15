package main

import (
	"fmt"

	"github.com/ItsJustVaal/HolodashGo3/database"
)

func main() {

	fmt.Println("Attempting to connect to DB")
	database.ConnectDatase()
	chs := TempLoadCSV()

	for _,ch := range chs {
		fmt.Println(ch)
	}


}