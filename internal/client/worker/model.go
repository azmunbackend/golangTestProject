package worker

type AddWorker struct {
	ID     int    `json:"id"`
	UserId int `json:"user_id"`
	Role   string `json:"role"`
	PageId int `json: "page_id"`
}
