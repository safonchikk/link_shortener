package application

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"hash/fnv"
	"link_shortener/internal/model"
	"link_shortener/util"
	"log"
	"net/http"
	"strconv"
)

type Repo interface {
	Insert(ctx context.Context, link model.Link) error
	FindByShort(ctx context.Context, short string) (model.Link, error)
}

type Link struct {
	Repository Repo
}

func (l *Link) MakeShort(w http.ResponseWriter, r *http.Request) {
	fmt.Println("make short called")
	var body struct {
		Long string `json:"long"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h := fnv.New32a()
	_, err := h.Write([]byte(body.Long))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	short := strconv.FormatInt(int64(h.Sum32()), 36)

	link := model.Link{Short: short, Long: body.Long}

	err = l.Repository.Insert(r.Context(), link)
	if err != nil {
		fmt.Println("failed to insert:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Error loading app.env file" + err.Error())
	}

	link.Short = config.Address + ":" + config.NginxPort + "/" + link.Short
	res, err := json.Marshal(link)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
	fmt.Println(short)
}

func (l *Link) MakeLong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("make long called")
	short := chi.URLParam(r, "id")
	o, err := l.Repository.FindByShort(r.Context(), short)
	if err != nil {
		fmt.Println("failed to find by short:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(o)
	if err := json.NewEncoder(w).Encode(o); err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
