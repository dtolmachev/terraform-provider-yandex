// Code generated by sdkgen. DO NOT EDIT.

// nolint
package iam

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

// YandexPassportUserAccountServiceClient is a iam.YandexPassportUserAccountServiceClient with
// lazy GRPC connection initialization.
type YandexPassportUserAccountServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ iam.YandexPassportUserAccountServiceClient = &YandexPassportUserAccountServiceClient{}

// GetByLogin implements iam.YandexPassportUserAccountServiceClient
func (c *YandexPassportUserAccountServiceClient) GetByLogin(ctx context.Context, in *iam.GetUserAccountByLoginRequest, opts ...grpc.CallOption) (*iam.UserAccount, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewYandexPassportUserAccountServiceClient(conn).GetByLogin(ctx, in, opts...)
}
