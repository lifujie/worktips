package main

import (
	"encoding/json"
	"log"
)

var (
	data = `[
		{
			"action": "set_bios",
			"category": "bios",
			"metadata": {
				"key": "CustomPowerPolicy",
				"value": "1",
				"custom": "YES"
			}
		},
		{
			"action": "set_bios",
			"category": "bios",
			"metadata": {
				"key": "MonitorMwaitEnable",
				"value": "1",
				"custom": "YES"
			}
		},
		{
			"action": "set_bios",
			"category": "bios",
			"metadata": {
				"key": "ProcessorCcxEnable",
				"value": "0",
				"custom": "YES"
			}
		},
		{
			"action": "reboot",
			"category": "reboot",
			"metadata": {}
		},
		{
			"action": "clear_settings",
			"category": "raid",
			"metadata": {
				"clear": "ON",
				"controller_index": "0"
			}
		},
		{
			"action": "create_array",
			"category": "raid",
			"metadata": {
				"level": "raid0",
				"drives": "32:0,32:1",
				"controller_index": "0"
			}
		},
		{
			"action": "set_global_hotspare",
			"category": "raid",
			"metadata": {
				"drives": "32:2",
				"controller_index": "0"
			}
		},
		{
			"action": "set_ip",
			"category": "oob",
			"metadata": {
				"ip": "<{manage_ip}>",
				"ip_src": "static",
				"gateway": "<{manage_gateway}>",
				"netmask": "<{manage_netmask}>"
			}
		},
		{
			"action": "reset_bmc",
			"category": "oob",
			"metadata": {
				"reset": "ON"
			}
		},
		{
			"action": "add_user",
			"category": "oob",
			"metadata": {
				"password": "calvin",
				"username": "root",
				"privilege_level": "4"
			}
		}
	]`
)

// OOBUserForHardwareData 带外用户信息
type OOBUserForHardwareData struct {
	Name           string `json:"privilege_level"`
	Password       string `json:"password"` //这个值采集不到
	PrivilegeLevel string `json:"privilege_level"`
}

type HardwareData struct {
	Action   string            `json:"action"`
	Category string            `json:"category"`
	Metadata map[string]string `json:"metadata"`
}

type HardwareDatas []HardwareData

// Tojson 序列化为JSON
func (oob HardwareDatas) Tojson() []byte {
	b, _ := json.Marshal(oob)
	return b
}

func toOOBHardwareData(oobs string) HardwareDatas {
	var oob HardwareDatas
	json.Unmarshal([]byte(oobs), &oob)
	return oob
}

func main() {
	oob := toOOBHardwareData(data)
	log.Printf("%v\n", data)
	for k := range oob {
		if oob[k].Action == "add_user" && oob[k].Category == "oob" {
			//log.Printf("%v\n", oob[k])
			oob[k].Metadata["password"] = "123414jk"
		}
	}
	//if strings.Compare(string(oob.Tojson()), data) == 0 {
	log.Printf("result: %s", string(oob.Tojson()))
	//}
}
