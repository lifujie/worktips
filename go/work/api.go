package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// APIMeta API元数据信息
type APIMeta struct {
	Method string
	URIReg *regexp.Regexp
}

// AuthorizationAPI 鉴权API
type AuthorizationAPI struct {
	API   *APIMeta
	Codes []string
}

// AuthorizationAPIModel 授权API配置
type AuthorizationAPIModel struct {
	API   *APIMetaModel `json:"api"`
	Codes []string      `json:"codes"`
}

// APIMetaModel API元信息
type APIMetaModel struct {
	Method string `json:"method"`
	URIReg string `json:"uri_regexp"`
	Desc   string `json:"desc"`
}

// AuthorizationAPIs 待鉴权的API元信息集合
type AuthorizationAPIs []*AuthorizationAPI

// 待鉴权API集合
var authorizationAPIs AuthorizationAPIs

// InitAuthorizationAPIs 初始化鉴权API
func InitAuthorizationAPIs(src string) error {
	var items []*AuthorizationAPIModel
	err := json.Unmarshal([]byte(src), &items)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	authAPIs := make([]*AuthorizationAPI, 0, len(items))
	for i := range items {
		fmt.Printf("%s\n", items[i].API.URIReg)
		reg, err := regexp.CompilePOSIX(items[i].API.URIReg)
		if err != nil {
			fmt.Printf("invalid regex(%s) expression in 'authorization' system setting", items[i].API.URIReg)
			return nil
		}
		authAPIs = append(authAPIs, &AuthorizationAPI{
			API: &APIMeta{
				Method: items[i].API.Method,
				URIReg: reg,
			},
			Codes: items[i].Codes,
		})
	}
	authorizationAPIs = AuthorizationAPIs(authAPIs)
	return nil
}

func main() {
	InitAuthorizationAPIs(items)
}

const (
	items = `[
		{\"api\":{\"desc\":\"新增数据中心\",\"method\":\"POST\",\"uri_regexp\":\"^/api/cloudboot/v1/idcs$\"},\"codes\":[\"button_idc_create\"]}`
)
