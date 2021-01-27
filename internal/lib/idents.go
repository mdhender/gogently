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

import "fmt"

type IDENTSTRUCT struct {
	str     string
	meaning int
}

var idtab map[string]*IDENTSTRUCT

func allocate_idtab() {
	idtab = make(map[string]*IDENTSTRUCT)
}

var initialized bool

func InitIdents() {
	allocate_idtab()
	initialized = true
}

func DefMeaning(id *IDENTSTRUCT, m int) {
	id.meaning = m
}

func ErrorI(str1 string, id *IDENTSTRUCT, str2 string, pos int) {
	Error(fmt.Sprintf("%s%s%s", str1, id.str, str2), pos)
}

func HasMeaning(id *IDENTSTRUCT) int {
	if id.meaning == 0 {
		return 0
	}
	return id.meaning
}

func UndefMeaning(id *IDENTSTRUCT) {
	id.meaning = 0
}

func id_to_string(id *IDENTSTRUCT) string {
	return id.str
}
func (id *IDENTSTRUCT) String() string {
	return id.str
}

func slice_to_id(b []byte) *IDENTSTRUCT {
	if !initialized {
		InitIdents()
	}

	str := string(b)
	if id, ok := idtab[str]; ok {
		return id
	}
	id := &IDENTSTRUCT{
		str:     str,
		meaning: 0,
	}
	idtab[id.str] = id
	return id
}

func string_to_id(str string) *IDENTSTRUCT {
	if id, ok := idtab[str]; ok {
		return id
	}
	return nil
}
