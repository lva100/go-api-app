package stat

import (
	"go/adv-demo/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentTime := datatypes.Date(time.Now())
	repo.Db.Find(&stat, "link_id=? and date=?", linkId, currentTime)
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentTime,
		})
	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}
