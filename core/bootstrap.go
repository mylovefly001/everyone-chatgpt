package core

import (
	"context"
	"everyone-chatgpt/core/base"
	"everyone-chatgpt/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kataras/iris/v12"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
	"time"
)

type IBootstrap interface {
	Start()
}

type bootstrap struct {
}

func Bootstrap() IBootstrap {
	global.Context = context.Background()
	global.RootPath, _ = os.Getwd()

	//mode=web时的参数
	flag.StringVar(&global.RunEnv, "env", "local", "默认：local")
	flag.IntVar(&global.RunPort, "port", 8080, "默认：8080")
	flag.Parse()

	//创建上传文件夹1
	_ = os.Mkdir(filepath.Join(global.RootPath, global.UploadDir), os.ModePerm)

	//设置日志的格式
	rl, _ := rotatelogs.New(filepath.Join(global.RootPath, global.LoggerDir, "%Y%m%d_info.log"))
	logrus.SetOutput(io.MultiWriter(os.Stdout, rl))
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: global.TimeFormat,
	})

	//绑定配置文件
	cfgPath := filepath.Join(global.RootPath, "config", fmt.Sprintf("%s.yaml", global.RunEnv))
	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Fatalf("读取配置文件：%s 失败", cfgPath)
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		logrus.WithError(err).Fatal("加载配置文件失败")
	}
	//监控配置文件变动
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Info("配置文件发生变动，重新加载")
		if err := viper.Unmarshal(&global.Config); err != nil {
			logrus.WithError(err).Fatal("重新加载配置文件失败")
		}
	})

	return &bootstrap{}
}

func (b bootstrap) Start() {
	b.startHttp()
}

func (b bootstrap) startHttp() {
	logrus.Info("Http服务启动中...")
	app := iris.New()
	app.Use(iris.Compression)
	app.RegisterView(iris.HTML(filepath.Join(global.RootPath, "app", "view"), ".html").Reload(true))
	app.HandleDir("/static", filepath.Join(global.RootPath, "static"))
	closed := make(chan bool, 1)
	//里面不可加东西，否则会导致无法优雅关机
	iris.RegisterOnInterrupt(func() {
		ctx, cancel := context.WithTimeout(global.Context, time.Duration(10)*time.Second)
		defer cancel()
		if err := app.Shutdown(ctx); err != nil {
			logrus.WithError(err).Fatal("Http关闭失败")
		}
		close(closed)
	})
	base.Router(app)
	if err := app.Listen(
		fmt.Sprintf(":%d", global.RunPort),
		iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed),
	); err != nil {
		logrus.WithError(err).Fatal("Http服务启动失败")
	}
	<-closed
}
