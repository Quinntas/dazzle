package sharedDomain

import (
	"github.com/quinntas/go-rest-template/internal/api/types"
)

type SharedDomain struct {
	ID        types.ID   `json:"id" `
	PID       types.UUID `json:"pid" `
	CreatedAt types.Date `json:"createdAt" `
	UpdatedAt types.Date `json:"updatedAt" `
}

func NewSharedDomain(
	id types.ID,
	pid types.UUID,
	createdAt types.Date,
	updatedAt types.Date,
) SharedDomain {
	return SharedDomain{
		ID:        id,
		PID:       pid,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
