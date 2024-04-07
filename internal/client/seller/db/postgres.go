package seller

import (
	"context"
	"fmt"
	"test-crm/internal/client/seller"
	"test-crm/pkg/client/postgresql"
	"test-crm/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client,  logger *logging.Logger) seller.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func(r repository) GetAllSeller(ctx context.Context) ([]seller.GetAllSeller, error) {
	var listget []seller.GetAllSeller
	q := "select * from seller"
	rows, err := r.client.Query(ctx,q)

	if err != nil {
		fmt.Println("err", err)
	}
	defer rows.Close()

	for rows.Next(){
		var list2 seller.GetAllSeller
		err = rows.Scan(&list2.ID, &list2.Name, &list2.Suname)
		if err != nil {
			fmt.Println(err)
		}
		listget = append(listget, list2)
	}

	return listget,nil
}

func (r repository) AddSeller(ctx context.Context, addsell  seller.AddSeller) (int, error) {
	var id int;
	q:= `insert into seller(name, suname) values($1, $2) returning id;`
	 err:= r.client.QueryRow(ctx, q, addsell.Name, addsell.Suname).Scan(&id)
	if err != nil {
			fmt.Println("r.client", err)
		}
	return id, nil
}

func (r repository)GetByIdSeller(ctx context.Context, id string) ([]seller.GetByIdSeller, error)  {
	var result []seller.GetByIdSeller
	q:= `select id, name from seller where id = $1`
	rows, err := r.client.Query(ctx, q, id)
	if err != nil {
		fmt.Println("error", err)
	}
	defer rows.Close()
	for rows.Next(){
		getSeller:= seller.GetByIdSeller{}
		err = rows.Scan(&getSeller.ID, &getSeller.Name)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, getSeller)
	}
	return result, nil
}

func ( r repository) UpdateSeller(ctx context.Context, updatesell seller.UpdateSeller) (int, error) {
	var id int
	q := `update seller set name =$1, suname = $2 where id = $3 returning id`
	err := r.client.QueryRow(ctx, q, updatesell.Name, updatesell.Suname, updatesell.ID).Scan(&id)

	if err != nil {
		fmt.Println("UpdateSeller => ", err)
		return 0, err
	}

	return updatesell.ID, nil
}

func (r repository) DeleteSeller (ctx context.Context, deletesell seller.DeleteSeller) (int, error){
	var id int
	q := `delete from seller where id = $1 returning id`
	err := r.client.QueryRow(ctx, q, deletesell.ID).Scan(&id)
	if err != nil {
		fmt.Println("DeleteSeller => ", err)
		return 0, err
	}

	return deletesell.ID, nil
}