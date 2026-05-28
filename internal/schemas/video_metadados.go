package schemas

type VideoMetadata struct {
	ID          int    `json:"id"`
	CameraID    int    `json:"camera_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Duration    int    `json:"duration"` // Duration in seconds
	Format      string `json:"format"`
	Size        int64  `json:"size"` // Size in bytes
	CreatedAt   string `json:"created_at"`
	FilePath    string `json:"file_path"`
}
