package common

import "time"

// Timestamp 共通時刻カラム
type Timestamp struct {
	CreatedAT time.Time `db:"created_at"`
	UpdatedAT time.Time `db:"updated_at"`

	common.Timestamp
}

// NullTime 0001-01-01T00:00:00Z の場合に nil にする
func NullTime(nullTime time.Time) *time.Time {
	if nullTime.IsZero() {
		return nil
	}
	return &nullTime
}
