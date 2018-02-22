go build -o mandelbrot mandelbrot_parallel.go

export GOMAXPROCS=1
echo "GOMAXPROCS=1"
time ./mandelbrot > 1.png

echo ""
export GOMAXPROCS=2
echo "GOMAXPROCS=2"
time ./mandelbrot > 2.png

echo ""
export GOMAXPROCS=4
echo "GOMAXPROCS=4"
time ./mandelbrot > 4.png

echo ""
export GOMAXPROCS=8
echo "GOMAXPROCS=8"
time ./mandelbrot > 8.png

echo ""
export GOMAXPROCS=16
echo "GOMAXPROCS=16"
time ./mandelbrot > 16.png
