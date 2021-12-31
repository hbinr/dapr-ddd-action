package daprhelp

import (
	"github.com/dapr-ddd-action/internal/pkg/constant"
	"github.com/dapr/go-sdk/client"
)

func BuildMySQLQueryBinding(name, sql string) *client.InvokeBindingRequest {
	req := &client.InvokeBindingRequest{Metadata: map[string]string{}}
	req.Name = name
	req.Operation = constant.MySQLOperationQuery
	req.Data = nil
	req.Metadata[constant.MySQLMetaDataKey] = sql
	return req
}

func BuildMySQLExecBinding(name, sql string) *client.InvokeBindingRequest {
	req := &client.InvokeBindingRequest{Metadata: map[string]string{}}
	req.Name = name
	req.Operation = constant.MySQLOperationExec
	req.Data = nil
	req.Metadata[constant.MySQLMetaDataKey] = sql
	return req
}
