package usercontroller

import (
	"gin-backend-api/confics"
	"gin-backend-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var users []models.User

	// confics.DB.Order("id DESC").Find(&users)

	//SQL
	confics.DB.Raw("SELECT * FROM users ORDER BY id DESC").Scan(&users)

	c.JSON(200, gin.H{
		"data": users,
	})
}

func Register(c *gin.Context) {
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	//เช็คอีเมลซ้ำ
	userExist := confics.DB.Where("email = ?", input.Email).First(&user)
	if userExist.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "มีผู้ใช้งานอีเมลนี้ในระบบแล้ว"})
		return
	}

	result := confics.DB.Debug().Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "สมัครสมาชิกสำเร็จ",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Login",
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := confics.DB.First(&user, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลนี้"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func SearchByFullname(c *gin.Context) {
	fullname := c.Query("fullname")
	c.JSON(200, gin.H{
		"data": fullname,
	})
}
