package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type PutStream struct {
	writer *io.PipeWriter
	c chan error	//如果发生错误，通过c传输错误给主线程
}

//NewPutStream 。。。
func NewPutStream(server,object string)*PutStream{
	//reader和writer 是互通的，在writer中写入就可以在reader中读出来
	reader,writer := io.Pipe()
	c := make(chan error)
	go func() {
		//NewRequest使用指定的方法、网址和可选的主题创建并返回一个新的*Request。
		request,_ := http.NewRequest("PUT","http://"+server+"/objects/"+object,reader)
		client := http.Client{}
		//将创建的request发送,返回响应r
		r,err := client.Do(request)
		if err != nil && r.StatusCode != http.StatusOK{
			err = fmt.Errorf("dataserver return http code %d",r.StatusCode)
		}
		//把错误传入channel
		c <- err
	}()
	return &PutStream{writer,c}
}


func (w *PutStream)Write(p []byte)(n int,err error){
	return w.writer.Write(p)
}

func (w *PutStream)Close() error {
	w.writer.Close()
	return <- w.c
}