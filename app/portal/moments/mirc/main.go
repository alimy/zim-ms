package main

import (
	"log"

	. "github.com/alimy/mir/v2/core"
	. "github.com/alimy/mir/v2/engine"

	_ "gitbus.com/exlab/zim-ms/app/portal/moments/mirc/routes"
)

//go:generate go run main.go
func main() {
	log.Println("generate code start")
	opts := Options{
		RunMode(InSerialDebugMode),
		GeneratorName(GeneratorGin),
		SinkPath("./gen"),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
