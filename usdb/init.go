package usdb

import (
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/kooksee/hashnet/cmn"
	"github.com/pingcap/tidb/store/tikv"
	"os"
	"time"
)

var Name = "txs"

var tdb *TikvStore

func Init() {
	tikv.MaxConnectionCount = 256

	url := os.Getenv("tikv_url")
	if url == "" {
		panic("please init tikv")
	}

	store, err := tikv.Driver{}.Open(fmt.Sprintf("tikv://%s/pd", url))
	cmn.MustNotErr("TikvStore Init Error", err)

	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Sprintf("init cache error: %s ", err.Error()))
	}

	tdb = &TikvStore{
		name:  []byte(Name),
		c:     store,
		cache: cache,
	}
}

func GetDb() *TikvStore {
	if tdb == nil {
		panic("please init usdb")
	}
	return tdb
}
