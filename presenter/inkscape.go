package presenter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

func ConvertSVGToPNG(
	svg string,
	bgColor string,
	width int,
	height int,
) (bytes.Buffer, error) {
	var buf bytes.Buffer

	svgData := []byte(svg)

	// Create a pipe
	pr, pw := io.Pipe()

	var wg sync.WaitGroup
	wg.Add(1)
	// Start GraphicsMagick in a separate goroutine
	go func() {
		defer wg.Done()

		cmd := exec.Command(
			"inkscape",
			"--pipe",
			"--export-filename=-",
			"--export-type=png",
			"--export-background="+bgColor,
			"--export-width="+fmt.Sprintf("%d", width),
			"--export-height="+fmt.Sprintf("%d", height),
		)
		cmd.Stdin = pr
		cmd.Stdout = &buf
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// Write SVG data to the pipe
	_, err := pw.Write(svgData)
	if err != nil {
		fmt.Println(err)
	}

	// Close the pipe writer to signal end of data
	err = pw.Close()
	if err != nil {
		fmt.Println(err)
	}

	wg.Wait()
	return buf, nil
}
