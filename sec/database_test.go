/*
TODO :
- Create DB
- CRUD user / password
- Save clés publiques / privées
*/

package sec

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseOpenCanPing(t *testing.T) {
	db, err := databaseOpen("./test.db")

	assert.Nil(t, err)
	assert.NotNil(t, db)
	assert.Nil(t, db.Ping())
}

func TestDatabaseOpenShouldThrowIfPathIsEmpty(t *testing.T) {
	db, err := databaseOpen("")

	assert.Nil(t, db)
	assert.NotNil(t, err)
}

func TestDatabaseInitShouldReturnNil(t *testing.T) {
	db, _ := databaseOpen("./test.db")
	err := databaseInit(db)

	assert.Nil(t, err)
}

/*func TestDatabase(t *testing.T) {

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	stmt1, err := db.Prepare("CREATE TABLE IF NOT EXISTS userinfo (uid INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(64) NULL, departname VARCHAR(64) NULL, created DATE NULL)")
	checkErr(err)

	stmt1.Exec()

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "yolo", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close() //good habit to close

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}*/
