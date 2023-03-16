package getcurrenttime

import (
	"time"

	"github.com/beevik/ntp"
)

// func GetCurrentTime() (currentTime time.Time, err error) {
// 	return ntp.Time("0.beevik-ntp.pool.ntp.org")
// }

func GetCurrentTime(timeServer string) (currentTime time.Time, err error) {
	return ntp.Time(timeServer)
}
