package rest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/timoth-y/kicksware-api/service-common/util"

	"go.kicksware.com/api/beta/core/meta"
	"go.kicksware.com/api/beta/core/model"
	"go.kicksware.com/api/beta/core/service"
	"go.kicksware.com/api/beta/env"
	"go.kicksware.com/api/beta/usecase/business"
	"go.kicksware.com/api/beta/usecase/serializer/json"
	"go.kicksware.com/api/beta/usecase/serializer/msg"
)

type Handler struct {
	service     service.BetaService
	auth        service.AuthService
	contentType string
}

func NewHandler(service service.BetaService, auth service.AuthService, config env.CommonConfig) *Handler {
	return &Handler{
		service,
		auth,
		config.ContentType,
	}
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var betas []*model.Beta
	var err error
	params := NewRequestParams(r)

	if r.Method == http.MethodPost {
		query, err := h.getRequestQuery(r); if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		betas, err = h.service.FetchQuery(query, params)
	} else if r.Method == http.MethodGet {
		codes := r.URL.Query()["betaID"]
		if codes != nil && len(codes) > 0 {
			betas, err = h.service.Fetch(codes, params)
		} else {
			betas, err = h.service.FetchAll(params)
		}
	}

	if err != nil {
		if errors.Cause(err) == business.ErrBetaNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, betas, http.StatusOK)
}

func (h *Handler) PostOne(w http.ResponseWriter, r *http.Request) {
	beta, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.StoreOne(beta)
	if err != nil {
		if errors.Cause(err) == business.ErrBetaNotValid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, beta, http.StatusOK)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	beta, err := h.getRequestBodies(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.service.Store(beta)
	if err != nil {
		if errors.Cause(err) == business.ErrBetaNotValid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, beta, http.StatusOK)
}


func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	beta, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.service.Modify(beta)
	if err != nil {
		if errors.Cause(err) == business.ErrBetaNotValid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) Count(w http.ResponseWriter, r *http.Request) {
	var count int
	var err error
	params := NewRequestParams(r)

	if r.Method == http.MethodPost {
		query, err := h.getRequestQuery(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		count, err = h.service.Count(query, params)
	} else if r.Method == http.MethodGet {
		query := r.URL.Query()
		if query != nil && len(query) > 0 {
			count, err = h.service.Count(util.ToQueryMap(query), params)
		} else {
			count, err = h.service.CountAll()
		}
	}

	if err != nil {
		if errors.Cause(err) == business.ErrBetaNotFound {
			h.setupResponse(w, 0, http.StatusOK)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, count, http.StatusOK)
}

func (h *Handler) setupResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", h.contentType)
	w.WriteHeader(statusCode)
	if body != nil {
		raw, err := h.serializer(h.contentType).Encode(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(raw); err != nil {
			log.Println(err)
		}
	}
}

func (h *Handler) getRequestQuery(r *http.Request) (meta.RequestQuery, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).DecodeMap(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *Handler) getRequestBody(r *http.Request) (*model.Beta, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}


func (h *Handler) getRequestBodies(r *http.Request) ([]*model.Beta, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	bodies, err := h.serializer(contentType).DecodeRange(requestBody)
	if err != nil {
		return nil, err
	}
	return bodies, nil
}

func (h *Handler) serializer(contentType string) service.BetaSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}