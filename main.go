package main

import (
	"fmt"
	"log"
	"os"

	"rsc.io/quote"
)

func readfile(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err.Error() + " will create a file")
		_, err := os.Create(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

func writefile(f string, text string) {
	if text == "" {
		fmt.Println("need to set string to be written")
		os.Exit(1)
	}

	err := os.Chmod(f, 0655)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	count, err := file.WriteString(text + "\n")
	if err != nil {
		fmt.Println(err.Error() + " will create a file")
		_, err := os.Create(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	file.Close()
	fmt.Printf("write %d bytes: %q\n", count, text)
}

func main() {
	file := "file.txt"

	if len(os.Args) > 1 {
		arg := os.Args[1]

		if arg == "wisdom" {
			fmt.Println(quote.Go())
		}

		if arg == "readfile" {
			readfile(file)
		}

		if arg == "writefile" {
			writefile(file, os.Args[2])
		}
	} else {
		fmt.Println("No command available")
	}
}
