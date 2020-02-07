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

import (
	"io"
	"testing"
)

// compile time test: Writer implements io.WriteSeeker
var _ io.WriteSeeker = &Writer{}

func TestWrite(t *testing.T) {
	w := &Writer{
		Buf: make([]byte, 100),
		Pos: 1,
	}
	n, err := w.Write([]byte{1})
	if n != 1 || err != nil {
		t.Error("write failed", n, err)
	}
	n, err = w.Write([]byte{2, 3, 4})
	if n != 3 || err != nil {
		t.Error("write failed", n, err)
	}

	for i, val := range []byte{0, 1, 2, 3, 4, 0} {
		if w.Buf[i] != val {
			t.Errorf("wrong contents Buf[%d] == %d != %d",
				i, w.Buf[i], val)
		}
	}
}

func TestSeek(t *testing.T) {
	w := &Writer{
		Buf: make([]byte, 100),
	}

	cases := []struct {
		before int
		offset int64
		whence int
		after  int
		err    error
	}{
		{0, 10, io.SeekStart, 10, nil}, // test 1
		{10, 0, io.SeekStart, 0, nil},
		{10, -1, io.SeekStart, 10, ErrOutOfBounds},
		{10, 100, io.SeekStart, 100, nil},
		{10, 101, io.SeekStart, 10, ErrOutOfBounds},

		{40, 0, io.SeekCurrent, 40, nil}, // test 6
		{40, -1, io.SeekCurrent, 39, nil},
		{40, +1, io.SeekCurrent, 41, nil},
		{40, -40, io.SeekCurrent, 0, nil},
		{40, -41, io.SeekCurrent, 40, ErrOutOfBounds}, // test 10
		{40, 60, io.SeekCurrent, 100, nil},
		{40, 61, io.SeekCurrent, 40, ErrOutOfBounds},

		{20, -40, io.SeekEnd, 60, nil}, // test 13
		{20, 0, io.SeekEnd, 100, nil},
		{20, 1, io.SeekEnd, 20, ErrOutOfBounds},
		{20, -100, io.SeekEnd, 0, nil},
		{20, -101, io.SeekEnd, 20, ErrOutOfBounds},
	}

	for i, test := range cases {
		w.Pos = test.before
		pos, err := w.Seek(test.offset, test.whence)
		if w.Pos != test.after || err != test.err {
			t.Errorf("test %d: expected (%d, %v), got (%d, %v)",
				i+1, test.after, test.err, w.Pos, err)
		}
		if int64(w.Pos) != pos {
			t.Errorf("test %d: expected position %d, got %d",
				i+1, w.Pos, pos)
		}
	}
}
