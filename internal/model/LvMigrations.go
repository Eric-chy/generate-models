package model

type LvMigrations struct {
	Id int `json:"id"` 
	Migration string `json:"migration"` 
	Batch int `json:"batch"` 
}
func (l LvMigrations) TableName() string {
	return "kbb_hotel_orders"
}