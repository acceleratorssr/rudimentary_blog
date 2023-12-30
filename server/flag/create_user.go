package flag

import (
	"bufio"
	"fmt"
	"os"
	"server/global"
	"server/models"
	"server/models/stype"
	"server/utils/pwd"
)

func CreateUser(permission string) {
	scanner := bufio.NewScanner(os.Stdin)
	// 创建用户
	// 用户名 昵称 密码 确认密码
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		avatar     string
		phoneNum   string
		path       string
		err        error
		userModel  models.UserModels
	)

	path = "/uploads/image/1702897857727_zXcxYWSC_2.png"

	for {
		fmt.Printf("请输入用户名:")
		_, err = fmt.Scanf("%s\n", &userName)
		if userName == "quit" {
			global.Log.Info("退出创建用户")
			return
		}

		if err != nil {
			global.Log.Warn("用户名不可为空:", err)
			continue
		}

		err = global.DB.Take(&userModel, "username = ?", userName).Error
		if err == nil {
			global.Log.Warn("用户名已存在:", err)
			continue
		}
		break
	}

	for {
		fmt.Printf("请输入昵称:")
		_, err = fmt.Scanf("%s\n", &nickName)
		if nickName == "quit" {
			global.Log.Info("退出创建用户")
			return
		}

		if err != nil {
			global.Log.Warn("昵称不可为空:", err)
			continue
		}
		break
	}

	for {
		fmt.Printf("请输入密码:")
		_, err = fmt.Scanf("%s\n", &password)
		if password == "quit" {
			global.Log.Info("退出创建用户")
			return
		}
		if err != nil {
			global.Log.Warn("密码不可为空:", err)
			continue
		}
		if len(password) < 6 {
			global.Log.Warn("密码长度不应低于六位")
			continue
		}

		// regex包不支持?=之类的格式
		//regex := `^(.*[0-9].*[a-z]|[0-9].*[A-Z]|[a-z].*[A-Z])(?=.*[!@#$%^&*()-_+=])[0-9a-zA-Z!@#$%^&*()-_+=]{6,}$`
		//isSucceed, e := regexp.MatchString(regex, password)
		//if e != nil {
		//	global.Log.Error("regexp.MatchString:", e)
		//	continue
		//}
		//if !isSucceed {
		//	global.Log.Warn("密码必须包含大小写字母、数字、特殊字符中至少两种，长度不能少于6位")
		//}

		break
	}

	for {
		fmt.Printf("请再次输入密码:")
		_, err = fmt.Scanf("%s\n", &rePassword)
		if rePassword == "quit" {
			global.Log.Info("退出创建用户")
			return
		}
		if err != nil {
			global.Log.Warn("密码不可为空:")
			continue
		}
		if rePassword != password {
			global.Log.Warn("与上次输入密码不一致:")
			continue
		}
		break
	}

	{
		fmt.Printf("请选择头像:")
		// 这种写法表示 必须输入，如果不输入则报错：unexpected newline
		//_, err = fmt.Scanf("%s\n", &avatar)
		//if err != nil {
		//	global.Log.Error("头像错误:", err)
		//	return
		//}
		scanner.Scan()
		avatar = scanner.Text()
		if avatar == "quit" {
			global.Log.Info("退出创建用户")
			return
		}
		if avatar == "" {
			avatar = path
			fmt.Println("已选择默认头像:", path)
		}
	}

	for {
		fmt.Printf("请输入绑定的手机号码（可选）:")
		scanner.Scan()
		phoneNum = scanner.Text()
		if phoneNum == "quit" {
			global.Log.Info("退出创建用户")
			return
		}
		if len(phoneNum) != 11 && len(phoneNum) != 0 {
			global.Log.Error("手机号码位数异常")
			continue
		}
		break
	}

	fmt.Println("正在创建用户...")
	var permissionNum stype.Permission
	// 暂时只有两种用户
	if permission == "admin" {
		permissionNum = 1
	} else {
		permissionNum = 2
	}
	password = pwd.HashAndSalt(password)
	err = global.DB.Create(&models.UserModels{
		Username:       userName,
		NickName:       nickName,
		Password:       password,
		Avatar:         avatar,
		Token:          "",
		IP:             "",
		PhoneNum:       phoneNum,
		Permission:     permissionNum,
		SignStatus:     0,
		ArticleModels:  nil,
		CollectsModels: nil,
	}).Error
	if err != nil {
		global.Log.Error("创建用户失败:", err)
	}
	switch permissionNum {
	case 1:
		fmt.Println("创建管理员账号成功")
		global.Log.Infof("创建管理员账号 %s 成功", userName)
	case 2:
		fmt.Println("创建普通用户成功")
		global.Log.Infof("创建普通用户 %s 成功", userName)
	default:
		fmt.Println("创建用户异常？？")
	}

}
