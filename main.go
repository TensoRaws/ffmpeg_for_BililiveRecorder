package main

import "github.com/TensoRaws/ffmpeg_for_BililiveRecorder/webhook"

func main() {
	wh := webhook.NewWebhookServer("test")
	wh.OnRecordStart(func(event *webhook.EventRecordStart) {
		//执行代码
	})

	webhook.StartServers(":8080")
}
