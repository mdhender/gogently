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
	"path/filepath"
)

func main() {
	progname, err := os.Executable()
	if err != nil {
		fmt.Printf("%s: %+v\n", os.Args[0], err)
	}
	progname = filepath.Base(progname)
	if err := lib.Run(len(os.Args), os.Args...); err != nil {
		fmt.Printf("%s: %+v\n", progname, err)
	}
	fmt.Printf("todo: implement %q\n", progname)
	os.Exit(2)
}
