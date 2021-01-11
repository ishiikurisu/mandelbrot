package main

import (
    "os"
    "github.com/ishiikurisu/mandelbrot"
    "github.com/ishiikurisu/mandelbrot/processing"
    "image/png"
    "image"
)

func main() {
    var height float64 = 900.0
    var width float64 = 1600.0
    frameInitX, frameInitY, frameEndX, frameEndY := mandelbrot.FirstSetting(height, width)
    fractal := mandelbrot.Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width)
    img := processing.FractalToImage(fractal)
    saveImage(img, "./out/m1.png")

    frameInitX, frameInitY, frameEndX, frameEndY = mandelbrot.ZoomAt(height / 10, width / 10, width, height,frameInitX, frameInitY, frameEndX, frameEndY)
    fractal = mandelbrot.Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width)
    img = processing.FractalToImage(fractal)
    saveImage(img, "./out/m2.png")
}

func saveImage(img image.Image, fn string) {
    fp, _ := os.Create(fn)
    defer fp.Close()
    png.Encode(fp, img)
}
