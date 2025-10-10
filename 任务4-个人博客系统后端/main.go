package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

var DB *gorm.DB

func init() {
	username := "root"
	password := "123456"
	host := "localhost"
	port := "3306"
	dabase := "testdb"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dabase)

	mysqlLogger := logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})

	if err != nil {
		fmt.Println("mysql connect fial, error: ", err)
	}

	DB = db
}

func AuthorizationReq(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		c.JSON(200, gin.H{
			"message": "没有认证令牌1",
		})
		c.Abort()
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(LOGIN_SECRET_KEY), nil
	})

	if err != nil {
		c.AbortWithStatus(401)
	}

	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
	}

	fmt.Println("登录请求解析：", claims)
	c.Set("user", claims)
	c.Set("userid", uint(claims["id"].(float64)))
	c.Set("username", claims["username"])
	c.Next()

}

func main() {
	// 创建表
	DB.AutoMigrate(&User{}, &Post{}, &Comments{})

	router := gin.Default()

	// 注册
	router.POST("/register", Register)

	// 登录
	router.POST("/login", Login)

	auth := router.Group("/api", AuthorizationReq)
	{
		auth.POST("/test", func(c *gin.Context) {
			user, exist := c.Get("user")
			if exist {
				fmt.Println(user)
			}

			c.JSON(200, gin.H{
				"message": "a success",
			})
		})

		// 文章创建
		auth.POST("/postCreate", PostCreate)

		// 文章更新
		auth.POST("/postUpdate", PostUpdate)

		// 文章删除
		auth.DELETE("/postDelete/:id", PostDelete)

		// 评论创建
		auth.POST("/commentsCreate", commentsCreate)
	}

	// 文章读取
	router.GET("/api/postQuery/:id", PostQuery)

	// 评论读取
	router.GET("/api/commentsQuery/:id", CommentsQuery)

	// 端口指定
	router.Run(":8080")
}
