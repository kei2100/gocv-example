package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	deviceID := 0
	// open webcam
	webcam, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("QR Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	points := gocv.NewMat()
	qr := gocv.NewMat()
	defer img.Close()
	defer points.Close()
	defer qr.Close()

	// prepare QR detector
	detector := gocv.NewQRCodeDetector()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}

		val := detector.DetectAndDecode(img, &points, &qr)
		if val == "" {
			continue
		}
		fmt.Println(val)
		break
	}
}
