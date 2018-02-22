package zentaoapi

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

//New create a new Bug.
//branch and moduleID default values is 0
func New(productID, branch, moduleID int, params url.Values) ([]byte, error) {
	createBugUrl := fmt.Sprintf("%s/index.php?m=bug&f=create&productID=%d&branch=%d&moduleID=%d&sid=%s", Conf.Host, productID, branch, moduleID, sessionObj.SessionID)
	postReader := strings.NewReader(params.Encode())

	resp, err := client.Post(createBugUrl, "application/x-www-form-urlencoded", postReader)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return nil, err
	}
	respStr, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return respStr, nil
}

//BuildParamsBrief build a brief bug params
func BuildParamsBrief(product, project, module, pri int, assignedTo, bugType, title, steps string) url.Values {
	steps = fmt.Sprintf("<pre class=\"prettyprint\">%s</pre>", steps)
	return url.Values{
		"product":       {strconv.Itoa(product)},
		"project":       {strconv.Itoa(project)},
		"module":        {strconv.Itoa(module)},
		"assignedTo":    {assignedTo},
		"type":          {bugType},
		"steps":         {steps},
		"openedBuild[]": {"trunk"},
		"pri":           {strconv.Itoa(pri)},
		"title":         {title},
	}
}

//BuildParamsFull
//Note deadline format 2018-02-23
func BuildParamsFull(product, project, module, pri, severity, story, task int, assignedTo, bugType, title, steps, os, browser, color, deadline, mailto string) url.Values {
	values := BuildParamsBrief(product, project, module, pri, assignedTo, bugType, title, steps)
	values.Set("task", strconv.Itoa(task))
	values.Set("story", strconv.Itoa(story))
	values.Set("severity", strconv.Itoa(severity))
	values.Set("os", os)
	values.Set("browser", browser)
	values.Set("color", color)
	values.Set("deadline", deadline)
	values.Set("mailto[]", mailto)
	return values
}
