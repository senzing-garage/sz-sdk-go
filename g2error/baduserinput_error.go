package g2error

type BadUserInputError struct {
	error
	IntA int
	IntB int
	Msg  string
}

func (err *BadUserInputError) Error() string {
	return err.Error()
}
