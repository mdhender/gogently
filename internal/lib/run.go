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

type FILE struct{}

var yyin io.ReadCloser

func Run(argc int, argv ...string) error {
	var err error

	/* INITIALIZE */
	if argc > 2 {
		return fmt.Errorf("too many arguments\n")
	}
	if argc == 2 {
		yyin, err = os.Open(argv[1])
		if err != nil {
			return err
		}
	}

	/* INVOKE GENERATED PROGRAM */
	if err = ROOT(); err != nil {
		return err
	}

	/* FINALIZE */

	return nil
}

func ROOT() error {
	return fmt.Errorf("ROOT not implemented")
}
