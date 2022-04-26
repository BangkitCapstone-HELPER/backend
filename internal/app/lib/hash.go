package lib

import "golang.org/x/crypto/bcrypt"

type Hash struct {
	cost int
}

func NewHash() Hash {
	return Hash{cost: 14}
}

func (h *Hash) Hash(payload string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload), h.cost)
	return string(bytes), err
}

func (h *Hash) Compare(payload, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(payload))
	return err == nil
}
