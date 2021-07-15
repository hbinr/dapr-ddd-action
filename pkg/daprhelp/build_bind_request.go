package daprhelp

import "github.com/dapr/go-sdk/client"

// BuildBindingRequest build request such as follow:
//type InvokeBindingRequest struct {
//	// Name is name of binding to invoke.
//	Name string
//	// Operation is the name of the operation type for the binding to invoke
//	Operation string
//	// Data is the input bindings sent
//	Data []byte
//	// Metadata is the input binding metadata
//	Metadata map[string]string
//}
func BuildBindingRequest(name, operation, metaDataKey, metaDataVal string, data []byte) *client.InvokeBindingRequest {
	req := &client.InvokeBindingRequest{Metadata: map[string]string{}}
	req.Name = name
	req.Operation = operation
	req.Data = data
	req.Metadata[metaDataKey] = metaDataVal
	return req
}
