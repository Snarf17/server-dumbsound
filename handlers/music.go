package handlers

import (
	dtomusic "dumbsound/dto/musics"
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var path_file = "http://localhost:8000/uploads/thumbnail/"
var path_song = "http://localhost:8000/uploads/song/"

type handleMusic struct {
	MusicRepository repositories.MusicRepository
}

func HandleMusic(MusicRepository repositories.MusicRepository) *handleMusic {
	return &handleMusic{MusicRepository}
}

func (h *handleMusic) ShowMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	thumbnail, err := h.MusicRepository.ShowMusics()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode((err.Error()))
	}
	for i, p := range thumbnail {
		thumbnail[i].Thumbnail = path_file + p.Thumbnail
		thumbnail[i].Attache = path_song + p.Attache
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: thumbnail}
	json.NewEncoder(w).Encode(response)

}

func (h *handleMusic) GetMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	thumbnail, err := h.MusicRepository.GetMusic(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	thumbnail.Thumbnail = path_file + thumbnail.Thumbnail
	thumbnail.Attache = path_song + thumbnail.Attache
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: thumbnail}
	json.NewEncoder(w).Encode(response)
}

func (h *handleMusic) CreateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	dataMusicContext := r.Context().Value("dataMusic") // add this code
	filename := dataContex.(string)
	musicFile := dataMusicContext.(string)
	// add this code
	artis_id, _ := strconv.Atoi(r.FormValue("artis_id"))
	request := dtomusic.MusicRequest{
		Title:     r.FormValue("title"),
		Year:      r.FormValue("year"),
		ArtistID:  artis_id,
		Thumbnail: filename,
		Attache:   musicFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	music := models.Music{
		Title:     request.Title,
		Year:      request.Year,
		Thumbnail: request.Thumbnail,
		Attache:   request.Attache,
		Artist_id: request.ArtistID,
	}

	data, err := h.MusicRepository.CreateMusic(music)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data2, err := h.MusicRepository.GetMusic(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data2}
	json.NewEncoder(w).Encode(response)
}
