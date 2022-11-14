package err

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "not found"
}

type NothingAvailableError struct{}

func (e NothingAvailableError) Error() string {
	return "nothing available"
}
