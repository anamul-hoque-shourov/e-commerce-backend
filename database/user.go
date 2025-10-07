package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

var userList []User

func (user User) Store() User {
	if user.ID != 0 {
		return user
	}
	user.ID = len(userList) + 1
	userList = append(userList, user)
	return user
}

func Find(email, password string) *User {
	for _, user := range userList {
		if user.Email == email && user.Password == password {
			return &user
		}
	}
	return nil
}
