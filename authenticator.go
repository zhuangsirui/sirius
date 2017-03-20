package main

import "sync"

type Authenticator struct {
	sync.RWMutex

	// Users simply mapping username to password
	Users map[string]string
}

func (a *Authenticator) SetUsers(users []User) {
	a.Lock()
	defer a.Unlock()
	a.Users = make(map[string]string)
	for _, user := range users {
		a.Users[user.Username] = user.Password
	}
}

func (a *Authenticator) Auth(username, password string) bool {
	a.RLock()
	defer a.RUnlock()
	return a.Users[username] == password
}
