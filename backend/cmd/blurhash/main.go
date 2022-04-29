package main

import (
	"flag"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/bbrks/go-blurhash"
)

func main() {
	flag.Parse()

	src := flag.Arg(0)
	if src == "" {
		log.Fatalln("usage: blurhash <path or url>")
	}

	var reader io.Reader

	if strings.HasPrefix(src, "https://") || strings.HasPrefix(src, "http://") {
		r, err := http.Get(src)
		if err != nil {
			log.Fatalln("cannot GET:", err)
		}
		defer r.Body.Close()
		reader = r.Body
	} else {
		f, err := os.Open(src)
		if err != nil {
			log.Fatalln("cannot open file:", err)
		}
		defer f.Close()
		reader = f
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalln("cannot decode image:", err)
	}

	hash, err := blurhash.Encode(5, 5, img)
	if err != nil {
		log.Fatalln("cannot hash:", err)
	}

	fmt.Println(hash)
}
