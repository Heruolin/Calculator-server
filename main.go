package main

import (
	"calculator-server/gen/elizav1/v1/elizav1connect" // 导入生成的服务连接代码
	"calculator-server/server"
	"log"
	"net/http"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// 创建服务实例
	calculator := &server.CalculatorServer{} // 确保 CalculatorServer 实现了 CalculatorServiceHandler 接口
	mux := http.NewServeMux()

	// 注册服务
	_, handler := elizav1connect.NewCalculatorServiceHandler(calculator) // 确保类型匹配
	mux.Handle("/connectrpc.eliza.v1.ElizaService/Say", handler)         // 修正服务路径

	// CORS 配置
	handlerWithCORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},                                     // 允许来自前端的请求
		AllowedMethods:   []string{"POST", "OPTIONS"},                                           // 支持的请求方法
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Connect-Protocol-Version"}, // 允许的请求头
		AllowCredentials: true,                                                                  // 支持跨域请求的凭证
		Debug:            true,                                                                  // 打印调试信息
	}).Handler(mux)

	// 启动服务，使用 HTTP/2 协议
	log.Println("Server is running at localhost:8080")
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(handlerWithCORS, &http2.Server{}), // 使用 h2c（HTTP/2 Clear Text）启用 HTTP/2
	)
}
