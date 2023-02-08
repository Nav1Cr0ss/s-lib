package user

func contains[T comparable](s []T, el T) bool {
	for _, v := range s {
		if v == el {
			return true
		}
	}

	return false
}

type User struct {
	Id          string `json:"id"`
	IsAdmin     bool   `json:"is_admin"`
	Permissions []string
}

func (u User) CheckPermission(handler string) bool {
	if u.IsAdmin {
		return true
	}
	if contains(u.Permissions, handler) {
		return true
	}

	return false
}
