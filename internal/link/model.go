package link

import (
	"go/adv-demo/internal/stat"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.generateHash()
	return link
}

func (link *Link) generateHash() {
	link.Hash = RandStringRunes(10)
}

var lettersRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range n {
		b[i] = lettersRunes[rand.Intn(len(lettersRunes))]
	}
	return string(b)
}
