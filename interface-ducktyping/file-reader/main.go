package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

// It is not necessary to modify this function because it take the Reader interface argument.
// This is the advanatage of interface argument.
func process(r io.Reader) error {
	buff := make([]byte, 2048)
	for {
		count, err := r.Read(buff)
		os.Stdout.Write(buff[:count])
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}
	return nil
}

func main() {

	file := os.Args[1]
	s := strings.Split(file, ".")
	ext := s[len(s)-1]

	//#1 Reading from text file
	r, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	switch ext {
	case "txt":
		process(r)
	case "gz":
		//#2 Reading from zip file
		gr, err := gzip.NewReader(r)
		if err != nil {
			log.Fatal(err)
		}
		defer gr.Close()
		process(gr)
	}
}
