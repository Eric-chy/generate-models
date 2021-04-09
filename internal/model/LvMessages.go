package model

type LvMessages struct {
	Id int `json:"id"` 
	From string `json:"from"` 
	Content string `json:"content"` 
	Status int `json:"status"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
	To string `json:"to"` 
	ArticleId int `json:"article_id"` 
	Email string `json:"email"` 
}
func (l LvMessages) TableName() string {
	return "kbb_hotel_orders"
}