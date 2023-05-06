package app

import "time"

func LocTime() time.Time {
	// loc, _ := time.LoadLocation(Timezone)
	// return time.Now().In(loc)
	return time.Now().Local()
}
