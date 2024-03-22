package entity

type Instance struct {
	Id   int64  `json:"instance_id" db:"id"`
	Name string `json:"instance_name" db:"instance_name"`
}
