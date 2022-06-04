package utils

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
)

var TempDomainName = "http://rcqzzkaz6.hd-bkt.clouddn.com"

func QiniuUpload(key string, data []byte) error {
	accessKey := "8VLtl-yhshJ_brzLSNx-WBvoaxw4jO_GO5UCKJ9s"
	secretKey := "vVoVZ6aH1AhrUDvJYTk2k9bojUJwYR-99H_bN2Aq"
	mac := qbox.NewMac(accessKey, secretKey)
	bucket := "noavideo"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
