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

package lib

import (
	"fmt"
	"io"
	"os"
)

func CloseOutput() {
	if outFile != nil {
		outFile.Close()
		outFile = nil
	}
}

func Nl() {
	if EMIT_CR {
		outFile.Write([]byte{'\r', '\n'})
		return
	}
	outFile.Write([]byte{'\n'})
}

func OpenOutput(name string) int {
	var err error
	CloseOutput()
	outFile, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("lib: %q: %+v\n", err)
		os.Exit(1)
	}
	return 1
}

func Put(str string) {
	_, _ = outFile.Write([]byte(str))
}

func PutI(i int) {
	_, _ = outFile.Write([]byte(fmt.Sprintf("%d", i)))
}

const EMIT_CR = false

var outFile io.WriteCloser
