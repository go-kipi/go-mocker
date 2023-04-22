package main

import (
	. "github.com/go-kipi/kipimanager"
	. "github.com/go-kipi/kipimanager/settings"
)

var serviceConf Configuration

func main() {
	InitalizeNewHTTPService(&serviceConf, "0.0.1", initRouters)
}
