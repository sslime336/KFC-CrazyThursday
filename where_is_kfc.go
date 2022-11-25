package kfc

type KFCPosition struct {
	Location string  // 可读表达
	X, Y     float64 // 经(X)纬(Y)度
}

// Position 返回附近 KFC 的位置
func Position() KFCPosition {
	return KFCPosition{}
}
