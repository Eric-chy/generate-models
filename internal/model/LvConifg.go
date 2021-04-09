package model

type LvConifg struct {
	Id int64 `json:"id"` 
	Name string `json:"name"` // 英文名称
	Title string `json:"title"` // 中文名称
	Config string `json:"config"` // 配置信息
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
}
func (l LvConifg) TableName() string {
	return "kbb_hotel_orders"
}