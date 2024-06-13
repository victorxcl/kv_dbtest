module example

go 1.22.4

require github.com/xlander-io/kv_leveldb v0.0.0-20240612131431-b28cf91a4944

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/xlander-io/kv v0.0.0-20240612131354-faf15072c970 // indirect
)

replace github.com/xlander-io/kv => ../kv.git

replace github.com/xlander-io/kv_leveldb => ../kv_leveldb.git
