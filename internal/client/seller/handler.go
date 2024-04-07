package seller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test-crm/internal/appresult"
	"test-crm/internal/handlers"
	"test-crm/pkg/logging"

	"github.com/gorilla/mux"
)

const (
	GetAllSellerurl = "/get-all-seller"
	AddSellerUrl = "/add-seller"
	GetByIdSellerUrl = "/get-by-id-seller/{id}"
	UpdateSellerUrl = "/update-seller"
	DeleteSellerUrl = "/delete-seller"
)

type handler struct{
	repository Repository
	logger     *logging.Logger
}

func NewHandler( repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(GetAllSellerurl, appresult.Middleware(h.getAllSeller)).Methods("GET")
	router.HandleFunc(AddSellerUrl, appresult.Middleware(h.AddSeller)).Methods("POST")
	router.HandleFunc(GetByIdSellerUrl, appresult.Middleware(h.GetByIdSeller)).Methods("GET")
	router.HandleFunc(UpdateSellerUrl, appresult.Middleware(h.UpdateSeller)).Methods("PUT")
	router.HandleFunc(DeleteSellerUrl, appresult.Middleware(h.DeleteSeller)).Methods("DELETE")
}

func (h *handler) getAllSeller(w http.ResponseWriter, r *http.Request) error {
	data, err := h.repository.GetAllSeller(context.TODO())
	if err != nil {
		fmt.Println("err on handler", err)
		return appresult.ErrInternalServer
	}

	successResult := appresult.Success
	successResult.Data = data
	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}
	return nil

 }

func (h *handler) AddSeller(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println(errBody)
		return appresult.ErrMissingParam
	}
	addSell := AddSeller{}
	errData := json.Unmarshal(body, &addSell)

	if errData != nil {
		fmt.Println(errData)
		return appresult.ErrInternalServer
	}

	data, errr1 := h.repository.AddSeller( context.TODO(), addSell)
	if errr1 != nil {
		fmt.Println(errr1)
	}

	successResult := appresult.Success
	successResult.Data = data

	w.Header().Add(appresult.HeaderContentTypeJson())
	err := json.NewEncoder(w).Encode(successResult)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (h handler) GetByIdSeller(w http.ResponseWriter, r *http.Request) error{
	id := mux.Vars(r)["id"]
	data, err := h.repository.GetByIdSeller(context.TODO(), id)
		if err != nil {
			fmt.Println("errr by id", err)
		}
		
	successResult := appresult.Success
	successResult.Data = data
	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) UpdateSeller(w http.ResponseWriter, r *http.Request) error  {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println(errBody)
		return appresult.ErrMissingParam
	}

	updateSell := UpdateSeller{}
	errData := json.Unmarshal(body, &updateSell)
	if errData != nil {
		fmt.Println(errData)
		return appresult.ErrInternalServer
	}
	data, err1 := h.repository.UpdateSeller(context.TODO(), updateSell)
	if err1 != nil {
		fmt.Println(err1)
	}

	successResult := appresult.Success
	successResult.Data = data

	w.Header().Add(appresult.HeaderContentTypeJson())
	err := json.NewEncoder(w).Encode(successResult)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ( h handler) DeleteSeller(w http.ResponseWriter, r *http.Request)error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("DeleteSellerHandler", errBody)
	}

	deleteSell := DeleteSeller{}
	errData := json.Unmarshal(body, &deleteSell)
	if errData != nil {
		fmt.Println(errData)
		return appresult.ErrInternalServer
	}
	data, err1 := h.repository.DeleteSeller( context.TODO(), deleteSell)
	if err1 != nil {
		fmt.Println(err1)
	}

	successResult := appresult.Success
	successResult.Data = data

	w.Header().Add(appresult.HeaderContentTypeJson())
	err := json.NewEncoder(w).Encode(successResult)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}