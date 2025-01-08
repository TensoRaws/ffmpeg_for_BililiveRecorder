package ffmpegrunner

import (
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

func Ffmpegrun(workpath string, outpath string, types int, comm string) {
	logrus.Info("开始处理" + workpath)
	//执行逻辑
	//自定义配置
	if types == 0 {
		//替换文件名
		outpath = strings.ReplaceAll(outpath, ".flv", ".mp4")
		comm = strings.ReplaceAll(comm, "{fullworkpath}", workpath)
		comm = strings.ReplaceAll(comm, "{fulloutpath}", outpath)
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
