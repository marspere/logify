package logify

func ExampleInfo() {
	EnableOutputToFile()
	Info("hello gopher!")

	// Output: 2019-12-11 14:33:10 [INFO] [example_test.go:5] hello gopher!
}
