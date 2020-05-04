package conf

import (
	"authen/eth"
	"authen/model"
	"os"

	log "github.com/alecthomas/log4go"
	"github.com/joho/godotenv"
)

func initDB() {
	log.Info("init DB")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWD")
	name := os.Getenv("DB_NAME")
	connStr := user + ":" + passwd + "@tcp(" + host + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	model.Init(connStr)
}

// Init 初始化
func Init() {
	// 读取.env环境变量
	godotenv.Load()
	initDB()
	eth.Init(os.Getenv("SWARM_HOST"))
}
