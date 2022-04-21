/**
    @Author: qiyou_wu
    @CreateDate: 2022/4/21
    @Description:
**/
package main

import (
	"fmt"
	"unsafe"
)

type MaleUser struct {
	Name string
	Age  int
	Sex  bool
}

func (user *MaleUser) SayHello() {
	fmt.Println("Hello!Im male")
}

type FemaleUser struct {
	Name string
	Age  int
	Sex  bool
}

func (user *FemaleUser) SayHello() {
	fmt.Println("Hello!Im female")
}

func main() {

	malePoint := &MaleUser{
		Name: "Jack",
		Age:  22,
		Sex:  true,
	}

	FemalePoint := &FemaleUser{
		Name: "Rose",
		Age:  22,
		Sex:  false,
	}

	point := unsafe.Pointer(malePoint)

}
