package main

import (
	"context"
	"fmt"
)

type contextKey string

var requestIdKey = contextKey("id")
var userKey = contextKey("user")

// выполняет функцию fn с учетом контекста ctx
func execute(ctx context.Context, fn func() int) int {
	reqId := ctx.Value(requestIdKey)
	if reqId != nil {
		fmt.Printf("Request ID = %d\n", reqId)
	} else {
		fmt.Println("Request ID unknown")
	}

	user := ctx.Value(userKey)
	if user != nil {
		fmt.Printf("Request user = %s\n", user)
	} else {
		fmt.Println("Request user unknown")
	}
	return fn()
}

func main() {
	work := func() int {
		return 42
	}

	// контекст с идентификатором запроса
	ctx := context.WithValue(context.Background(), requestIdKey, 1234)
	// и пользователем
	ctx = context.WithValue(ctx, userKey, "admin")
	res := execute(ctx, work)
	fmt.Println(res)

	// пустой контекст
	ctx = context.Background()
	res = execute(ctx, work)
	fmt.Println(res)
}
