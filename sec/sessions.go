/*
TODO :
- Create token
- Invalid old token
- Check token
*/

package sec

import (
	"fmt"
	"time"
)

// CreateNewSession create
// TODO Create convertToTimestamp method
func CreateNewSession(context *SecurityContext, userUID int64) (string, error) {

	actualTime := time.Now()
	actualTimestamp := actualTime.UnixNano() / 1000000
	expiredTimestamp := actualTime.AddDate(0, 0, 7).UnixNano() / 1000000
	message := fmt.Sprintf("%d%d", userUID, actualTimestamp)
	token := Sha512(message)

	insertTokenStatement, err := context.connection.Prepare("INSERT INTO tokens values(?,?,?)")
	if err != nil {
		return "", err
	}

	insertTokenStatement.Exec(userUID, token, expiredTimestamp)

	return token, nil
}
