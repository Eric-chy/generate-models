package model

type LvFriends struct {
	Id int64 `json:"id"` 
	Title string `json:"title"` 
	Url string `json:"url"` 
	Sort int `json:"sort"` 
	Key string `json:"key"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
	Logo string `json:"logo"` 
}
func (l LvFriends) TableName() string {
	return "kbb_hotel_orders"
}