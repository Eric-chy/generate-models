package model

type LvUsers struct {
	Id int64 `json:"id"` 
	Name string `json:"name"` 
	Email string `json:"email"` 
	EmailVerifiedAt time.Time `json:"email_verified_at"` 
	Password string `json:"password"` 
	RememberToken string `json:"remember_token"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
}
func (l LvUsers) TableName() string {
	return "kbb_hotel_orders"
}