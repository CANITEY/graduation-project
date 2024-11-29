package database

import (
	"database/sql"
)

func (d *dbConnector) AddToken(userid uint, publicKey string) error {
	stmt := `SELECT id from employees where id=$1`
	_, err := d.db.Exec(stmt, userid)
	if err != nil {
		return err
	}

	stmt = `INSERT INTO securitytokens(user_id, publickey) values($1, $2)`
	_, err = d.db.Exec(stmt, userid, publicKey)
	if err != nil {
		return err
	}
	return nil
}

func (d *dbConnector) GetToken(userid uint) (string, error) {
	stmt := `SELECT publickey from securitytokens where user_id=$1`
	var publicKey string

	row := d.db.QueryRow(stmt, userid)
	if err := row.Scan(&publicKey); err != nil {
		return "", nil
	}
	
	return publicKey, nil
}

func AddChallenge(d *sql.DB, userid uint, challUUID string, challBody []byte) error {
	stmt := `UPDATE securitytokens SET
	chal_token=$1,
	challenge=$2
	WHERE user_id=$3`

	_, err := d.Exec(stmt, challUUID, challBody, userid)
	return err
}

func GetChallengeStatus(d *sql.DB, challUUID string) (bool, error) {
	var solved bool
	stmt := `SELECT solved from securitytokens where chal_token=$1`
	row := d.QueryRow(stmt, challUUID)
	if err := row.Scan(&solved); err != nil {
		return false, err
	}

	return solved, nil
}

func DeleteChallenge(d *sql.DB, challUUID string) (error) {
	stmt := `UPDATE securitytokens SET
	chal_token='',
	challenge='',
	solved=false
	WHERE chal_token=$1`
	_, err := d.Exec(stmt, challUUID)
	if err != nil {
		return err
	}

	return nil
}
