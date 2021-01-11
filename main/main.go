package main

import (
    "github.com/ishiikurisu/mandelbrot"
    "github.com/ishiikurisu/mandelbrot/processing"
)

func main() {
    var height float64 = 900.0
    var width float64 = 1600.0
    frameInitX, frameInitY, frameEndX, frameEndY := mandelbrot.FirstSetting(height, width)
    fractal := mandelbrot.Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width)
    img := processing.FractalToImage(fractal)
    processing.SaveImage(img, "./out/m1.png")

    frameInitX, frameInitY, frameEndX, frameEndY = mandelbrot.ZoomAt(width / 10, height / 2, width, height,frameInitX, frameInitY, frameEndX, frameEndY)
    fractal = mandelbrot.Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width)
    img = processing.FractalToImage(fractal)
    processing.SaveImage(img, "./out/m2.png")
}
