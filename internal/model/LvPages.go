package model

type LvPages struct {
	Id int64 `json:"id"` 
	Title string `json:"title"` 
	Type int `json:"type"` // 1：公司简介，2：招贤纳士，3：发展历程
	Remark string `json:"remark"` 
	Pic string `json:"pic"` 
	Content string `json:"content"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
}
func (l LvPages) TableName() string {
	return "kbb_hotel_orders"
}