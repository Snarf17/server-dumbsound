package handlers

import (
	"context"
	dtomusic "dumbsound/dto/musics"
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

// var path_file = "http://localhost:8000/uploads/thumbnail/"
// var path_song = "http://localhost:8000/uploads/song/"

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
		thumbnail[i].Thumbnail = p.Thumbnail
		thumbnail[i].Attache = p.Attache
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
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: thumbnail}
	json.NewEncoder(w).Encode(response)
}

func (h *handleMusic) CreateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	dataMusicContext := r.Context().Value("dataMusic")
	filepath := dataContex.(string)
	musicFile := dataMusicContext.(string)

	ctx := context.Background()
	CLOUD_NAME := "dr8ts6upb"
	API_KEY := "724231161762166"
	API_SECRET := "eUWKbh-weqN_ErRcV5vdsYH81-s"

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resImg, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "fileimage"})
	if err != nil {
		fmt.Println("gagal Upload Image", err.Error())
	}
	resMusic, err := cld.Upload.Upload(ctx, musicFile, uploader.UploadParams{Folder: "filemusic"})
	if err != nil {
		fmt.Println("gagal Upload Music", err.Error())
	}
	// resMusic, err := cld.Upload.Upload(ctx, musicFile, uploader.UploadParams)
	// add this code
	artis_id, _ := strconv.Atoi(r.FormValue("artis_id"))
	request := dtomusic.MusicRequest{
		Title:     r.FormValue("title"),
		Year:      r.FormValue("year"),
		ArtistID:  artis_id,
		Thumbnail: resImg.SecureURL,
		Attache:   resMusic.SecureURL,
	}

	validation := validator.New()
	err = validation.Struct(request)
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
