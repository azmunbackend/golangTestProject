package worker

import "context"


type Repository interface {

	AddWorker(ctx context.Context, worker AddWorker, userID int) (int, error)

}