package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type GetStream struct {
	reader io.Reader	//记录http返回的io.reader
}

//newGetStream 向url 发送get请求，拿到返回的响应和错误码，如果没有错误，新建一个GetStream返回
func newGetStream(url string)(*GetStream,error){
	//get 向url（数据服务节点）发送一个get请求，返回响应和错误码
	r,err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("dataServer return http code %d",r.StatusCode)
	}
	return &GetStream{r.Body},nil
}

//NewGetStream server：数据服务节点的地址，object：要get的存储对象；函数将两者拼接成一个完整的url，调用内部的newGetStream函数
//返回的*GetStream里的reader就是要接口服务要get到的目标存储对象
func NewGetStream(server,object string)(*GetStream,error){
	//验证数据真实性
	if server =="" || object ==""{
		return nil,fmt.Errorf("invalid server %s object %s",server,object)
	}
	//调用内部函数获取GetStream
	return newGetStream("http://"+server+"/objects/"+object)
}

//Read 方法用于读取GetStream内部信息
func (r *GetStream)Read(p []byte)(n int, err error){
	//reader本质是一个io.reader
	return r.reader.Read(p)
}