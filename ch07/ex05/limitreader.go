package main

import (
	"io"
)

type limitReader struct {
	r io.Reader
	n int64
}

// MEMO:
// A LimitedReader reads from R but limits the amount of data returned to just N bytes.
// Each call to Read updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.

func (lr *limitReader) Read(p []byte) (int, error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > lr.n {
		p = p[:lr.n]
	}

	n, err := lr.r.Read(p)
	lr.n -= int64(n)

	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}
