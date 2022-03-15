package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// 创建zap包内的log
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	// 开始监听8081端口
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	// 连接数据库
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("cannot context mongodb", zap.Error(err))
	}

	// 从本地文件里获取private key.
	f, err := os.Open("auth/private-key.txt")
	if err != nil {
		logger.Fatal("cannot open private key file", zap.Error(err))
	}
	// 接着读取
	pkBytes, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Fatal("cannot read all private key file.", zap.Error(err))
	}

	// 格式转化
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot transfer []byte read from file to a *rsa.PrivateKey.")
	}

	// 注册grpc服务
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID: 		"wx2aed1cedd66881d3",
			AppSecret: 	"e62790d4227d4409ca5765e9e4259ef5",
		},
		Mongo: dao.NewMongo(mc.Database("coolcar")),
		Logger: logger,
		TokenGenerator:	token.NewJWTTokenGen("coolcar/auth", privateKey),
		TokenExpire:	2 * time.Hour,
	})

	// grpc开始在8081端口Serve
	err = s.Serve(lis)
	logger.Fatal("cannot serve", zap.Error(err))
}
