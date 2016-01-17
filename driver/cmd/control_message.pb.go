// Code generated by protoc-gen-go.
// source: control_message.proto
// DO NOT EDIT!

/*
Package cmd is a generated protocol buffer package.

It is generated from these files:
	control_message.proto

It has these top-level messages:
	ControlMessage
	NetChan
	ComputeResource
	StartRequest
	StartResponse
	StopRequest
	StopResponse
	GetStatusRequest
	ChannelStatus
	GetStatusResponse
	DeleteDatasetShardRequest
	DeleteDatasetShardResponse
	LocalStatusReportRequest
	LocalStatusReportResponse
*/
package cmd

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ControlMessage_Type int32

const (
	ControlMessage_NoOp                       ControlMessage_Type = 1
	ControlMessage_StartRequest               ControlMessage_Type = 2
	ControlMessage_StartResponse              ControlMessage_Type = 3
	ControlMessage_StopRequest                ControlMessage_Type = 4
	ControlMessage_StopResponse               ControlMessage_Type = 5
	ControlMessage_GetStatusRequest           ControlMessage_Type = 6
	ControlMessage_GetStatusResponse          ControlMessage_Type = 7
	ControlMessage_DeleteDatasetShardRequest  ControlMessage_Type = 8
	ControlMessage_DeleteDatasetShardResponse ControlMessage_Type = 9
	ControlMessage_LocalStatusReportRequest   ControlMessage_Type = 10
	ControlMessage_LocalStatusReportResponse  ControlMessage_Type = 11
)

var ControlMessage_Type_name = map[int32]string{
	1:  "NoOp",
	2:  "StartRequest",
	3:  "StartResponse",
	4:  "StopRequest",
	5:  "StopResponse",
	6:  "GetStatusRequest",
	7:  "GetStatusResponse",
	8:  "DeleteDatasetShardRequest",
	9:  "DeleteDatasetShardResponse",
	10: "LocalStatusReportRequest",
	11: "LocalStatusReportResponse",
}
var ControlMessage_Type_value = map[string]int32{
	"NoOp":                       1,
	"StartRequest":               2,
	"StartResponse":              3,
	"StopRequest":                4,
	"StopResponse":               5,
	"GetStatusRequest":           6,
	"GetStatusResponse":          7,
	"DeleteDatasetShardRequest":  8,
	"DeleteDatasetShardResponse": 9,
	"LocalStatusReportRequest":   10,
	"LocalStatusReportResponse":  11,
}

func (x ControlMessage_Type) Enum() *ControlMessage_Type {
	p := new(ControlMessage_Type)
	*p = x
	return p
}
func (x ControlMessage_Type) String() string {
	return proto.EnumName(ControlMessage_Type_name, int32(x))
}
func (x *ControlMessage_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ControlMessage_Type_value, data, "ControlMessage_Type")
	if err != nil {
		return err
	}
	*x = ControlMessage_Type(value)
	return nil
}

type ControlMessage struct {
	Type                       *ControlMessage_Type        `protobuf:"varint,1,req,name=type,enum=cmd.ControlMessage_Type" json:"type,omitempty"`
	StartRequest               *StartRequest               `protobuf:"bytes,2,opt,name=startRequest" json:"startRequest,omitempty"`
	StartResponse              *StartResponse              `protobuf:"bytes,3,opt,name=startResponse" json:"startResponse,omitempty"`
	StopRequest                *StopRequest                `protobuf:"bytes,4,opt,name=stopRequest" json:"stopRequest,omitempty"`
	StopResponse               *StopResponse               `protobuf:"bytes,5,opt,name=stopResponse" json:"stopResponse,omitempty"`
	GetStatusRequest           *GetStatusRequest           `protobuf:"bytes,6,opt,name=getStatusRequest" json:"getStatusRequest,omitempty"`
	GetStatusResponse          *GetStatusResponse          `protobuf:"bytes,7,opt,name=getStatusResponse" json:"getStatusResponse,omitempty"`
	DeleteDatasetShardRequest  *DeleteDatasetShardRequest  `protobuf:"bytes,8,opt,name=deleteDatasetShardRequest" json:"deleteDatasetShardRequest,omitempty"`
	DeleteDatasetShardResponse *DeleteDatasetShardResponse `protobuf:"bytes,9,opt,name=deleteDatasetShardResponse" json:"deleteDatasetShardResponse,omitempty"`
	LocalStatusReportRequest   *LocalStatusReportRequest   `protobuf:"bytes,10,opt,name=localStatusReportRequest" json:"localStatusReportRequest,omitempty"`
	LocalStatusReportResponse  *LocalStatusReportResponse  `protobuf:"bytes,11,opt,name=localStatusReportResponse" json:"localStatusReportResponse,omitempty"`
	XXX_unrecognized           []byte                      `json:"-"`
}

func (m *ControlMessage) Reset()         { *m = ControlMessage{} }
func (m *ControlMessage) String() string { return proto.CompactTextString(m) }
func (*ControlMessage) ProtoMessage()    {}

func (m *ControlMessage) GetType() ControlMessage_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ControlMessage_NoOp
}

func (m *ControlMessage) GetStartRequest() *StartRequest {
	if m != nil {
		return m.StartRequest
	}
	return nil
}

func (m *ControlMessage) GetStartResponse() *StartResponse {
	if m != nil {
		return m.StartResponse
	}
	return nil
}

func (m *ControlMessage) GetStopRequest() *StopRequest {
	if m != nil {
		return m.StopRequest
	}
	return nil
}

func (m *ControlMessage) GetStopResponse() *StopResponse {
	if m != nil {
		return m.StopResponse
	}
	return nil
}

func (m *ControlMessage) GetGetStatusRequest() *GetStatusRequest {
	if m != nil {
		return m.GetStatusRequest
	}
	return nil
}

func (m *ControlMessage) GetGetStatusResponse() *GetStatusResponse {
	if m != nil {
		return m.GetStatusResponse
	}
	return nil
}

func (m *ControlMessage) GetDeleteDatasetShardRequest() *DeleteDatasetShardRequest {
	if m != nil {
		return m.DeleteDatasetShardRequest
	}
	return nil
}

func (m *ControlMessage) GetDeleteDatasetShardResponse() *DeleteDatasetShardResponse {
	if m != nil {
		return m.DeleteDatasetShardResponse
	}
	return nil
}

func (m *ControlMessage) GetLocalStatusReportRequest() *LocalStatusReportRequest {
	if m != nil {
		return m.LocalStatusReportRequest
	}
	return nil
}

func (m *ControlMessage) GetLocalStatusReportResponse() *LocalStatusReportResponse {
	if m != nil {
		return m.LocalStatusReportResponse
	}
	return nil
}

type NetChan struct {
	Server           *string `protobuf:"bytes,1,req,name=server" json:"server,omitempty"`
	Port             *int32  `protobuf:"varint,2,req,name=port" json:"port,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NetChan) Reset()         { *m = NetChan{} }
func (m *NetChan) String() string { return proto.CompactTextString(m) }
func (*NetChan) ProtoMessage()    {}

func (m *NetChan) GetServer() string {
	if m != nil && m.Server != nil {
		return *m.Server
	}
	return ""
}

func (m *NetChan) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

type ComputeResource struct {
	CpuCount         *int32 `protobuf:"varint,1,req,name=cpuCount" json:"cpuCount,omitempty"`
	CpuLevel         *int32 `protobuf:"varint,2,req,name=cpuLevel" json:"cpuLevel,omitempty"`
	Memory           *int32 `protobuf:"varint,3,req,name=memory" json:"memory,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ComputeResource) Reset()         { *m = ComputeResource{} }
func (m *ComputeResource) String() string { return proto.CompactTextString(m) }
func (*ComputeResource) ProtoMessage()    {}

func (m *ComputeResource) GetCpuCount() int32 {
	if m != nil && m.CpuCount != nil {
		return *m.CpuCount
	}
	return 0
}

func (m *ComputeResource) GetCpuLevel() int32 {
	if m != nil && m.CpuLevel != nil {
		return *m.CpuLevel
	}
	return 0
}

func (m *ComputeResource) GetMemory() int32 {
	if m != nil && m.Memory != nil {
		return *m.Memory
	}
	return 0
}

type StartRequest struct {
	Path             *string          `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Args             []string         `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	Envs             []string         `protobuf:"bytes,3,rep,name=envs" json:"envs,omitempty"`
	Dir              *string          `protobuf:"bytes,4,req,name=dir" json:"dir,omitempty"`
	Resource         *ComputeResource `protobuf:"bytes,5,req,name=resource" json:"resource,omitempty"`
	HashCode         *uint32          `protobuf:"varint,6,req,name=hashCode" json:"hashCode,omitempty"`
	Host             *string          `protobuf:"bytes,7,opt,name=host" json:"host,omitempty"`
	Port             *int32           `protobuf:"varint,8,opt,name=port" json:"port,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}

func (m *StartRequest) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *StartRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *StartRequest) GetEnvs() []string {
	if m != nil {
		return m.Envs
	}
	return nil
}

func (m *StartRequest) GetDir() string {
	if m != nil && m.Dir != nil {
		return *m.Dir
	}
	return ""
}

func (m *StartRequest) GetResource() *ComputeResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *StartRequest) GetHashCode() uint32 {
	if m != nil && m.HashCode != nil {
		return *m.HashCode
	}
	return 0
}

func (m *StartRequest) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *StartRequest) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

type StartResponse struct {
	Pid              *int32     `protobuf:"varint,1,req,name=pid" json:"pid,omitempty"`
	Error            *string    `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Stderr           *NetChan   `protobuf:"bytes,3,opt,name=stderr" json:"stderr,omitempty"`
	Outputs          []*NetChan `protobuf:"bytes,4,rep,name=outputs" json:"outputs,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}

func (m *StartResponse) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *StartResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *StartResponse) GetStderr() *NetChan {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *StartResponse) GetOutputs() []*NetChan {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type StopRequest struct {
	StartRequestHash *uint32 `protobuf:"varint,1,req,name=startRequestHash" json:"startRequestHash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}

func (m *StopRequest) GetStartRequestHash() uint32 {
	if m != nil && m.StartRequestHash != nil {
		return *m.StartRequestHash
	}
	return 0
}

type StopResponse struct {
	StartRequestHash *uint32 `protobuf:"varint,1,req,name=startRequestHash" json:"startRequestHash,omitempty"`
	Error            *string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}

func (m *StopResponse) GetStartRequestHash() uint32 {
	if m != nil && m.StartRequestHash != nil {
		return *m.StartRequestHash
	}
	return 0
}

func (m *StopResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

type GetStatusRequest struct {
	StartRequestHash *uint32 `protobuf:"varint,1,req,name=startRequestHash" json:"startRequestHash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetStatusRequest) Reset()         { *m = GetStatusRequest{} }
func (m *GetStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatusRequest) ProtoMessage()    {}

func (m *GetStatusRequest) GetStartRequestHash() uint32 {
	if m != nil && m.StartRequestHash != nil {
		return *m.StartRequestHash
	}
	return 0
}

type ChannelStatus struct {
	Length           *int64 `protobuf:"varint,1,req,name=length" json:"length,omitempty"`
	StartTime        *int64 `protobuf:"varint,2,opt,name=startTime" json:"startTime,omitempty"`
	StopTime         *int64 `protobuf:"varint,3,opt,name=stopTime" json:"stopTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChannelStatus) Reset()         { *m = ChannelStatus{} }
func (m *ChannelStatus) String() string { return proto.CompactTextString(m) }
func (*ChannelStatus) ProtoMessage()    {}

func (m *ChannelStatus) GetLength() int64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *ChannelStatus) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *ChannelStatus) GetStopTime() int64 {
	if m != nil && m.StopTime != nil {
		return *m.StopTime
	}
	return 0
}

type GetStatusResponse struct {
	StartRequestHash *uint32          `protobuf:"varint,1,req,name=startRequestHash" json:"startRequestHash,omitempty"`
	Error            *string          `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	InputStatuses    []*ChannelStatus `protobuf:"bytes,3,rep,name=inputStatuses" json:"inputStatuses,omitempty"`
	OutputStatus     *ChannelStatus   `protobuf:"bytes,4,opt,name=outputStatus" json:"outputStatus,omitempty"`
	RequestTime      *int64           `protobuf:"varint,5,opt,name=requestTime" json:"requestTime,omitempty"`
	StartTime        *int64           `protobuf:"varint,6,opt,name=startTime" json:"startTime,omitempty"`
	StopTime         *int64           `protobuf:"varint,7,opt,name=stopTime" json:"stopTime,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GetStatusResponse) Reset()         { *m = GetStatusResponse{} }
func (m *GetStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatusResponse) ProtoMessage()    {}

func (m *GetStatusResponse) GetStartRequestHash() uint32 {
	if m != nil && m.StartRequestHash != nil {
		return *m.StartRequestHash
	}
	return 0
}

func (m *GetStatusResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *GetStatusResponse) GetInputStatuses() []*ChannelStatus {
	if m != nil {
		return m.InputStatuses
	}
	return nil
}

func (m *GetStatusResponse) GetOutputStatus() *ChannelStatus {
	if m != nil {
		return m.OutputStatus
	}
	return nil
}

func (m *GetStatusResponse) GetRequestTime() int64 {
	if m != nil && m.RequestTime != nil {
		return *m.RequestTime
	}
	return 0
}

func (m *GetStatusResponse) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *GetStatusResponse) GetStopTime() int64 {
	if m != nil && m.StopTime != nil {
		return *m.StopTime
	}
	return 0
}

type DeleteDatasetShardRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteDatasetShardRequest) Reset()         { *m = DeleteDatasetShardRequest{} }
func (m *DeleteDatasetShardRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDatasetShardRequest) ProtoMessage()    {}

func (m *DeleteDatasetShardRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type DeleteDatasetShardResponse struct {
	Error            *string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteDatasetShardResponse) Reset()         { *m = DeleteDatasetShardResponse{} }
func (m *DeleteDatasetShardResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteDatasetShardResponse) ProtoMessage()    {}

func (m *DeleteDatasetShardResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

type LocalStatusReportRequest struct {
	StartRequestHash *uint32          `protobuf:"varint,1,req,name=startRequestHash" json:"startRequestHash,omitempty"`
	Error            *string          `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	InputStatuses    []*ChannelStatus `protobuf:"bytes,3,rep,name=inputStatuses" json:"inputStatuses,omitempty"`
	OutputStatus     *ChannelStatus   `protobuf:"bytes,4,opt,name=outputStatus" json:"outputStatus,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *LocalStatusReportRequest) Reset()         { *m = LocalStatusReportRequest{} }
func (m *LocalStatusReportRequest) String() string { return proto.CompactTextString(m) }
func (*LocalStatusReportRequest) ProtoMessage()    {}

func (m *LocalStatusReportRequest) GetStartRequestHash() uint32 {
	if m != nil && m.StartRequestHash != nil {
		return *m.StartRequestHash
	}
	return 0
}

func (m *LocalStatusReportRequest) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *LocalStatusReportRequest) GetInputStatuses() []*ChannelStatus {
	if m != nil {
		return m.InputStatuses
	}
	return nil
}

func (m *LocalStatusReportRequest) GetOutputStatus() *ChannelStatus {
	if m != nil {
		return m.OutputStatus
	}
	return nil
}

type LocalStatusReportResponse struct {
	Error            *string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LocalStatusReportResponse) Reset()         { *m = LocalStatusReportResponse{} }
func (m *LocalStatusReportResponse) String() string { return proto.CompactTextString(m) }
func (*LocalStatusReportResponse) ProtoMessage()    {}

func (m *LocalStatusReportResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("cmd.ControlMessage_Type", ControlMessage_Type_name, ControlMessage_Type_value)
}
