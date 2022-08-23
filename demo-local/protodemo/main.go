package main

import (
	"fmt"
	"log"
	"protodemo/student"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("here start proto")
	test := &student.Student{
		Name:   "geektutu",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	newTest := &student.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}
	fmt.Println("here end proto")
}
