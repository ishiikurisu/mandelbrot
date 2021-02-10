package main

import (
    "fmt"
    "github.com/ishiikurisu/mandelbrot"
    "github.com/ishiikurisu/mandelbrot/processing"
    "sync"
    "time"
)

func main() {
    height := mandelbrot.NewFloat(900.0)
    width := mandelbrot.NewFloat(1600.0)
    targetX := mandelbrot.Atof("-1.62917")
    targetY := mandelbrot.Atof("-0.0203968")
    factor := mandelbrot.NewFloat(0.9)
    frameInitX, frameInitY, frameEndX, frameEndY := mandelbrot.FirstSetting(height, width)

    const maxItr int = 1024
    const noIterations int = 100
    const noSkip int = 0
    const totalGroups int = 3
    var groups int = 0
    var wg sync.WaitGroup

    for i := 0; i < noSkip; i++ {
        fmt.Printf("s%03d\n", i)
        frameInitX, frameInitY, frameEndX, frameEndY = mandelbrot.Follow(
            targetX, targetY, factor, frameInitX, frameInitY, frameEndX, frameEndY)
    }

    startTime := time.Now()
    groupStartTime := time.Now()

    for i := noSkip; i < noIterations; i++ {
        // allocating memory
        f0x := mandelbrot.NewFloat(0.0)
        f0y := mandelbrot.NewFloat(0.0)
        fx := mandelbrot.NewFloat(0.0)
        fy := mandelbrot.NewFloat(0.0)
        f0x.Copy(frameInitX)
        f0y.Copy(frameInitY)
        fx.Copy(frameEndX)
        fy.Copy(frameEndY)

        // saving image
        wg.Add(1)
        go func(i int, wg *sync.WaitGroup) {
            defer wg.Done()
            fmt.Printf("m%03d\n", i)
            fractal := mandelbrot.Mandelbrot(f0x, f0y, fx, fy, height, width, maxItr)
            img := processing.FractalToImage(fractal, maxItr)
            fn := fmt.Sprintf("./out/m%03d.png", i)
            processing.SaveImage(img, fn)
        }(i, &wg)
        groups++

        if groups == totalGroups {
            wg.Wait()
            groups = 0
            duration := time.Since(groupStartTime)
            fmt.Println(duration)
            groupStartTime = time.Now()
        }

        // zooming
        frameInitX, frameInitY, frameEndX, frameEndY = mandelbrot.Follow(
            targetX, targetY, factor, frameInitX, frameInitY, frameEndX, frameEndY)
    }

    if groups > 0 {
        wg.Wait()
    }

    totalTime := time.Since(startTime)
    fmt.Println(totalTime)
}
