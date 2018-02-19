echo "mandelbrot"
go build ./mandelbrot.go
time ./mandelbrot > mandelbrot.png

echo "mandelbrot_parallel"
go build ./mandelbrot_parallel.go
time ./mandelbrot_parallel > mandelbrot_parallel.png
