package main

import (
	"fmt"
	"os"
)

func tcheck(in_error error) {
	if in_error != nil {
		panic(in_error)
	}
}

func tdrop[T any](in_list []T, in_index int) []T {
	return append(in_list[:in_index], in_list[in_index+1:]...)
}

func tload(in_path string) string {
	file_data, file_error := os.ReadFile(in_path)
	tcheck(file_error)
	return string(file_data)
}

func tprint[T any](in_list []T) {
	fmt.Printf("%#v", in_list)
}

func tuse(in_list ...any) {
	for _, temp_var := range in_list {
		_ = temp_var
	}
}