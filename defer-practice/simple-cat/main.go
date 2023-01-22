package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
	Reading a file in chunks
	https://kgrz.io/reading-files-in-go-an-overview.html
*/

func main() {

	//Using os.Args to check command line args
	if len(os.Args) < 2 {
		log.Fatal("no file spcified")
	}
	//Getting the specified file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	//Cleaning up
	defer file.Close()
	//Buffer on a memory (1 bytes)
	//Reading a file in chunks of some size
	const bufferSize = 3
	buffer := make([]byte, bufferSize)
	i := 0
	for {
		i++
		fmt.Printf("Read: %v time(s)\n", i)
		//Reading a file into buffer
		count, err := file.Read(buffer)
		fmt.Printf("count: %v bytes\n", count)
		fmt.Printf("buffer: %v\n", buffer)
		os.Stdout.Write(buffer[:count])
		fmt.Println()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			} else {
				fmt.Println("EOF")
			}
			break
		}
	}
}
