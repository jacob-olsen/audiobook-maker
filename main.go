package main

import (
	soundmaker "audiobook-maker/soundMaker"
)

func main() {
	soundmaker.MakeAudio("this is a test and i hope it works", "./", "test sound")
}
