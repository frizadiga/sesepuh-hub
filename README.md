# sesepuh-gpt

In Go, instead of `npm install`, you use `go get` or the newer `go mod` system to manage dependencies. Here's how you can add OpenAI's Go client library:

```bash
# Add openai dependency to your project
go get github.com/sashabaranov/go-openai
go get -u 'github.com/openai/openai-go@v0.1.0-alpha.59'

# Update go.mod and go.sum
go mod tidy
```

Then in your openai.go file, you can import and use it:

```go
package main

import (
    "fmt"
    "github.com/sashabaranov/go-openai"
)

func ModOpenAI() {
    client := openai.NewClient("your-api-key-here")
    fmt.Println("-> openai")
    fmt.Println("vendor:", vendor)
}
```

Key differences from npm:
- Go uses go.mod file (similar to package.json)
- Dependencies are stored in $GOPATH/pkg/mod (not node_modules)
- No need for separate install command - `go run` or `go build` will download missing dependencies automatically
- `go mod tidy` cleans up unused dependencies (similar to npm prune)
