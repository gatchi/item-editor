package main

import (
	"flag"
	"github.com/dcrodman/archon/util"
	"os"
	"strings"
)

var (
	ip = flag.Int("num", 1, "the num to print")
)

func main() {
	flag.Parse()

	// Open bin file for reading
	file, err := os.Open("ItemPMT.bin")
	if err != nil {
		println("Can't find ItemPMT.bin.")
		os.Exit(1)
	}

	// Maybe this should be a daemon that gets
	// requests from the PHP website?
	// I guess either way we're running commands.
	// Is there a reason for the program to continue
	// running when not getting requests?  Not particularly.
	// 
	// Itd be better probably to make a command-line utility
	// that takes in flags and responds based on flag.
	// Man why cant this shit be in MySQL....
	if *ip == 2 {
		println("ahh yes perfect")
	} else {
		println("num is", *ip)
	}

	// Ok, lets convert the bin contents to bytes
	r := strings.NewReader()
	var buf [87000]byte
	io.ReadFull(r, buf)
	var table ItemPMT
	util.StructFromBytes(buf, table)

	file.Close()
}
