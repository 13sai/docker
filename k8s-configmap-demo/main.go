package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	conf = pflag.StringP("config", "c", "", "config filepath")
)


type Config struct {
	Name string
}

// 对外的初始化配置方法
func configRun(cfg string) error {
	c := Config{
			Name: cfg,
	}

	if err := c.init(); err != nil {
			return err
	}

	return nil
}

func (c *Config) init() error {
	if c.Name != "" {
			viper.SetConfigFile(c.Name)
	} else {
			// 默认配置文件是./config.yaml
			viper.AddConfigPath(".")
			viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	// viper解析配置文件
	err := viper.ReadInConfig() 
	if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}


	return nil
}

func main() {
	pflag.Parse()

	// 初始化配置
	if err := configRun(*conf); err != nil {
			panic(err)
	}


	gin.SetMode(viper.GetString("mode"))
	g := gin.New()
	g = LoadRoute(g)

	g.Run(viper.GetString("addr"))
}

func LoadRoute(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func (c *gin.Context)  {
			c.String(http.StatusNotFound, "404 not found");
	})

	g.GET("/", Index)

	return g
}

// 返回
type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// api返回结构
func ApiResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
			Code: code,
			Message: message,
			Data: data,
	})
}

func Index(c *gin.Context) {
	ApiResponse(c, 0, "success", viper.GetString("hi"))
}