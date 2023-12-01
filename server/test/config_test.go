package test

import (
	"blog-server/config"
	"blog-server/router"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	router.InitGlobalVariable()
	fmt.Println(config.Conf.Server.BackPort)
}
