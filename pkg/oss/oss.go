package oss

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"

	"github.com/Mutezebra/tiktok/consts"

	"github.com/qiniu/go-sdk/v7/storage"
)

type Model struct {
}

func NewOssModel() *Model {
	return &Model{}
}

// UploadAvatar from the oss
func (m *Model) UploadAvatar(ctx context.Context, name string, data []byte) (err error, path string) {
	mac, bucket, _ := getOSS()
	path = conf.AvatarPath + name

	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, path),
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region: &storage.ZoneHuadong,
	}
	formLoader := storage.NewFormUploader(&cfg)

	ret := &PutRet{}
	putExtra := &storage.PutExtra{}

	length := int64(len(data))

	err = formLoader.Put(ctx, ret, upToken, path, bytes.NewReader(data), length, putExtra)
	return errors.Wrap(err, "upload avatar failed"), path
}

func (m *Model) DownloadAvatar(ctx context.Context, path string) string {
	_, _, domain := getOSS()

	publicURL := storage.MakePublicURLv2(domain, path)
	return publicURL
}

func (m *Model) UploadVideo(ctx context.Context, name string, data []byte) (err error, path string) {
	mac, bucket, _ := getOSS()
	path = conf.VideoPath + name

	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, path),
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region: &storage.ZoneHuadong,
	}
	resumeUploader := storage.NewResumeUploaderV2(&cfg)

	ret := PutRet{}
	putExtra := storage.RputV2Extra{PartSize: 2 * consts.MB}

	err = resumeUploader.Put(ctx, &ret, upToken, path, bytes.NewReader(data), int64(len(data)), &putExtra)
	return errors.Wrap(err, "upload video failed"), path
}

func (m *Model) DownloadVideo(ctx context.Context, path string) string {
	_, _, domain := getOSS()

	publicURL := storage.MakePublicURLv2(domain, path)
	return publicURL
}

func (m *Model) VideoFeed(name string) ([]byte, error) {
	mac, bucket, domain := getOSS()
	key := conf.VideoPath + name
	bm := storage.NewBucketManager(mac, &storage.Config{
		Region: &storage.ZoneHuadong,
	})

	resp, err := bm.Get(bucket, key, &storage.GetObjectInput{
		DownloadDomains: []string{domain},
		PresignUrl:      true,
		Range:           "bytes=0-",
	})
	if err != nil {
		return nil, errors.Wrap(err, "get video feed failed")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read video feed failed")
	}

	return data, nil
}

func (m *Model) UploadVideoCover(ctx context.Context, name string, data []byte) (err error, path string) {
	mac, bucket, _ := getOSS()
	path = conf.CoverPath + name

	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, path),
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region: &storage.ZoneHuadong,
	}
	formLoader := storage.NewFormUploader(&cfg)

	ret := &PutRet{}
	putExtra := &storage.PutExtra{}

	length := int64(len(data))

	err = formLoader.Put(ctx, ret, upToken, path, bytes.NewReader(data), length, putExtra)
	return errors.Wrap(err, "upload video cover failed"), path
}

func (m *Model) DownloadVideoCover(ctx context.Context, path string) string {
	_, _, domain := getOSS()

	publicURL := storage.MakePublicURLv2(domain, path)
	return publicURL
}
