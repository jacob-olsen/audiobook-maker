package soundmaker

import (
	"fmt"
	"os/exec"
	"strings"
)

func MakeAudio(text string, outputPath string, fileName string) {
	comand := exec.Command("piper", "--model", "soundMaker/voice/en_US-danny-low.onnx", "--output_file", outputPath+"/"+fileName+".wav")
	comand.Stdin = strings.NewReader(text)
	err := comand.Run()
	if err != nil {
		fmt.Println(err)
	}
}
