package entity

type Item struct {
	name     string
	category string
	price    uint64
	quantity uint32

	//relation for user
	user *User
}
