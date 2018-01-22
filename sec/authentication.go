/*
TODO :
- CRUD user / pass
- Validate user / pass
- Check password strength
- Avoid brute forcing
*/

package sec

import (
	_ "github.com/mattn/go-sqlite3"
)

// RegisterUser is the main function to add new user in database
func RegisterUser(context *SecurityContext, username string, password string) (int64, error) {

	newUserStatement, err := context.connection.Prepare("INSERT INTO users(username, password) values(?, ?)")
	if err != nil {
		return -1, err
	}

	res, err := newUserStatement.Exec(username, password)
	if err != nil {
		return -1, err
	}

	newID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return newID, nil
}

// Login check username and password in database
func Login(context *SecurityContext, username string, challengeResponse string) {

}

func GetChallenge() {

}
