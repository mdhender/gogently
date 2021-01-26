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
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var OUTFILE io.Writer

type info struct {
	name        string
	replacement string
	used        bool
	buffer      []byte // only set if file has been loaded
}

var info_list []*info

func main() {
	var err error

	outputFile := "gen.l"

	// expect arguments to be a list of name=replacement pairs.
	// the replacement value must be a filename with an extension of ".t"
	for i, arg := range os.Args[1:] {
		fields := strings.SplitN(arg, "=", 2)
		if len(fields) != 2 {
			fatal("args: %d: missing '=' in argument %d\n", i)
		}
		opt, value := fields[0], fields[1]
		if opt == "--output" {
			if value == "" {
				fatal("main: --output flag requires filename\n")
			}
			outputFile = value
			continue
		}
		if filepath.Ext(value) == "" {
			fatal("main: %d: filename %q does not have extension\n", i+1, value)
		}
		fmt.Printf("main: %d: aliasing %q to %q\n", i+1, opt, value)
		buffer, err := ioutil.ReadFile(value)
		if err != nil {
			fatal("main: %d: %q: %+v\n", i+1, value, err)
		}
		info_list = append(info_list, &info{
			name:        opt,
			replacement: value,
			buffer:      buffer,
		})
	}

	switch outputFile {
	case "stderr":
		OUTFILE = os.Stdout
	case "stdout":
		OUTFILE = os.Stderr
	default:
		if OUTFILE, err = os.Create(outputFile); err != nil {
			fatal("main: %q: %+v\n", outputFile, err)
		}
	}

	/* ( 1) %{ */
	emit(leftpar...)
	/* ( 2) YYSTYPE block */
	copy_or_text("YYSTYPE.b", yystype...)
	/* ( 3) SETPOS block */
	copy_or_text("SETPOS.b", setpos...)
	/* ( 4) LITBLOCK block */
	copy_or_text("LITBLOCK.b", litblock...)
	/* ( 5) %} */
	emit(rightpar...)
	/* ( 6) LEXDEF block */
	copy_or_text("LEXDEF.b", lexdef...)
	/* ( 7) %% */
	emit(separator...)
	/* ( 8) gen.lit */
	F, err := OPEN("gen.lit")
	if err != nil {
		fatal("main: %q: %+v\n", "gen.lit", err)
	}
	copy_buffer(F)
	/* ( 9) <token>.t for each <token> in gen.tkn */
	filelist()
	/* (10) COMMENTS block */
	copy_or_text("COMMENTS.b", comments...)
	/* (11) LAYOUT block */
	copy_or_text("LAYOUT.b", layout...)
	/* (12) ILLEGAL block */
	copy_or_text("ILLEGAL.b", illegal...)
	/* (13) %% */
	emit(separator...)
	/* (14) LEXFUNC block */
	copy_or_text("LEXFUNC.b", lexfunc...)
	/* (15) YYWRAP block */
	copy_or_text("YYWRAP.b", yywrap...)

	for _, cur := range info_list {
		if !cur.used {
			fatal("main: %q: not used\n", cur.name)
		}
	}
}

func copy_buffer(INFILE []byte) {
	if _, err := OUTFILE.Write(INFILE); err != nil {
		fatal("copy_buffer: %+v\n", err)
	}
	if len(INFILE) != 0 && INFILE[len(INFILE)-1] != '\n' {
		if _, err := OUTFILE.Write([]byte{'\n'}); err != nil {
			fatal("copy_buffer: %+v\n", err)
		}
	}
}

func copy_or_text(filename string, lines ...string) {
	INFILE, err := OPEN(filename)
	if err != nil {
		emit(lines...)
	} else {
		copy_buffer(INFILE)
	}
}

func emit(lines ...string) {
	for _, line := range lines {
		if _, err := OUTFILE.Write([]byte(line)); err != nil {
			fatal("emit: %+v\n", err)
		}
	}
}

func fatal(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

func filelist() {
	LISTFILE, err := OPEN("gen.tkn")
	if err != nil {
		fatal("filelist: %q: %+v\n", "gen.tkn", err)
	}

	for i, line := range bytes.Split(LISTFILE, []byte{'\n'}) {
		line = bytes.Trim(line, " \t\r\n")
		if len(line) == 0 || line[0] == '#' {
			// empty line or comment, so skip
			continue
		}

		var end int
		for end = 0; end < len(line); end++ {
			ch := line[end]
			if !(('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z') || ('0' <= ch && ch <= '9') || ch == '_') {
				break
			}
		}
		var name string
		if end < len(line) {
			name = string(line[:end]) + ".t"
		} else {
			name = string(line) + ".t"
		}

		INFILE, err := OPEN(name)
		if err != nil {
			fatal("filelist: %q:%d: cannot open %q: %+v\n", "gen.tkn", i+1, name, err)
		}
		copy_buffer(INFILE)
	}
}

// Change OPEN to be ioutil.ReadFile
func OPEN(name string) ([]byte, error) {
	for _, cur := range info_list {
		if name == cur.name {
			if cur.used {
				return cur.buffer, nil
			}
			if b, err := ioutil.ReadFile(cur.replacement); err != nil {
				fatal("open: %q: %q: %+v\n", cur.name, cur.replacement)
			} else {
				cur.used, cur.buffer = true, b
			}
			return cur.buffer, nil
		}
	}
	return ioutil.ReadFile(name)
}
