package main

import (
	"github.com/zserge/lorca"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ui, err := lorca.New("", "", 950, 650)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Bind("goHttpRequest", httpRequest)
	ui.Load("data:text/html," + url.PathEscape(getHtmlText()))
	<-ui.Done()
}

// 获取界面html
func getHtmlText() string {
	data, err := ioutil.ReadFile("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// 发送http请求
func httpRequest(method, url string, headerMap map[string]string) string {
	log.Println(method, url, "header", headerMap)

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	if headerMap != nil {
		for k, v := range headerMap {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return resp.Status
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
