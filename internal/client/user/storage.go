package user

import "context"

type Repository interface{
	UserRegister(ctx context.Context, user UserRegister) (UserRegister,error)
	UserLogin(ctx context.Context, username string) (UserData, error)
	
}