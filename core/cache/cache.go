package cache

var m = make(map[uint32]interface{})

func SetCache(key uint32, value interface{}) {
	m[key] = value
}

func GetCache(key uint32) interface{} {
	return m[key]
}

func DeleteCache(key uint32) {
	delete(m, key)
}
