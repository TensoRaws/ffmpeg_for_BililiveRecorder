package ffmpegrunner

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func Ffmpegrun(workpath string, outpath string, types int, comms string) {
	logrus.Info("开始处理" + workpath)
	//执行逻辑
	//自定义配置
	var comm string
	if types == 0 {
		comm = comms
	} else {
		comm = getdefconf(types)
	}
	//替换文件名
	outpath = strings.ReplaceAll(outpath, ".flv", ".mp4")
	comm = strings.ReplaceAll(comm, "{fullworkpath}", workpath)
	comm = strings.ReplaceAll(comm, "{fulloutpath}", outpath)
	//提交执行
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		comm += "-n 19 ffmpeg"
		cmd = exec.Command("nice", strings.Split(comm, " ")...)
	} else {
		cmd = exec.Command("ffmpeg", strings.Split(comm, " ")...)
	}
	logrus.Debug("执行命令" + comm)
	err := cmd.Run()
	if err != nil {
		logrus.Fatalf("cmd.Run() failed with %s\n", err)
	}

	logrus.Info("完成处理" + workpath)
}
