package main

import (
	"fmt"

	"github.com/terobox/go-s3"
)

func main() {
	fmt.Println("ULID:", s3.GenerateULID())
}
