package heartBeat

import "math/rand"

func ChooseRandomDataServer ()string{
	ds := GetDataServers()
	n := len(ds)
	//如果Get到的dataServers里没有数据就返回空串
	if n == 0 {
		return ""
	}
	//如果有数据就随机返回一个
	return ds[rand.Intn(n)]
}
