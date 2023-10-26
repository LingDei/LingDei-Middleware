package utils

import (
	"LingDei_Middleware/config"
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func getBucketManager() (*storage.BucketManager, error) {
	// 初始化mac
	mac := qbox.NewMac(config.QINIU_ACCESS_KEY, config.QINIU_SECRET_KEY)

	// 初始化cfg
	cfg := storage.Config{}

	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong

	// 是否使用https域名
	cfg.UseHTTPS = true

	return storage.NewBucketManager(mac, &cfg), nil
}

// UploadFile 文件上传（分片）
func UploadFileByPieces(key string, file io.Reader) error {
	bucket_manager, err := getBucketManager()
	if err != nil {
		return err
	}

	// 获取上传凭证
	putPolicy := storage.PutPolicy{
		Scope: config.QINIU_BUCKET,
	}
	upload_token := putPolicy.UploadToken(bucket_manager.Mac)

	// 分片上传
	resumeUploader := storage.NewResumeUploaderV2(bucket_manager.Cfg)
	upHost, err := resumeUploader.UpHost(config.QINIU_ACCESS_KEY, config.QINIU_BUCKET)
	if err != nil {
		return err
	}

	// 初始化分块上传
	initPartsRet := storage.InitPartsRet{}
	err = resumeUploader.InitParts(context.TODO(), upload_token, upHost, config.QINIU_BUCKET, key, true, &initPartsRet)
	if err != nil {
		return err
	}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	fileLen := len(fileContent)
	chunkSize2 := 2 * 1024 * 1024

	num := fileLen / chunkSize2
	if fileLen%chunkSize2 > 0 {
		num++
	}

	// 分块上传
	var uploadPartInfos []storage.UploadPartInfo
	for i := 1; i <= num; i++ {
		partNumber := int64(i)

		var partContentBytes []byte
		endSize := i * chunkSize2
		if endSize > fileLen {
			endSize = fileLen
		}
		partContentBytes = fileContent[(i-1)*chunkSize2 : endSize]
		partContentMd5 := fmt.Sprintf("%x", md5.Sum(partContentBytes))
		uploadPartsRet := storage.UploadPartsRet{}
		err = resumeUploader.UploadParts(context.TODO(), upload_token, upHost, config.QINIU_BUCKET, key, true,
			initPartsRet.UploadID, partNumber, partContentMd5, &uploadPartsRet, bytes.NewReader(partContentBytes),
			len(partContentBytes))
		if err != nil {
			return err
		}
		uploadPartInfos = append(uploadPartInfos, storage.UploadPartInfo{
			Etag:       uploadPartsRet.Etag,
			PartNumber: partNumber,
		})
	}

	// 完成上传
	rPutExtra := storage.RputV2Extra{Progresses: uploadPartInfos}
	comletePartRet := storage.PutRet{}
	err = resumeUploader.CompleteParts(context.TODO(), upload_token, upHost, &comletePartRet, config.QINIU_BUCKET, key,
		true, initPartsRet.UploadID, &rPutExtra)
	if err != nil {
		return err
	}

	fmt.Println(comletePartRet.Key, comletePartRet.Hash)

	return nil
}

// DeleteFile 删除文件
func DeleteFile(key string) error {
	bucket_manager, err := getBucketManager()
	if err != nil {
		return err
	}

	err = bucket_manager.Delete(config.QINIU_BUCKET, key)
	if err != nil {
		return err
	}

	return nil
}
