package ffmpegrunner

import "github.com/sirupsen/logrus"

func getdefconf(types int) string {
	def := map[int]string{
		10: "-i {fullworkpath} -c:v libsvtav1 -c:a copy -crf 40 -preset 6 {fulloutpath}",
		11: "-i {fullworkpath} -c:v libx265 -c:a copy -preset slow {fulloutpath}",
	}
	res, ok := def[types]
	if !ok {
		logrus.Fatal("配置序号不存在")
	}
	return res
}
