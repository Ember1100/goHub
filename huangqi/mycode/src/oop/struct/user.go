package model

type User struct {
	Name string
	Age  int32
}

// 定义成员方法
func (u User) Run() bool {
	// 实现User的Run方法
	// 这里可以编写User类型的具体运行逻辑
	return true
}

type Student struct {
	Name  string
	Age   int32
	Grade int32
}

// 定义成员方法
func (s Student) Run() bool {
	return true
}
func NewStudent(Name string, Age int32, Grade int32) *Student {
	return &Student{Name, Age, Grade}
}

// 指针方法
func (s *Student) SetName(name string) {
	s.Name = name
}
