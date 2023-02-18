package main

type Dog struct {
	tableName struct{} `pg:"dogs"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	Breed     bool     `json:"breed" pg:"breed"`
	Color     string   `json:"color" pg:"color"`
}

// FindAllDogs Получить список собак.
func FindAllDogs() []Dog {
	var dogs []Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dogs).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dogs
}

// CreateDog Создать собаку.
func CreateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()
	_, err := pgConnect.Model(&dog).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

// FindDogById Получить собачек по id.
func FindDogById(id string) Dog {
	var dog Dog

	pgConnect := PostgresConnect()
	err := pgConnect.Model(&dog).
		Where("id = ?", id).
		First()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog

}

// DeleteDogById Удалить собачку по id.
func DeleteDogById(id string) Dog {
	var dog Dog

	pgConnect := PostgresConnect()
	_, err := pgConnect.Model(&dog).
		Where("id = ?", id).
		Delete()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog

}

func UpdateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	oldDog := FindDogById(dog.ID)

	oldDog.Name = dog.Name
	oldDog.Breed = dog.Breed
	oldDog.Color = dog.Color

	_, err := pgConnect.Model(&oldDog).
		Set("name = ?", oldDog.Name).
		Set("breed = ?", oldDog.Breed).
		Set("color = ?", oldDog.Color).
		Where("id = ?", oldDog.ID).
		Update()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}
