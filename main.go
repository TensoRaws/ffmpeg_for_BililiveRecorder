package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/TensoRaws/ffmpeg_for_BililiveRecorder/webhook"
)

func main() {
	workdir, outputdir, ffmpegtype, ffmpegcom := Getconf()
	fmt.Println("录播姬工作目录" + workdir)
	fmt.Println("输出文件" + outputdir)
	fmt.Println("配置类型" + strconv.Itoa(ffmpegtype))
	fmt.Println("自定义命令" + ffmpegcom)
	wh := webhook.NewWebhookServer("test")
	wh.OnFileClose(func(event *webhook.EventFileClose) {
		fmt.Println("已接收待处理文件" + event.EventData.RelativePath)
		go ffmpegrun(event.EventData.RelativePath)
		fmt.Println("已接提交文件" + event.EventData.RelativePath)
	})

	webhook.StartServers(":8080")
}

func ffmpegrun(path string) {
	fmt.Println("开始处理" + path)
	//执行逻辑
	time.Sleep(time.Second * 60)
	fmt.Println("完成处理" + path)
}
