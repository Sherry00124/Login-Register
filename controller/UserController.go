package controller

//http的处理函数
import (
	"Register-Login-Project/common"
	"Register-Login-Project/dto"
	"Register-Login-Project/model"
	"Register-Login-Project/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	//获取前端数据：方法3
	var requestUser = model.User{}
	c.Bind(&requestUser)
	//获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//log.Println("name:", name)
	//log.Println("telephone:", telephone)
	//log.Println("password:", password)
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号为11位")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号为11位"})
		//log.Println("telephone", telephone)
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码大于6")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码大于6"})
		return
	}
	if len(name) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "请输入名称")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入名称"})
		return
	}
	//log.Println(name, telephone, password)
	//判断手机号是否存在
	if isNameExist(DB, name) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户名已经存在")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名已经存在"})
		return
	}
	//创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "加密错误")
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	DB.Create(&newUser)
	//response.Success(c, nil, "注册成功")
	//c.JSON(200, gin.H{
	//	"message": "注册成功",
	//})

	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Println("token generate error: %v", err)
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "注册成功")

}
func Login(c *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	c.Bind(&requestUser)
	//获取参数
	name := requestUser.Name
	//telephone := requestUser.Telephone
	password := requestUser.Password
	log.Println("name:", name)
	//log.Println("telephone:", telephone)
	log.Println("password:", password)
	//数据验证
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码大于6")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码大于6"})
		return
	}
	if len(name) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码大于6")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入名称"})
		return
	}
	//判断用户名是否存在
	var user model.User
	DB.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断密码是否正确
	//如果验证失败就会有错误
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码错误")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码错误"})
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Println("token generate error: %v", err)
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": gin.H{"token": token},
	//	"msg":  "登录成功",
	//})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	//})
}
func isNameExist(db *gorm.DB, name string) bool {
	var user model.User
	db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
