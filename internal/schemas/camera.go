package schemas

import (
	"time"
)

type Cameras struct {
	ID int64 `db:"id"`

	Name string `db:"name"`

	IP   string `db:"ip"`
	Port int    `db:"port"`

	RTSPUrl string `db:"rtsp_url"`

	Username          string `db:"username"`
	PasswordEncrypted string `db:"password_encrypted"`

	Brand string `db:"brand"`
	Model string `db:"model"`

	Location string `db:"location"`

	FPS int `db:"fps"`

	ResolutionWidth  int `db:"resolution_width"`
	ResolutionHeight int `db:"resolution_height"`

	IsActive bool `db:"is_active"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
