package file

import (
	"bytes"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
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

// parse template to html file
func ParseTmpl(writer io.Writer, data interface{}, file ...string) error {
	tmpl, err := template.ParseGlob("view/html/**/*")
	if err != nil {
		return err
	}

	if len(file) != 0 {
		tmpl, err = tmpl.ParseFiles(file...)
	}
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(writer, "index", data)
	if err != nil {
		return err
	}
	return nil
}

// SaveUploadedFile uploads the form file to specific dst.
func SaveFile(file *multipart.FileHeader, dir, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	srcByte, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	out, err := os.OpenFile(dir+"/"+dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.WriteString("{{define \"content\"}}\n")
	if err != nil {
		return err
	}

	_, err = out.Write(srcByte)

	_, err = out.WriteString("{{end}}")
	if err != nil {
		return err
	}

	return err
}

func SaveMarkdown2Tmpl(file *multipart.FileHeader, dir, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	srcByte, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	data := bytes.Replace(srcByte, []byte("\r"), nil, -1)

	out, err := os.OpenFile(dir+"/"+dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.WriteString("{{define \"content\"}}")
	if err != nil {
		return err
	}

	_, err = out.Write(blackfriday.Run(data,
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.HardLineBreak),
	))
	if err != nil {
		return err
	}

	_, err = out.WriteString("{{end}}")
	if err != nil {
		return err
	}

	return nil
}