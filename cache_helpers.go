package cache

import "time"

const (
	defaultLifetime time.Duration = time.Second * 3
)

func getExpiration(lifetime time.Duration) time.Time {
	return time.Now().Add(lifetime)
}

func isExpired(item cacheItem) bool {
	return time.Now().After(item.whenExpired)
}
