package models

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type TaskInfoDTO struct {
	Finished   bool          `json:"finished"`
	Path       string        `json:"path"`
	TaskRecord TaskRecordDTO `json:"taskRecord"`
}

type TaskRecordDTO struct {
	SysUploadTask            // 嵌入SysUploadTask结构体，实现继承的效果
	ExitPartList  []*s3.Part `json:"exitPartList"`
}

func ConvertFromEntity(task SysUploadTask) TaskRecordDTO {
	dto := TaskRecordDTO{
		SysUploadTask: task,
	}
	return dto
}

type InitTaskParam struct {
	Identifier string `validate:"required"`
	TotalSize  int64  `json:"totalSize"`
	ChunkSize  int64  `json:"chunkSize"`
	FileName   string `validate:"required"`
}
