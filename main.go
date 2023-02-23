package main

import (
	"fmt"
	"kelindan/pkg/settings"
)

func init() {
	settings.Setup()
}

func main() {

	fmt.Println(settings.AppConfigSetting.Server.RunMode)
}
