package images_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils"
	"strconv"
	"strings"
	"time"
)

var (
	//// 黑名单，不允许的文件后缀
	//blacklist = []string{".exe", ".bat", ".sh"}

	// 白名单，允许的文件后缀
	whitelist = []string{".jpg", ".png"}
)

type FileUploadResponse struct {
	FileName string `json:"file_name"`
	Msg      string `json:"msg"`
}

func (ImagesApi) ImageUploadView(c *gin.Context) {
	// 这里的image名字要对应form-data中的key
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Multipart form error: %s", err.Error()))
		return
	}
	fileList, ok := form.File["image"]
	if !ok {
		res.FailWithMessage("找不到表单中上传文件的字段名 -> image", c)
		return
	}

	// 判断路径是否存在
	filePath := global.Config.Upload.Path
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			res.FailWithMessage("目录创建失败 -> ImageUpload", c)
			return
		} else {
			global.Log.Infof("%s 目录创建成功 -> ImageUpload", filePath)
		}
	}
	var resList []FileUploadResponse

	for _, file := range fileList {
		fileName := generateUniqueFileName(file.Filename)
		filePathWithName := path.Join(filePath, fileName)

		size := float64(file.Size) / float64(1024*1024)
		if size > global.Config.Upload.Size {
			resList = append(resList, FileUploadResponse{
				FileName: file.Filename,
				Msg:      "图片大小不能超过" + strconv.FormatFloat(global.Config.Upload.Size, 'f', -1, 64) + "M",
			})
			continue
		}

		if !isAllowedFile(file.Filename, whitelist) {
			resList = append(resList, FileUploadResponse{
				FileName: file.Filename,
				Msg:      "是不合法文件",
			})
			continue
		}

		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		bytes, err := io.ReadAll(fileObj)
		if err != nil {
			return
		}
		imageHash := utils.Md5(bytes)
		fmt.Println(imageHash)

		// 在数据库中查看图片是否存在
		var imageModel models.ImageModel
		if err = global.DB.Take(&imageModel, "`key` = ?", imageHash).Error; err == nil {
			// 找到了匹配的记录
			resList = append(resList, FileUploadResponse{
				FileName: file.Filename,
				Msg:      "该图片已存在",
			})
			continue
		} else {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 没有找到匹配的记录
				fmt.Println("图片不存在于数据库中")
			} else {
				// 其他数据库查询错误
				fmt.Println("数据库查询错误:", err)
				return
			}
		}
		// 保存通过HTTP请求上传的文件的函数
		// 通常，用于处理multipart/form-data类型的POST请求
		// 该请求类型通常用于文件上传。
		// 第二个参数可以直接写："./uploads/"+file.Filename
		err = c.SaveUploadedFile(file, filePathWithName)
		if err != nil {
			global.Log.Error(err)
			return
		}
		resList = append(resList, FileUploadResponse{
			FileName: file.Filename,
			Msg:      "上传成功",
		})
		global.DB.Create(&models.ImageModel{
			MODEL: models.MODEL{},
			Path:  path.Join(filePath, fileName),
			Key:   imageHash,
			Name:  fileName,
		})
	}
	res.OKWithData(resList, c)
	// 仅上传单张
	//fileList, err := c.FormFile("image")
	//if err != nil {
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}
}

func generateUniqueFileName(originalName string) string {
	// 使用时间戳和随机字符串生成唯一文件名
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomString := generateRandomString(8)
	// originalName包括后缀名和.
	if len(originalName) > 5 {
		originalName = originalName[len(originalName)-5:]
	}
	return fmt.Sprintf("%d_%s_%s", timestamp, randomString, originalName)
}

func generateRandomString(length int) string {
	// 生成指定长度的随机字符串
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func isAllowedFile(fileName string, allowedList []string) bool {
	// 获取文件后缀，并转换为小写
	fileExt := strings.ToLower(fileName[strings.LastIndex(fileName, "."):])
	//也可以使用split取最后一块
	return utils.InList(fileExt, allowedList)
}
