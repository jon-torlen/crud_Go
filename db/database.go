package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(host:port)/database?charset=utf8
const url = "root:@tcp(localhost:3306)/goweb_db"

//Guarda la conexion
var db *sql.DB

//Realizar lac conexion
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection
}

//Cerrar la Conexion
func Close() {
	db.Close()
}

//Verificar la conexion
func Ping(){
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//Verifica si una tabla existe o no
func ExistsTable(tableName string) bool{
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil{
		fmt.Println("Error:", err)
	}
	return rows.Next()
}

//Reiniciar el registro de ua tabla
func TruncateTable(tableName string){
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

// Crear una tabla
func CreateTable(schema string, name string) {
	if !ExistsTable(name){
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
	
}

//Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	resust, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return resust, err
}

//Polimorfismo de de Query
func Query(query string, args ...interface{}) (*sql.Rows, error){
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}