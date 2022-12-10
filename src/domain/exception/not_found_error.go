package exception

type NotFoundError struct {
	msg string
}

func (n NotFoundError) Error() string {
	return n.msg
}

func (n NotFoundError) Msg() string {
	return n.msg
}

func CreateNotFoundError(description string) NotFoundError {
	return NotFoundError{
		msg: description,
	}
}
