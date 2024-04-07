package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test-crm/internal/appresult"
	"test-crm/internal/config"
	"test-crm/internal/handlers"
	"test-crm/pkg/logging"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserLoginURL = "/login"
	UserRegisterURL = "/register"
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

func (h handler) Register(router *mux.Router) {
	router.HandleFunc(UserRegisterURL, appresult.Middleware(h.UserRegister)).Methods("POST")
	router.HandleFunc(UserLoginURL, appresult.Middleware(h.UserLogin)).Methods("GET")
}

func (h handler) UserRegister(w http.ResponseWriter, r *http.Request) error{
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("errrr UserEgister: ", errBody)
	}
	userRegister := UserRegister{}
	
	errData := json.Unmarshal(body, &userRegister)
	if errData != nil {
		fmt.Println(errData)
		return appresult.ErrInternalServer
	}

	data , err1 := h.repository.UserRegister(context.TODO(), userRegister)
	
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

func (h handler) UserLogin(w http.ResponseWriter, r *http.Request) error{
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("error login body:", err)
		return appresult.ErrMissingParam
	}

	login := userLogin{}

	errData := json.Unmarshal(body, &login)
	if errData != nil {
		fmt.Println("error :, ", errData)
		return appresult.ErrMissingParam
	}
	data, err := h.repository.UserLogin(context.TODO(), login.Name)
	if err != nil {
		fmt.Println("Error h.repository.GetEmployeeData :", err)
		return appresult.ErrInternalServer
	}
	if data.Name == "" {
		fmt.Println("Username not correct")
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(login.Password))
	if err != nil {
		fmt.Println("Password not correct", err)
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}
	tokenDTO := ""
	atClaims := jwt.MapClaims{}
	atClaims["id"] = data.ID
	atClaims["name"] = data.Name
	atClaims["surname"] = data.Surname
	atClaims["exp"] = time.Now().Add(time.Minute * 60 * 12).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	cfg := config.GetConfig()
	tokenDTO, err = at.SignedString([]byte(cfg.JwtKey))
	if err != nil {
		fmt.Println("Token can not generete")
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	successResult := appresult.Success
	successResult.Data = ResultTokenDTO{
		ID:  data.ID,
		Name: data.Name,
		Surname: data.Surname,
		Token: tokenDTO,
	}

	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}
	return nil

}
