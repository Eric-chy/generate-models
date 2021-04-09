package model

type LvPasswordResets struct {
	Email string `json:"email"` 
	Token string `json:"token"` 
	CreatedAt time.Time `json:"created_at"` 
}
func (l LvPasswordResets) TableName() string {
	return "kbb_hotel_orders"
}