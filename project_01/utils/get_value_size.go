package utils

import (
	"reflect"
	"unsafe"
)

// GetValueSize
// @Description 获取value的大小
// @Author Oberl-Fitzgerald 2024-07-08 16:22:03
// @Param  value interface{}
// @Return int64
func GetValueSize(value interface{}) int64 {
	return getValueSize(reflect.ValueOf(value), make(map[uintptr]bool))
}

// getValueSize
// @Description 获取value的大小
// @Author Oberl-Fitzgerald 2024-07-09 14:48:23
// @Param  v reflect.Value
// @Param  visited map[uintptr]bool
// @Return int64
func getValueSize(v reflect.Value, visited map[uintptr]bool) int64 {
	if !v.IsValid() {
		return 0
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return 0
		}
		ptr := v.Pointer()
		if visited[ptr] {
			return 0
		}
		visited[ptr] = true
		return int64(unsafe.Sizeof(ptr)) + getValueSize(v.Elem(), visited)
	case reflect.Slice:
		if v.IsNil() {
			return 0
		}
		size := int64(unsafe.Sizeof(v.Pointer())) + int64(v.Cap())*int64(v.Type().Elem().Size())
		for i := 0; i < v.Len(); i++ {
			size += getValueSize(v.Index(i), visited)
		}
		return size
	case reflect.Map:
		if v.IsNil() {
			return 0
		}
		size := int64(unsafe.Sizeof(v.Pointer()))
		for _, key := range v.MapKeys() {
			size += getValueSize(key, visited)
			size += getValueSize(v.MapIndex(key), visited)
		}
		return size
	case reflect.String:
		return int64(unsafe.Sizeof(v.String())) + int64(v.Len())
	case reflect.Struct:
		size := int64(0)
		for i := 0; i < v.NumField(); i++ {
			size += getValueSize(v.Field(i), visited)
		}
		return size
	case reflect.Array:
		size := int64(0)
		for i := 0; i < v.Len(); i++ {
			size += getValueSize(v.Index(i), visited)
		}
		return size
	case reflect.Interface:
		if v.IsNil() {
			return 0
		}
		return int64(unsafe.Sizeof(v.Interface())) + getValueSize(v.Elem(), visited)
	default:
		return int64(v.Type().Size())
	}
}
