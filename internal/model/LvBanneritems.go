package model

type LvBanneritems struct {
	Id int64 `json:"id"` 
	BannerId int `json:"banner_id"` 
	Pic string `json:"pic"` 
	Title string `json:"title"` 
	Digest string `json:"digest"` 
	Sort int `json:"sort"` 
	Isshow int `json:"isshow"` // 1：显示，0：不显示
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
}
func (l LvBanneritems) TableName() string {
	return "kbb_hotel_orders"
}