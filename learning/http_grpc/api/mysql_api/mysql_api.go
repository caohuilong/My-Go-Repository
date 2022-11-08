package mysql_api

import (
	"context"
	"database/sql"
	"fmt"

	pb "http_grpc/api/proto"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// pb.StaffServiceServer的实现
type StaffServiceServer struct {
	db *sql.DB
}

// 创建staffService服务
func NewStaffServiceServer(db *sql.DB) pb.StaffServiceServer {
	return &StaffServiceServer{db}
}

//connect从池中返回SQL数据库连接
func (s *StaffServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// 创建新的员工信息事务
func (s *StaffServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "INSERT INTO ecf_staff(`name`, `phonenumber`, `position`) values (?, ?, ?)", req.Staffinfo.Name, req.Staffinfo.Phonenumber, req.Staffinfo.Position)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ecf_staff->"+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created StaffInfo->"+err.Error())
	}

	return &pb.CreateResponse{
		Id: id,
	}, nil
}

// 读取StaffInfo信息
func (s *StaffServiceServer) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// 按照ID查询ecf_staff
	rows, err := c.QueryContext(ctx, "SELECT `id`, `name`, `phonenumber`, `position` FROM ecf_staff WHERE `name`=?", req.Name)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ecf_staff->"+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ecf_staff->"+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ecf_staff with name='%s' is not found", req.Name))
	}

	var staff pb.StaffInfo
	if err := rows.Scan(&staff.Id, &staff.Name, &staff.Phonenumber, &staff.Position); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ecf_staff row->"+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ecf_staff rows with name='%s'", req.Name))
	}

	return &pb.ReadResponse{
		Staffinfo: &staff,
	}, nil
}

// 更新员工信息
func (s *StaffServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE ecf_staff SET `phonenumber`=?, `position`=? WHERE `name`=?", req.Staffinfo.Phonenumber, req.Staffinfo.Position, req.Staffinfo.Name)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "update request has invalid format-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ecf_staff with name='%s' is not found", req.Staffinfo.Name))
	}

	return &pb.UpdateResponse{
		Updated: rows,
	}, nil
}

// 删除员工信息
func (s *StaffServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM ecf_staff WHERE `name`=?", req.Name)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ecf_staff-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ecf_staff with name='%s' is not found", req.Name))
	}

	return &pb.DeleteResponse{
		Deleted: rows,
	}, nil
}

// 读取所有员工信息
func (s *StaffServiceServer) ReadAll(ctx context.Context, req *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	// 从池中获取sql连接
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// 获取StaffInfo列表
	rows, err := c.QueryContext(ctx, "SELECT `id`, `name`, `phonenumber`, `position` FROM ecf_staff")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ecf_staff-> "+err.Error())
	}

	list := []*pb.StaffInfo{}
	for rows.Next() {
		info := new(pb.StaffInfo)
		if err := rows.Scan(&info.Id, &info.Name, &info.Phonenumber, &info.Position); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ecf_staff row-> "+err.Error())
		}
		list = append(list, info)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ecf_staff-> "+err.Error())
	}

	return &pb.ReadAllResponse{
		Staffinfos: list,
	}, nil
}
