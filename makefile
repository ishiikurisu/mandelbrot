default: run

clean:
	rm out/*.png
	rm out/*.mp4

run:
	go run main/main.go

video: clean run
	ffmpeg -framerate 1 -pattern_type glob -i 'out/*.png' -c:v libx264 -r 30 -pix_fmt yuv420p out/mandelbrot.mp4
