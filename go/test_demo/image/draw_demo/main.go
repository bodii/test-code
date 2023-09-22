package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	// demo()
	img := demo1()
	fmt.Println(img.Bounds())
}

func demo1() image.Image {
	file, err := os.Open("dst.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	fmt.Println(format)
	if err != nil {
		panic(err)
	}

	return img
}

func demo() {
	canvasWidth := 300
	spacingWidth := 5
	columnNum := 3

	// 图片列表
	images := []string{
		"https://cdn.pixabay.com/photo/2017/04/30/11/42/hijab-2272708_960_720.png",
		"https://cdn.pixabay.com/photo/2016/08/20/05/38/avatar-1606916_960_720.png",
		"https://cdn.pixabay.com/photo/2017/04/30/11/42/hijab-2272708_960_720.png",
		"https://cdn.pixabay.com/photo/2016/08/20/05/38/avatar-1606916_960_720.png",
		"https://cdn.pixabay.com/photo/2017/04/30/11/42/hijab-2272708_960_720.png",
		"https://cdn.pixabay.com/photo/2016/08/20/05/38/avatar-1606916_960_720.png",
		"https://cdn.pixabay.com/photo/2017/04/30/11/42/hijab-2272708_960_720.png",
		"https://cdn.pixabay.com/photo/2016/08/20/05/38/avatar-1606916_960_720.png",
		"https://cdn.pixabay.com/photo/2017/04/30/11/42/hijab-2272708_960_720.png",
	}

	mergeImage("dst.jpg", canvasWidth, spacingWidth, columnNum, image.White, images)
}

func mergeImage(dst string, canvasWidth, spacingWidth, columnNum int,
	background color.Color, images []string) {

	// 创建要保存的文件
	file, _ := os.Create(dst)
	defer file.Close()

	// 创建画布
	jpg := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasWidth))
	// 设置背景色为白色
	for x := 0; x < jpg.Bounds().Dx(); x++ {
		for y := 0; y < jpg.Bounds().Dy(); y++ {
			jpg.Set(x, y, image.White)
		}
	}

	// imagesCount := len(images)
	imgWidth := (canvasWidth - (spacingWidth * (columnNum + 1))) / columnNum
	for i, img := range images {
		imgFile := readInternetImage(img)
		m := resizeImage(uint(imgWidth), 0, imgFile)

		x := i % columnNum
		y := i / columnNum

		startX := x*imgWidth + (x+1)*spacingWidth

		startY := y*imgWidth + (y+1)*spacingWidth

		draw.Draw(jpg, jpg.Bounds(), m,
			m.Bounds().Min.Sub(image.Pt(startX, startY)), draw.Over)
	}

	jpeg.Encode(file, jpg, nil)
}

func resizeImage(width, height uint, img image.Image) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
}

func readInternetImage(src string) image.Image {
	rsp, err := http.Get(src)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	// file, err := os.Open(src)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// img, format, err := image.Decode(rsp.Body)
	// fmt.Println(format)
	// if err != nil {
	// 	panic(err)
	// }

	imageType, ok := rsp.Header["Content-Type"]
	if !ok {
		panic("file type is not exists")
	}
	str := strings.Split(imageType[0], "/")
	format := str[1]

	var img image.Image
	if format == "png" {
		img, err = png.Decode(rsp.Body)
	} else if format == "jpeg" {
		img, err = jpeg.Decode(rsp.Body)
	} else if format == "gif" {
		img, err = gif.Decode(rsp.Body)
	} else {
		panic(errors.New("not image format"))
	}
	if err != nil {
		panic(err)
	}

	return img
}
