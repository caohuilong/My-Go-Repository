package mysql_api

import (
	"context"
	"database/sql"

	pb "api/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StaffServiceServer struct {
	db *sql.DB
}

func NewStaffServiceServer(db *sql.DB) pb.StaffServiceServer {
	return &StaffServiceServer{db}
}

func (s *StaffServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
