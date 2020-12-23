package db

import(
	"github.com/rekuza/chadnet"
	badger "github.com/timshannon/badgerhold"
)

func AddBlock(key,data interface{}) (bool) {
	var opt = badger.DefaultOptions
	opt.Dir = "/tmp/chaddb"
	opt.ValueDir = "/tmp/chaddb"
	var db,err = badger.Open(badger.DefaultOptions)
	err = db.Insert(key,data)
	if err != nil {
		return false
	}
	return true
}
