package db

const (
	insertImageQuery = `INSERT INTO people
				   (name, surname, patronymic, age, gender, nationality) 
				   VALUES($1, $2, $3, $4, $5, $6)`

	getPeopleSort = `SELECT id, name, surname, patronymic, age, gender, nationality 
					 FROM people 
					 WHERE age >= 20
					 ORDER BY age ASC`

	deleteImageByDateQuery = `DELETE FROM people WHERE "id" = $1`

	updateInfoPeople = `UPDATE people
						SET
							name = $2,
							surname = $3,
							patronymic = $4,
							age = $5,
							gender = $6,
							nationality = $7
						WHERE
							id = $1;`
)
