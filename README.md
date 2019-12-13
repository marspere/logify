# logify

Simple, structured, pluggable logging for Go.

Whether on Linux or Windows, nicely color-coded in development.

## Installation

If you don't have the Go development environment installed, visit the [Getting Started](http://golang.org/doc/install.html) 
document and follow the instructions. Once you're ready, execute the following command:

```bash
go get -u github.com/marspere/logify
```

### Quick Start

The simplest way to use Logify is simply the package-level exported logger:

```go
package main

import (
	log "github.com/marspere/logify"
)

func main() {
	log.Info("hello gopher!")
}
```

### Level logging

Logify has four logging levels: Debug, Info, Warn, Error.

```gotemplate
logify.Debug("Useful debugging information.")
logify.Info("Something noteworthy happened!")
logify.Warn("You should probably take a look at this.")
logify.Error("Something failed but I'm not quitting.")
```

### Print filename and line number

The logify output the called file name and file line number by default.
If your application encapsulates the logify package, you need to set ```SetLogCallDepth```.
The default value is 2, which is the level of direct calls. If you encapsulate
multiple layers, you need to adjust according to your needs.

```gotemplate
logify.SetLogCallDepth(3)
```

### Output log to file

The logify does not output to files by default. If you need output log, you 
need to set ```EnableOutputToFile```. And the default maximum number of days
saved is 7 days. Of course, you can use ```SetMaxSaveDays``` to reset. The log
file will be saved to the current directory by default, you can use ```SetLogLocation```
to set where the files are saved.

```gotemplate
EnableOutputToFile()
SetMaxSaveDays(10)
SetLogLocation("./temp")
```

### Issues

Feel free to push issues that could make Logify better: [https://github.com/marspere/logify/issues](https://github.com/marspere/logify/issues)