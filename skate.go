package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type SkateResponse struct {
	Html string
}

func parseSkateResponse(inlineVersion string) string {
	str := inlineVersion
	res := &SkateResponse{}
	json.Unmarshal([]byte(str), &res)
	return res.Html
}

func getInlineVersion(html string) string {
	data := url.Values{}
	data.Set("source", html)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://zurb.com/ink/skate-proxy.php", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Referer", "http://github.com/aitor/skate")

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return parseSkateResponse(string(body))
}

func checkReading(e error) {
	if e != nil {
		panic(e)
	}
}

func readSourceFromStdin() string {
	content, err := ioutil.ReadAll(os.Stdin)
	checkReading(err)
	return string(content)
}

func readSourceFromFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	checkReading(err)
	return string(content)
}

func main() {
	var source string

	flag.Parse()
	switch name := flag.Arg(0); {
	case name == "":
		source = readSourceFromStdin()
	default:
		source = readSourceFromFile(name)
	}
	fmt.Print(getInlineVersion(source))
}
