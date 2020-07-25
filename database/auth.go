package database

import "app/logger"

// generate a new access token
func NewAccessToken(uid, code string) error {
	return checkAffect(db.Exec("call new_token(?, ?)", uid, code))
}

// generate a new access token
func CheckAccessToken(uid, code, oid string) (bool, error) {
	log.Debug("", uid, code, oid);
	return db.SQL("call check_token(?, ?, ?)", uid, code, oid).Exist()
}