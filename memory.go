// Copyright (c) 2017-2020 Claudio Guarnieri.
//
// This file is part of Snoopdigg.
//
// Snoopdigg is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Snoopdigg is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Snoopdigg.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"os"
	"path/filepath"
)

var binPath = filepath.Join(getCwd(), "bin")

func initBinFolder() error {
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		err = os.MkdirAll(binPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
