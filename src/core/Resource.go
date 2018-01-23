package core

type Resource struct {
	id          int    //资源id
	name        string //资源名称
	finalTime   int64  //资源完成时间
	description string //资源描述
}

func (r *Resource) Id() int {
	return r.id
}

func (r *Resource) Name() string {
	return r.name
}

func (r *Resource) FinalTIme() int64 {
	return r.finalTime
}

func (r *Resource) Description() string {
	return r.description
}

func (r *Resource) SetId(id int) {
	r.id = id
}

func (r *Resource) SetName(name string) {
	r.name = name
}

func (r *Resource) SetFinalTime(finalTime int64) {
	r.finalTime = finalTime
}

func (r *Resource) SetDescription(description string) {
	r.description = description
}
