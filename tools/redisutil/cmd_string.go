package redisutil

func MGet() {
	redis := GetRedisClient()
	defer redis.CloseRedisClient()
}

func MSet() {
	redis := GetRedisClient()
	defer redis.CloseRedisClient()
}
func Incr() {
	redis := GetRedisClient()
	defer redis.CloseRedisClient()
}

func IncrBy() {
	redis := GetRedisClient()
	defer redis.CloseRedisClient()
}
func Set(key, val string) bool {
	redis := GetRedisClient()
	defer redis.CloseRedisClient()

	exec, err := redis.Exec("SET", key, val)
	if err != nil {
		return false
	}
	// 检查操作是否成功
	return exec == "OK"
}

func Get(key string) string {
	redis := GetRedisClient()
	defer redis.CloseRedisClient()

	// 执行 GET 操作
	reply, err := redis.String(redis.Exec("GET", key))
	if err != nil {
		return ""
	}

	// 检查操作是否成功
	return reply
}

func Exists() {
	redis := GetRedisClient()
	defer redis.CloseRedisClient()
}

func Expire() {
	redis := GetRedisClient()
	defer redis.CloseRedisClient()
}
