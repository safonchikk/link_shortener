package application

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"hash/fnv"
	"link_shortener/internal/model"
	"net/http"
	"os"
	"strconv"
)

type Link struct {
	Repo *RedisRepo
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

	err = l.Repo.Insert(r.Context(), link)
	if err != nil {
		fmt.Println("failed to insert:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	link.Short = os.Getenv("ADDRESS") + "/" + link.Short
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
	o, err := l.Repo.FindByShort(r.Context(), short)
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
