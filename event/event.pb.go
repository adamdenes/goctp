// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: event.proto

package event

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Kline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType                string  `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
	EventTime                int64   `protobuf:"varint,2,opt,name=eventTime,proto3" json:"eventTime,omitempty"`
	Symbol                   string  `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	StartTime                int64   `protobuf:"varint,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	CloseTime                int64   `protobuf:"varint,5,opt,name=closeTime,proto3" json:"closeTime,omitempty"`
	SymbolInner              string  `protobuf:"bytes,6,opt,name=symbolInner,proto3" json:"symbolInner,omitempty"`
	Interval                 string  `protobuf:"bytes,7,opt,name=interval,proto3" json:"interval,omitempty"`
	FirstTradeID             int64   `protobuf:"varint,8,opt,name=firstTradeID,proto3" json:"firstTradeID,omitempty"`
	LastTradeID              int64   `protobuf:"varint,9,opt,name=lastTradeID,proto3" json:"lastTradeID,omitempty"`
	OpenPrice                float64 `protobuf:"fixed64,10,opt,name=openPrice,proto3" json:"openPrice,omitempty"`
	ClosePrice               float64 `protobuf:"fixed64,11,opt,name=closePrice,proto3" json:"closePrice,omitempty"`
	HighPrice                float64 `protobuf:"fixed64,12,opt,name=highPrice,proto3" json:"highPrice,omitempty"`
	LowPrice                 float64 `protobuf:"fixed64,13,opt,name=lowPrice,proto3" json:"lowPrice,omitempty"`
	BaseAssetVolume          int64   `protobuf:"varint,14,opt,name=baseAssetVolume,proto3" json:"baseAssetVolume,omitempty"`
	NumberOfTrades           int64   `protobuf:"varint,15,opt,name=numberOfTrades,proto3" json:"numberOfTrades,omitempty"`
	IsKlineClosed            bool    `protobuf:"varint,16,opt,name=isKlineClosed,proto3" json:"isKlineClosed,omitempty"`
	QuoteAssetVolume         float64 `protobuf:"fixed64,17,opt,name=quoteAssetVolume,proto3" json:"quoteAssetVolume,omitempty"`
	TakerBuyBaseAssetVolume  int64   `protobuf:"varint,18,opt,name=takerBuyBaseAssetVolume,proto3" json:"takerBuyBaseAssetVolume,omitempty"`
	TakerBuyQuoteAssetVolume float64 `protobuf:"fixed64,19,opt,name=takerBuyQuoteAssetVolume,proto3" json:"takerBuyQuoteAssetVolume,omitempty"`
	Ignore                   int64   `protobuf:"varint,20,opt,name=ignore,proto3" json:"ignore,omitempty"`
}

func (x *Kline) Reset() {
	*x = Kline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kline) ProtoMessage() {}

func (x *Kline) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kline.ProtoReflect.Descriptor instead.
func (*Kline) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *Kline) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Kline) GetEventTime() int64 {
	if x != nil {
		return x.EventTime
	}
	return 0
}

func (x *Kline) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Kline) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Kline) GetCloseTime() int64 {
	if x != nil {
		return x.CloseTime
	}
	return 0
}

func (x *Kline) GetSymbolInner() string {
	if x != nil {
		return x.SymbolInner
	}
	return ""
}

func (x *Kline) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

func (x *Kline) GetFirstTradeID() int64 {
	if x != nil {
		return x.FirstTradeID
	}
	return 0
}

func (x *Kline) GetLastTradeID() int64 {
	if x != nil {
		return x.LastTradeID
	}
	return 0
}

func (x *Kline) GetOpenPrice() float64 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *Kline) GetClosePrice() float64 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

func (x *Kline) GetHighPrice() float64 {
	if x != nil {
		return x.HighPrice
	}
	return 0
}

func (x *Kline) GetLowPrice() float64 {
	if x != nil {
		return x.LowPrice
	}
	return 0
}

func (x *Kline) GetBaseAssetVolume() int64 {
	if x != nil {
		return x.BaseAssetVolume
	}
	return 0
}

func (x *Kline) GetNumberOfTrades() int64 {
	if x != nil {
		return x.NumberOfTrades
	}
	return 0
}

func (x *Kline) GetIsKlineClosed() bool {
	if x != nil {
		return x.IsKlineClosed
	}
	return false
}

func (x *Kline) GetQuoteAssetVolume() float64 {
	if x != nil {
		return x.QuoteAssetVolume
	}
	return 0
}

func (x *Kline) GetTakerBuyBaseAssetVolume() int64 {
	if x != nil {
		return x.TakerBuyBaseAssetVolume
	}
	return 0
}

func (x *Kline) GetTakerBuyQuoteAssetVolume() float64 {
	if x != nil {
		return x.TakerBuyQuoteAssetVolume
	}
	return 0
}

func (x *Kline) GetIgnore() int64 {
	if x != nil {
		return x.Ignore
	}
	return 0
}

type HistoricalKline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kline *Kline `protobuf:"bytes,1,opt,name=kline,proto3" json:"kline,omitempty"`
}

func (x *HistoricalKline) Reset() {
	*x = HistoricalKline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoricalKline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoricalKline) ProtoMessage() {}

func (x *HistoricalKline) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoricalKline.ProtoReflect.Descriptor instead.
func (*HistoricalKline) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *HistoricalKline) GetKline() *Kline {
	if x != nil {
		return x.Kline
	}
	return nil
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0xc5, 0x05, 0x0a, 0x05, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x44,
	0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x44, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x61,
	0x73, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x69, 0x73, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x38,
	0x0a, 0x17, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x75, 0x79, 0x42, 0x61, 0x73, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x17, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x75, 0x79, 0x42, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x74, 0x61, 0x6b, 0x65,
	0x72, 0x42, 0x75, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x18, 0x74, 0x61, 0x6b, 0x65,
	0x72, 0x42, 0x75, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x22, 0x35, 0x0a, 0x0f,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x05, 0x6b, 0x6c,
	0x69, 0x6e, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_event_proto_goTypes = []interface{}{
	(*Kline)(nil),           // 0: event.Kline
	(*HistoricalKline)(nil), // 1: event.HistoricalKline
}
var file_event_proto_depIdxs = []int32{
	0, // 0: event.HistoricalKline.kline:type_name -> event.Kline
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoricalKline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
