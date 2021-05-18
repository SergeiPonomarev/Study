package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SergeiPonomarev/Study"
	"google.golang.org/grpc"
)

func main() {

	//Открываем соединение, grpc.WithInsecure() означает,
	//что шифрование не используется
	conn, err := grpc.Dial("localhost:20100", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	/*

	   Создаём нового клиента, используя соединение conn
	   Обратим внимание на название клиента и на название сервиса,
	   которое мы определили в proto-файле:

	   service Mailer {
	   rpc SendPass(MsgRequest) returns (MsgReply) {}
	   rpc RetrievePass(MsgRequest) returns (MsgReply) {}
	   }

	*/

	c := pb.NewMailerClient(conn)

	//Определяем контекст с таймаутом в 1 секунду
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	/*
	   Шлём запрос 1, ожидаем получение true в структуру rply
	   типа MsgReply, определённую в прото-файле как:

	   message MsgReply {
	   bool sent = 1;
	   }

	*/

	rply, err := c.SendPass(ctx, &pb.MsgRequest{"first", "test"})
	if err != nil {
		log.Println("something went wrong", err)
	}
	log.Println(rply.Sent)

	//Шлём запрос 2, ожидаем false
	rply, err = c.RetrievePass(ctx, &pb.MsgRequest{"second", "test"})
	if err != nil {
		log.Println("something went wrong", err)
	}
	log.Println(rply.Sent)

}
