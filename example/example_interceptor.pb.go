// Code generated by protoc-gen-mixgo. DO NOT EDIT!

package example

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"github.com/idle-ape/mixgo/metrics"
	"google.golang.org/grpc/status"
)

var (
	// DefaultServerCounterName default name for server counter
	DefaultServerCounterName = "server_counter"
	// DefaultServerHistogramName default name for server histogram
	DefaultServerHistogramName = "server_histogram"
	// DefaultClientCounterName default name for client counter
	DefaultClientCounterName = "client_counter"
	// DefaultClientHistogramName default name for client histogram
	DefaultClientHistogramName = "client_histogram"
	// DefaultQPSName default species name to collect qps metrics
	DefaultQPSName = "QPS"
)

// HelloWorldMetricsUnaryServerInterceptor server side unary interceptor for metrics collection
func HelloWorldMetricsUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 统计接口耗时
	start := time.Now()
	defer func() {
		metrics.Histogram(DefaultServerHistogramName).WithLabelValues(info.FullMethod).Observe(time.Since(start).Seconds())
	}()
	// 统计QPS
	metrics.Counter(DefaultServerCounterName).WithLabelValues(info.FullMethod, DefaultQPSName).Add(1)
	if resp, err = handler(ctx, req); err != nil {
		// 统计错误码分布
		if s, ok := status.FromError(err); ok {
			metrics.Counter(DefaultServerCounterName).WithLabelValues(info.FullMethod, s.Code().String()).Add(1)
		}
		return nil, err
	}
	return resp, nil
}

// HelloWorldMetricsUnaryClientInterceptor client side unary interceptor for metrics collection
func HelloWorldMetricsUnaryClientInterceptor(ctx context.Context, fullMethod string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 统计接口耗时
	start := time.Now()
	defer func() {
		metrics.Histogram(DefaultClientHistogramName).WithLabelValues(fullMethod).Observe(time.Since(start).Seconds())
	}()
	// 统计QPS
	metrics.Counter(DefaultClientCounterName).WithLabelValues(fullMethod, DefaultQPSName).Add(1)
	if err := invoker(ctx, fullMethod, req, reply, cc, opts...); err != nil {
		// 统计错误码分布
		if s, ok := status.FromError(err); ok {
			metrics.Counter(DefaultClientCounterName).WithLabelValues(fullMethod, s.Code().String()).Add(1)
		}
		return err
	}
	return nil
}
