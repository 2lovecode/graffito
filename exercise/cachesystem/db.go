package main

type Db struct {
	data map[string]string
	count int
}

func (db *Db) Select(key string) string {
	value, ok := db.data[key]

	db.count++

	if ok {
		return value
	}

	return ""
}


func (db *Db) Insert(key string, value string) {
	db.data[key] = value
}

func (db *Db) Count() int {
	return db.count
}

func (db *Db) SetCount(count int) {
	db.count = count
}

func NewDb(data map[string]string) *Db {
	return &Db{data, 0}
}