package main
import (
	"errors"
	"sync"
	"time"
)

// DBData 資料庫撈出的字串快取
type DBData struct {
	data    string
	version int
	dealine time.Time
}

// DBCacheMap 資料庫Map, Key: cache game id , Value: cache data
type DBCacheMap map[int]DBData

// DefaultDBCacheMap 預設的資料庫管理快取
var (
	DefaultDBCacheMap = make(DBCacheMap, 0)
	mapRWLocker       = new(sync.RWMutex)
)

// SetCache 設定快取資料
func (dc *DBCacheMap) SetCache(data string, gameID, version int) {

	mapRWLocker.Lock()
	DefaultDBCacheMap[gameID] = DBData{
		data:    data,
		version: version,
		dealine: time.Now().Add(1 * time.Minute),
	}
	mapRWLocker.Unlock()
}

// GetCache 取得快取資料
func (dc *DBCacheMap) GetCache(gameID int) (string, int, error) {
	mapRWLocker.Lock()
	defer mapRWLocker.Unlock()
	if dbData, exist := DefaultDBCacheMap[gameID]; exist {

		// 若目前時間已經在deadline之後，無效
		if time.Now().After(dbData.dealine) {
			return "", 0, errors.New("cache timeout ")
		}

		return DefaultDBCacheMap[gameID].data, DefaultDBCacheMap[gameID].version, nil
	}

	return "", 0, errors.New("no data")
}

