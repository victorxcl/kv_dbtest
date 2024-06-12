module example

go 1.22.4

require (
	github.com/xlander-io/kv_interface v0.0.0-20240612061642-b23d8dc85c41
	github.com/xlander-io/kv_leveldb v0.0.0-20240612061659-c5d5e575736d
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
)

replace github.com/xlander-io/kv_interface => ../kv_interface.git

replace github.com/xlander-io/kv_leveldb => ../kv_leveldb.git
