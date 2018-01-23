package core

import  "container/list"

type Answer struct {
	res Resource //资源

	taskQueue  list.List//任务队列
}

func (a *Answer) Res() Resource {
	return a.res
}

func (a *Answer) TaskQueue() list.List {
	return a.taskQueue
}

func (a *Answer) SetRes(res Resource) {
	a.res = res
}

func (a *Answer) AddTask(task Task){
	a.taskQueue.PushBack(task)
}
