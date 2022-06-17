# go-habitify

[![Go Reference](https://pkg.go.dev/badge/github.com/kentakozuka/go-habitify.svg)](https://pkg.go.dev/github.com/kentakozuka/go-habitify)

A [Habitify API](https://docs.habitify.me/) client for Golang

## Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/kentakozuka/go-habitify"
)

func main() {
	cli := habitify.New("YOUR_API_KEY")
	habits, err := cli.ListHabits(context.Background())
	if err != nil {
		panic(err)
	}
	for _, habit := range habits {
		fmt.Println(habit.Name)
	}
}

```