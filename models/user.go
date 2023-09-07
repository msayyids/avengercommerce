package models

type Users struct {
	Id            int     `json:"id"`
	UserName      string  `json:"user_name" binding:"required,unique"`
	Password      string  `json:"password" binding:"required,min=5"`
	DepositAmount float32 `json:"deposit_amount"`
}

type JWtClaims struct {
	UserId string `json:"user_id"`
	// jwt.registered for exprd
}
