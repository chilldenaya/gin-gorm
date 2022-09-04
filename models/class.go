package models

type Class struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func GetClasses() []Class {
	return []Class{
		{
			Name: "Calculus 1",
			Code: "CA100",
		},
		{
			Name: "Linear Algebra",
			Code: "CA101",
		},
	}
}

func GetClassById(id string) Class {
	return Class{
		Name: "Calculus 1",
		Code: "CA100",
	}
}
