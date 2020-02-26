package file

import (
	"bytes"
	"gopkg.in/russross/blackfriday.v2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

func IsExist(paths ...string) bool {
	if _, err := os.Stat(filepath.Join(paths...)); err == nil {
		return true
	} else {
		return false
	}
}

func Checkdir(dir string) error {
	var err error
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0700)
	}
	return err
}

// SaveUploadedFile uploads the form file to specific dst.
func SaveMulipart(file *multipart.FileHeader, dir, dst string) error {
	// check exist and create
	if err := Checkdir(dir); err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.OpenFile(dir+"/"+dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func SaveMarkdown2Html(file *multipart.FileHeader, dir, dst string) error {
	// check exist and create
	if err := Checkdir(dir); err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	srcByte, err := ioutil.ReadAll(src)

	data := bytes.Replace(srcByte, []byte("\r"), nil, -1)

	out, err := os.OpenFile(dir+"/"+dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(blackfriday.Run(data,
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.HardLineBreak),
	))
	if err != nil {
		return err
	}

	return nil
}
