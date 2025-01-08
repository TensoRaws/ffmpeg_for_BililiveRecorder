package main

import (
	"flag"
	"path"
	"strconv"

	"github.com/TensoRaws/ffmpeg_for_BililiveRecorder/ffmpegrunner"
	"github.com/TensoRaws/ffmpeg_for_BililiveRecorder/webhook"
	"github.com/sirupsen/logrus"
)

func main() {
	//启动参数
	var ports int
	var debugon bool
	flag.IntVar(&ports, "port", 8080, "端口")
	flag.BoolVar(&debugon, "debug", false, "调试模式")
	flag.Parse()
	//设置日志级别
	if debugon {
		logrus.SetLevel(logrus.DebugLevel)
	}
	//读取配置
	workdir, outputdir, ffmpegtype, ffmpegcom := Getconf()
	logrus.Debug("录播姬工作目录" + workdir)
	logrus.Debug("输出文件" + outputdir)
	logrus.Debug("配置类型" + strconv.Itoa(ffmpegtype))
	logrus.Debug("自定义命令" + ffmpegcom)
	//构建webhook
	wh := webhook.NewWebhookServer("ffmpeg")
	wh.OnFileClose(func(event *webhook.EventFileClose) {
		logrus.Info("已接收待处理文件" + event.EventData.RelativePath)
		fullworkpath := workdir + event.EventData.RelativePath
		fulloutpath := outputdir + path.Base(fullworkpath)
		go ffmpegrunner.Ffmpegrun(fullworkpath, fulloutpath, ffmpegtype, ffmpegcom)
		logrus.Info("已接提交文件" + event.EventData.RelativePath)
	})

	webhook.StartServers(":" + strconv.Itoa(ports))
}
