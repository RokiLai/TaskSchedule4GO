package core

type Task struct {
	id          int
	name        string
	time        int64
	description string
	status		int
	

	fatherTask  []Task
	childTask	[]Task
	res 		[]Resource

}
