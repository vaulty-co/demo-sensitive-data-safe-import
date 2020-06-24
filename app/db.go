package main

import "sync"

type DB struct {
	users []user
	mux   sync.Mutex
}

func (d *DB) Save(u user) {
	d.mux.Lock()
	d.users = append(d.users, u)
	d.mux.Unlock()
}

func (d *DB) AddList(list []user) {
	d.mux.Lock()
	d.users = append(d.users, list...)
	d.mux.Unlock()
}

func (d *DB) List() []user {
	return d.users
}
