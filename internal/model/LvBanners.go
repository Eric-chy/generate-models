package model

type LvBanners struct {
	Id int64 `json:"id"` 
	Title string `json:"title"` 
	Entitle string `json:"entitle"` 
}
func (l LvBanners) TableName() string {
	return "kbb_hotel_orders"
}