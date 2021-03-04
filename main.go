package main

import (
	"github.com/gogo/protobuf/vanity/command"
	"github.com/suutaku/protoc-gen-gorm/plugin"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	op := &plugin.OrmPlugin{}
	response := command.GeneratePlugin(command.Read(), op, ".pb.gorm.go")
	op.CleanFiles(response)
	command.Write(response)

}
