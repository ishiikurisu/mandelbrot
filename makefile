default: run

clean:
	rm out/*.png
	rm out/*.mp4

build:
	go build -o main.exe main/main.go

run: build
	./main.exe

video: run
	ffmpeg -framerate 12 -pattern_type glob -i 'out/*.png' -c:v libx264 -r 30 -pix_fmt yuv420p out/mandelbrot.mp4
