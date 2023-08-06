# im
im系统
### 引入gin
```go
$ go get -u github.com/gin-gonic/gin
```
### 初如化配置文件
```go
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败:%s", err.Error())
		return
	}
	fmt.Println("读取配置文件成功", viper.GetString("app_name"))
}

```
### 获取mysql读取文件
```go
func InitMySql() {
	var err error
	port := viper.GetInt("mysql.port")
	user := viper.GetString("mysql.user")
	host := viper.GetString("mysql.host")
	dbname := viper.GetString("mysql.dbname")
	password := viper.GetString("mysql.password")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)
	fmt.Println("dsn", dsn)

	//dsn := "root:701XTAY1993@tcp(127.0.0.1:3306)/im?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&model.User{})
}
```
### 集成swagger
【集成swagger】(https://github.com/swaggo/gin-swagger)
```go
package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"im/controller"
	"im/docs"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	v1 := r.Group("/api/v1")
	{
		v1.GET("/user", controller.GetUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
```
