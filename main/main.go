package main

import (
    "fmt"
    "github.com/ishiikurisu/mandelbrot"
    "github.com/ishiikurisu/mandelbrot/processing"
    "sync"
    "time"
    "encoding/json"
    "os"
)

type Configuration struct {
    Height string
    Width string
    TargetX string
    TargetY string
    Factor string
    MaxItr int
    NumberOfIterations int
    NumberOfSkips int
    TotalGroups int
}

func main() {
    config, oops := loadConfig("./config.json")
    if oops != nil {
        panic(oops)
    }

    height := mandelbrot.Atof(config.Height)
    width := mandelbrot.Atof(config.Width)
    targetX := mandelbrot.Atof(config.TargetX)
    targetY := mandelbrot.Atof(config.TargetY)
    factor := mandelbrot.Atof(config.Factor)
    maxItr := config.MaxItr
    noIterations := config.NumberOfIterations
    noSkip := config.NumberOfSkips
    totalGroups := config.TotalGroups

    frameInitX, frameInitY, frameEndX, frameEndY := mandelbrot.FirstSetting(height, width)
    groups := 0

    for i := 0; i < noSkip; i++ {
        fmt.Printf("s%03d\n", i)
        frameInitX, frameInitY, frameEndX, frameEndY = mandelbrot.Follow(
            targetX, targetY, factor, frameInitX, frameInitY, frameEndX, frameEndY)
    }

    var wg sync.WaitGroup
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

func loadConfig(path string) (Configuration, error) {
    file, _ := os.Open(path)
    defer file.Close()
    config := Configuration{}
    decoder := json.NewDecoder(file)
    err := decoder.Decode(&config)
    return config, err
}
