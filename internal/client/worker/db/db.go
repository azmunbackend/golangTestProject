package worker

import (
	"context"
	"fmt"

	"test-crm/internal/client/worker"
	"test-crm/pkg/client/postgresql"
	"test-crm/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client,  logger *logging.Logger) worker.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r repository) AddWorker(ctx context.Context, worker worker.AddWorker, userID int) (int, error) {
    var id int
    fmt.Println("dddddddd", worker)
    
    if worker.PageId != 0 {
        q := `INSERT INTO workers (user_id, role, page_id) VALUES ($1, $2, $3) RETURNING id`
        err := r.client.QueryRow(ctx, q, userID, worker.Role, worker.PageId).Scan(&id)
        if err != nil {
            fmt.Println("ERRRRRRR", err)
            return 0, err
        }
    } 
    
    return id, nil
}
