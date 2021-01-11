package main

import (
    "os"
    "github.com/ishiikurisu/mandelbrot"
    "github.com/ishiikurisu/mandelbrot/processing"
    "image/png"
)

func main() {
    var height float64 = 600.0
    var width float64 = 800.0
    frameInitX, frameInitY, frameEndX, frameEndY := mandelbrot.FirstSetting(height, width)
    fractal := mandelbrot.Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width)
    img := processing.FractalToImage(fractal)
    outputFile := "./out/mandelbrot.png"
    fp, _ := os.Create(outputFile)
    png.Encode(fp, img)
}
