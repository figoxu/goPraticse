package main

import (
	"testing"
	"github.com/icrowley/fake"
	"math/rand"
	"time"
	"github.com/quexer/utee"
)

func TestBitUserDao_Save(t *testing.T) {
	bitUserDao := NewBitUserDao(pg_plsql)
	for i := 0; i < 10000; i++ {
		buser := BitUser{
			Name: fake.FullName(),
			Age:  35,
			Sex:  fake.Gender(),
		}
		bitUserDao.Save(&buser)
	}
}

func TestBitUserFriendDao_Save(t *testing.T) {
	bitUserFriendDao := NewBitUserFriendDao(pg_plsql)

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10000; i++ {
		fids := make([]int, 0)
		for j := 0; j < 100; j++ {
			fids = append(fids, rand.Intn(10000))
		}
		fids = utee.UniqueInt(fids)
		for _, fid := range fids {
			buf := BitUserFriend{
				Uid:    i,
				Friend: fid,
			}
			bitUserFriendDao.Save(&buf)
			buf2 := BitUserFriend{
				Uid:    fid,
				Friend: i,
			}
			bitUserFriendDao.Save(&buf2)
		}
	}
}
