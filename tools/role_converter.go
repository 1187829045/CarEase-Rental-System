package tools

import (
	"strconv"
)

// ConvertStringRolesToInt8 将字符串角色数组转换为整数角色数组
// 特别处理：将"1"转换为0
func ConvertStringRolesToInt8(roleStrings []string) []int8 {
	result := make([]int8, 0, len(roleStrings))
	for _, roleStr := range roleStrings {
		// 把"1"转换为0
		if roleStr == "1" {
			result = append(result, 0)
		} else {
			// 尝试转换为整数
			if roleID, err := strconv.ParseInt(roleStr, 10, 8); err == nil {
				result = append(result, int8(roleID))
			}
		}
	}
	return result
}

// ParseRoleString 解析角色字符串，支持单个字符串和JSON数组字符串
func ParseRoleString(roleStr string) []string {
	// 这里可以添加JSON数组解析逻辑
	// 暂时简单返回单个字符串
	return []string{roleStr}
}
