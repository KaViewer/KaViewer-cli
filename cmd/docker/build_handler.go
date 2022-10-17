package docker

import (
	"container/list"
	"fmt"
)

type buildHandler interface {
	support(buildType string) bool
	handle()
}

type localBuild struct {
}

func (lb localBuild) support(buildType string) (support bool) {
	support = buildType == "local"
	return
}

func (lb localBuild) handle() {
	fmt.Println("local handle")
}

type classicBuild struct {
}

func (cb classicBuild) support(buildType string) (support bool) {
	support = buildType == ""
	return
}

func (cb classicBuild) handle() {
	fmt.Println("classic handle")
}

var builders = list.New()

func Handle(buildType string) {
	if builders.Len() == 0 {
		builders.PushBack(&classicBuild{})
		builders.PushBack(&localBuild{})
	}

	for builder := builders.Front(); builder != nil; builder = builder.Next() {
		b := (builder.Value).(buildHandler)
		doBreak := doHandle(b, buildType)
		if doBreak {
			break
		}
	}
}

func doHandle(builder buildHandler, buildType string) (flag bool) {
	flag = false
	if builder.support(buildType) {
		builder.handle()
		flag = true
		return
	}
	return
}
