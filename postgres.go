package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresOrderIntegration struct {
	innerDB *gorm.DB
}

func NewTestPostgresOrderStorage() (*PostgresOrderIntegration, error) {
	internalDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost port=5432 user=postgres password=Qqqwwweee12321 dbname=postgres sslmode=disable",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbManager := &PostgresOrderIntegration{innerDB: internalDB}
	dbManager.CleanDBAfterTest()
	err = dbManager.makeAllTables()
	if err != nil {

		panic(err)
	}
	return dbManager, nil
}

func (storage *PostgresOrderIntegration) CleanDBAfterTest() {
	err := storage.innerDB.Migrator().DropTable(&Requisite{})
	if err != nil {
		panic(err)
	}
	err = storage.innerDB.Migrator().DropTable(&Order{})
	if err != nil {
		panic(err)
	}
	err = storage.innerDB.Migrator().DropTable("order_requisite_relation")
	if err != nil {
		panic(err)
	}
}

func (storage *PostgresOrderIntegration) makeAllTables() (err error) {
	err = storage.innerDB.AutoMigrate(&Order{}, &Requisite{})
	if err != nil {
		return
	}
	err = storage.innerDB.Exec("ALTER TABLE order_requisite_relation ADD FOREIGN KEY (order_id)" +
		"REFERENCES orders (id) ON DELETE CASCADE ON UPDATE CASCADE;").Error
	err = storage.innerDB.Exec("ALTER TABLE order_requisite_relation ADD FOREIGN KEY (requisite_id)" +
		"REFERENCES requisites (id) ON DELETE CASCADE ON UPDATE CASCADE;").Error
	if err != nil {
		return
	}
	return
}

var (
	strings         = []string{"1", "2", "3", "4"}
	ids             = []string{uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()}
	validRequisites = []Requisite{
		{
			ID:      &ids[0],
			Comment: &strings[0],
		},
		{
			ID:      &ids[1],
			Comment: &strings[1],
		},
		{
			ID:      &ids[2],
			Comment: &strings[2],
		},
	}

	//validOrders = []Order{
	//	{
	//		ID:            &ids[0],
	//		Comment:       &strings[0],
	//		RequisiteList: []Requisite{validRequisites[0]},
	//	},
	//	{
	//		ID:            &ids[1],
	//		Comment:       &strings[1],
	//		RequisiteList: []string{*validRequisites[0].ID},
	//	},
	//	{
	//		ID:            &ids[2],
	//		Comment:       &strings[2],
	//		RequisiteList: []string{*validRequisites[0].ID},
	//	},
	//}
	validOrders = []Order{
		{
			ID:            &ids[0],
			Comment:       &strings[0],
			RequisiteList: []Requisite{{ID: validRequisites[0].ID}},
		},
		{
			ID:            &ids[1],
			Comment:       &strings[1],
			RequisiteList: []Requisite{{ID: validRequisites[0].ID}},
		},
		{
			ID:            &ids[2],
			Comment:       &strings[2],
			RequisiteList: []Requisite{{ID: validRequisites[0].ID}},
		},
	}
)
