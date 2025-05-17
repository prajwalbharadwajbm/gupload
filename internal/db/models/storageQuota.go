package models

type StorageQuota struct {
	ID        string `json:"user_id"`
	MaxBytes  int    `json:"max_bytes"`
	UsedBytes int    `json:"used_bytes"`
}
