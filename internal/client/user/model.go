package user

type MyError struct {
    Message string
}

func (e MyError) Error() string {
    return e.Message
}

type UserRegister struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Password string `json: "password"`
}

type userLogin struct{
	Name string `json:"name"`
	Password string `json: "password"`
}

type UserData struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Password string `json: "password"`
}

type ResultTokenDTO struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Token string `json:"token"`
}