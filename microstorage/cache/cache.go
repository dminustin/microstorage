package cache

import (
	"encoding/json"
	"errors"
	"microstorage"
	"microstorage/consumers"
	"time"
)

type SCache struct {
	Get    func(key string) (CachedObject, error)
	put    func(key string, data CachedObject) //Have to use the PutCache channel
	Exists func(key string) bool
	Delete func(key string)
}

var Cache = SCache{}

type CachedObjectChan struct {
	Uri    string
	Object CachedObject
}

var PutCache = make(chan CachedObjectChan, 100)

type CachedObject struct {
	Data consumers.InputData
	File []byte
}

func CacheRoutine() {
	for {
		select {
		case x := <-PutCache:
			Cache.put(x.Uri, x.Object)
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func transformFromJson(data interface{}) []byte {
	r, _ := json.MarshalIndent(data, "", " ")
	return r
}

func transformToJson(data []byte) (result CachedObject) {
	//todo handle error
	json.Unmarshal(data, &result)
	return result
}

func DefaultCacheSetup(config microstorage.TCache, cache *SCache) {
	cache.Get = func(key string) (CachedObject, error) {
		return CachedObject{}, errors.New("Key not exists")
	}
	cache.put = func(key string, data CachedObject) {

	}
	cache.Exists = func(key string) bool {
		return false
	}

	cache.Delete = func(key string) {

	}
}

func Setup(config microstorage.TCache) {
	switch config.Engine {
	case "redis":
		RedisCacheSetup(config, &Cache)
		break
	default:
		DefaultCacheSetup(config, &Cache)
	}
}
