package exception

type InvalidArgumentError struct {
	msg string
}

func (i InvalidArgumentError) Error() string {
	return i.msg
}

func (i InvalidArgumentError) Msg() string {
	return i.msg
}

func CreateInvalidArgumentError(description string) InvalidArgumentError {
	return InvalidArgumentError{
		msg: description,
	}
}
