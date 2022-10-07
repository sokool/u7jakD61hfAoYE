package dictionary

import (
	"encoding/json"
	"net/http"
)

type HTTP struct{ dictionary *Storage }

func NewHTTP(s *Storage) *HTTP {
	return &HTTP{s}
}

func (h *HTTP) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		h.Add(res, req)
	case http.MethodGet:
		h.Get(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *HTTP) Add(res http.ResponseWriter, req *http.Request) {
	var w Word
	if err := json.NewDecoder(req.Body).Decode(&w); err != nil {
		h.error(res, err)
		return
	}
	h.dictionary.Put(w)
}

func (h *HTTP) Get(res http.ResponseWriter, req *http.Request) {
	w, err := NewWord(req.URL.Query().Get("prefix"))
	if err != nil {
		h.error(res, err)
		return
	}

	if err = json.NewEncoder(res).Encode(h.dictionary.Search(w).Recent()); err != nil {
		h.error(res, err)
	}
}

func (h *HTTP) error(res http.ResponseWriter, err error) {
	http.Error(res, err.Error(), http.StatusBadRequest)
}
