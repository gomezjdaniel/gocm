package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var configFlag = flag.String("config", "", "config path")

type Params struct {
	ApiKey string
	Port   int
}

func main() {
	flag.Parse()

	if *configFlag == "" {
		fmt.Println("argument required")
		flag.Usage()
		return
	}

	configFile, err := os.Open(*configFlag)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer configFile.Close()

	d := json.NewDecoder(configFile)
	var config *Params

	if err := d.Decode(config); err != nil {
		fmt.Println(err)
		return
	}

	if config.ApiKey == "" || config.Port < 8000 {
		fmt.Println("ApiKey is empty or the port name is incorrect (Port >= 8000)")
		return
	}

}
