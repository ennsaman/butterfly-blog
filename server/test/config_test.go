package test

import (
	"blog-server/config"
	"blog-server/routes"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	routes.InitGlobalVariable()
	fmt.Println(config.Conf.Server.BackPort)
}
