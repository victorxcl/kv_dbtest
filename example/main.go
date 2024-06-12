package main

import (
	"fmt"

	"github.com/xlander-io/kv/kv_leveldb"
)

func main() {

	kvdb, err := kv_leveldb.NewDB("./test.db")
	if err != nil {
		fmt.Println("new db err:", err)
		return
	}
	defer kvdb.Close()

	// //simple test
	// kvdb.Put([]byte("key1"), []byte("val1"), false)
	// kvdb.Put([]byte("key2"), []byte("val2"), true)
	// kvdb.Put([]byte("key1"), []byte("val11"), false)

	// //simple get
	// val1_, err := kvdb.Get([]byte("key1"))
	// fmt.Println("kvdb key1 get :", string(val1_))

	// val2_, err := kvdb.Get([]byte("key2"))
	// fmt.Println("kvdb key2 get :", string(val2_))

	// //batch test
	// b := kv_interface.NewBatch()
	// b.Put([]byte("key1"), []byte("val1"))
	// b.Put([]byte("key1"), []byte("val11"))
	// b.Put([]byte("key2"), []byte("val2"))
	// b.Put([]byte("key3"), []byte("val3"))
	// b.Put([]byte("key4"), []byte("val4"))
	// b.Delete([]byte("key3"))

	// kvdb.WriteBatch(b, true)

	// val1_, err := kvdb.Get([]byte("key1"))
	// fmt.Println("kvdb key1 get :", string(val1_))

	// val2_, err := kvdb.Get([]byte("key2"))
	// fmt.Println("kvdb key2 get :", string(val2_))

	// val3_, err := kvdb.Get([]byte("key3"))
	// fmt.Println("kvdb key3 get :", string(val3_))

	// val4_, err := kvdb.Get([]byte("key4"))
	// fmt.Println("kvdb key4 get :", string(val4_))

	//iterator test

	kvdb.Put([]byte("key1"), []byte("content1"), false)
	kvdb.Put([]byte("key11"), []byte("content2"), true)
	kvdb.Put([]byte("key111"), []byte("content3"), false)
	kvdb.Put([]byte("key1111"), []byte("content4"), true)

	iter := kvdb.NewIterator([]byte("key111"), []byte("key1111"))
	iter.Seek([]byte("key1")) //"keyaaa" "keybbb"

	key := iter.Key()
	value := iter.Value()

	fmt.Println(string(key), string(value))

	iter.Next()
	key2 := iter.Key()
	value2 := iter.Value()

	fmt.Println(string(key2), string(value2))

	iter.Next()
	key3 := iter.Key()
	value3 := iter.Value()

	fmt.Println(string(key3), string(value3))

}
