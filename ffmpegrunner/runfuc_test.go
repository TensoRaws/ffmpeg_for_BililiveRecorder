package ffmpegrunner

import "testing"

func TestGetdefconf(t *testing.T) {
	if getdefconf(10) != "-i {fullworkpath} -c:v libsvtav1 -c:a copy -crf 40 -preset 6 {fulloutpath}" {
		t.Error(10)
	}
	if getdefconf(11) != "-i {fullworkpath} -c:v libx265 -c:a copy -preset slow {fulloutpath}" {
		t.Error(11)
	}
}
