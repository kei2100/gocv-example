package main

import (
	"fmt"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"

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
	webcamimg := gocv.NewMat()
	grayscale := gocv.NewMat()
	binimg := gocv.NewMat()
	points := gocv.NewMat()
	qr := gocv.NewMat()
	defer webcamimg.Close()
	defer grayscale.Close()
	defer binimg.Close()
	defer points.Close()
	defer qr.Close()

	for {
		if ok := webcam.Read(&webcamimg); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if webcamimg.Empty() {
			continue
		}
		// show the image in the window, and wait 1 millisecond

		// >> use adaptive threshold
		//gocv.CvtColor(webcamimg, &grayscale, gocv.ColorRGBToGray)
		//gocv.AdaptiveThreshold(grayscale, &binimg, 255, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, 83, 2)
		// >> use threshold
		//gocv.CvtColor(webcamimg, &grayscale, gocv.ColorRGBToGray)
		//gocv.Threshold(grayscale, &binimg, 0, 255, gocv.ThresholdBinary|gocv.ThresholdOtsu)

		// >> show raw webcam image
		window.IMShow(webcamimg)
		// >> show processed image
		//window.IMShow(binimg)
		if window.WaitKey(1) >= 0 {
			break
		}

		// >> decode raw webcom image
		img, err := webcamimg.ToImage()
		// >> decode processed image
		//img, err := binimg.ToImage()
		if err != nil {
			fmt.Printf("to image: %v\n", err)
			return
		}
		bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
		qrReader := qrcode.NewQRCodeReader()
		result, err := qrReader.Decode(bmp, nil)
		if err != nil {
			continue
		}
		fmt.Println(result)
		break
	}
}
