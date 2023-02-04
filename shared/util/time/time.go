package time

import (
	"os"
	"sync"
	"time"
)

var (
	loc   *time.Location
	muLoc = &sync.Mutex{}
)

const (
	defaultTZ = "Asia/Jakarta"
	envKeyTZ  = "TZ"
)

// TZ return timezone
func TZ() string {
	if os.Getenv(envKeyTZ) == "" {
		return defaultTZ
	}
	return os.Getenv(envKeyTZ)
}

// Location return location used for time calculation
func Location() *time.Location {
	muLoc.Lock()
	defer muLoc.Unlock()

	if loc == nil {
		loc, _ = time.LoadLocation(TZ())
	}
	return loc
}

// Now return time.Now with spesific location
func Now() time.Time {
	return time.Now().In(Location())
}
