package err

const CODE_NOTHING_AVAILABLE = 10
const CODE_CANNOT_RECEIVE_MORE_MESSAGES = 20
const CODE_CANNOT_WRITE_MORE_MESSAGES = 30

type GenericServerError struct{}

func (e GenericServerError) Error() string {
	return "something went wrong"
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "not found"
}

type NothingAvailableError struct{}

func (e NothingAvailableError) Error() string {
	return "nothing available"
}

type CannotReceiveMoreMessagesError struct{}

func (e CannotReceiveMoreMessagesError) Error() string {
	return "cannot receive more messages"
}

type CannotWriteMoreMessagesError struct{}

func (e CannotWriteMoreMessagesError) Error() string {
	return "cannot write more messages"
}
