package main

import (
	"context"
	"fmt"
	"strings"
)

// TL;DR - don't export keys nor the type that's being used for context
//
// using custom type as underlying 'any' type this ensures that no one
// else has access to or inadvertently sets a context outside of your control
type ctxKey string

const (
	favoriteColor ctxKey = "favorite-color"
)

func main() {
	//Exploring contexts
	// using string as a key
	ctx := context.Background()
	ctx = context.WithValue(ctx, "favorite-color", "blue")
	value := ctx.Value("favorite-color")
	fmt.Println("Value from 'favorite-color':", value)

	// value will now return nil
	// "favoriteColor" is NOT the same type as the "favorite-key" above
	ctx = context.WithValue(ctx, "favorite-color", "blue")
	value = ctx.Value(favoriteColor)
	fmt.Println("Value from ctxKey:", value)

	// value will now return the correct value
	// ctx was set with "favoriteColor" and value pulls it out
	ctx = context.WithValue(ctx, favoriteColor, "blue")
	value = ctx.Value(favoriteColor)
	fmt.Println("Value from correct ctxKey:", value)

	// value1 and value2 ensures that even if other libraries
	// inadvertently set "favorite-color", the custom context
	// "favoriteColor" is under your/the dev's control
	ctx = context.WithValue(ctx, "favorite-color", 123)
	value1 := ctx.Value("favorite-color")
	value2 := ctx.Value(favoriteColor)
	fmt.Printf("---\nValue1 (int): %d\nValue2 (str): %s\n", value1, value2)

	// Using type assertions
	strValue, ok := value2.(string)
	if !ok {
		fmt.Println("not cool bro")
	}
	fmt.Printf("---\nAsserted value: %s\n", strValue)
	fmt.Println("Having prefix value (str)", strings.HasPrefix(strValue, "b"))

	// type assert to int - will fail
	fmt.Println("\n--- typing to int and checking 'ok'")
	value = ctx.Value(favoriteColor)
	intValue, ok := value.(int)
	if !ok {
		fmt.Println("it's not an int, bro")
		return
	}
	fmt.Printf("---\nAsserted value: %q\n", intValue)
	//fmt.Println("Having prefix value (str)", strings.HasPrefix(strValue, "b"))
}
