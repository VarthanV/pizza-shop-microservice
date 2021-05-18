package handlers

import "errors"

func (u *UserLoginRequest) Validate() (err error) {
	if u.Email == "" {
		return errors.New("email must not be empty")
	}
	if u.Password == "" {
		return errors.New("password must not be empty")
	}
	return nil
}

func (u *UserSignupRequest) Validate() (err error) {
	if u.Email == "" {
		return errors.New("email must not be empty")
	}
	if u.Password == "" {
		return errors.New("password must not be empty")
	}
	if u.PhoneNumber == "" {
		return errors.New("phone_number must not be empty")
	}
	if u.Name == "" {
		return errors.New("name must not be empty")
	}

	return nil
}
