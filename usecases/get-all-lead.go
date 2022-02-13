package usecases

import (
	"context"

	"github.com/chalfel/go-lead-store/models"
	"github.com/chalfel/go-lead-store/repositories"
)

type GetAllLeads interface {
	Handle(ctx context.Context) ([]models.Lead, error)
}

type GetAllLeadsUseCase struct {
	Repo repositories.LeadRepository
}

func (g *GetAllLeadsUseCase) Handle(ctx context.Context) ([]models.Lead, error) {
	return g.Repo.GetLeads()
}
