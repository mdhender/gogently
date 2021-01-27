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
	"io"
	"os"
)

func SetOption_SUBDIR() {
	SUBDIR = true
}

func Tell(Name string) {
	Told()
	var err error
	if OutFile, err = os.OpenFile(Name, os.O_WRONLY|os.O_CREATE, 0664); err != nil {
		fmt.Printf("cannot open %q: %+v\n", Name)
		os.Exit(1)
	}
}

func TellFile(Name string) {
	if SUBDIR {
		Tell(fmt.Sprintf("_G_/%s", Name))
		return
	}
	Tell(fmt.Sprintf("%s", Name))
}

func TellClauseFile() {
	TellFile(fmt.Sprintf("%s.c", SourceName()))
}

func TellSymbolFile() {
	TellFile(fmt.Sprintf("%s.if", SourceName()))
}

func TellXRefFile() {
	TellFile(fmt.Sprintf("%s.nst", SourceName()))
}

func Told() {
	if OutFile != nil {
		_ = OutFile.Close()
		OutFile = nil
	}
}

const EMIT_CR = false

var OutFile io.WriteCloser
var SUBDIR = false

func doublequote() {
	_, _ = fmt.Fprint(OutFile, "\"")
}

func i(N int) {
	_, _ = fmt.Fprintf(OutFile, "%d", N)
}

func nl() {
	if EMIT_CR {
		_, _ = fmt.Fprintln(OutFile, "\r")
		return
	}
	_, _ = fmt.Fprintln(OutFile, "")
}

func qu_s(Str string) {
	_, _ = fmt.Fprintf(OutFile, "%q", Str)
}

func s(Str string) {
	_, _ = fmt.Fprint(OutFile, s)
}
