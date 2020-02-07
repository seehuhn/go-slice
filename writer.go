package slice

// Writer allows to modify a byte slice using the io.WriteSeeker interface.
type Writer struct {
	// Buf is the byte slice modififed by the Writer.
	Buf []byte

	// Pos is the current writing position.  This must always be in the
	// range 0, ..., len(Buf).
	Pos int
}

// NewWriter creates a new Writer which modifies the given byte slice.
func NewWriter(buf []byte) *Writer {
	return &Writer{
		Buf: buf,
	}
}

// Clear sets all bytes in the slice to 0 and moves the write position to the
// start of the slice.
func (w *Writer) Clear() {
	for i := 0; i < len(w.Buf); i++ {
		w.Buf[i] = 0
	}
	w.Pos = 0
}

// Write writes len(p) bytes from p to the underlying slice.  If the write
// would go beyond the end of the underlying slice, only part of the data is
// written, and ErrFull is returned.
func (w *Writer) Write(p []byte) (n int, err error) {
	n = len(w.Buf) - w.Pos
	if n > len(p) {
		n = len(p)
	}
	copy(w.Buf[w.Pos:w.Pos+n], p[:n])
	w.Pos += n
	if len(p) > n {
		err = ErrFull
	}
	return
}

// Seek sets the offset for the next Write to offset, interpreted according to
// whence: io.SeekStart means relative to the start of the file, io.SeekCurrent
// means relative to the current offset, and io.SeekEnd means relative to the
// end.  Seek returns the new offset relative to the start of the file.
//
// If the new position would be outside the range 0, ..., len(Buf), the current
// position is not modified and ErrOutOfBounds is returned.
func (w *Writer) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case 0: // io.SeekStart
		if offset < 0 || offset > int64(len(w.Buf)) {
			return int64(w.Pos), ErrOutOfBounds
		}
		w.Pos = int(offset)
	case 1: // io.SeekCurrent
		if offset < int64(-w.Pos) || offset > int64(len(w.Buf)-w.Pos) {
			return int64(w.Pos), ErrOutOfBounds
		}
		w.Pos += int(offset)
	case 2: // io.SeekEnd
		if offset < int64(-len(w.Buf)) || offset > 0 {
			return int64(w.Pos), ErrOutOfBounds
		}
		w.Pos = len(w.Buf) + int(offset)
	}
	return int64(w.Pos), nil
}
