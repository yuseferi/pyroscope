// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: storegateway/v1/storegateway.proto

package storegatewayv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/grafana/pyroscope/api/gen/proto/go/ingester/v1"
	_ "github.com/grafana/pyroscope/api/gen/proto/go/storegateway/v1"
	v11 "github.com/grafana/pyroscope/api/gen/proto/go/types/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// StoreGatewayServiceName is the fully-qualified name of the StoreGatewayService service.
	StoreGatewayServiceName = "storegateway.v1.StoreGatewayService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StoreGatewayServiceMergeProfilesStacktracesProcedure is the fully-qualified name of the
	// StoreGatewayService's MergeProfilesStacktraces RPC.
	StoreGatewayServiceMergeProfilesStacktracesProcedure = "/storegateway.v1.StoreGatewayService/MergeProfilesStacktraces"
	// StoreGatewayServiceMergeProfilesLabelsProcedure is the fully-qualified name of the
	// StoreGatewayService's MergeProfilesLabels RPC.
	StoreGatewayServiceMergeProfilesLabelsProcedure = "/storegateway.v1.StoreGatewayService/MergeProfilesLabels"
	// StoreGatewayServiceMergeProfilesPprofProcedure is the fully-qualified name of the
	// StoreGatewayService's MergeProfilesPprof RPC.
	StoreGatewayServiceMergeProfilesPprofProcedure = "/storegateway.v1.StoreGatewayService/MergeProfilesPprof"
	// StoreGatewayServiceMergeSpanProfileProcedure is the fully-qualified name of the
	// StoreGatewayService's MergeSpanProfile RPC.
	StoreGatewayServiceMergeSpanProfileProcedure = "/storegateway.v1.StoreGatewayService/MergeSpanProfile"
	// StoreGatewayServiceLabelNamesProcedure is the fully-qualified name of the StoreGatewayService's
	// LabelNames RPC.
	StoreGatewayServiceLabelNamesProcedure = "/storegateway.v1.StoreGatewayService/LabelNames"
	// StoreGatewayServiceSeriesProcedure is the fully-qualified name of the StoreGatewayService's
	// Series RPC.
	StoreGatewayServiceSeriesProcedure = "/storegateway.v1.StoreGatewayService/Series"
)

// StoreGatewayServiceClient is a client for the storegateway.v1.StoreGatewayService service.
type StoreGatewayServiceClient interface {
	MergeProfilesStacktraces(context.Context) *connect_go.BidiStreamForClient[v1.MergeProfilesStacktracesRequest, v1.MergeProfilesStacktracesResponse]
	MergeProfilesLabels(context.Context) *connect_go.BidiStreamForClient[v1.MergeProfilesLabelsRequest, v1.MergeProfilesLabelsResponse]
	MergeProfilesPprof(context.Context) *connect_go.BidiStreamForClient[v1.MergeProfilesPprofRequest, v1.MergeProfilesPprofResponse]
	MergeSpanProfile(context.Context) *connect_go.BidiStreamForClient[v1.MergeSpanProfileRequest, v1.MergeSpanProfileResponse]
	LabelNames(context.Context, *connect_go.Request[v11.LabelNamesRequest]) (*connect_go.Response[v11.LabelNamesResponse], error)
	Series(context.Context, *connect_go.Request[v1.SeriesRequest]) (*connect_go.Response[v1.SeriesResponse], error)
}

// NewStoreGatewayServiceClient constructs a client for the storegateway.v1.StoreGatewayService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStoreGatewayServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) StoreGatewayServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &storeGatewayServiceClient{
		mergeProfilesStacktraces: connect_go.NewClient[v1.MergeProfilesStacktracesRequest, v1.MergeProfilesStacktracesResponse](
			httpClient,
			baseURL+StoreGatewayServiceMergeProfilesStacktracesProcedure,
			opts...,
		),
		mergeProfilesLabels: connect_go.NewClient[v1.MergeProfilesLabelsRequest, v1.MergeProfilesLabelsResponse](
			httpClient,
			baseURL+StoreGatewayServiceMergeProfilesLabelsProcedure,
			opts...,
		),
		mergeProfilesPprof: connect_go.NewClient[v1.MergeProfilesPprofRequest, v1.MergeProfilesPprofResponse](
			httpClient,
			baseURL+StoreGatewayServiceMergeProfilesPprofProcedure,
			opts...,
		),
		mergeSpanProfile: connect_go.NewClient[v1.MergeSpanProfileRequest, v1.MergeSpanProfileResponse](
			httpClient,
			baseURL+StoreGatewayServiceMergeSpanProfileProcedure,
			opts...,
		),
		labelNames: connect_go.NewClient[v11.LabelNamesRequest, v11.LabelNamesResponse](
			httpClient,
			baseURL+StoreGatewayServiceLabelNamesProcedure,
			opts...,
		),
		series: connect_go.NewClient[v1.SeriesRequest, v1.SeriesResponse](
			httpClient,
			baseURL+StoreGatewayServiceSeriesProcedure,
			opts...,
		),
	}
}

// storeGatewayServiceClient implements StoreGatewayServiceClient.
type storeGatewayServiceClient struct {
	mergeProfilesStacktraces *connect_go.Client[v1.MergeProfilesStacktracesRequest, v1.MergeProfilesStacktracesResponse]
	mergeProfilesLabels      *connect_go.Client[v1.MergeProfilesLabelsRequest, v1.MergeProfilesLabelsResponse]
	mergeProfilesPprof       *connect_go.Client[v1.MergeProfilesPprofRequest, v1.MergeProfilesPprofResponse]
	mergeSpanProfile         *connect_go.Client[v1.MergeSpanProfileRequest, v1.MergeSpanProfileResponse]
	labelNames               *connect_go.Client[v11.LabelNamesRequest, v11.LabelNamesResponse]
	series                   *connect_go.Client[v1.SeriesRequest, v1.SeriesResponse]
}

// MergeProfilesStacktraces calls storegateway.v1.StoreGatewayService.MergeProfilesStacktraces.
func (c *storeGatewayServiceClient) MergeProfilesStacktraces(ctx context.Context) *connect_go.BidiStreamForClient[v1.MergeProfilesStacktracesRequest, v1.MergeProfilesStacktracesResponse] {
	return c.mergeProfilesStacktraces.CallBidiStream(ctx)
}

// MergeProfilesLabels calls storegateway.v1.StoreGatewayService.MergeProfilesLabels.
func (c *storeGatewayServiceClient) MergeProfilesLabels(ctx context.Context) *connect_go.BidiStreamForClient[v1.MergeProfilesLabelsRequest, v1.MergeProfilesLabelsResponse] {
	return c.mergeProfilesLabels.CallBidiStream(ctx)
}

// MergeProfilesPprof calls storegateway.v1.StoreGatewayService.MergeProfilesPprof.
func (c *storeGatewayServiceClient) MergeProfilesPprof(ctx context.Context) *connect_go.BidiStreamForClient[v1.MergeProfilesPprofRequest, v1.MergeProfilesPprofResponse] {
	return c.mergeProfilesPprof.CallBidiStream(ctx)
}

// MergeSpanProfile calls storegateway.v1.StoreGatewayService.MergeSpanProfile.
func (c *storeGatewayServiceClient) MergeSpanProfile(ctx context.Context) *connect_go.BidiStreamForClient[v1.MergeSpanProfileRequest, v1.MergeSpanProfileResponse] {
	return c.mergeSpanProfile.CallBidiStream(ctx)
}

// LabelNames calls storegateway.v1.StoreGatewayService.LabelNames.
func (c *storeGatewayServiceClient) LabelNames(ctx context.Context, req *connect_go.Request[v11.LabelNamesRequest]) (*connect_go.Response[v11.LabelNamesResponse], error) {
	return c.labelNames.CallUnary(ctx, req)
}

// Series calls storegateway.v1.StoreGatewayService.Series.
func (c *storeGatewayServiceClient) Series(ctx context.Context, req *connect_go.Request[v1.SeriesRequest]) (*connect_go.Response[v1.SeriesResponse], error) {
	return c.series.CallUnary(ctx, req)
}

// StoreGatewayServiceHandler is an implementation of the storegateway.v1.StoreGatewayService
// service.
type StoreGatewayServiceHandler interface {
	MergeProfilesStacktraces(context.Context, *connect_go.BidiStream[v1.MergeProfilesStacktracesRequest, v1.MergeProfilesStacktracesResponse]) error
	MergeProfilesLabels(context.Context, *connect_go.BidiStream[v1.MergeProfilesLabelsRequest, v1.MergeProfilesLabelsResponse]) error
	MergeProfilesPprof(context.Context, *connect_go.BidiStream[v1.MergeProfilesPprofRequest, v1.MergeProfilesPprofResponse]) error
	MergeSpanProfile(context.Context, *connect_go.BidiStream[v1.MergeSpanProfileRequest, v1.MergeSpanProfileResponse]) error
	LabelNames(context.Context, *connect_go.Request[v11.LabelNamesRequest]) (*connect_go.Response[v11.LabelNamesResponse], error)
	Series(context.Context, *connect_go.Request[v1.SeriesRequest]) (*connect_go.Response[v1.SeriesResponse], error)
}

// NewStoreGatewayServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStoreGatewayServiceHandler(svc StoreGatewayServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	storeGatewayServiceMergeProfilesStacktracesHandler := connect_go.NewBidiStreamHandler(
		StoreGatewayServiceMergeProfilesStacktracesProcedure,
		svc.MergeProfilesStacktraces,
		opts...,
	)
	storeGatewayServiceMergeProfilesLabelsHandler := connect_go.NewBidiStreamHandler(
		StoreGatewayServiceMergeProfilesLabelsProcedure,
		svc.MergeProfilesLabels,
		opts...,
	)
	storeGatewayServiceMergeProfilesPprofHandler := connect_go.NewBidiStreamHandler(
		StoreGatewayServiceMergeProfilesPprofProcedure,
		svc.MergeProfilesPprof,
		opts...,
	)
	storeGatewayServiceMergeSpanProfileHandler := connect_go.NewBidiStreamHandler(
		StoreGatewayServiceMergeSpanProfileProcedure,
		svc.MergeSpanProfile,
		opts...,
	)
	storeGatewayServiceLabelNamesHandler := connect_go.NewUnaryHandler(
		StoreGatewayServiceLabelNamesProcedure,
		svc.LabelNames,
		opts...,
	)
	storeGatewayServiceSeriesHandler := connect_go.NewUnaryHandler(
		StoreGatewayServiceSeriesProcedure,
		svc.Series,
		opts...,
	)
	return "/storegateway.v1.StoreGatewayService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StoreGatewayServiceMergeProfilesStacktracesProcedure:
			storeGatewayServiceMergeProfilesStacktracesHandler.ServeHTTP(w, r)
		case StoreGatewayServiceMergeProfilesLabelsProcedure:
			storeGatewayServiceMergeProfilesLabelsHandler.ServeHTTP(w, r)
		case StoreGatewayServiceMergeProfilesPprofProcedure:
			storeGatewayServiceMergeProfilesPprofHandler.ServeHTTP(w, r)
		case StoreGatewayServiceMergeSpanProfileProcedure:
			storeGatewayServiceMergeSpanProfileHandler.ServeHTTP(w, r)
		case StoreGatewayServiceLabelNamesProcedure:
			storeGatewayServiceLabelNamesHandler.ServeHTTP(w, r)
		case StoreGatewayServiceSeriesProcedure:
			storeGatewayServiceSeriesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStoreGatewayServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStoreGatewayServiceHandler struct{}

func (UnimplementedStoreGatewayServiceHandler) MergeProfilesStacktraces(context.Context, *connect_go.BidiStream[v1.MergeProfilesStacktracesRequest, v1.MergeProfilesStacktracesResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.MergeProfilesStacktraces is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) MergeProfilesLabels(context.Context, *connect_go.BidiStream[v1.MergeProfilesLabelsRequest, v1.MergeProfilesLabelsResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.MergeProfilesLabels is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) MergeProfilesPprof(context.Context, *connect_go.BidiStream[v1.MergeProfilesPprofRequest, v1.MergeProfilesPprofResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.MergeProfilesPprof is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) MergeSpanProfile(context.Context, *connect_go.BidiStream[v1.MergeSpanProfileRequest, v1.MergeSpanProfileResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.MergeSpanProfile is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) LabelNames(context.Context, *connect_go.Request[v11.LabelNamesRequest]) (*connect_go.Response[v11.LabelNamesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.LabelNames is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) Series(context.Context, *connect_go.Request[v1.SeriesRequest]) (*connect_go.Response[v1.SeriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.Series is not implemented"))
}
