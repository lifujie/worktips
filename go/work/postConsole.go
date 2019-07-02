package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/httplib"
)

// PostOobConsole 发送post请求
func postOobConsole(ip, name, password string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	userInfo := struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}{
		User:     name,
		Password: password,
	}

	userInfoStr, err := json.Marshal(userInfo)
	fmt.Printf("%s\n", string(userInfoStr))
	//
	url := fmt.Sprintf("https://%s/console", ip)
	resp, err := httplib.Post(url).Header("Content-Type", "application/json").Header("cache-control", "no-cache").Header("Connection", "keep-alive").Body(userInfoStr).SetTransport(tr).DoRequest()
	if err != nil {
		fmt.Printf("resp: %s\n", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("post log failed: %s", http.StatusText(resp.StatusCode))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	for _, c := range resp.Cookies() {
		//req.AddCookie(c)
		fmt.Printf("cookie: %v\n", c)
	}
	_ = body
	fmt.Println(string(body))
	return nil
}

func main() {
	err := postOobConsole("10.0.10.100", "root", "calvin")
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
	}
}
