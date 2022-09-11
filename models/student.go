package models

type Student struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetStudents() ([]Student, error) {
	var students []Student

	err := db.Find(&students).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}

func GetStudentById(id int) (Student, error) {
	var student Student

	err := db.First(&student, id).Error
	if err != nil {
		return student, err
	}

	return student, nil
}

func CreateStudent(student Student) (Student, error) {
	err := db.Create(&student).Error
	if err != nil {
		return student, err
	}

	return student, nil
}

func DeleteStudentById(id int) error {
	err := db.Delete(&Student{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
