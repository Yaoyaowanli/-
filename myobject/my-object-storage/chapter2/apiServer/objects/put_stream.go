package objects

import (
	"fmt"
	"my-object-storage/chapter2/apiServer/heartBeat"
	"my-object-storage/src/objectstream"
)

func putStream (object string)(*objectstream.PutStream,error){
	//寻找可用的心跳信息正常的数据服务节点
	server := heartBeat.ChooseRandomDataServer()
	if server ==""{
		return nil,fmt.Errorf("cannot find any dataserver ")
	}
	return objectstream.NewPutStream(server,object),nil
}
