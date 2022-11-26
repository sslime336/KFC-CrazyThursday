package kfc

import (
	"strconv"
	"strings"

	"github.com/sslime336/kfct/pos"
)

type KFCPosition struct {
	Location string  // 可读表达
	X, Y     float64 // 经(X)纬(Y)度
}

// Position 返回附近 KFC 的位置
func Position() (kpos []KFCPosition) {
	kpos = make([]KFCPosition, 0)
	res := pos.End()
	for _, p := range res.Pois {
		l := KFCPosition{}
		l.Location = p.Adname + " " + p.Address + " " + p.Name
		fields := strings.FieldsFunc(p.Location, func(r rune) bool {
			return r == ','
		})
		l.X, _ = strconv.ParseFloat(fields[0], 64)
		l.Y, _ = strconv.ParseFloat(fields[1], 64)
		kpos = append(kpos, l)
	}
	return
}

func PosRaw() *pos.KFCPosResponse {
	return pos.End()
}
