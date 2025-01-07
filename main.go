package main

import (
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/TensoRaws/ffmpeg_for_BililiveRecorder/webhook"
	"github.com/sirupsen/logrus"
)

func main() {
	//设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
	//读取配置
	workdir, outputdir, ffmpegtype, ffmpegcom := Getconf()
	logrus.Debug("录播姬工作目录" + workdir)
	logrus.Debug("输出文件" + outputdir)
	logrus.Debug("配置类型" + strconv.Itoa(ffmpegtype))
	logrus.Debug("自定义命令" + ffmpegcom)
	//构建webhook
	wh := webhook.NewWebhookServer("test")
	wh.OnFileClose(func(event *webhook.EventFileClose) {
		logrus.Info("已接收待处理文件" + event.EventData.RelativePath)
		fullworkpath := workdir + event.EventData.RelativePath
		fulloutpath := outputdir + path.Base(fullworkpath)
		go ffmpegrun(fullworkpath, fulloutpath, ffmpegtype, ffmpegcom)
		logrus.Info("已接提交文件" + event.EventData.RelativePath)
	})

	webhook.StartServers(":8080")
}

func ffmpegrun(workpath string, outpath string, types int, comm string) {
	logrus.Info("开始处理" + workpath)
	//执行逻辑
	//自定义配置
	if types == 0 {
		//替换文件名
		outpath = strings.ReplaceAll(outpath, ".flv", ".mp4")
		comm = strings.ReplaceAll(comm, "{fullworkpath}", workpath)
		comm = strings.ReplaceAll(comm, "{fullwoutpath}", outpath)
		//提交执行
		logrus.Debug("执行命令" + comm)
		cmd := exec.Command("ffmpeg", strings.Split(comm, " ")...)
		err := cmd.Run()
		if err != nil {
			logrus.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}
	logrus.Info("完成处理" + workpath)
}
