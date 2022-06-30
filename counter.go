package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tfacedetect [videoFile] [classifier XML file]")
		return
	}

	// parse args
	videoFile := os.Args[1]
	xmlFile := os.Args[2]

	// open webcam
	video, err := gocv.OpenVideoCapture(videoFile)
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", videoFile)
		return
	}
	defer video.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	imgFG := gocv.NewMat()
	defer imgFG.Close()

	imgCleaned := gocv.NewMat()
	defer imgCleaned.Close()

	mog2 := gocv.NewBackgroundSubtractorMOG2()
	defer mog2.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}
	count := 0
	axis := "y"
	line := 530
	width := 10
	fmt.Printf("Start reading device: %v\n", videoFile)
	for {
		if ok := video.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", videoFile)
			return
		}
		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScaleWithParams(img, 1.1, 3, 0, image.Pt(100, 100), image.Pt(140, 140))
		//gocv.Rectangle(&img, quad, blue, 3)

		mog2.Apply(img, &imgFG)
		kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
		gocv.Erode(imgFG, &imgCleaned, kernel)
		kernel.Close()

		// calculate the image moment based on the cleaned frame
		moments := gocv.Moments(imgCleaned, true)
		area := moments["m00"]
		if area >= 1 {
			x := int(moments["m10"] / area)
			y := int(moments["m01"] / area)

			if axis == "y" {
				if x > 0 && x < img.Cols() && y > line && y < line+width {
					count++
				}
				gocv.Line(&img, image.Pt(100, line), image.Pt(1000, line), color.RGBA{255, 0, 0, 0}, 2)
			}
		}

		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)
		}

		gocv.PutText(&img, fmt.Sprintf("Qtd de Carros: %d", count), image.Pt(10, 20),
			gocv.FontHersheyPlain, 1.2, color.RGBA{0, 255, 0, 0}, 2)

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

}
