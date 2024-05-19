package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/OscarG356/web_server_GO/Repository"
)

type HeroHandler struct {
	BD *repository.BaseData
}

func NewHandlerHero(bd *repository.BaseData) *HeroHandler {
	return &HeroHandler{
		BD: bd,
	}
}

func (ph *HeroHandler) GetHero() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extraer los parámetros de la URL
		
		heroName := r.PathValue("superhero")


		if heroName == "" {
			http.Error(w, "Falta el nombre", http.StatusBadRequest)
		}

		for i := 1; i < 6; i++ {

			hero := ph.BD.Storage[i]

			if hero.Name == heroName {
				payload, err := json.Marshal(hero)
				if err != nil {
					http.Error(w, "Fallo la codificacion del JSON", http.StatusInternalServerError)
					return
				}else{
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(payload)
					return
				}
			}
		}
		http.Error(w, "Heroe no encontrado", http.StatusNotFound)
	}
}
