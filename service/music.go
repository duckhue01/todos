package service

import (
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func RunMP3(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	d, err := mp3.NewDecoder(f)
	if err != nil {
		log.Fatal(err)
	}
	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()
	p := c.NewPlayer()
	defer p.Close()
	if _, err := io.Copy(p, d); err != nil {
		log.Fatal(err)
	}
}
