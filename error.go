// seehuhn.de/go/slice - an io.WriteSeeker interface for byte slices
// Copyright (C) 2020  Jochen Voss <voss@seehuhn.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package slice

// Error is the type used for error constants in the slice package.
type Error string

func (err Error) Error() string {
	return string(err)
}

// Error constants used by the slice package.
const (
	ErrFull        = Error("write beyond end of slice attempted")
	ErrOutOfBounds = Error("position out of bounds")
)
