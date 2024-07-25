package main

import (
	"flag"
	"log"
	"os"

	"github.com/Richtermnd/game/cmd/mapconvertor/convertor"
	"github.com/Richtermnd/game/internal/config"
)

var (
	filename string
)

func init() {
	flag.StringVar(&filename, "filename", "maps/map.txt", "path to file with map")
	flag.Parse()
}

func main() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	content, err := convertor.Read(f, config.FieldWidth, config.FieldHeight)
	if err != nil {
		log.Fatal(err)
	}

	convertedData, err := convertor.Convert(content)
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.Create(convertor.OutputFilename(filename))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	n, err := out.Write(convertedData)
	if err != nil || n != len(convertedData) {
		log.Fatal(err)
	}
}
