package main

import (
	"flag"
	services "hajime/services"
	"log"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	h := services.RequestHandler
	if *compress {
		color.Cyan("with compress")
		h = fasthttp.CompressHandler(h)
	} else {
		color.Cyan("without compress")
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
		color.Red("Error in ListenAndServe: %s", err)
	}
}
