package server

import (
	"calculator-server/gen/elizav1/v1" // 使用正确的导入路径
	"connectrpc.com/connect"
	"context"
	"fmt"
)

type CalculatorServer struct{}

// 实现 CalculatorServiceHandler 接口
func (s *CalculatorServer) Calculate(
	ctx context.Context,
	req *connect.Request[elizav1.CalculateRequest],
) (*connect.Response[elizav1.CalculateResponse], error) {
	var result float64
	switch req.Msg.Operator {
	case "+":
		result = req.Msg.A + req.Msg.B
	case "-":
		result = req.Msg.A - req.Msg.B
	case "*":
		result = req.Msg.A * req.Msg.B
	case "/":
		if req.Msg.B == 0 {
			return nil, connect.NewError(connect.CodeInvalidArgument,
				fmt.Errorf("division by zero"))
		}
		result = req.Msg.A / req.Msg.B
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument,
			fmt.Errorf("unknown operator: %s", req.Msg.Operator))
	}

	return connect.NewResponse(&elizav1.CalculateResponse{
		Result: result,
	}), nil
}
