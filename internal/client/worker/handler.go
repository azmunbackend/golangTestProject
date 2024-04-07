package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test-crm/internal/appresult"
	"test-crm/internal/handlers"
	"test-crm/pkg/logging"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

const (
	AddWorkerURL = "/add-worker"
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
	

func ( h handler) Register(router *mux.Router)  {
	router.HandleFunc(AddWorkerURL, appresult.MiddTokenChkSupAdmin((appresult.Middleware(h.AddWorker)))).Methods("POST")
}

func (h *handler) AddWorker(w http.ResponseWriter, r *http.Request) error {

	token1 := r.Header.Get("Authorization")
    token, err := jwt.Parse(token1, func(token *jwt.Token) (interface{}, error) {
        return []byte("normalnybol!!!"), nil
    })
    if err != nil {
		fmt.Println("lllllllllllllllllllllllllllll")
        panic(err.Error())
    }

    var userID int
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    userID = int(claims["id"].(float64))
}


    body, errBody := ioutil.ReadAll(r.Body)
    if errBody != nil {
        fmt.Println(errBody)
        return appresult.ErrMissingParam
    }

    var worker AddWorker
    errData := json.Unmarshal(body, &worker)
    if errData != nil {
        fmt.Println(errData)
        return appresult.ErrInternalServer
    }

    data, errr1 := h.repository.AddWorker(context.TODO(), worker, userID)
    if errr1 != nil {
        fmt.Println(errr1)
        return errr1
    }

    successResult := appresult.Success
    successResult.Data = data

    w.Header().Add(appresult.HeaderContentTypeJson())
    err1 := json.NewEncoder(w).Encode(successResult)
    if err1 != nil {
        fmt.Println(err1)
        return err1
    }
    return nil
}
