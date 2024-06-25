package user

import "time"

type UserFormatter struct {
	IdUser    int       `json:"id_user"`
	Nama      string    `json:"nama"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FormatUser formats a UserTable instance to UserFormatter
func FormatUser(user UserTable, token string) UserFormatter {
	return UserFormatter{
		IdUser:    user.IdUser,
		Nama:      user.Nama,
		Username:  user.Username,
		Role:      user.Role,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

}

// FormatUsers formats a slice of UserTable instances to a slice of UserFormatter