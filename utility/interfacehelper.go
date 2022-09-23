package utility

import (
	"strconv"
)

func Interface2String(value interface{}) (string, bool) {
	success := false
	switch value.(type) {
	case string:
		success = true
		return value.(string), success
	}
	return "", false
}

func Interface2uint64(value interface{}) (uint64, bool) {
	success := false
	switch value.(type) {
	case uint64:
		success = true
		return value.(uint64), success
	case string:
		i, err := strconv.ParseUint(value.(string), 10, 64)
		if err == nil {
			success = true
		}
		return i, success
	}
	return 0, false
}

func Interface2uint32(value interface{}) (uint32, bool) {
	success := false
	switch value.(type) {
	case uint32:
		success = true
		return value.(uint32), success
	case string:
		i, err := strconv.ParseUint(value.(string), 10, 32)
		if err == nil {
			success = true
		}
		return uint32(i), success
	}
	return 0, false
}

func Interface2int64(value interface{}) (int64, bool) {
	success := false
	switch value.(type) {
	case int64:
		success = true
		return value.(int64), success
	case string:
		i, _ := strconv.ParseInt(value.(string), 10, 64)
		return i, success
	}
	return 0, false
}

func Interface2int32(value interface{}) (int32, bool) {
	success := false
	switch value.(type) {
	case int32:
		success = true
		return value.(int32), success
	case string:
		i, _ := strconv.ParseInt(value.(string), 10, 32)
		return int32(i), success
	}
	return 0, false
}

func Interface2int(value interface{}) (int, bool) {
	success := false
	switch value.(type) {
	case int:
		success = true
		return value.(int), success
	case string:
		i, _ := strconv.Atoi(value.(string))
		if i != 0 {
			success = true
		}
		return i, success
	}
	return 0, false
}
