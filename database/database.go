package database

import (
	"database/sql"
	"time"

	"anshulgithub.com/anshul/usermangement/helper"
	"anshulgithub.com/anshul/usermangement/models"
	_ "github.com/go-sql-driver/mysql"
)

var db1 *sql.DB

func init() {

	db, err := sql.Open("mysql", "JayShreeRam:JayShreeRam#9557@tcp(localhost:3306)/userdatabase")
	db1 = db
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	helper.ErrCheck(err)
	db1.Ping()
	//fmt.Println(db1)
}

func GetAllUser() ([]models.User, error) {
	users, err := db1.Query("SELECT * FROM usertable")
	var user models.User
	slcUsers := []models.User{}
	for users.Next() {
		// for each row, scan the result into our user composite object
		err = users.Scan(&user.UserId, &user.Name, &user.Address, &user.GmailId, &user.PhoneNo)
		helper.ErrCheck(err)
		// print out the todo's Name attribute
		slcUsers = append(slcUsers, user)
		//log.Printf(user.Name)
	}
	return slcUsers, err
}

func CreateUser(user models.User) error {
	_, err := db1.Exec("INSERT INTO usertable (userid, name, address,emailid,number) VALUES (?, ?, ?,?,?)", user.UserId, user.Name, user.Address, user.GmailId, user.PhoneNo)
	return err
}

func ReadUser(id int) (*models.User, error) {
	var user models.User
	var empuser models.User

	err := db1.QueryRow("SELECT userid, name, address,emailid,number FROM usertable WHERE userid = ?", id).Scan(&user.UserId, &user.Name, &user.Address, &user.GmailId, &user.PhoneNo)
	if err != nil {
		return &empuser, err
	}
	return &user, err
}

func UpdateUser(user *models.User) error {
	_, err := db1.Exec("UPDATE usertable SET  name = ?, address = ?, emailid= ?,number= ? WHERE userid = ?", &user.Name, &user.Address, &user.GmailId, &user.PhoneNo, &user.UserId)
	return err
}

func DeleteUser(id int) error {
	_, err := db1.Exec("DELETE FROM usertable WHERE userid = ?", id)
	return err
}
