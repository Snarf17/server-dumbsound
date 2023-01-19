package authdto

type RegisterResponse struct {
	Message string `json:"msg"`
}

type LoginResponse struct {
	Message string `json:"msg"`
	Email   string `json:"email"`
	Token   string `json:"token"`
	Role    string `gorm:"type: varchar(50)" json:"role"`
}
type CheckAuthResponse struct {
	Id        int    `gorm:"type: int" json:"id"`
	FullName  string `gorm:"type: varchar(255)" json:"name"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Role      string `gorm:"type: varchar(50)" json:"role"`
	Subscribe string `gorm:"type: varchar(50)" json:"subscribe"`
}
