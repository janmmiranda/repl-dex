package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Pokecache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

func NewPokecache(timeout time.Duration) Pokecache {
	cache := Pokecache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}

	go cache.reapLoop(timeout)

	return cache
}

func (pc *Pokecache) Get(key string) ([]byte, bool) {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	if key == "" {
		return []byte{}, false
	}

	val, ok := pc.cache[key]
	return val.val, ok
}

func (pc *Pokecache) Add(key string, val []byte) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (pc *Pokecache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		pc.reap(time.Now().UTC(), interval)
	}
}

func (pc *Pokecache) reap(now time.Time, last time.Duration) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	for k, v := range pc.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(pc.cache, k)
		}
	}
}
