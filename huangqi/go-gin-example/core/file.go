package core

import (
	"fmt"
	"go-gin-example/conf"
	"go-gin-example/models"
	"math"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/google/uuid"
)

func UploadToMinIO(objectKey, filePath string) error {
	// 获取 AWS 会话
	svc, err := conf.GetAwsS3()
	if err != nil {
		return err
	}
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建上传对象的输入参数
	uploadInput := &s3.PutObjectInput{
		Bucket: aws.String(conf.BucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	}

	// 执行上传对象操作
	_, err = svc.PutObject(uploadInput)
	if err != nil {
		return err
	}

	return nil
}

func GetStsUploadCredentials() (*sts.Credentials, error) {
	// 获取 AWS 会话
	svc, err := conf.GetAwsSts()
	if err != nil {
		return nil, err
	}

	// 指定 STS 角色 ARN
	roleARN := "arn:aws:iam::YOUR_ACCOUNT_ID:role/YOUR_ROLE_NAME"

	// 指定临时凭证的有效期（秒）
	durationSeconds := int64(3600)

	// 创建 AssumeRole 输入参数
	assumeRoleInput := &sts.AssumeRoleInput{
		RoleArn:         aws.String(roleARN),
		RoleSessionName: aws.String("s3-upload-session"),
		DurationSeconds: aws.Int64(durationSeconds),
	}

	// 获取临时凭证
	assumeRoleOutput, err := svc.AssumeRole(assumeRoleInput)
	if err != nil {
		return nil, err
	}

	// 返回临时凭证
	return assumeRoleOutput.Credentials, nil
}

// 初始化任务
func InitTask(param models.InitTaskParam) (models.TaskInfoDTO, error) {
	svc, _ := conf.GetAwsS3()
	bucketName := conf.BucketName
	fileName := param.FileName
	suffix := fileName[strings.LastIndex(fileName, ".")+1:]
	key := fmt.Sprintf("%s/%s.%s", time.Now().Format("2006-01-02"), uuid.New().String(), suffix)
	contentType := getMediaType(key)
	fmt.Println(key)
	fmt.Println(contentType)
	input := &s3.CreateMultipartUploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}
	result, err := svc.CreateMultipartUpload(input)
	if err != nil {
		return models.TaskInfoDTO{}, fmt.Errorf("无法创建分片上传任务: %v", err)
	}
	uploadID := *result.UploadId
	task := models.SysUploadTask{
		BucketName:     conf.BucketName,
		ChunkNum:       int(math.Ceil(float64(param.TotalSize) / float64(param.ChunkSize))),
		ChunkSize:      param.ChunkSize,
		TotalSize:      param.TotalSize,
		FileIdentifier: param.Identifier,
		FileName:       fileName,
		ObjectKey:      key,
		UploadID:       uploadID,
	}
	InsertUploadTask(task)
	return models.TaskInfoDTO{
		Finished:   false,
		TaskRecord: models.ConvertFromEntity(task),
		Path:       GetPath(conf.Endpoint, conf.BucketName, key),
	}, nil
}

// 生成预签名上传 URL
func GenPreSignUploadURL(objectKey, uploadID string, partNumber int64) (string, error) {
	currentDate := time.Now()
	expireDate := currentDate.Add(time.Duration(60*10*1000) * time.Millisecond)
	svc, _ := conf.GetAwsS3()
	request, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(conf.BucketName),
		Key:    aws.String(objectKey),
	})
	// 添加参数
	params := map[string][]string{
		"partNumber": {strconv.FormatInt(partNumber, 10)},
		"uploadId":   {uploadID},
	}
	request.HTTPRequest.URL.RawQuery = url.Values(params).Encode()
	// 计算时间间隔
	duration := expireDate.Sub(currentDate)
	url, err := request.Presign(duration)
	if err != nil {
		return "", fmt.Errorf("无法生成预签名上传 URL: %v", err)
	}
	return url, nil
}

// 根据文件扩展名获取媒体类型
func getMediaType(key string) string {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(key), "."))
	mimeType := mime.TypeByExtension("." + ext)
	if mimeType == "" {
		// 如果无法通过扩展名获取到媒体类型，则使用 http.DetectContentType 函数进行判断
		file, err := os.Open(key)
		if err != nil {
			panic(fmt.Errorf("无法打开文件：%v", err))
		}
		defer file.Close()
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			panic(fmt.Errorf("无法读取文件：%v", err))
		}
		mimeType = http.DetectContentType(buffer)
	}
	return mimeType
}

func GetPath(endpoint, bucket, objectKey string) string {
	return endpoint + "/" + bucket + "/" + objectKey
}

func GetTaskInfo(identifier string) *models.TaskInfoDTO {
	task, err := GetByIdentifier(identifier)
	emptyStruct := models.SysUploadTask{}
	if err != nil || task == emptyStruct {
		return nil
	}
	result := &models.TaskInfoDTO{
		Finished:   true,
		TaskRecord: models.ConvertFromEntity(task),
		Path:       GetPath(conf.Endpoint, task.BucketName, task.ObjectKey),
	}
	input := &s3.HeadObjectInput{
		Bucket: aws.String(task.BucketName),
		Key:    aws.String(task.ObjectKey),
	}
	svc, _ := conf.GetAwsS3()
	_, err = svc.HeadObject(input)
	if err != nil {
		// 未上传完，返回已上传的分片
		listInput := &s3.ListPartsInput{
			Bucket:   aws.String(task.BucketName),
			Key:      aws.String(task.ObjectKey),
			UploadId: aws.String(task.UploadID),
		}
		listOutput, err := svc.ListParts(listInput)
		fmt.Println(listOutput)
		fmt.Println(err != nil)
		if err != nil {
			return nil
		}
		result.Finished = false
		result.TaskRecord.ExitPartList = listOutput.Parts
	}
	return result
}

func Merge(identifier string) {
	// 获取分片任务信息
	task, err := GetByIdentifier(identifier)
	fmt.Println("开始合并分片")
	fmt.Println("***********************************************")
	emptyStruct := models.SysUploadTask{}
	if err != nil || task == emptyStruct {
		fmt.Println(err)
		return
	}
	svc, _ := conf.GetAwsS3()
	// 列出已上传的分块
	listPartsInput := &s3.ListPartsInput{
		Bucket:   aws.String(task.BucketName),
		Key:      aws.String(task.ObjectKey),
		UploadId: aws.String(task.UploadID),
	}
	partListing, err := svc.ListParts(listPartsInput)
	if err != nil {
		fmt.Println("列出分块失败:", err)
		return
	}
	parts := partListing.Parts
	if len(parts) != int(task.ChunkNum) {
		// 已上传分块数量与记录中的数量不对应，不能合并分块
		fmt.Println("分片缺失，请重新上传")
		return
	}
	// 构造CompleteMultipartUpload请求
	completeUploadInput := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(task.BucketName),
		Key:      aws.String(task.ObjectKey),
		UploadId: aws.String(task.UploadID),
	}
	// 构造PartETags
	var partETags []*s3.CompletedPart
	for _, part := range parts {
		partETag := &s3.CompletedPart{
			ETag:       part.ETag,
			PartNumber: part.PartNumber,
		}
		partETags = append(partETags, partETag)
	}
	completeUploadInput.SetMultipartUpload(&s3.CompletedMultipartUpload{
		Parts: partETags,
	})
	// 完成合并分块操作
	_, err = svc.CompleteMultipartUpload(completeUploadInput)
	if err != nil {
		fmt.Println("合并分块失败:", err)
		return
	}
	fmt.Println("***********************************************")
	fmt.Println("合并分块成功")
}
