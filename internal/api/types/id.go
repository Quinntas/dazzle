package types

type ID struct {
	Value int
}

func NewID(id int) ID {
	return ID{
		Value: id,
	}
}
