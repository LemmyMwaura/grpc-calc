package calcserver

import (
	"context"

	"github.com/LemmyMwaura/grpccalc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CalcServer struct {
	pb.UnimplementedCalculatorServer
}

func NewCalcServer() *CalcServer {
	return &CalcServer{}
}

func (s *CalcServer) Add(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.A + in.B,
	}, nil
}

func (s *CalcServer) Subtract(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.A - in.B,
	}, nil
}

func (s *CalcServer) Divide(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	if in.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot divide by zero")
	}

	return &pb.CalculationResponse{
		Result: in.A / in.B,
	}, nil
}

func (s *CalcServer) Multiply(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.A * in.B,
	}, nil
}

func (s *CalcServer) Sum(ctx context.Context, in *pb.NumbersRequest) (*pb.CalculationResponse, error) {
	var sum int64 = 0

	for _, num := range in.Numbers {
		sum += num
	}

	return &pb.CalculationResponse{
		Result: sum,
	}, nil
}
