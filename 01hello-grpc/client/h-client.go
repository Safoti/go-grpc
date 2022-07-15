package client

import (
     "context"
     "google.golang.org/grpc"
     "log"
    ps "go-grpc-demo/01hello-grpc/pb"
     "time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/14
 * @Description   grpc 客户端编写、00 吧  0
 **/
func main() {
   conn ,err:=grpc.Dial("localhost:8090",grpc.WithInsecure(),grpc.WithBlock())
     if err != nil {
          log.Fatalf("did not connect: %v", err)
     }

     // main方法执行完毕后关闭远程连接
     defer conn.Close()
     c:=ps.NewGreeterClient(conn)
     name:="world"
     // 超时设置

     ctx, cancel := context.WithTimeout(context.Background(), time.Second)
     defer cancel()
     r, err :=  c.SayHello(ctx,&ps.HelloRequest{Name: name})
     if err != nil {
          log.Fatalf("could not greet: %v", err)
     }
     log.Panicf("greeting :%s",r.GetMessage())
}