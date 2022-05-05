package objects

import (
	"fmt"
	"io"
	"my-object-storage/chapter2/apiServer/locate"
	"my-object-storage/src/objectstream"
)

func getStream(object string)(io.Reader,error){
	server := locate.Locate(object)
	if server ==""{
		return nil,fmt.Errorf("object %s locate fail",object)
	}
	//getStream 实现了io.reader接口
	return objectstream.NewGetStream(server,object)
}
