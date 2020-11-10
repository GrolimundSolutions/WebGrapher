package main

import (
	"flag"
	"fmt"
)

func main() {
	urlPtr := flag.String("url", "https:google.com", "Link to the first Website")
	deepPtr := flag.Int("deep", 2, "Number of layers in graph")
	outPtr := flag.String("out", "./", "Output file location")
	flag.Parse()
	fmt.Println("url:", *urlPtr)
	fmt.Println("deep:", *deepPtr)
	fmt.Println("out:", *outPtr)

}
