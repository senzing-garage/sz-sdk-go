package g2error

type UnrecoverableInputError struct {
	error
	IntA int
	IntB int
	Msg  string
}

func (err *UnrecoverableInputError) Error() string {
	return err.Error()
}
