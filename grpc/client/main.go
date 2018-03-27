package main
import (
	"google.golang.org/grpc"
	pb "github.com/evilabandon/gonetstudy/grpc/protobuf"
	"github.com/evilabandon/gonetstudy/grpc/connect"
	"fmt"
	"sync"
)

const (
	ADDRESS = "127.0.0.1:10023"
)

func main() {
	conn,err := grpc.Dial(ADDRESS,grpc.WithInsecure())
	if err !=nil{
		fmt.Println("die not connect:%v",err)
	}
	defer conn.Close()

	factory := func()(interface{},error) {
		return pb.NewDataClient(conn),nil
	}

	close := func(v interface{} ) error {return conn.Close()}

	p,err := connect.InitThread(10,30,factory,close)
	if err!=nil{
		fmt.Println("init error")
		return
	}
	var wg sync.WaitGroup
	for i:=0;i<50;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			v,_:=p.Get()
			client := v.(pb.DataClient)
			info := &pb.UserInfoRequest{
				Uid:10012,
			}
			connect.GetUserInfo(client,info)
			p.Put(v)
		}()
		wg.Wait()
	}


	for i := 0;i < 50;i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			//获取连接
			v,_ := p.Get()
			client := v.(pb.DataClient)
			connect.ChangeUserInfo(client)
			//归还链接
			p.Put(v)
		}()
		wg.Wait()
	}
	//获取链接池大小
	current := p.Len()
	fmt.Println("len=", current)
}