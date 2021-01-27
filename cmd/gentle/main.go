/*
 * gogently - a universal toolkit for interpreters
 * Copyright (C) 2021  Michael D Henderson
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"github.com/mdhender/gogently/internal/lib"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing file name\n")
		os.Exit(1)
	}
	scanargs(os.Args[1:]...)
	init_scanner()
	init_idtab()
	lib.ROOT()
}

var TraceFlag = false

func TraceOption() bool {
	return TraceFlag
}

var SymbolFileFlag = false

func SymbolFileOption() bool {
	return SymbolFileFlag
}

func scanargs(args ...string) {
	source_defined := false
	for _, arg := range args {
		if arg == "-subdir" {
			SetOption_SUBDIR()
		} else if arg == "-alert" {
			SetOption_ALERT()
		} else if arg == "-if" {
			SymbolFileFlag = true
		} else if arg == "-trace" {
			TraceFlag = true
		} else {
			if length := len(arg); length > 0 && arg[0] == '-' {
				fmt.Printf("invalid option: %q\n", arg)
				os.Exit(1)
			} else if length <= 2 || arg[length-2] != '.' || arg[length-1] != 'g' {
				fmt.Printf("invalid filename: %q\n", arg)
				os.Exit(1)
			}
			DefSourceName(arg)
			source_defined = true
		}
	}
	if !source_defined {
		fmt.Printf("missing file name\n")
		os.Exit(1)
	}
}
