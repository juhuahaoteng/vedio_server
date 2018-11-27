package dbops

import "testing"

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", TestDeleteUser)
	t.Run("Reget", testRegetUser)

}

func testAddUser(t *testing.T) {
	err := AddUserCredential("chore", "123")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("chore")
	if pwd != "123" || err != nil {
		t.Errorf("%v", err)
	}
}
func TestDeleteUser(t *testing.T) {
	err := DeleteUser("chore", "123")
	if err != nil {
		t.Errorf("%v", err)
	}
}
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("chore")
	if pwd != "" && err != nil {
		t.Errorf("%v", err)
	}
}
