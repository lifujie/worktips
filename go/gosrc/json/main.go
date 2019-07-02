package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	data := `[{
		"ipArr": [{
			"vlan": "",
			"relationDescribe": "",
			"port": "",
			"Ip": "192.168.1.52",
			"Info": "",
			"switch": ""
		}, {
			"vlan": "",
			"relationDescribe": "",
			"port": "",
			"Ip": "",
			"Info": "",
			"switch": ""
		}, {
			"vlan": "",
			"relationDescribe": "",
			"port": "",
			"Ip": "",
			"Info": "",
			"switch": ""
		}, {
			"vlan": "",
			"relationDescribe": "",
			"port": "",
			"Ip": "",
			"Info": "",
			"switch": ""
		}, {
			"vlan": "",
			"relationDescribe": "",
			"port": "",
			"Ip": "192.168.1.53",
			"Info": "",
			"switch": ""
		}],
		"installOS": "ture",
		"installSoft": "ture",
		"ResPoolName": "",
		"assetNumber": "yy00000007",
		"MachineName": "马场电信机房",
		"CabinetNumber": "马场电信机房马场仓库B01",
		"ResPoolID": "",
		"powerAccess": "ture",
		"CabinetId": "299",
		"sn": "4f6e8e10fds",
		"remarks": "",
		"markedTag": "ture",
		"EquipLocation": "ddddd"
	}, 
	{
		"ipArr": [{
			"vlan": "",
			"relationDescribe": "",
			"port": "",
			"Ip": "192.168.1.43",
			"Info": "",
			"switch": ""
		}],
		"installOS": "ture",
		"installSoft": "ture",
		"ResPoolName": "",
		"assetNumber": "",
		"MachineName": "",
		"CabinetNumber": "马场电信机房马场仓库B01",
		"ResPoolID": "",
		"powerAccess": "ture",
		"CabinetId": "299",
		"sn": "3Q28132",
		"remarks": "",
		"markedTag": "ture",
		"EquipLocation": "ddddd"
	}]`
	type IPS struct {
		IpArr []struct {
			Info             string
			Ip               string
			Port             string
			RelationDescribe string
			Switch           string
			Vlan             string
		} `json:"ipArr"`
		SN            string `json:"sn"`
		CabinetNumber string
		CabinetID     string
		MachineName   string
		ResPoolName   string
		EquipLocation string `json:"EquipLocation"`
	}
	var inputParams []IPS
	err := json.Unmarshal([]byte(data), &inputParams)
	if err != nil {
		fmt.Printf("%v", err)
	}
	for k := range inputParams {
		leng := len(inputParams[k].IpArr)
		fmt.Printf("IP: %s, MIP: %s, SN: %s, Location: %s\n",
			inputParams[k].IpArr[0].Ip,
			inputParams[k].IpArr[leng-1].Ip,
			inputParams[k].SN,
			inputParams[k].CabinetID)
	}

	//
	input := []string{"02/08/2015", "143123412", "2018-08-08"}
	local, _ := time.LoadLocation("UTC")
	for k := range input {
		t, err := time.ParseInLocation("2006-01-02", input[k], local)
		if err != nil {
			fmt.Printf("turn time error: %s\n", err.Error())
		} else {
			fmt.Printf("turn result %v", t)
		}
	}

	fmt.Printf("len(\"do something\"): %d\n", len("do something"))
}
