package types

type Client struct {
	ID        string `gorm:"primary_key"`
	Connected bool
	DeviceInfoStruct
}
