package types

type Parameter struct {
	ParamID       string      `gorm:"primary_key;column:param_id"`
	Type          string      `gorm:"column:type"`
	Name          string      `gorm:"column:name"`
	Value         string      `gorm:"column:value"`
	Placeholder   string      `gorm:"column:placeholder"`
	JoinDirection string      `gorm:"column:join_direction"`
	JoinItems     []Parameter `gorm:"-"`
}

type Task struct {
	ID         string      `gorm:"primary_key;column:id"`
	Name       string      `gorm:"column:name"`
	Module     string      `gorm:"column:module"`
	Status     string      `gorm:"column:status"`
	Parameters []Parameter `gorm:"foreignkey:ParamID"`
}
