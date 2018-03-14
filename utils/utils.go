package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"math/rand"
	"path/filepath"
	"reflect"
	"time"

	"github.com/angao/gin-xorm-admin/router/multitemplate"
)

// FormatDate 格式化时间
func FormatDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

// LoadTemplates 加载资源文件
func LoadTemplates(templatesDir string) multitemplate.Render {
	r := multitemplate.New()

	layouts, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}

	commons, err := filepath.Glob(templatesDir + "/common/*.html")
	if err != nil {
		panic(err.Error())
	}

	systems, err := filepath.Glob(templatesDir + "/system/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append([]string{layout}, systems...)
		files = append(files, commons...)
		r.Add(filepath.Base(layout), template.Must(template.ParseFiles(files...)))
	}
	return r
}

// Encrypt is encrypt the data with salt
func Encrypt(data string, salt string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(salt))
	if err != nil {
		return "", err
	}
	cipher := hash.Sum(nil)

	buf := new(bytes.Buffer)
	buf.Write(cipher)
	buf.WriteString(data)
	_, err = hash.Write(buf.Bytes())
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// StructToMap will change struct to map
func StructToMap(inter interface{}) map[string]interface{} {
	param := make(map[string]interface{})

	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	for i := 0; i < t.NumField(); i++ {
		param[t.Field(i).Name] = v.Field(i).Interface()
	}
	return param
}

// RandomString generate random string
func RandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// Contains slice contain sub
func Contains(slice []string, sub string) bool {
	for _, str := range slice {
		if str == sub {
			return true
		}
	}
	return false
}
