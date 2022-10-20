// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qanpb/collector.proto

package qanv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"

	_ "github.com/percona/pmm/api/inventorypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *CollectRequest) Validate() error {
	for _, item := range this.MetricsBucket {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MetricsBucket", err)
			}
		}
	}
	return nil
}

func (this *MetricsBucket) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

func (this *CollectResponse) Validate() error {
	return nil
}
