package pwdless

import (
	"time"

	"github.com/OmBiller/go-client-api/logging"
)

func (rs *Resource) choresTicker() {
	ticker := time.NewTicker(time.Hour * 1)
	go func() {
		for range ticker.C {
			if err := rs.Store.PurgeExpiredToken(); err != nil {
				logging.Logger.WithField("chore", "purgeExpiredToken").Error(err)
			}
		}
	}()
}
