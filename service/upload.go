package service

import (
	"io/ioutil"
	"mginmall/conf"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (fileName string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.Config.Path.AvataPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg" //todo 提取file后缀
	connect, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, connect, 0666)
	if err != nil {
		return
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	return err == nil
}
