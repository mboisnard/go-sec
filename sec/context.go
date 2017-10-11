package sec

import (
	"database/sql"
)

// SecurityContext holding database connection and stuff.
type SecurityContext struct {
	connection     *sql.DB
	pwdConstraints *PasswordConstraints
	salt           string
}

// ContextInit is used to initialize the database connection and create the necessary
// tables, storing a pointer of the database connection in the returned context.
func ContextInit(dbPath string, constraints *PasswordConstraints, salt string) (*SecurityContext, error) {
	db, err := databaseOpen(dbPath)
	if err != nil {
		return nil, err
	}

	if err := databaseInit(db); err != nil {
		return nil, err
	}

	context := new(SecurityContext)
	context.connection = db
	context.salt = salt

	initPwdConstraints(context, constraints)

	return context, nil
}

func initPwdConstraints(context *SecurityContext, constraints *PasswordConstraints) {
	if constraints == nil {
		context.pwdConstraints = NewDefaultPasswordConstraints()
	} else {
		context.pwdConstraints = constraints
	}
}
