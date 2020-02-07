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
