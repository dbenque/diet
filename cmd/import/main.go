package main

import (
	"fmt"
	"os"

	"github.com/dbenque/diet/db"
)

func main() {
	database, err := db.Open("/home/david/code/diet/db/diet.db3")
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	fmt.Println("---")
	database.Import()
}
