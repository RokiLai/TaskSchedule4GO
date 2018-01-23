package core

import "container/list"

type Task struct {
	Id             int    //任务id
	Name           string //任务名称
	Time           int64  //执行时间
	Description    string //描述
	Status         int    //状态
	FirstStartTime int64  //最早开始时间
	FinishTime     int64  //完成时间

	FatherTasks list.List//父任务
	ChildTasks  list.List //子任务
	//res        []Resource //资源集
}

//func NewTask() Task{
//	task := Task{}
//	return task
//}

