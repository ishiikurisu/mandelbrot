package main

import (
    "fmt"
    "github.com/ishiikurisu/mandelbrot"
    "github.com/ishiikurisu/mandelbrot/processing"
)

func main() {
    var height float64 = 900.0
    var width float64 = 1600.0
    var noIterations int = 100
    frameInitX, frameInitY, frameEndX, frameEndY := mandelbrot.FirstSetting(height, width)

    for i := 0; i < noIterations; i++ {
        fractal := mandelbrot.Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width)
        img := processing.FractalToImage(fractal)
        fn := fmt.Sprintf("./out/m%03d.png", i)
        processing.SaveImage(img, fn)
        frameInitX, frameInitY, frameEndX, frameEndY = mandelbrot.ZoomAt(width / 10, height / 2, width, height,frameInitX, frameInitY, frameEndX, frameEndY)
    }
}
