package entity

type Project struct {
	Id   int64  `json:"project_id" db:"id"`
	Name string `json:"project_name" db:"project_name"`
}
