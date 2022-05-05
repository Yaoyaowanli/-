package heartBeat

import (
	"my-object-storage/src/rabbitmq"
	"os"
	"strconv"
	"sync"
	"time"
)
//dataServers 记录了当前所有节点的心跳信息和地址
var dataServers = make(map[string]time.Time)
var rwMutex sync.RWMutex

func ListenHeatBeat(){
	//建立rabbitmq连接
	mq:=rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer mq.Close()

	//给当前mq的队列绑定交换机
	mq.Bind("apiServers")

	//消费数据服务节点的心跳消息
	msg := mq.Consume()
	go removeExpiredDataServer()
	for d := range msg{
		dataServer,err := strconv.Unquote(string(d.Body))
		if err != nil {
			panic(err)
		}
		//并发操作map需要加锁
		rwMutex.Lock()
		//更新map
		dataServers[dataServer]=time.Now()
		rwMutex.Unlock()
	}
}

//removeExpiredDataServer 每隔5秒取dataServers里遍历检查，如果有超过10秒的dataServer就删除它
func removeExpiredDataServer(){
	for {
		time.Sleep(5*time.Second)
		rwMutex.RLock()
		for s,t := range dataServers{
			if t.Add(10*time.Second).Before(time.Now()){
				rwMutex.Lock()
				delete(dataServers,s)
				rwMutex.Unlock()
			}
		}
		rwMutex.RUnlock()
	}
}

//GetDataServers 拿到当前dataServers所有的dataServer以字符串切片的形式返回
func GetDataServers()[]string{
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	ds := make([]string,0)
	for s,_ := range dataServers{
		ds=append(ds,s)
	}
	return ds
}