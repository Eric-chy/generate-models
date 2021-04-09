package model

type LvManagers struct {
	Id int64 `json:"id"` 
	Username string `json:"username"` 
	Password string `json:"password"` 
	Status int `json:"status"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
}
func (l LvManagers) TableName() string {
	return "kbb_hotel_orders"
}