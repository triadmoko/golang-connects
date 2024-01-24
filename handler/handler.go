package handler

import (
	"context"
	"fmt"
	todov1 "golang-connects/gen/proto/todo/v1"
	"golang-connects/gen/proto/todo/v1/todov1connect"

	"connectrpc.com/connect"
)

type TodoServer struct {
	todov1connect.UnimplementedTodoServiceHandler
}

func (s *TodoServer) Create(ctx context.Context, req *connect.Request[todov1.CreateRequest]) (*connect.Response[todov1.CreateResponse], error) {
	fmt.Println("data terikirim")
	res := connect.NewResponse(&todov1.CreateResponse{
		Todo: req.Msg.Todo,
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
func (s *TodoServer) Read(context.Context, *connect.Request[todov1.ReadRequest]) (*connect.Response[todov1.ReadResponse], error) {
	res := connect.NewResponse(&todov1.ReadResponse{
		Todo: &todov1.Todo{
			Title:       "ini dari client",
			Id:          "1",
			Description: "World",
			Completed:   false,
			CreatedAt:   nil,
			UpdatedAt:   nil,
		},
	})
	return res, nil
}
func (s *TodoServer) Update(context.Context, *connect.Request[todov1.UpdateRequest]) (*connect.Response[todov1.UpdateResponse], error) {
	return nil, nil
}
func (s *TodoServer) Delete(context.Context, *connect.Request[todov1.DeleteRequest]) (*connect.Response[todov1.DeleteResponse], error) {
	return nil, nil
}

// func (s *TodoServer) List(ctx context.Context, stream *connect.ClientStream[todov1.ListRequest]) (*connect.Response[todov1.ListResponse], error) {

// 	var response *connect.Response[todov1.ListResponse]
// 	for stream.Receive() {
// 		res := connect.NewResponse(&todov1.ListResponse{
// 			Page:  1,
// 			Limit: 10,
// 			Todos: []*todov1.Todo{
// 				{
// 					Title:       "ini dari client",
// 					Id:          "1",
// 					Description: "World",
// 					Completed:   false,
// 					CreatedAt:   timestamppb.Now(),
// 					UpdatedAt:   timestamppb.Now(),
// 				},
// 			},
// 		})
// 		response.Header().Set("Greet-Version", "v1")
// 		response = res
// 	}
// 	if err := stream.Err(); err != nil {
// 		return nil, connect.NewError(connect.CodeUnknown, err)
// 	}
// 	fmt.Println("data terikirim")
// 	return response, nil
// }

func (s *TodoServer) List(ctx context.Context, in *connect.BidiStream[todov1.ListRequest, todov1.ListResponse]) error {
	for {
		_, err := in.Receive()
		if err != nil {
			return err
		}
		err = in.Send(&todov1.ListResponse{
			Page:  1,
			Total: 10,
		})
		if err != nil {
			return err
		}
	}
}
