package models

import (
	"github.com/qgweb/gossdb"
)

type SDBType struct {
	gossdb.Value
}

type SDB struct {
}

func (this SDB) Set(db string, key string, value interface{}) error {
	client, err := dbpool.NewClient()
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Hset(db, key, value)
}

func (this SDB) Incr(db string, key string, num int64) (int64, error) {
	client, err := dbpool.NewClient()
	if err != nil {
		return 0, err
	}
	defer client.Close()
	return client.Hincr(db, key, num)
}

func (this SDB) Size(db string) (int64, error) {
	client, err := dbpool.NewClient()
	if err != nil {
		return 0, err
	}
	defer client.Close()
	return client.Hsize(db)
}

func (this SDB) Get(db string, key string) (SDBType, error) {
	client, err := dbpool.NewClient()
	if err != nil {
		return SDBType{}, err
	}
	defer client.Close()
	v, err := client.Hget(db, key)
	return SDBType{v}, err
}

func (this SDB) Del(db string, key string) error {
	client, err := dbpool.NewClient()
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Hdel(db, key)
}

func (this SDB) DBDel(db string) error {
	client, err := dbpool.NewClient()
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Hclear(db)
}

func (this SDB) Scan(db string, bkey string, ekey string, limit int64) (map[string]SDBType, error) {
	client, err := dbpool.NewClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	res := make(map[string]SDBType)
	v, err := client.Hscan(db, bkey, ekey, limit)
	if v != nil {
		for k, vv := range v {
			res[k] = SDBType{vv}
		}
	}
	return res, err
}
