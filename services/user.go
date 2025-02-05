package services

import (
	"context"
	"grpc-practise/models"
	pb "grpc-practise/proto"
	repositories "grpc-practise/repository"
)

var repository = repositories.NewUserRepository()

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (service *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	resp, err := repository.GetUser(req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Id: resp.Id, Name: resp.Name, Location: resp.Location, Title: resp.Title}, nil
}

func (service *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser := models.User{Name: req.Name, Location: req.Location, Title: req.Title}
	_, err := repository.CreateUser(newUser)

	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Data: "User created successfully!"}, nil
}

func (service *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	newUser := models.User{Name: req.Name, Location: req.Location, Title: req.Title}
	_, err := repository.UpdateUser(req.Id, newUser)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{Data: "User updated successfully!"}, nil
}

func (service *UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := repository.DeleteUser(req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{Data: "User details deleted successfully!"}, nil
}

func (service *UserServiceServer) GetAllUsers(context.Context, *pb.Empty) (*pb.GetAllUsersResponse, error) {
	resp, err := repository.GetAllUsers()
	var users []*pb.UserResponse

	if err != nil {
		return nil, err
	}

	for _, v := range resp {
		var singleUser = &pb.UserResponse{
			Id:       v.Id,
			Name:     v.Name,
			Location: v.Location,
			Title:    v.Title,
		}
		users = append(users, singleUser)
	}

	return &pb.GetAllUsersResponse{Users: users}, nil
}
