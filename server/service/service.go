package service

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Service struct {
	storage map[string][]byte
}

func NewService() *Service {
	return &Service{
		storage: make(map[string][]byte),
	}
}

func (srv *Service) delete(response http.ResponseWriter, request *http.Request) {
	key := strings.TrimPrefix(request.URL.Path, "/")
	srv.storage[key] = nil
	response.WriteHeader(http.StatusOK)
}

func (srv *Service) get(response http.ResponseWriter, request *http.Request) {
	key := strings.TrimPrefix(request.URL.Path, "/")
	value := srv.storage[key]
	response.WriteHeader(http.StatusOK)
	response.Write(value)
}

func (srv *Service) put(response http.ResponseWriter, request *http.Request) {
	key := strings.TrimPrefix(request.URL.Path, "/")
	value, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	srv.storage[key] = value
	response.WriteHeader(http.StatusOK)
}

func (srv *Service) Handle(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodDelete:
		srv.delete(response, request)
		break
	case http.MethodGet:
		srv.get(response, request)
		break
	case http.MethodPut:
		srv.put(response, request)
		break
	default:
		response.Header().Set("Allow", fmt.Sprint(http.MethodGet, ",", http.MethodPut))
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		break
	}
}
