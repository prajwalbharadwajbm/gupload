package validator

import "errors"

func IsValidUsername(username string) (bool, error) {
	if !(len(username) > 3 && len(username) < 30) {
		return false, errors.New("GUPLD103")
	}
	return true, nil

}

func IsValidPassword(username, password string) (bool, error) {
	if !(len(password) >= 8) {
		return false, errors.New("GUPLD104")
	}
	if password == username {
		return false, errors.New("GUPLD105")
	}
	return true, nil
}
