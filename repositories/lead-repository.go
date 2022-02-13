package repositories

import (
	"context"

	"github.com/chalfel/go-lead-store/db"
	"github.com/chalfel/go-lead-store/models"
	"go.mongodb.org/mongo-driver/bson"
)

type LeadRepository interface {
	GetLeads() ([]models.Lead, error)
	InsertLead(lead models.Lead) error
}

type MongoLeadRepository struct {
	Db *db.DatabaseConnection
}

func (ml *MongoLeadRepository) InsertLead(lead models.Lead) error {
	_, err := ml.Db.Conn.Database("lead_store").Collection("leads").InsertOne(context.Background(), interface{}(lead))
	if err != nil {
		return err
	}
	return nil
}

func (ml *MongoLeadRepository) GetLeads() ([]models.Lead, error) {
	response, err := ml.Db.Conn.Database("lead_store").Collection("leads").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var decodedLeads []models.Lead

	if err = response.All(context.Background(), &decodedLeads); err != nil {
		return nil, err
	}

	return decodedLeads, nil
}
