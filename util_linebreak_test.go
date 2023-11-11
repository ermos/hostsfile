package hostsfile

import (
	"runtime"
	"testing"
)

func TestLinebreakWindows(t *testing.T) {
	if linebreakFromOS("windows") != "\r\n" {
		t.Errorf("linebreakFromOS returned incorrect linebreak")
	}
}

func TestLinebreakDefault(t *testing.T) {
	if linebreakFromOS("linux") != "\n" {
		t.Errorf("linebreakFromOS returned incorrect linebreak")
	}
}

func TestLinebreak(t *testing.T) {
	if linebreak() != linebreakFromOS(runtime.GOOS) {
		t.Errorf("linebreak returned incorrect linebreak")
	}
}
