package handlers

import (
	artistdto "dumbsound/dto/artists"
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type handleArtist struct {
	ArtistRepository repositories.ArtistRepository
}

func HandleArtist(ArtistRepository repositories.ArtistRepository) *handleArtist {
	return &handleArtist{ArtistRepository}
}

func (h *handleArtist) ShowArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	artist, err := h.ArtistRepository.ShowArtists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode((err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: artist}
	json.NewEncoder(w).Encode(response)

}

func (h *handleArtist) GetArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	artist, err := h.ArtistRepository.GetArtist(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: artist}
	json.NewEncoder(w).Encode(response)
}

func (h *handleArtist) AddArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// rName, _ := r.FormValue("name")
	rOld, _ := strconv.Atoi(r.FormValue("old"))
	// rType, _ := r.FormValue("type")
	// rStartCarer, _ := r.FormValue("startcarer")
	request := artistdto.ArtistRequest{
		Name:       r.FormValue("name"),
		Old:        rOld,
		Type:       r.FormValue("type"),
		StartCarer: r.FormValue("startcarer"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	artist := models.Artist{
		Name:       request.Name,
		Old:        request.Old,
		Type:       request.Type,
		StartCarer: request.StartCarer,
	}

	dataArtist, err := h.ArtistRepository.AddArtist(artist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	getArtist, err := h.ArtistRepository.GetArtist(dataArtist.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArtist(getArtist)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseArtist(u models.Artist) artistdto.ArtistResponse {
	return artistdto.ArtistResponse{
		ID:         u.ID,
		Name:       u.Name,
		Old:        u.Old,
		Type:       u.Type,
		StartCarer: u.StartCarer,
	}
}
