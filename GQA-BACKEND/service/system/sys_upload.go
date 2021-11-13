package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/utils"
	"mime/multipart"
	"strconv"
	"time"
)

type ServiceUpload struct {
}

func (s *ServiceUpload) UploadAvatar(username string, avatar multipart.File, avatarHeader *multipart.FileHeader) (err error, avatarUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("avatarMaxSize")
	if maxSizeString == ""{
		return errors.New("找不到头像大小配置！"), ""
	}
	if !utils.CheckFileSize(avatar, maxSizeString){
		return errors.New("头像大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("avatarExt")
	if extListString == ""{
		return errors.New("找不到头像后缀配置！"), ""
	}
	if !utils.CheckFileExt(avatarHeader, extListString){
		return errors.New("头像后缀不被允许！"), ""
	}
	createPath := global.GqaConfig.Upload.AvatarSavePath + "/" + username
	createUrl := global.GqaConfig.Upload.AvatarUrl + "/" + username
	filename, err := utils.UploadFile(createPath, avatarHeader)
	if err != nil{
		return err, ""
	}
	avatarUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, avatarUrl
}

func (s *ServiceUpload) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (err error, fileUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("fileMaxSize")
	if maxSizeString == ""{
		return errors.New("没有找到文件大小配置！"), ""
	}
	if !utils.CheckFileSize(file, maxSizeString){
		return errors.New("文件大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("fileExt")
	if extListString == ""{
		return errors.New("没有找到文件后缀配置！"), ""
	}
	if !utils.CheckFileExt(fileHeader, extListString){
		return errors.New("文件后缀不被允许！"), ""
	}
	createPath := global.GqaConfig.Upload.FileSavePath + "/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String()
	createUrl := global.GqaConfig.Upload.FileUrl  + "/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String()
	filename, err := utils.UploadFile(createPath, fileHeader)
	if err != nil{
		return err, ""
	}
	fileUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, fileUrl
}
