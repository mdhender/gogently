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
	"os"
)

var THIS_RUNTIME_SYSTEM = "Gentle 3.0 01100401 (C) 1992, 1997"

const yyCntlMax = 500

type yyt [yyCntlMax+200]int

var yyh []int
var yyhx yyt

const HEAPPIECE = 20000

func yyExtend() {
	yyh = make([]int, cap(yyh)+HEAPPIECE, cap(yyh)+HEAPPIECE)
}

var CURBLOCK *yyt
var CURPOS *yyt
var FIRSTBLOCK *yyt

var FREELIST []yyt

func NEWBLOCK() yyt {
	var p yyt
	switch len(FREELIST) {
	case 0:
		//p = new(yyt)
	case 1:
		p, FREELIST = FREELIST[0], nil
	default:
		p, FREELIST = FREELIST[len(FREELIST)-1], FREELIST[:len(FREELIST)-1]
	}
	return p
}

func FREEBLOCK(p yyt) {
	FREELIST = append(FREELIST, p)
}

func yyAllocCntl(n int) yyt {
	p := CURPOS
	CURPOS += n
	if CURPOS >= CURBLOCK+yyCntlMax {
		b := NEWBLOCK()
		*CURBLOCK = b
		CURBLOCK = b
		CURPOS = CURBLOCK + 1
		p = CURPOS
		CURPOS += n
	}
	return p
}

type yysave struct {
	firstblock yyt
	curblock   yyt
	curpos     yyt
}

func yyBeginChoice(ref_saved yysave) {
	ref_saved.curblock = CURBLOCK
	ref_saved.curpos = CURPOS
	ref_saved.firstblock = FIRSTBLOCK

	FIRSTBLOCK = NEWBLOCK()
	*FIRSTBLOCK = 0
	CURBLOCK = FIRSTBLOCK
	CURPOS = CURBLOCK + 1
}

func yyEndChoice(saved yysave) {
	p := FIRSTBLOCK
	for p != yyt(0) {
		next := yyt(p)
		FREEBLOCK(p)
		p = next
	}

	CURBLOCK = saved.curblock
	CURPOS = saved.curpos
	FIRSTBLOCK = saved.firstblock
}

func yyAbort(Code int, FileName string, Line int) {
	switch Code {
	case 1:
		fmt.Printf("Undefined value in %q, line %d\n", FileName+".g", Line)
	case 2:
		fmt.Printf("No rule applicable in %q, line %d\n", FileName+".g", Line)
	case 3:
		fmt.Printf("Selected grammar rule failed in %q, line %d\n", FileName+".g", Line)
	case 4:
		fmt.Printf("Selected CHOICE rule failed in %q, line %d\n", FileName+".g", Line)
	default:
		fmt.Printf("Error %d (?) in %q, line %d\n", Code, FileName+".g", Line)
	}
	os.Exit(1)
}

func yyPrintOpaque(i int) {
	fmt.Printf("<<%d>>", i)
}

func yyPrintIndex(i int) {
	fmt.Printf("#%d", i)
}

func yyPrint_INT(i int) {
	fmt.Printf("%d", i)
}

func yyPrint_POS(i int) {
	fmt.Printf("%d", i)
}

const STRINGLENGTH = 40

func yyPrint_STRING(Str string) {
	fmt.Printf("%q", Str)
}

var yyIndentation = 0

func yyIndent() {
	for i := 1; i <= yyIndentation; i++ {
		fmt.Printf("   ")
	}
}

func yyTerm(f string) {
	fmt.Print(f)
}

func yyFirstArg() {
	fmt.Printf("(\n")
	yyIndentation++
	yyIndent()
}

func yyNextArg() {
	fmt.Printf(",\n")
	yyIndent()
}

func yyEndArgs() {
	yyIndentation--
	fmt.Printf("\n")
	yyIndent()
	fmt.Printf(")")
}

func yyNoArgs() {
}

func yyEndPrint() {
	fmt.Printf("\n")
}
