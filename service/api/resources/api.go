package resources

import (
	"context"
	"log"

	"github.com/chimano/water-those-service/api/resources/plants"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// API provides application resources and handlers.
type API struct {
	Plant *plants.PlantResource
}

// NewAPI configures and returns application API.
func NewAPI() *API {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://service-db:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	plantCollection := client.Database("water-those-db").Collection("plants")

	plant := plants.NewPlantResource(plantCollection)
	api := &API{
		Plant: plant,
	}
	return api
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/plant", a.Plant.Router())

	return r
}
