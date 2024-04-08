package main

import (
	"GoDemo/src/main/mapper"
	"GoDemo/src/main/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	dsn := "qy:2278037785@Qy@tcp(rm-8vb1r34u74wguakfw2o.mysql.zhangbei.rds.aliyuncs.com:3306)/godemo?charset=utf8&parseTime=True&loc=Local"
	err := mapper.InitDatabase(dsn)
	if err != nil {
		panic("failed to connect to database")
	}
	router.InitRouter(engine)

	engine.Run("localhost:8080")
}
