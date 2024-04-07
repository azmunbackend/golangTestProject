package seller

type GetAllSeller struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Suname string `json:"suname"`
}

type AddSeller struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Suname string `json:"suname"`
}

type GetByIdSeller struct{
	ID int `json:"id"`
	Name string `json:"name"`
}

type UpdateSeller struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Suname string `json:"suname"`
}

type DeleteSeller struct{
	ID int `json:"id"`
}