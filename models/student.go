package models

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetStudents() []Student {
	return []Student{
		{
			Name: "Raj",
			Age:  20,
		},
		{
			Name: "Ossas",
			Age:  21,
		},
	}
}

func GetStudentById(id string) Student {
	return Student{
		Name: "Raj",
		Age:  20,
	}
}
