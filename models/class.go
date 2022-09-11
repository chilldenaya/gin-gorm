package models

type Class struct {
	ID   uint `gorm:"primarykey"`
	Name string
	Code string
}

func GetClasses() ([]Class, error) {
	var classes []Class

	err := db.Find(&classes).Error
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func GetClassById(id int) (Class, error) {
	var class Class

	err := db.First(&class, id).Error
	if err != nil {
		return class, err
	}

	return class, nil
}

func CreateClass(class Class) (Class, error) {
	err := db.Create(&class).Error
	if err != nil {
		return class, err
	}

	return class, nil
}

func DeleteClassById(id int) error {
	err := db.Delete(&Class{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
