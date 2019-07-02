package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego/httplib"
)

// Collected co
type Collected struct {
	SN string
}

// Agent ag
type Agent struct {
	logPath    string
	ServerAddr string
	collected  Collected
}

// loadData 加载日志
func (agent *Agent) loadData() ([]byte, error) {
	data, err := ioutil.ReadFile(agent.logPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data, nil
}

// APIVersion ap
var APIVersion = "v1"

// PostLog 向服务端发送日志
func (agent *Agent) PostLog() (err error) {
	data, err := agent.loadData()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/api/cloudboot/%s/devices/%s/components/%s/logs?lang=en-US",
		agent.ServerAddr, APIVersion, agent.collected.SN, "cloudboot-agent")

	resp, err := httplib.Post(url).Header("Accept", "application/json").Body(data).DoRequest()
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("post log failed: %s", http.StatusText(resp.StatusCode))
	}
	return nil
}

func main() {
	agent := Agent{
		logPath:    "./agent_log.go",
		ServerAddr: "http://10.0.106.27:8083",
		collected: Collected{
			SN: "3Q28132",
		},
	}
	agent.PostLog()
}
