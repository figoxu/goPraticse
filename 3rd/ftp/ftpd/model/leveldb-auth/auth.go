package ldbauth

import (
	"fmt"
	"strings"

	"github.com/figoxu/goPraticse/3rd/ftp/ftpd/model/web"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type LDBAuth struct {
	DB *leveldb.DB
}

func (db *LDBAuth) CheckPasswd(user, pass string) (bool, error) {
	p, err := db.GetUser(user)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return p == pass, nil
}

func (db *LDBAuth) GetUser(user string) (string, error) {
	v, err := db.DB.Get([]byte(fmt.Sprintf("pass:%s", user)), nil)
	if err != nil {
		return "", err
	}
	return string(v), err
}

func (db *LDBAuth) AddUser(user, pass string) error {
	return db.DB.Put([]byte(fmt.Sprintf("pass:%s", user)), []byte(pass), nil)
}

func (db *LDBAuth) DelUser(user string) error {
	return db.DB.Delete([]byte(fmt.Sprintf("pass:%s", user)), nil)
}

func (db *LDBAuth) ChgPass(user, pass string) error {
	return db.DB.Put([]byte(fmt.Sprintf("pass:%s", user)), []byte(pass), nil)
}

func (db *LDBAuth) UserList(users *[]web.User) error {
	it := db.DB.NewIterator(&util.Range{[]byte("pass:"), nil}, nil)
	defer it.Release()
	for it.Next() {
		key := string(it.Key())
		if !strings.HasPrefix(key, "pass:") {
			break
		}
		*users = append(*users, web.User{
			key[5:],
			string(it.Value()),
		})
	}
	return nil
}

func (db *LDBAuth) GroupList(groups *[]string) error {
	it := db.DB.NewIterator(&util.Range{[]byte("group:"), nil}, nil)
	defer it.Release()
	for it.Next() {
		key := string(it.Key())
		if !strings.HasPrefix(key, "group:") {
			break
		}
		*groups = append(*groups, key[6:])
	}
	return nil
}

func (db *LDBAuth) AddGroup(group string) error {
	return db.DB.Put([]byte(fmt.Sprintf("group:%s", group)), []byte(""), nil)
}

func (db *LDBAuth) DelGroup(group string) error {
	start := fmt.Sprintf("groupuser:%s:", group)
	it := db.DB.NewIterator(&util.Range{[]byte(start), nil}, nil)
	defer it.Release()
	keys := make([]string, 0)
	for it.Next() {
		key := string(it.Key())
		if !strings.HasPrefix(key, start) {
			break
		}
		keys = append(keys, key)
		user := strings.TrimLeft(key, start)
		err := db.DB.Delete([]byte(fmt.Sprintf("usergroup:%s:%s", user, group)), nil)
		if err != nil {
			return err
		}
	}
	for _, k := range keys {
		err := db.DB.Delete([]byte(k), nil)
		if err != nil {
			return err
		}
	}
	return db.DB.Delete([]byte(fmt.Sprintf("group:%s", group)), nil)
}

func (db *LDBAuth) AddUserGroup(user, group string) error {
	err := db.DB.Put([]byte(fmt.Sprintf("usergroup:%s:%s", user, group)), []byte(""), nil)
	if err != nil {
		return err
	}
	return db.DB.Put([]byte(fmt.Sprintf("groupuser:%s:%s", group, user)), []byte(""), nil)
}

func (db *LDBAuth) DelUserGroup(user, group string) error {
	err := db.DB.Delete([]byte(fmt.Sprintf("usergroup:%s:%s", user, group)), nil)
	if err != nil && err != leveldb.ErrNotFound {
		return err
	}
	err = db.DB.Delete([]byte(fmt.Sprintf("groupuser:%s:%s", group, user)), nil)
	if err != nil && err != leveldb.ErrNotFound {
		return err
	}
	return nil
}

func (db *LDBAuth) GroupUser(group string, users *[]string) error {
	prefix := fmt.Sprintf("groupuser:%s:", group)
	it := db.DB.NewIterator(&util.Range{[]byte(prefix), nil}, nil)
	defer it.Release()
	for it.Next() {
		key := string(it.Key())
		if !strings.HasPrefix(key, prefix) {
			break
		}
		*users = append(*users, key[len(prefix):])
	}
	return nil
}
