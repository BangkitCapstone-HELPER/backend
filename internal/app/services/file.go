package services

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/config"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/repo"
	"go.uber.org/fx"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type fileServiceParams struct {
	fx.In
	Upload       lib.ClientUploader
	BucketConfig config.BucketConfig
	FileRepo     repo.FileRepo
	FileConfig   config.FileConfig
}

type FileService interface {
	GetFile(code string) ([]byte, dao.File, error)
	UploadFile(f multipart.FileHeader, folder string) (dao.File, error)
	PredictImage(f multipart.FileHeader) (map[string]interface{}, error)
}

func NewFileService(params fileServiceParams) FileService {
	return &params
}

func (u *fileServiceParams) GetFile(code string) ([]byte, dao.File, error) {
	file, err := u.FileRepo.GetFile(code)
	if err != nil {
		return nil, dao.File{}, err
	}
	fileBytes, err := ioutil.ReadFile(path.Join(u.FileConfig.Path(), file.OriginalFileName))
	if err != nil {
		return nil, dao.File{}, err
	}
	return fileBytes, file, err

}

//func (u *fileServiceParams) UploadFile(f multipart.FileHeader, folder string) (dao.File, error) {
//
//	createdFilename := path.Join(u.FileConfig.Path(), folder, f.Filename)
//	i := 1
//	for {
//		_, err := os.Stat(createdFilename)
//		isFileNotExist := os.IsNotExist(err)
//
//		if isFileNotExist {
//			break
//		}
//		f.Filename = appendToFilename(f.Filename, fmt.Sprintf("_(%d)", i))
//		createdFilename = path.Join(u.FileConfig.Path(), f.Filename)
//		i += 1
//	}
//	src, err := f.Open()
//	if err != nil {
//		return dao.File{}, err
//	}
//	defer src.Close()
//
//	dst, err := os.Create(createdFilename)
//	if err != nil {
//		return dao.File{}, err
//	}
//	defer dst.Close()
//
//	if _, err := io.Copy(dst, src); err != nil {
//		return dao.File{}, err
//	}
//
//	hashedFilename := md5Hash(f.Filename)
//
//	file, err := u.FileRepo.AddFile(dao.File{
//		Extension:        filepath.Ext(f.Filename),
//		OriginalFileName: folder + "/" + f.Filename,
//		FileCode:         hashedFilename,
//	})
//	if err != nil {
//		return dao.File{}, err
//	}
//
//	return file, err
//}

func (u *fileServiceParams) UploadFile(f multipart.FileHeader, folder string) (dao.File, error) {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*5)
	defer cancel()
	blobFile, err := f.Open()
	if err != nil {
		return dao.File{}, err
	}
	hashedFilename := md5Hash(f.Filename)
	extension := filepath.Ext(f.Filename)
	// Upload an object with storage.Writer.
	wc := u.Upload.Cl.Bucket(u.BucketConfig.BucketName()).Object(folder + "/" + hashedFilename + extension).NewWriter(ctx)
	if _, err := io.Copy(wc, blobFile); err != nil {
		return dao.File{}, err
	}
	if err := wc.Close(); err != nil {
		return dao.File{}, err
	}

	uploaded, err := u.FileRepo.AddFile(dao.File{
		Extension:        extension,
		OriginalFileName: f.Filename,
		FileCode:         folder + "/" + hashedFilename,
	})

	return uploaded, err
}

func (u *fileServiceParams) PredictImage(f multipart.FileHeader) (map[string]interface{}, error) {
	var result map[string]interface{}
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*5)
	defer cancel()
	blobFile, err := f.Open()
	if err != nil {
		return result, err
	}
	defer blobFile.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", f.Filename)
	if err != nil {
		return result, err
	}
	_, err = io.Copy(part, blobFile)
	if err != nil {
		return result, err
	}
	writer.Close()

	r, err := http.NewRequest("POST", "http://34.143.187.239", bytes.NewReader(body.Bytes()))
	if err != nil {
		return result, err
	}
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(r)
	fmt.Println(resp)
	if err != nil {
		return result, err
	}
	byte_blob, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(byte_blob, &result)
	return result, err
}

func md5Hash(source string) string {
	hash := md5.Sum([]byte(source))
	return hex.EncodeToString(hash[:])
}

func appendToFilename(filename, extra string) string {
	dotIdx := strings.LastIndex(filename, ".")
	if dotIdx == -1 {
		return filename + extra
	}
	extension := filepath.Ext(filename)
	trimmed := filename[0:dotIdx]
	trimmed += extra + extension

	return trimmed
}
