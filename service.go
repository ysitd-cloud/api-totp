package totp

import (
	"context"

	"google.golang.org/grpc"
)

//go:generate protoc -I . -I vendor --gogo_out=plugins=grpc:. ./service.proto

type Client struct {
	Endpoint string
}

func (c *Client) connect(ctx context.Context) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, c.Endpoint, grpc.WithInsecure())
}

func (c *Client) IssueKey(ctx context.Context, in *IssueKeyRequest, opts ...grpc.CallOption) (*IssueKeyReply, error) {
	cc, err := c.connect(ctx)
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	out := new(IssueKeyReply)
	err = cc.Invoke(ctx, "/totp.Totp/issueKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateReply, error) {
	cc, err := c.connect(ctx)
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	out := new(ValidateReply)
	err = cc.Invoke(ctx, "/totp.Totp/validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) RecoverKey(ctx context.Context, in *RecoverRequest, opts ...grpc.CallOption) (*RecoverReply, error) {
	cc, err := c.connect(ctx)
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	out := new(RecoverReply)
	err = cc.Invoke(ctx, "/totp.Totp/recoverKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) RemoveKey(ctx context.Context, in *RemoveKeyRequest, opts ...grpc.CallOption) (*RemoveKeyReply, error) {
	cc, err := c.connect(ctx)
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	out := new(RemoveKeyReply)
	err = cc.Invoke(ctx, "/totp.Totp/removeKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
