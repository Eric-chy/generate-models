package model

type LvCases struct {
	Id int64 `json:"id"` 
	Title string `json:"title"` 
	Remark string `json:"remark"` 
	Pic string `json:"pic"` 
	Sort int `json:"sort"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
}
func (l LvCases) TableName() string {
	return "kbb_hotel_orders"
}