package zentaoapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ZenTaoConfig struct {
	Host     string
	Account  string
	Password string
}

var conf = &ZenTaoConfig{}
var defaultLog = log.New(os.Stdout, "zentaoapi", 1)

func init() {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return
	}
	err = json.Unmarshal(raw, conf)
	if err != nil {
		fmt.Printf("err when unmarshal= %+v\n", err)
		return
	}
	fmt.Println("Config inited")
}
