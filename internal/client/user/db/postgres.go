package user

import (
	"context"
	"fmt"
	"test-crm/internal/client/user"
	"test-crm/pkg/client/postgresql"
	"test-crm/pkg/logging"

	"golang.org/x/crypto/bcrypt"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client,  logger *logging.Logger) user.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) UserRegister(ctx context.Context,user user.UserRegister )(user.UserRegister, error) {
	pass, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	
	if err1 != nil {
		fmt.Println("errr hash passs")
	}

	q:= `insert into "user"(name, surname, password) values($1, $2, $3) returning id;`
	 err:= r.client.QueryRow(ctx, q, user.Name, user.Surname, pass).Scan(&user.ID)
	if err != nil {
		fmt.Println("errrr bar : ->", err.Error())
		}
	return user, nil
}

func (r *repository) UserLogin(ctx context.Context, username string) (user.UserData, error){
	
	var result user.UserData
	fmt.Println(username)
	q:= `select id, password, name, surname from "user" where name =$1;`
	err := r.client.QueryRow(ctx, q, username).Scan(&result.ID, &result.Password, &result.Name, &result.Surname)

	if err != nil  {
		fmt.Println("Get USERLOGIN  postgres err:   ", err)
		return result, err
	}
	return result, nil
}