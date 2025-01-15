package ffmpegrunner

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func Ffmpegrun(workpath string, outpath string, types int, comms string, rmtime int) {
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
		comm = "-n 19 ffmpeg " + comm
		cmd = exec.Command("nice", strings.Split(comm, " ")...)
		logrus.Debug("执行命令" + "nice " + comm)
	} else {
		cmd = exec.Command("ffmpeg", strings.Split(comm, " ")...)
		logrus.Debug("执行命令" + "ffmpeg " + comm)
	}
	err := cmd.Run()
	if err != nil {
		logrus.Fatalf("cmd.Run() failed with %s\n", err)
	}

	logrus.Info("完成处理" + workpath)
	if rmtime >= 0 {
		time.Sleep(time.Duration(rmtime) * time.Hour)
		err := os.Remove(workpath)
		if err != nil {
			logrus.Fatal("删除" + workpath + "失败")
		}
		logrus.Info("删除" + workpath + "成功")
	}
}
