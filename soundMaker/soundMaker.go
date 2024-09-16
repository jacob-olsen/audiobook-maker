package soundmaker

import (
	"fmt"
	"os/exec"
	"strings"
)

type SoundTaskInfo struct {
	Text       string
	OutputPath string
	FileName   string
}

func MakeAudio(data SoundTaskInfo) {
	comand := exec.Command("piper", "--model", "soundMaker/voice/en_US-danny-low.onnx", "--output_file", data.OutputPath+"/"+data.FileName+".wav")
	comand.Stdin = strings.NewReader(data.Text)
	err := comand.Run()
	if err != nil {
		fmt.Println(err)
	}
}
