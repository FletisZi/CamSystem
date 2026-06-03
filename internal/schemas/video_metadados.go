package schemas

// VideoMetadata
// type VideoRecordings struct {
// 	CameraID    int    `json:"camera_id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	Duration    int    `json:"duration"` // Duration in seconds
// 	Format      string `json:"format"`
// 	Size        int64  `json:"size"` // Size in bytes
// 	CreatedAt   string `json:"created_at"`
// 	FilePath    string `json:"file_path"`
// }

type VideoRecordings struct {
	CameraID       int     `json:"camera_id"`
	FileName       string  `json:"file_name"`
	FilePath       string  `json:"file_path"`
	Event_type     string  `json:"event_type"`
	TicketID       *string `json:"ticket_id"`
	PlacaDetectada string  `json:"placa_detectada"`
}
