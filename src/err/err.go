package err

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
