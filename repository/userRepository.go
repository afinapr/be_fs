package repository

import (
	"be_fs/driver"
	"be_fs/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type ResponseModel struct {
	Code    int    `json: "code" validate: "required"`
	Message string `json: "message" validate: "required"`
}

//createFunction
func CreateUser(U *models.User) *ResponseModel {
	Response := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return Response
	}

	//validationCharacter
	userName := len(U.UserName)
	if userName < 3 {
		Response = &ResponseModel{400, "Username min 3 Character"}
		return Response
	}

	password := len(U.Password)
	if password < 7 {
		Response = &ResponseModel{400, "Password min 7 Character"}
		return Response
	}

	name := len(U.Name)
	if name < 3 {
		Response = &ResponseModel{400, "Name min 3 Character"}
		return Response
	}

	//hashPassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(U.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO users (users.username, users.password, users.name) VALUES (?, ?, ?)`, U.UserName, hashedPassword, U.Name)
	if err != nil {
		fmt.Println(err.Error())
		Response = &ResponseModel{400, "Failed save Data"}
		return Response
	}

	defer db.Close()
	fmt.Println("Insert success!")
	Response = &ResponseModel{200, "Success save Data"}

	return Response
}

//readAllFunction
func ReadAllUserNoLimit() map[string]interface{} {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()
	var result = make(map[string]interface{})
	var data []models.User

	items, err := db.Query(`SELECT * FROM users`)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = models.User{}
		var err = items.Scan(&each.Id, &each.UserName, &each.Password, &each.Name)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		data = append(data, each)
	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	result["data"] = data
	return result
}

//readAllFunction
func ReadAllUser(limit int, offset int) []models.User {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()
	var result []models.User

	items, err := db.Query(`SELECT * FROM users LIMIT ? OFFSET ?`, limit, offset)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = models.User{}
		var err = items.Scan(&each.Id, &each.UserName, &each.Password, &each.Name)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)
	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

//readByIdFunction
func ReadByIdUser(userId int) []models.User {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()
	var result []models.User

	items, err := db.Query(`SELECT * FROM users where id = ?`, userId)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = models.User{}
		var err = items.Scan(&each.Id, &each.UserName, &each.Password, &each.Name)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)
	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

//updateFunction
func UpdateUser(U *models.User, userId int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	//validationCharacter
	userName := len(U.UserName)
	if userName < 3 {
		Res = &ResponseModel{400, "Username min 3 Character"}
		return Res
	}

	password := len(U.Password)
	if password < 7 {
		Res = &ResponseModel{400, "Password min 7 Character"}
		return Res
	}

	name := len(U.Name)
	if name < 3 {
		Res = &ResponseModel{400, "Name min 3 Character"}
		return Res
	}

	//hashPassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(U.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("UPDATE users SET username = ?, password = ?, name = ? WHERE id = ?", U.UserName, hashedPassword, U.Name, userId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}

	defer db.Close()
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

//deleteFunction
func DeleteUser(userId int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE id = ?", userId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed delete Data"}
		return Res
	}
	fmt.Println("Delete success!")
	Res = &ResponseModel{200, "Success delete Data"}
	return Res
}
