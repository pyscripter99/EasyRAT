package types

type Parameter struct {
	Type          string
	Name          string
	Value         string
	Placeholder   string
	JoinDirection string
	JoinItems     []Parameter
}

type Task struct {
	ID         string `gorm:"primary_key"`
	Name       string
	Module     string
	Status     string
	Parameters []Parameter
}
