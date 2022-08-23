package main

import (
	"fmt"
	"runtime"
	"sync"
)

// RegisterService 注册服务
type RegisterService struct {
	session *string
	resp    int
}

// NewRegisterService 初始化注册服务
func NewRegisterService(session *string) *RegisterService {
	return &RegisterService{
		session: session,
		resp:    0,
	}
}

// Register 账号注册
func (s *RegisterService) Register(abc, efg string) string {
	// ...
	return "Register"
}

// FastRegisterRegister 快速账号注册
func (s *RegisterService) FastRegister(abc, efg string) string {
	// ...
	return "FastRegister"
}

type RegisterService2 struct {
}

// Register 账号注册
func (s *RegisterService2) Register(session *string, resp int, abc, efg string) string {
	// ...
	return "Register"
}

// FastRegisterRegister 快速账号注册
func (s *RegisterService2) FastRegister(session *string, resp int, abc, efg string) string {
	// ...
	return "FastRegister"
}

var instance *RegisterService3
var once sync.Once

type RegisterService3 struct {
}

// NewRegisterService 初始化注册服务
func NewRegisterService3() *RegisterService3 {
	once.Do(func() {
		instance = &RegisterService3{}
	})
	return instance
}

// Register 账号注册
func (s *RegisterService3) Register(abc, efg string) string {
	// ...
	return "Register"
}

// FastRegisterRegister 快速账号注册
func (s *RegisterService3) FastRegister(abc, efg string) string {
	// ...
	return "FastRegister"
}

func main() {
	var m0 runtime.MemStats
	runtime.ReadMemStats(&m0)
	fmt.Printf("mem0 %+v\n", m0)
	//
	abcd := "abcd"
	for i := 0; i < 10000000; i++ {
		obj := NewRegisterService(&abcd)
		obj.Register("abcd", "abcd")
		//fmt.Printf(res + "\n")
		obj.FastRegister("abcd", "abcd")
		//fmt.Printf(res2 + "\n")
	}
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	fmt.Printf("mem1 %+v\n", m1)
	obj2 := RegisterService2{}
	for i := 0; i < 10000000; i++ {
		obj2.Register(&abcd, i, "abcd", "abcd")
		//fmt.Printf(res + "\n")
		obj2.FastRegister(&abcd, i, "abcd", "abcd")
		//fmt.Printf(res2 + "\n")
	}
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	fmt.Printf("mem2 %+v\n", m2)

	obj3 := RegisterService3{}
	for i := 0; i < 10000000; i++ {
		obj3.Register("abcd", "abcd")
		//fmt.Printf(res + "\n")
		obj3.FastRegister("abcd", "abcd")
		//fmt.Printf(res2 + "\n")
	}
	var m3 runtime.MemStats
	runtime.ReadMemStats(&m3)
	fmt.Printf("mem3 %+v\n", m3)

	fmt.Printf("mem1-mem0 Alloc; %+v, mem1-mem0 TotalAlloc; %+v, mem1-mem0 Mallocs; %+v \n", m1.Alloc-m0.Alloc, m1.TotalAlloc-m0.TotalAlloc, m1.Mallocs-m0.Mallocs)

	fmt.Printf("mem2-mem1 Alloc; %+v, mem2-mem1 TotalAlloc; %+v, mem2-mem1 Mallocs; %+v \n", m2.Alloc-m1.Alloc, m2.TotalAlloc-m1.TotalAlloc, m2.Mallocs-m1.Mallocs)

	fmt.Printf("mem3-mem1 Alloc; %+v, mem2-mem1 TotalAlloc; %+v, mem2-mem1 Mallocs; %+v \n", m3.Alloc-m2.Alloc, m3.TotalAlloc-m2.TotalAlloc, m3.Mallocs-m2.Mallocs)
}
