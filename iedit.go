package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/gatchi/item-editor/itempmt"
	"os"
)

const (
	itemPmtFile = "bin/ItemPMT.bin"
)

var (
	ip = flag.Int("num", 1, "the num to print")
)

func main() {
	flag.Parse()

	// Maybe this should be a daemon that gets
	// requests from the PHP website?
	// I guess either way we're running commands.
	// Is there a reason for the program to continue
	// running when not getting requests?  Not particularly.
	// 
	// Itd be better probably to make a command-line utility
	// that takes in flags and responds based on flag.
	// Man why cant this shit be in MySQL....

	// Open ItemPMT.
	file, err := os.Open(itemPmtFile)
	if err != nil {
		fmt.Println("Can't find ItemPMT.bin.")
		os.Exit(1)
	}

	// Seek to start offset
	offset, err := file.Seek(itempmt.START, 2)
	if err != nil {
		fmt.Println("Can't seek.")
	} else {
		fmt.Printf("%X\n", offset)
	}
	
	// Read the address
	var startAddr int32
	binary.Read(file, binary.LittleEndian, &startAddr)
	fmt.Printf("Addr: %X\n", startAddr)
}
