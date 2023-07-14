package errors

// Public wraps the original error with a new error that has a
// `Public() string` method. Returns a "public-facing" error
// to the end user. Can also be unwrapped using traditional
// `errors` package approach.
func Public(err error, msg string) error {
	return publicError{err, msg}
}

type publicError struct {
	err error
	msg string
}

func (pe publicError) Error() string {
	return pe.err.Error()
}

func (pe publicError) Public() string {
	return pe.msg
}

func (pe publicError) Unwrap() error {
	return pe.err
}
