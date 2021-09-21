package model

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(tpe string, length int, value string) Result {
	return Result{tpe, length, value}
}
