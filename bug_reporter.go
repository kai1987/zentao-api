package zentaoapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//host/index.php?m=bug&f=create&productID=4&branch=0&extra=moduleID=0
func New(productID, branch, moduleID int, params url.Values) ([]byte, error) {
	resp, err := http.PostForm(fmt.Sprintf("%s/index.php?m=bug&f=create&productID=%d&branch=%d&moduleID=%d&sid=%s", conf.Host, productID, branch, moduleID, sessionObj.SessionID), params)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return nil, err
	}
	respStr, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return respStr, nil
}

func BuildParams(product, project, module int, assignedTo, bugType, steps string) url.Values {
	return url.Values{
		"product":    {strconv.Itoa(product)},
		"project":    {strconv.Itoa(project)},
		"module":     {strconv.Itoa(module)},
		"assignedTo": {assignedTo},
		"type":       {bugType},
		"steps":      {steps},
	}
}
