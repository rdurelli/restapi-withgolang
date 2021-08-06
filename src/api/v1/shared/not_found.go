package shared

type NotFound struct {
	Msg string
}

func (nF *NotFound) Error() string {
	return nF.Msg
}
