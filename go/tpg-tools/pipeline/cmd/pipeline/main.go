package main

import "github.com/rwx-yxu/lab/go/tpg-tools/pipeline"

func main() {
	pipeline.FromFile("testdata/hello.txt").Column(1).Stdout()
}
