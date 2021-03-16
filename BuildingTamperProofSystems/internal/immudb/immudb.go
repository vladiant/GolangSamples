package immudb

import (
	"context"
	"encoding/json"
	"log"

	immuclient "github.com/codenotary/immudb/pkg/client"
	"google.golang.org/grpc/metadata"
)

// Engineer -
type Engineer struct {
	ID     string
	Salary int
}

type DB struct {
	Ctx    context.Context
	Client immuclient.ImmuClient
}

func NewImmuDB() (DB, error) {
	client, err := immuclient.NewImmuClient(immuclient.DefaultOptions())
	if err != nil {
		log.Print(err.Error())
		return DB{}, err
	}

	ctx := context.Background()
	lr, err := client.Login(ctx, []byte(`immudb`), []byte(`immudb`))
	if err != nil {
		log.Print(err.Error())
		return DB{}, err
	}

	md := metadata.Pairs("authorization", lr.Token)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	return DB{
		Ctx:    ctx,
		Client: client,
	}, nil
}

func (db *DB) UpdateSalary(id string, salary int) error {
	engineer := Engineer{
		ID:     id,
		Salary: salary,
	}

	jsonEngineer, err := json.Marshal(engineer)
	if err != nil {
		return err
	}

	vtx, err := db.Client.VerifiedSet(db.Ctx, []byte(id), jsonEngineer)
	if err != nil {
		return err
	}

	log.Printf("Successfully commited and verified the engineer's salary update: %v\n", vtx)
	return nil
}

func (db *DB) GetVerifiedSalary(id string) (int, error) {
	ventry, err := db.Client.VerifiedGet(db.Ctx, []byte(id))
	if err != nil {
		log.Fatal(err)
	}

	var eng Engineer
	if err := json.Unmarshal(ventry.Value, &eng); err != nil {
		return 0, err
	}

	return eng.Salary, nil
}
