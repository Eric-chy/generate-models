package model

type LvLink struct {
	Id int `json:"id"` 
	Name string `json:"name"` 
	Link string `json:"link"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
	Status int `json:"status"` 
	Logo string `json:"logo"` 
}
func (l LvLink) TableName() string {
	return "kbb_hotel_orders"
}