// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/department/department.proto

package department

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Department service

type DepartmentService interface {
	// 销售总额
	Sale(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type departmentService struct {
	c    client.Client
	name string
}

func NewDepartmentService(name string, c client.Client) DepartmentService {
	return &departmentService{
		c:    c,
		name: name,
	}
}

func (c *departmentService) Sale(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Department.Sale", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Department service

type DepartmentHandler interface {
	// 销售总额
	Sale(context.Context, *Request, *Response) error
}

func RegisterDepartmentHandler(s server.Server, hdlr DepartmentHandler, opts ...server.HandlerOption) error {
	type department interface {
		Sale(ctx context.Context, in *Request, out *Response) error
	}
	type Department struct {
		department
	}
	h := &departmentHandler{hdlr}
	return s.Handle(s.NewHandler(&Department{h}, opts...))
}

type departmentHandler struct {
	DepartmentHandler
}

func (h *departmentHandler) Sale(ctx context.Context, in *Request, out *Response) error {
	return h.DepartmentHandler.Sale(ctx, in, out)
}
