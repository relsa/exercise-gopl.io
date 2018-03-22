go build -o mandelbrot gopl.io/ch3/mandelbrot
go build -o convert main.go

./mandelbrot > mandelbrot.png
./convert -format jpeg < mandelbrot.png > mandelbrot.jpeg
./convert -format gif < mandelbrot.png > mandelbrot.gif
