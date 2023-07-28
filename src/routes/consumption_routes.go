package routes

import (
	"github.com/cantillo16/bia_energy/src/handlers"
	"github.com/cantillo16/bia_energy/src/repositories"
	"github.com/cantillo16/bia_energy/src/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {

	repo := repositories.NewConsumptionRepository(db)
	service := services.NewConsumptionService(repo)
	handler := handlers.NewHandler(service)

	router.HandleFunc("/consumption", handler.GetConsumption).Methods("GET")
	router.HandleFunc("/consumption/latest", handler.GetLast).Methods("GET")

}
