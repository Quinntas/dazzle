package sharedModel

type SharedModel struct {
	ID        int    `json:"id" db:"id"`
	PID       string `json:"pid" db:"pid"`
	CreatedAt string `json:"createdAt" db:"createdAt"`
	UpdatedAt string `json:"updatedAt" db:"updatedAt"`
}

func NewSharedModel(
	id int,
	pid string,
	createdAt string,
	updatedAt string,
) SharedModel {
	return SharedModel{
		ID:        id,
		PID:       pid,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
