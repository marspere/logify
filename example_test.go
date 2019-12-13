package logify

func ExampleInfo() {
	EnableOutputToFile()
	SetMaxSaveDays(10)
	SetLogLocation("./temp")
	Info("hello gopher!")

	// Output: 2019-12-11 14:33:10 [INFO] [example_test.go:5] hello gopher!
}

func ExampleSetLogCallDepth() {
	SetLogCallDepth(3)
}