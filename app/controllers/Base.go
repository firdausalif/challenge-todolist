package controllers

import "github.com/jellydator/ttlcache/v2"

var cacheNotFound = ttlcache.ErrNotFound
var cache ttlcache.SimpleCache = ttlcache.NewCache()
