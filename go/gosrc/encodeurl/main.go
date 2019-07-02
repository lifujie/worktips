package main

import (
	"fmt"
	"net/url"
)

func main() {
	oldpassword := url.QueryEscape("7JhjFAA80QW9+uXGxkRMmswq6BCyF+YoUQ==")
	newpassword := url.QueryEscape("7JhjFABmhkzwThULoRCnnPoOuGEtZ4ReK3T2038=")
	fmt.Printf("%s, %s\n", oldpassword, newpassword)
}
