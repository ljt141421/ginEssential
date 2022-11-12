package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {

	initConfig()
	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080

}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
