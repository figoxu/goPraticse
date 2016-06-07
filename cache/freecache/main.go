package main

import (
	"github.com/coocood/freecache"
	"github.com/quexer/utee"
	"log"
)

func main() {
	log.Println("hello")
	cache := NewDvExistCache(100, 3600)
	cache.Put("Hello", true)
	v := cache.Get("Hello")
	log.Println("Hello must be true @v:", v)
	v = cache.Get("NotExist")
	log.Println("NotExist must be false @v:", v)
	cache.Put("Word", false)
	v = cache.Get("Word")
	log.Println("Word must be false @v:", v)
}

type SimpleCache struct {
	cache *freecache.Cache
	ttl   int
}

func NewDvExistCache(mb, ttl int) *SimpleCache {
	lc := SimpleCache{
		cache: freecache.NewCache(mb * 1024 * 1024),
		ttl:   ttl,
	}
	return &lc
}

func (p *SimpleCache) Put(key string, val bool) {
	if val {
		p.cache.Set([]byte(key), []byte{byte(1)}, p.ttl)
	} else {
		p.cache.Set([]byte(key), []byte{byte(0)}, p.ttl)
	}
}

func (p *SimpleCache) Get(key string) bool {
	v, e := p.cache.Get([]byte(key))
	if freecache.ErrNotFound == e {
		return false
	}
	utee.Chk(e)
	return len(v) > 0 && v[0] == byte(1)
}
