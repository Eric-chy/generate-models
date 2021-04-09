package model

type LvSpiderbots struct {
	Id int64 `json:"id"` 
	Spidername string `json:"spidername"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
}
func (l LvSpiderbots) TableName() string {
	return "kbb_hotel_orders"
}