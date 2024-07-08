package utils

import "encoding/json"

// GetValueSize
// @Description 获取value的大小（大概）
// @Author Oberl-Fitzgerald 2024-07-08 16:22:03
// @Param  value interface{}
// @Return int64
func GetValueSize(value interface{}) int64 {
	bytes, _ := json.Marshal(value)
	size := int64(len(bytes))
	return size
}
