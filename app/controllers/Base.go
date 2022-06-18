package controllers

import (
	"github.com/firdausalif/challenge-todolist/app/models"
	"github.com/jellydator/ttlcache/v2"
	"time"
)

var cacheNotFound = ttlcache.ErrNotFound
var cache ttlcache.SimpleCache = ttlcache.NewCache()

type WorkRequest struct {
	Service any
	Delay   time.Duration
}

type Result struct {
	Resp any
	Err  error
}

var respsChan = make(chan []*models.Todo)
var respChan = make(chan *models.Todo)
