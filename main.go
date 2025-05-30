package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"
	"strings"
	"terminal-display/display"
	"terminal-display/utils"
	"time"
)

//go:embed resources/font.ttf
var fontData []byte

var screen *display.Display
var (
	cl         chan any
	port       string
	test       bool
	brightness uint8 = maxBrightness
	ttyNum     uint
)

const (
	maxBrightness = 20
	minBrightness = 2
	timeout       = 10 * time.Second
)

func init() {
	flag.StringVar(&port, "tty", "/dev/ttyACM0", "display tty port")
	flag.BoolVar(&test, "test", false, "Test the display")
	flag.UintVar(&ttyNum, "n", 1, "tty to render")
	flag.Parse()
	cl = make(chan any)
	if ttyNum > 7 || ttyNum < 1 {
		log.Fatal("Target tty must be between 1 and 7")
	}
	var err error
	screen, err = display.NewDisplay(cl, port, 480, 320, fontData, test)
	for err != nil {
		time.Sleep(time.Second)
		screen, err = display.NewDisplay(cl, port, 480, 320, fontData, test)
	}
	if test {
		screen.Demo()
		os.Exit(0)
	}
}

func main() {
	chst := make(chan string)
	go func(chan string) {
		var data string
		for {
			data = <-chst
			screen.Fill(0, 0, 0)
			data = strings.Replace(data, "\t", "  ", -1)
			screen.WriteTextChunked(data, color.White, 0, 0, 16, 0, 0, 78)
			screen.Update()
		}
	}(chst)
	var last string
	tc := time.NewTicker(time.Second / 2)
	tm := time.NewTimer(timeout)
	for {
		select {
		case <-tc.C:
			cmd := exec.Command("tail", fmt.Sprintf("/dev/vcs%d", ttyNum))
			output, err := cmd.Output()
			utils.Check(err)
			data := string(output)
			screen.SetBrightness(brightness)
			if data == last {
				continue
			}
			tm.Reset(timeout)
			brightness = maxBrightness
			last = data
			chst <- data
		case <-tm.C:
			brightness = minBrightness
		}
	}
}
