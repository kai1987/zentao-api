package zentaoapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

type SessionObj struct {
	SessionName string
	SessionID   string
}

var sessionObj SessionObj
var client *http.Client

func getSession() {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}
	client = &http.Client{
		Jar: jar,
	}

	resp, err := client.Get(fmt.Sprintf("%s/?m=api&f=getSessionID&t=json", conf.Host))
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return
	}

	respJson, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	jsonObj := struct {
		Status string
		Data   string
		Md5    string
	}{}

	err = json.Unmarshal(respJson, &jsonObj)
	if err != nil {
		defaultLog.Printf("err:%v", err)
	}
	err = json.Unmarshal([]byte(jsonObj.Data), &sessionObj)
	if err != nil {
		defaultLog.Printf("err:%v", err)
	}

}

func Login() {
	getSession()
	u, err := url.Parse(conf.Host)

	cookies := []*http.Cookie{
		&http.Cookie{Name: "keepLogin", Value: "on"},
		&http.Cookie{Name: "za", Value: "liangkai"},
		&http.Cookie{Name: "zp", Value: "1db9171473f8cd6d3712895b1ea3e06c262c5a5e"},
	}

	client.Jar.SetCookies(u, cookies)

	resp, err := client.Get(fmt.Sprintf("%s/index.php?m=user&f=login&t=json&sid=%s", conf.Host, sessionObj.SessionID))
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return
	}

	respJson, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("respJson = %s", respJson)

	for _, cookie := range client.Jar.Cookies(u) {
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}
}
