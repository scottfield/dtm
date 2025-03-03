/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/yedf/dtm/dtmcli/dtmimp"
)

var rdb *redis.Client
var once sync.Once

func RedisGet() *redis.Client {
	once.Do(func() {
		dtmimp.Logf("connecting to redis: %v", Config.Store)
		rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", Config.Store.Host, Config.Store.Port),
			Username: Config.Store.User,
			Password: Config.Store.Password,
		})
	})
	return rdb
}
