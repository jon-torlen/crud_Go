package models

import (
	"gomysql/db"
)

type User struct {
	Id			int64
	Username	string
	Password	string
	Email		string
}

type Users []User

const UserSchema string = `CREATE TABLE users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(64) NOT NULL,
	email VARCHAR(50),
	create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

	//Construir usuario
	func NewUsuer(username, password, email string) *User {
		user := &User{Username: username, Password: password, Email: email}
		return user
	}

	//Cear usuario e inserta bd
	func CreateUser(username, password, email string) *User{
		user := NewUsuer(username, password, email)
		user.Save()
		return user
	}

	//Insertar Registro
	func (user *User) insert() {
		sql := "INSERT users SET username=?, password=?, email=?"
		result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
		user.Id, _ = result.LastInsertId()
	}

	//Listar todo el registro
	func ListUsers() Users {
		sql := "SELECT id, username, password, email FROM users"
		users := Users{}
		rows, _ := db.Query(sql)

		for rows.Next(){
			user := User{}
			rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
			users = append(users, user)
		}

		return users
	}

	//Obtener un registro
	func GetUser(id int) *User {
		user := NewUsuer("","","")

		sql := "SELECT id, username, password, email FROM users WHERE id=?"
		rows, _:= db.Query(sql, id)

		for rows.Next(){
			rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		}
		return user
	}

	//Actualizar registro
	func (user *User) update(){
		sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
		db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
	}

	//Guardar o editar registro
	func (user *User) Save(){
		if user.Id == 0 {
			user.insert()
		}else {
			user.update()
		}
	}
	//Eliminar un registro
	func (user *User) Delete() {
		sql := "DELETE FROM users WHERE id=?"
		db.Exec(sql, user.Id)
	}
