package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err := stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}
func GetUserCredential(loginName string) (string,error) {
	stmtout, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	stmtout.QueryRow(loginName).Scan(&pwd)
	stmtout.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FORM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}
	stmtDel.Exec(loginName,pwd)
	stmtDel.Close()
	return nil

}