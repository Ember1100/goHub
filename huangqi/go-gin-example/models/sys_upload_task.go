package models

type SysUploadTask struct {
	ID             int64  `gorm:"column:id" json:"id"`
	UploadID       string `gorm:"column:upload_id" json:"uploadId"`
	FileIdentifier string `gorm:"column:file_identifier" json:"fileIdentifier"`
	FileName       string `gorm:"column:file_name" json:"fileName"`
	BucketName     string `gorm:"column:bucket_name" json:"bucketName"`
	ObjectKey      string `gorm:"column:object_key" json:"objectKey"`
	TotalSize      int64  `gorm:"column:total_size" json:"totalSize"`
	ChunkSize      int64  `gorm:"column:chunk_size" json:"chunkSize"`
	ChunkNum       int    `gorm:"column:chunk_num" json:"chunkNum"`
}

func (SysUploadTask) TableName() string {
	return "sys_upload_task"
}
