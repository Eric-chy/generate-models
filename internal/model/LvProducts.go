package model

type LvProducts struct {
	Id int64 `json:"id"` 
	Cid int `json:"cid"` 
	Title string `json:"title"` 
	Keyword string `json:"keyword"` 
	Desc string `json:"desc"` 
	Remark string `json:"remark"` 
	Pic string `json:"pic"` 
	Tuijian int `json:"tuijian"` 
	Views int `json:"views"` 
	Alt string `json:"alt"` 
	Content string `json:"content"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
}
func (l LvProducts) TableName() string {
	return "kbb_hotel_orders"
}