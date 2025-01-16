package object

// getDefaultValue 获取默认值，如果未提供则返回零值
func getDefaultValue[V any](defaultValue ...V) V {
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	var zero V
	return zero
}
