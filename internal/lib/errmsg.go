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
	"os"
)

func Error(msg string, pos int) {
	fmt.Printf("line %d, col %d: %s\n", yyLineAtPos(pos), yyColAtPos(pos), msg)
	os.Exit(1)
}

func yyGetPos(ref_pos *int) {
	*ref_pos = yypos - 1
}

func yyPosToNextLine() {
	yyLineCount++
	yypos = yyFileCount*yyFCODE + yyLineCount*yyLCODE + 1
}

func yyPosToNextFile() {
	yyLineCount = 1
	yyFileCount++
	yypos = yyFileCount*yyFCODE + yyLineCount*yyLCODE + 1
}

func yyFileAtPos(pos int) int {
	return pos / yyFCODE
}

func yyLineAtPos(pos int) int {
	return (pos % yyFCODE) / yyLCODE
}

func yyColAtPos(pos int) int {
	return pos % yyLCODE
}

func yyerror(msg string) {
	var pos int
	yyGetPos(&pos)
	Error(msg, pos)
}

/*
func yylexerror (msg string) {
    var pos int
    yyGetPos(&pos);
    Error(msg, pos);
}
*/

/*--------------------------------------------------------------------*/
/* Source Positions                                                   */
/*--------------------------------------------------------------------*/

const yyLCODE = 1000
const yyFCODE = 1000000000

var yypos int = (yyFCODE + yyLCODE + 1)

var yyLineCount int = 1
var yyFileCount int = 1
