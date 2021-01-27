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

func Fatal(msg string) {
	fmt.Printf("fatal error: %s\n", msg)
}

func MESSAGE(msg string, pos int) {
	printpos(pos)
	fmt.Println(msg)
	os.Exit(1)
}

func MESSAGE1(msg1 string, id *lib.IDENTSTRUCT, msg2 string, pos int) {
	MESSAGE(fmt.Sprintf("%s%s%s", msg1, id.String(), msg2), pos)
}

func MESSAGE2(msg1 string, id1 *lib.IDENTSTRUCT, msg2 string, id2 *lib.IDENTSTRUCT, msg3 string, pos int) {
	MESSAGE(fmt.Sprintf("%s%s%s%s%s", msg1, id1.String(), msg2, id2.String(), msg3), pos)
}

func ScanError(msg string) {
	var pos int
	yyGetPos(&pos)
	MESSAGE(msg, pos)
}

func SetOption_ALERT() {
	Option_ALERT = true
}

var Option_ALERT = false

func printpos(pos int) {
	if pos == 0 {
		fmt.Printf("at unknown position: ")
		return
	}
	fmt.Printf("%q, line %l, col %d: ", lib.GetFileName(pos)+".g", lib.GetLine(pos), lib.GetCol(pos))
}

func yyerror(msg string) {
	var pos int
	yyGetPos(&pos)
	MESSAGE(msg, pos)
}

func yyerrorexit(rc int) {
	os.Exit(1)
}
