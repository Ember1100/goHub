package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func main() {
	endpoint := "http://1.12.42.181:9001"
	bucket := "minio03"
	objectKey := "2023-08-24/e8b3b971-244e-4982-89b3-adf4268e8fc2.xlsx"
	fileName := "e8b3b971-244e-4982-89b3-adf4268e8fc2.xlsx"
	suffix := fileName[strings.LastIndex(fileName, ".")+1:]
	fmt.Println(endpoint + bucket + objectKey)
	key := fmt.Sprintf("%s/%s.%s", time.Now().Format("2006-01-02"), uuid.New().String(), suffix)
	fmt.Print(key)
}
