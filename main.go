package main

import (
	"gocv.io/x/gocv"
)

func main() {
	window := gocv.NewWindow("Hello")
	master := gocv.IMRead("DSC_0098-Edit-1.jpg", gocv.IMReadColor)
	target := gocv.IMRead("DSC_0098-Edit-2.jpg", gocv.IMReadColor)
	img := gocv.NewMat()
	gocv.CvtColor(master, &master, gocv.ColorRGBToGray)
	gocv.CvtColor(target, &target, gocv.ColorRGBToGray)
	for {
		window.WaitKey(5)
		gocv.AbsDiff(master, target, &img)
		gocv.Threshold(img, &img, 75, 100, gocv.ThresholdBinary)
		window.IMShow(img)
	}
}
