/* PSO ItemPMT Browser Editor
 * Copyright 2017 Christen Gottschlich
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/gatchi/item-editor/itempmt"
	"os"
)

const itemPmtFile = "bin/ItemPMT.bin"

var ip = flag.Int("num", 1, "the num to print")

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
