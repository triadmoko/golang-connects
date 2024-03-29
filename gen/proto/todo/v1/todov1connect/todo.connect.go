// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/todo/v1/todo.proto

package todov1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "golang-connects/gen/proto/todo/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TodoServiceName is the fully-qualified name of the TodoService service.
	TodoServiceName = "todo.v1.TodoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TodoServiceCreateProcedure is the fully-qualified name of the TodoService's Create RPC.
	TodoServiceCreateProcedure = "/todo.v1.TodoService/Create"
	// TodoServiceReadProcedure is the fully-qualified name of the TodoService's Read RPC.
	TodoServiceReadProcedure = "/todo.v1.TodoService/Read"
	// TodoServiceUpdateProcedure is the fully-qualified name of the TodoService's Update RPC.
	TodoServiceUpdateProcedure = "/todo.v1.TodoService/Update"
	// TodoServiceDeleteProcedure is the fully-qualified name of the TodoService's Delete RPC.
	TodoServiceDeleteProcedure = "/todo.v1.TodoService/Delete"
	// TodoServiceListProcedure is the fully-qualified name of the TodoService's List RPC.
	TodoServiceListProcedure = "/todo.v1.TodoService/List"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	todoServiceServiceDescriptor      = v1.File_proto_todo_v1_todo_proto.Services().ByName("TodoService")
	todoServiceCreateMethodDescriptor = todoServiceServiceDescriptor.Methods().ByName("Create")
	todoServiceReadMethodDescriptor   = todoServiceServiceDescriptor.Methods().ByName("Read")
	todoServiceUpdateMethodDescriptor = todoServiceServiceDescriptor.Methods().ByName("Update")
	todoServiceDeleteMethodDescriptor = todoServiceServiceDescriptor.Methods().ByName("Delete")
	todoServiceListMethodDescriptor   = todoServiceServiceDescriptor.Methods().ByName("List")
)

// TodoServiceClient is a client for the todo.v1.TodoService service.
type TodoServiceClient interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
	List(context.Context) *connect.BidiStreamForClient[v1.ListRequest, v1.ListResponse]
}

// NewTodoServiceClient constructs a client for the todo.v1.TodoService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTodoServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TodoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &todoServiceClient{
		create: connect.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+TodoServiceCreateProcedure,
			connect.WithSchema(todoServiceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		read: connect.NewClient[v1.ReadRequest, v1.ReadResponse](
			httpClient,
			baseURL+TodoServiceReadProcedure,
			connect.WithSchema(todoServiceReadMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+TodoServiceUpdateProcedure,
			connect.WithSchema(todoServiceUpdateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+TodoServiceDeleteProcedure,
			connect.WithSchema(todoServiceDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+TodoServiceListProcedure,
			connect.WithSchema(todoServiceListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// todoServiceClient implements TodoServiceClient.
type todoServiceClient struct {
	create *connect.Client[v1.CreateRequest, v1.CreateResponse]
	read   *connect.Client[v1.ReadRequest, v1.ReadResponse]
	update *connect.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
	list   *connect.Client[v1.ListRequest, v1.ListResponse]
}

// Create calls todo.v1.TodoService.Create.
func (c *todoServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Read calls todo.v1.TodoService.Read.
func (c *todoServiceClient) Read(ctx context.Context, req *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error) {
	return c.read.CallUnary(ctx, req)
}

// Update calls todo.v1.TodoService.Update.
func (c *todoServiceClient) Update(ctx context.Context, req *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls todo.v1.TodoService.Delete.
func (c *todoServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// List calls todo.v1.TodoService.List.
func (c *todoServiceClient) List(ctx context.Context) *connect.BidiStreamForClient[v1.ListRequest, v1.ListResponse] {
	return c.list.CallBidiStream(ctx)
}

// TodoServiceHandler is an implementation of the todo.v1.TodoService service.
type TodoServiceHandler interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
	List(context.Context, *connect.BidiStream[v1.ListRequest, v1.ListResponse]) error
}

// NewTodoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTodoServiceHandler(svc TodoServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	todoServiceCreateHandler := connect.NewUnaryHandler(
		TodoServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(todoServiceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	todoServiceReadHandler := connect.NewUnaryHandler(
		TodoServiceReadProcedure,
		svc.Read,
		connect.WithSchema(todoServiceReadMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	todoServiceUpdateHandler := connect.NewUnaryHandler(
		TodoServiceUpdateProcedure,
		svc.Update,
		connect.WithSchema(todoServiceUpdateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	todoServiceDeleteHandler := connect.NewUnaryHandler(
		TodoServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(todoServiceDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	todoServiceListHandler := connect.NewBidiStreamHandler(
		TodoServiceListProcedure,
		svc.List,
		connect.WithSchema(todoServiceListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/todo.v1.TodoService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TodoServiceCreateProcedure:
			todoServiceCreateHandler.ServeHTTP(w, r)
		case TodoServiceReadProcedure:
			todoServiceReadHandler.ServeHTTP(w, r)
		case TodoServiceUpdateProcedure:
			todoServiceUpdateHandler.ServeHTTP(w, r)
		case TodoServiceDeleteProcedure:
			todoServiceDeleteHandler.ServeHTTP(w, r)
		case TodoServiceListProcedure:
			todoServiceListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTodoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTodoServiceHandler struct{}

func (UnimplementedTodoServiceHandler) Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.TodoService.Create is not implemented"))
}

func (UnimplementedTodoServiceHandler) Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.TodoService.Read is not implemented"))
}

func (UnimplementedTodoServiceHandler) Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.TodoService.Update is not implemented"))
}

func (UnimplementedTodoServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.TodoService.Delete is not implemented"))
}

func (UnimplementedTodoServiceHandler) List(context.Context, *connect.BidiStream[v1.ListRequest, v1.ListResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.TodoService.List is not implemented"))
}
