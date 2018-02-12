package zentaoapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var sessionID, sessionName string

func getSession() {

	resp, err := http.Get(fmt.Sprintf("%s/?m=api&f=getSessionID&t=json", conf.Host))
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return
	}

	respJson, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("respJson = %+v\n", respJson)
	defer resp.Body.Close()

	jsonObj := struct {
		status string
		data   map[string]string
		md5    string
	}{}

	err = json.Unmarshal(respJson, &jsonObj)
	if err != nil {
		defaultLog.Printf("err:%v", err)
	}
	fmt.Printf("m = %+v\n", jsonObj)
}
