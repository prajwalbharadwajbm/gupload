package models

type Files struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Filename    string `json:"filename"`
	FilePath    string `json:"file_path"`
	SizeBytes   int    `json:"size_bytes"`
	ContentType string `json:"content_type"`
	IsDeleted   string `json:"is_deleted"`
}
