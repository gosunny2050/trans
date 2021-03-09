package interface_utils

import (
	"encoding/json"
	"errors"
	"strconv"
)

func InterfaceToStruct(value interface{}, obj interface{}) error {
	if value == nil {
		return errors.New("value is nil")
	}
	switch value.(type) {
		case string:
			value1 := value.(string)
			jsonResErr:= json.Unmarshal([]byte(value1), obj)
			if jsonResErr != nil {
				return jsonResErr
			}
		default:
			retJson, err:= json.Marshal(value)
				if err != nil {
					return err
				}
			jsonResErr:= json.Unmarshal(retJson, obj)
			if jsonResErr != nil {
				return jsonResErr
			}
	}
	return nil
}
func InterfaceToBool(value interface{}) bool {
	if value == nil {
		return false
	}
	switch retValue := value.(type) {
		case bool:
			return retValue
		case *bool:
			if retValue == nil {
				return false
			}
			return *retValue
		case int,int8,int16,int32,int64:
			return retValue != 0
		case *int:
			if retValue == nil {
				return false
			}
			return *retValue != 0
		case *int8:
			if retValue == nil {
				return false
			}
			return *retValue != 0
		case *int16:
			if retValue == nil {
				return false
			}
			return *retValue != 0
		case *int32:
			if retValue == nil {
				return false
			}
			return *retValue != 0
		case *int64:
			if retValue == nil {
				return false
			}
			return *retValue != 0
		case string:
			return retValue == "true"
		case *string:
			if retValue == nil {
				return false
			}
			return *retValue == "true"
		default:
			return false
	}
}
func InterfaceToString(value interface{}) string {
	if value == nil {
		return ""
	}
	switch retValue := value.(type) {
	case string:
		return retValue
	case *string:
		if retValue == nil {
			return ""
		}
		return *retValue
	case int:
		return strconv.Itoa(retValue)
	case int8:
		return strconv.FormatInt(int64(retValue),10)
	case int16:
		return strconv.FormatInt(int64(retValue), 10)
	case int32:
		return strconv.FormatInt(int64(retValue), 10)
	case int64:
		return strconv.FormatInt(retValue, 10)
	case *int:
		if retValue == nil {
			return ""
		}
		return strconv.Itoa(*retValue)
	case *int64:
		return strconv.FormatInt(*retValue, 10)
	case *int32:
		return strconv.Itoa(int(*retValue))
	case *int16:
		return strconv.Itoa(int(*retValue))
	case *int8:
		return strconv.Itoa(int(*retValue))
	case float32:
		return strconv.FormatFloat(float64(retValue), 'f',  -1,32)
	case float64:
		return strconv.FormatFloat(float64(retValue), 'f',  -1,64)
	case *float32:
		return strconv.FormatFloat(float64(*retValue), 'f',  -1,32)
	case *float64:
		return strconv.FormatFloat(float64(*retValue), 'f',  -1,64)
	case bool:
		return  strconv.FormatBool(retValue)
	case *bool:
		if retValue == nil {
			return ""
		}
		return  strconv.FormatBool(*retValue)
	default:
		return ""
	}
}
func InterfaceToInt(value interface{}) int {
	if value == nil {
		return 0
	}
	switch retValue := value.(type) {
		case int,int8,int16,int32,int64:
			return retValue.(int)
		case *int64:
			if retValue == nil {
				return 0
			}
			return int(*retValue)
		case *int32:
			if retValue == nil {
				return 0
			}
			return int(*retValue)
		case *int16:
			if retValue == nil {
				return 0
			}
			return int(*retValue)
		case *int8:
			if retValue == nil {
				return 0
			}
			return int(*retValue)
		case float32,float64:
			return int(retValue.(float64))
		case bool:
			return 0
		case *bool:
			if retValue == nil {
				return 0
			}
			return 0
		case string:
			intRet, _ := strconv.Atoi(retValue)
			return intRet
		case *string:
			if retValue == nil {
				return 0
			}
			intRet, _ := strconv.Atoi(*retValue)
			return intRet
		case *int:
			if retValue == nil {
				return 0
			}
			return *retValue
		default:
			return 0
	}
}
func InterfaceToInt64(value interface{}) int64 {
	if value == nil {
		return 0
	}
	switch retValue := value.(type) {
		case int64:
			return retValue
		case *int64:
			if retValue == nil {
				return 0
			}
			return *retValue
		case int:
			return int64(retValue)
		case int8:
			return int64(retValue)
		case int16:
			return int64(retValue)
		case int32:
			return int64(retValue)
		case *int32:
			if retValue == nil {
				return 0
			}
			return int64(*retValue)
		case *int16:
			if retValue == nil {
				return 0
			}
			return int64(*retValue)
		case *int8:
			if retValue == nil {
				return 0
			}
			return int64(*retValue)
		case string:
			intRet, _ := strconv.ParseInt(retValue, 10, 64)
			return intRet
		case *string:
			if retValue == nil {
				return 0
			}
			intRet, _ := strconv.ParseInt(*retValue, 10, 64)
			return intRet
		case bool:
			return 0
		case *bool:
			return 0
		default:
			return 0
	}
}
func InterfaceToFloat64(value interface{}) float64 {
	if value == nil {
		return 0
	}
	switch retValue := value.(type) {
	case float64:
		return retValue
	case float32:
		return float64(retValue)
	case *int64:
		if retValue == nil {
			return 0
		}
		return float64(*retValue)
	case int:
		return float64(retValue)
	case int8:
		return float64(retValue)
	case int16:
		return float64(retValue)
	case int32:
		return float64(retValue)
	case int64:
		return float64(retValue)
	case *int32:
		if retValue == nil {
			return 0
		}
		return float64(*retValue)
	case *int16:
		if retValue == nil {
			return 0
		}
		return float64(*retValue)
	case *int8:
		if retValue == nil {
			return 0
		}
		return float64(*retValue)
	case string:
		ret, _ := strconv.ParseFloat(retValue, 64)
		return ret
	case *string:
		if retValue == nil {
			return 0
		}
		ret, _ := strconv.ParseFloat(*retValue, 64)
		return ret
	case bool:
		return 0
	case *bool:
		return 0
	default:
		return 0
	}
}

func InterfaceToFloat32(value interface{}) float32 {
	if value == nil {
		return 0
	}
	switch retValue := value.(type) {
	case float32:
		return retValue
	case float64:
		return float32(retValue)
	case *int64:
		if retValue == nil {
			return 0
		}
		return float32(*retValue)
	case int:
		return float32(retValue)
	case int8:
		return float32(retValue)
	case int16:
		return float32(retValue)
	case int32:
		return float32(retValue)
	case int64:
		return float32(retValue)
	case *int32:
		if retValue == nil {
			return 0
		}
		return float32(*retValue)
	case *int16:
		if retValue == nil {
			return 0
		}
		return float32(*retValue)
	case *int8:
		if retValue == nil {
			return 0
		}
		return float32(*retValue)
	case string:
		ret, _ := strconv.ParseFloat(retValue, 32)
		return float32(ret)
	case *string:
		if retValue == nil {
			return 0
		}
		ret, _ := strconv.ParseFloat(*retValue, 32)
		return float32(ret)
	case bool:
		return 0
	case *bool:
		return 0
	default:
		return 0
	}
}
