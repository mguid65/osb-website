package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mwalto7/osb-website/database"
)

func main() {
	data, err := ioutil.ReadFile("./test.json")
	if err != nil {
		log.Fatal(err)
	}

	var specs database.Specs
	if err := json.Unmarshal(data, &specs); err != nil {
		log.Fatal(err)
	}
	fmt.Println(specs.SysInfo)
}
