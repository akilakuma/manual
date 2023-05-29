package server

import (
	"context"
	"errors"

	glog "pika.rdtech.vip/genesis-lib/log"
	"pika.rdtech.vip/genesis-lib/snowflake"

	"google.golang.org/grpc"
	"pika.rdtech.vip/genesis-lib/goworker"
	httpClient "pika.rdtech.vip/genesis-lib/http-client"
	gorm "pika.rdtech.vip/genesis-lib/mysql"
	nsq "pika.rdtech.vip/genesis-lib/nsq"
	"pika.rdtech.vip/genesis-lib/redis"
	"pika.rdtech.vip/genesis-lib/scylla"
)

var serverObject *Server

// Server server object to use
type Server struct {
	Constant      map[string]interface{}
	logger        *glog.Log
	nsqProducer   nsq.Producer
	nsqConsumer   map[string]nsq.Consumer
	nsqLookUpAddr string //nolint (may not use)
	redisCacher   map[string]*redis.Cacher
	mySQLDB       map[string]*gorm.Orm
	scyllaDB      map[string]*scylla.Cqlx
	goworker      map[string]*goworker.Pool
	grpcClient    map[string]*grpc.ClientConn
	httpClient    map[string]httpClient.Methods
	gracefulCtx   *context.Context
	snowflake     snowflake.Node
}

// NewServer create server object to use
func NewServer() (newServerObject *Server, err error) {
	initServer()
	newServerObject = serverObject
	return
}

func initServer() {
	serverObject = &Server{}
	serverObject.nsqConsumer = make(map[string]nsq.Consumer)
	serverObject.redisCacher = make(map[string]*redis.Cacher)
	serverObject.mySQLDB = make(map[string]*gorm.Orm)
	serverObject.scyllaDB = make(map[string]*scylla.Cqlx)
	serverObject.goworker = make(map[string]*goworker.Pool)
	serverObject.grpcClient = make(map[string]*grpc.ClientConn)
	serverObject.httpClient = make(map[string]httpClient.Methods)
}

// GetServerInstance 取得Server struct pointer
func GetServerInstance() *Server {
	if serverObject == nil {
		initServer()
	}
	return serverObject
}

//SetLogger SetLogger
func (s *Server) SetLogger(logger *glog.Log) {
	s.logger = logger
}

//SetNSQProducer SetNSQProducer
func (s *Server) SetNSQProducer(producer nsq.Producer) {
	s.nsqProducer = producer
}

//SetNSQConsumer SetNSQConsumer
func (s *Server) SetNSQConsumer(consumer map[string]nsq.Consumer) {
	s.nsqConsumer = consumer
}

//SetRedisCacher SetRedisCacher
func (s *Server) SetRedisCacher(redisCacher map[string]*redis.Cacher) {
	s.redisCacher = redisCacher
}

//SetMySQLDB SetMySQLDB
func (s *Server) SetMySQLDB(mysqlOrm map[string]*gorm.Orm) {
	s.mySQLDB = mysqlOrm
}

//SetScyllaDB SetScyllaDB
func (s *Server) SetScyllaDB(scyllaDB map[string]*scylla.Cqlx) {
	s.scyllaDB = scyllaDB
}

//SetWorker SetWorker
func (s *Server) SetWorker(worker map[string]*goworker.Pool) {
	s.goworker = worker
}

//SetGrpcClient SetGrpcClient
func (s *Server) SetGrpcClient(grpcClient map[string]*grpc.ClientConn) {
	s.grpcClient = grpcClient
}

// SetHTTPClient setHttpClient to ServerStruct
func (s *Server) SetHTTPClient(httpClient map[string]httpClient.Methods) {
	s.httpClient = httpClient
}

// SetSnowflake SetSnowflake to ServerStruct
func (s *Server) SetSnowflake(snowflakeNode snowflake.Node) {
	s.snowflake = snowflakeNode
}

//AddHTTPClient add a new client to server.httpClient
func (s *Server) AddHTTPClient(name string, client httpClient.Methods) (err error) {
	if _, ok := s.httpClient[name]; ok {
		return errors.New("Client name " + name + " already esists.")
	}

	s.httpClient[name] = client
	return
}

// SetGracefulCtx 設定Graceful Shutdown 的 context
func (s *Server) SetGracefulCtx(c *context.Context) {
	s.gracefulCtx = c
}

//GetLogger GetLogger
func (s *Server) GetLogger() *glog.Log {
	return s.logger
}

//GetNSQProducer GetNSQProducer
func (s *Server) GetNSQProducer() (producer nsq.Producer) {
	producer = s.nsqProducer
	return
}

//GetNSQConsumer GetNSQConsumer
func (s *Server) GetNSQConsumer() (consumer map[string]nsq.Consumer) {
	consumer = s.nsqConsumer
	return
}

//GetRedisCacher GetRedisCacher
func (s *Server) GetRedisCacher() (redisCacher map[string]*redis.Cacher) {
	redisCacher = s.redisCacher
	return
}

//GetMySQL GetMySQL
func (s *Server) GetMySQL() (mysqlOrm map[string]*gorm.Orm) {
	mysqlOrm = s.mySQLDB
	return
}

//GetScylla GetScylla
func (s *Server) GetScylla() (scyllaDB map[string]*scylla.Cqlx) {
	scyllaDB = s.scyllaDB
	return
}

//GetWorker GetWorker
func (s *Server) GetWorker() (worker map[string]*goworker.Pool) {
	worker = s.goworker
	return
}

//GetGrpcClient GetGrpcClient
func (s *Server) GetGrpcClient() (grpcClient map[string]*grpc.ClientConn) {
	grpcClient = s.grpcClient
	return
}

// GetHTTPClient setHttpClient to ServerStruct
func (s *Server) GetHTTPClient() (httpClient map[string]httpClient.Methods) {
	httpClient = s.httpClient
	return
}

// GetGracefulCtx 取得Graceful Shutdown 的 context
func (s *Server) GetGracefulCtx() *context.Context {
	return s.gracefulCtx
}

// GetSnowflake 取得snowflakeNode
func (s *Server) GetSnowflake() snowflake.Node {
	return s.snowflake
}
