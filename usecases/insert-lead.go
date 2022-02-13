package usecases

import (
	"context"

	"github.com/chalfel/go-lead-store/models"
	"github.com/chalfel/go-lead-store/repositories"
)

type InsertLead interface {
	Handle(ctx context.Context, lead models.Lead) error
}

type InsertLeadUseCase struct {
	Repo repositories.LeadRepository
}

func (g *InsertLeadUseCase) Handle(ctx context.Context, lead models.Lead) error {
	return g.Repo.InsertLead(lead)
}
