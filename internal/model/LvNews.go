package model

type LvNews struct {
	Id int64 `json:"id"` 
	Title string `json:"title"` 
	Keyword string `json:"keyword"` 
	Pic string `json:"pic"` 
	Desc string `json:"desc"` 
	Remark string `json:"remark"` 
	Category string `json:"category"` 
	Views int `json:"views"` 
	Content string `json:"content"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
	Original int `json:"original"` 
	Top int `json:"top"` 
	Recommend int `json:"recommend"` 
	Child int `json:"child"` 
}
func (l LvNews) TableName() string {
	return "kbb_hotel_orders"
}