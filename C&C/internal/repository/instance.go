package repository

import (
	"candc/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type InstanceDB struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func (i InstanceDB) CreateInstance(instance entity.Instance) error {
	query := fmt.Sprintf("INSERT INTO %s (instance_name) values($1)", instancesTable)
	_, err := i.db.Exec(query, instance.Name)
	return err
}

func (i InstanceDB) GetInstance(id int64) ([]entity.Instance, error) {
	var items []entity.Instance

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", instancesTable)
	err := i.db.Select(&items, query, id)

	return items, err
}

func (i InstanceDB) GetAllInstances() ([]entity.Instance, error) {
	var orders []entity.Instance

	query := fmt.Sprintf("SELECT * FROM %s", instancesTable)
	err := i.db.Select(&orders, query)

	return orders, err
}

func (i InstanceDB) UpdateInstance(instance entity.Instance) error {
	query := fmt.Sprintf("UPDATE %s SET instance_name=$1 WHERE id=$10", instancesTable)
	_, err := i.db.Exec(query, instance.Name)
	return err
}

func (i InstanceDB) DeleteInstance(id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", instancesTable)
	_, err := i.db.Exec(query, id)
	return err
}

func NewInstanceDB(db *sqlx.DB, logger *zap.Logger) *InstanceDB {
	return &InstanceDB{
		db:     db,
		logger: logger,
	}
}
