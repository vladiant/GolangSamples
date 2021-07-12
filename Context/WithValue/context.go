package main

import (
	"context"
	"fmt"
)

// User defined structures should be used for key
type TestContextKey string

func RunWithContext(ctx context.Context, key TestContextKey) {
	if v := ctx.Value(key); v != nil {
		fmt.Println("Key", key, "found")
		return
	}

	fmt.Println("Key", key, "not found")
}

func Run() {
	testKey := TestContextKey("TestKey")
	testValue := "Value"
	ctx := context.WithValue(context.Background(), testKey, testValue)

	RunWithContext(ctx, testKey)

	fmt.Println("Context error: ", ctx.Err())
	fmt.Print("Context deadline: ")
	fmt.Println(ctx.Deadline())

	RunWithContext(ctx, TestContextKey("AnotherTestKey"))

	fmt.Println("Context error: ", ctx.Err())
	fmt.Print("Context deadline: ")
	fmt.Println(ctx.Deadline())

	fmt.Println("Done")
}

func main() {
	Run()
}
