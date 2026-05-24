package schemas

type CameraRequest struct {
	ID     int    `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	URL    string `json:"url" binding:"required"`
	Active *bool  `json:"active" binding:"required"`
}
