package g2error

type RetryableError struct {
	error
	IntA int
	IntB int
	Msg  string
}

func (err *RetryableError) Error() string {
	return err.Error()
}
