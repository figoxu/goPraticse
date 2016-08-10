package ldbperm

import (
	"fmt"
	"os"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

type LDBPerm struct {
	db           *leveldb.DB
	defaultUser  string
	defaultGroup string
	defaultMode  os.FileMode
}

func NewLDBPerm(db *leveldb.DB, user, group string, mode os.FileMode) *LDBPerm {
	return &LDBPerm{db, user, group, mode}
}

func (db *LDBPerm) GetOwner(rPath string) (string, error) {
	v, err := db.db.Get([]byte(fmt.Sprintf("p:owner:%s", rPath)), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return db.defaultUser, nil
		}
		return "", err
	}
	return string(v), nil
}

func (db *LDBPerm) GetGroup(rPath string) (string, error) {
	v, err := db.db.Get([]byte(fmt.Sprintf("p:group:%s", rPath)), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return db.defaultGroup, nil
		}
		return "", err
	}
	return string(v), nil
}

func (db *LDBPerm) GetMode(rPath string) (os.FileMode, error) {
	v, err := db.db.Get([]byte(fmt.Sprintf("p:mode:%s", rPath)), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return db.defaultMode, nil
		}
		return 0, err
	}
	mode, err := strconv.Atoi(string(v))
	if err != nil {
		return 0, err
	}
	return os.FileMode(mode), nil
}

func (db *LDBPerm) ChOwner(rPath, owner string) error {
	return db.db.Put([]byte(fmt.Sprintf("p:owner:%s", rPath)),
		[]byte(owner), nil)
}

func (db *LDBPerm) ChGroup(rPath, group string) error {
	return db.db.Put([]byte(fmt.Sprintf("p:group:%s", rPath)),
		[]byte(group), nil)
}

func (db *LDBPerm) ChMode(rPath string, mode os.FileMode) error {
	return db.db.Put([]byte(fmt.Sprintf("p:mode:%s", rPath)),
		[]byte(strconv.Itoa(int(mode))), nil)
}
