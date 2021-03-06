// Code generated by protoc-gen-goext. DO NOT EDIT.

package iam

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

func (m *GetApiKeyRequest) SetApiKeyId(v string) {
	m.ApiKeyId = v
}

func (m *ListApiKeysRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *ListApiKeysRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListApiKeysRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListApiKeysResponse) SetApiKeys(v []*ApiKey) {
	m.ApiKeys = v
}

func (m *ListApiKeysResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateApiKeyRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateApiKeyRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateApiKeyResponse) SetApiKey(v *ApiKey) {
	m.ApiKey = v
}

func (m *CreateApiKeyResponse) SetSecret(v string) {
	m.Secret = v
}

func (m *DeleteApiKeyRequest) SetApiKeyId(v string) {
	m.ApiKeyId = v
}

func (m *DeleteApiKeyMetadata) SetApiKeyId(v string) {
	m.ApiKeyId = v
}

func (m *ListApiKeyOperationsRequest) SetApiKeyId(v string) {
	m.ApiKeyId = v
}

func (m *ListApiKeyOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListApiKeyOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListApiKeyOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListApiKeyOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
