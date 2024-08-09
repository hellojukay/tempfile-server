package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"sync"

	"github.com/hellojukay/tempfile-server/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	once sync.Once
)

func GetMinIOClient() *minio.Client {
	var c *minio.Client
	once.Do(func() {
		client, err := minio.New(config.S3EndPoint, &minio.Options{
			Creds:  credentials.NewStaticV4(config.AccessKeyID, config.AccessKeySecret, ""),
			Secure: config.UseSSL,
		})
		if err != nil {
			panic(err)
		}
		c = client
	})
	return c
}

type MinIOServer struct {
}

// implemnts http.Handler for MinIOServer
func (s *MinIOServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		s.Get(w, r)
	case http.MethodPut, http.MethodPost:
		s.PostOrPut(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (s *MinIOServer) Get(w http.ResponseWriter, r *http.Request) {
	// get file name from url path
}

func (s *MinIOServer) PostOrPut(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	// parse file from request
	r.ParseMultipartForm(32 << 20)
	file, fileName, err := getFilefromRequest(r)
	if err != nil {
		log.Printf("upload file %s\n", path.Base(r.URL.Path))
		return err
	}
	targetDir := path.Join(config.BucketName, path.Dir(r.URL.Path))
	if e := s.createPathIfNotExist(config.BucketName, targetDir); e != nil {
		return e
	}
	minioClient := GetMinIOClient()
	info, err := minioClient.PutObject(context.Background(), config.BucketName, fileName, file, -1, minio.PutObjectOptions{})
	log.Printf(info.ChecksumSHA1)
	return err
}

func (s *MinIOServer) createPathIfNotExist(bukect string, objectPath string) error {
	minioClient := GetMinIOClient()
	if minioClient == nil {
		return fmt.Errorf("minio client is nil")
	}
	// check if bucket exists
	exist, err := minioClient.BucketExists(context.Background(), bukect)
	if err != nil {
		log.Printf("check bucket %s exist failed: %v\n", bukect, err)
		return err
	}
	if exist {
		return nil
	}
	err = minioClient.MakeBucket(context.Background(), bukect, minio.MakeBucketOptions{})
	if err != nil {
		log.Printf("create bucket %s failed: %v\n", bukect, err)
		return err
	}
	return nil
}
