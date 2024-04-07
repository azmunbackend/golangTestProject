package seller

import "context"

type Repository interface{
	GetAllSeller(ctx context.Context) ([]GetAllSeller, error)
	AddSeller( ctx context.Context, addsell AddSeller) (int, error)
	GetByIdSeller(ctx context.Context, id string) ([]GetByIdSeller, error)
	UpdateSeller(ctx context.Context, updatesell UpdateSeller) (int, error)
	DeleteSeller(ctx context.Context, deleteSell DeleteSeller) (int, error)
}
