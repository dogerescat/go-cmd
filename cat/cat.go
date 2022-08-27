package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	buf := make([]byte, 4096)
	flag.Parse()
	args := flag.Args()
	var file_exist bool = false
	for _, v := range args {
		f, err := os.Open(v)
		if err != nil {
			fmt.Printf("cannot open file %s\n", v)
			os.Exit(1)
		}
		io.Copy(os.Stdout, f)
		file_exist = true
		f.Close()
	}
	if file_exist {
		return
	}
	_, err := os.Stdin.Read(buf)
	if err == io.EOF {
		fmt.Println("EOF")
		os.Exit(1)
	}
	os.Stdout.Write(buf)
}
