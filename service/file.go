package service

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type FileService interface {
	DownloadAndCompress(url string, fileName string) error
	GetFilePath(fileName string) (string, error)
}

type fileService struct {
}

func NewFileService() FileService {
	return &fileService{}
}

func (f *fileService) DownloadAndCompress(url string, fileName string) error {
	log.Print("In service - function DownloadAndCompress")
	zipFileName := "sample.zip"

	res, resErr := http.Get(url)
	if resErr != nil {
		return resErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	fileHeader := &zip.FileHeader{
		Name:     fileName,
		Modified: time.Now(),
		Method:   8,
	}

	writer, zipWriteErr := zipWriter.CreateHeader(fileHeader)
	if zipWriteErr != nil {
		return zipWriteErr
	}

	if _, err := writer.Write(body); err != nil {
		return err
	}

	closeErr := zipWriter.Close()
	if closeErr != nil {
		return closeErr
	}

	file, createErr := os.Create(zipFileName)
	if createErr != nil {
		return createErr
	}

	if _, cpyErr := io.Copy(file, buf); cpyErr != nil {
		return cpyErr
	}

	syncErr := file.Sync()
	if syncErr != nil {
		return syncErr
	}

	closeErr = file.Close()
	if closeErr != nil {
		return closeErr
	}

	log.Print("Out service - function DownloadAndCompress")
	return nil
}

func (f *fileService) GetFilePath(fileName string) (string, error) {
	log.Print("In service - function GetFilePath")

	cmd := exec.Command("pwd")
	path, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	log.Print("Out service - function GetFilePath")
	return string(path) + fileName, nil
}
