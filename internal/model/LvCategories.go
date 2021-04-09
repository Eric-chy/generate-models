package model

type LvCategories struct {
	Id int64 `json:"id"` 
	Name string `json:"name"` 
	Pid int `json:"pid"` 
	Sort int `json:"sort"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
}
func (l LvCategories) TableName() string {
	return "kbb_hotel_orders"
}