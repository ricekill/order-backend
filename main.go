package main

import (
	"order-backend/common"
	_ "order-backend/docs" // docs is generated by Swag CLI, you have to import it.
	"order-backend/router"
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"log"
	"os"
	"runtime"
)
// @title Project Order Backend API
// @version 1.0
// @description 订单模块微服务.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 0.0.0.0:5070
// @BasePath /api/v1
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	os.Mkdir("runtime",os.ModeDir|os.ModePerm)

	common.CheckErr(common.LoadConfig())
	common.CheckErr(common.SetupLogger())
	common.CheckErr(common.OpenDb())
	defer common.DB.Close()

	common.CheckErr(common.OpenRedis())
	common.Log.Infoln("[JOB]","Every 5 sec send reminder emails")

	uri := "tcp://" + common.Config.Listen + "/"
	fmt.Println("url === ", uri)
	server := rpc.NewTCPServer(uri)
	server.Debug = common.Config.System.Debug
	router.TcpRouter(server)
	server.Start()

	log.Println("Server exiting")

	defer server.Close()
}