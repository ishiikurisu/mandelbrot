package main

import (
    "fmt"
    "github.com/ishiikurisu/mandelbrot"
    "github.com/ishiikurisu/mandelbrot/processing"
)

func main() {
    height := mandelbrot.NewFloat(900.0)
    width := mandelbrot.NewFloat(1600.0)
    targetX := mandelbrot.Atof("-0.6567347481")
    targetY := mandelbrot.Atof("0.3416144335")
    factor := mandelbrot.NewFloat(0.9)
    var noIterations int = 1
    frameInitX, frameInitY, frameEndX, frameEndY := mandelbrot.FirstSetting(height, width)

    for i := 0; i < noIterations; i++ {
        // saving image
        fractal := mandelbrot.Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width)
        img := processing.FractalToImage(fractal)
        fn := fmt.Sprintf("./out/m%03d.png", i)
        processing.SaveImage(img, fn)

        // zooming
        frameInitX, frameInitY, frameEndX, frameEndY = mandelbrot.Follow(
            targetX, targetY, factor, frameInitX, frameInitY, frameEndX, frameEndY)
    }
}
