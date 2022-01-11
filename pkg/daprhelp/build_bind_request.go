package daprhelp

import (
	"github.com/dapr/go-sdk/client"
)

const MySQLOperationQuery = "query"
const MySQLOperationExec = "exec"
const MySQLMetaDataKey = "sql"

func BuildMySQLQueryBinding(name, sql string) *client.InvokeBindingRequest {
	req := &client.InvokeBindingRequest{Metadata: map[string]string{}}
	req.Name = name
	req.Operation = MySQLOperationQuery
	req.Data = nil
	req.Metadata[MySQLMetaDataKey] = sql
	return req
}

func BuildMySQLExecBinding(name, sql string) *client.InvokeBindingRequest {
	req := &client.InvokeBindingRequest{Metadata: map[string]string{}}
	req.Name = name
	req.Operation = MySQLOperationExec
	req.Data = nil
	req.Metadata[MySQLMetaDataKey] = sql
	return req
}
