package entities

import (
	// _playerCoinModel "github.com/Kotravel095/golang-backend-gameshop/pkg/playerCoin/"

	"time"
)

type (
	PlayerCoin struct {
		ID        uint64    `gorm:"primaryKey;autoIncrement;"`
		PlayerID  string    `gorm:"type:varchar(64);not null;"`
		Amount    int64     `gorm:"not null;"`
		CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	}
)

// func (p *PlayerCoin) ToPlayerCoinModel() *_playerCoinModel.PlayerCoin {
// 	return &_playerCoinModel.PlayerCoin{
// 		ID:        p.ID,
// 		PlayerID:  p.PlayerID,
// 		Amount:    p.Amount,
// 		CreatedAt: p.CreatedAt,
// 	}
// }