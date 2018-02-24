package zentaoapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type SessionObj struct {
	SessionName string
	SessionID   string
}

var sessionObj SessionObj
var client *http.Client

func getSession() {
	//jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		log.Fatal(err)
	}
	client = &http.Client{
		Jar:     jar,
		Timeout: time.Second * 5,
	}

	resp, err := client.Get(fmt.Sprintf("%s/?m=api&f=getSessionID&t=json", Conf.Host))
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
	u, err := url.Parse(Conf.Host)

	cookies := []*http.Cookie{
		&http.Cookie{Name: "keepLogin", Value: "on"},
	}

	client.Jar.SetCookies(u, cookies)
	loginUrl := fmt.Sprintf("%s/index.php?m=user&f=login&t=json&sid=%s", Conf.Host, sessionObj.SessionID)

	postReader := strings.NewReader(url.Values{"account": []string{"liangkai"}, "password": []string{"123456"}}.Encode())

	resp, err := client.Post(loginUrl, "application/x-www-form-urlencoded", postReader)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return
	}

	//respJson, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	//fmt.Printf("login respJson = %s", respJson)

	//for _, cookie := range client.Jar.Cookies(u) {
	//fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	//}
}
