// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: carbonapi.proto

/*
	Package carbonapipb is a generated protocol buffer package.

	It is generated from these files:
		carbonapi.proto

	It has these top-level messages:
		AccessLogDetails
*/
package carbonapipb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AccessLogDetails struct {
	Handler                       string   `protobuf:"bytes,1,opt,name=handler,proto3" json:"handler,omitempty"`
	CarbonapiUuid                 string   `protobuf:"bytes,2,opt,name=carbonapi_uuid,json=carbonapiUuid,proto3" json:"carbonapi_uuid,omitempty"`
	Username                      string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Url                           string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	PeerIp                        string   `protobuf:"bytes,5,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	PeerPort                      string   `protobuf:"bytes,6,opt,name=peer_port,json=peerPort,proto3" json:"peer_port,omitempty"`
	Host                          string   `protobuf:"bytes,7,opt,name=host,proto3" json:"host,omitempty"`
	Referer                       string   `protobuf:"bytes,8,opt,name=referer,proto3" json:"referer,omitempty"`
	Format                        string   `protobuf:"bytes,9,opt,name=format,proto3" json:"format,omitempty"`
	UseCache                      bool     `protobuf:"varint,10,opt,name=use_cache,json=useCache,proto3" json:"use_cache,omitempty"`
	Targets                       []string `protobuf:"bytes,11,rep,name=targets" json:"targets,omitempty"`
	CacheTimeout                  int32    `protobuf:"varint,12,opt,name=cache_timeout,json=cacheTimeout,proto3" json:"cache_timeout,omitempty"`
	Metrics                       []string `protobuf:"bytes,13,rep,name=metrics" json:"metrics,omitempty"`
	HaveNonFatalErrors            bool     `protobuf:"varint,14,opt,name=have_non_fatal_errors,json=haveNonFatalErrors,proto3" json:"have_non_fatal_errors,omitempty"`
	Runtime                       float64  `protobuf:"fixed64,15,opt,name=runtime,proto3" json:"runtime,omitempty"`
	HttpCode                      int32    `protobuf:"varint,16,opt,name=http_code,json=httpCode,proto3" json:"http_code,omitempty"`
	CarbonzipperResponseSizeBytes int64    `protobuf:"varint,17,opt,name=carbonzipper_response_size_bytes,json=carbonzipperResponseSizeBytes,proto3" json:"carbonzipper_response_size_bytes,omitempty"`
	CarbonapiResponseSizeBytes    int64    `protobuf:"varint,18,opt,name=carbonapi_response_size_bytes,json=carbonapiResponseSizeBytes,proto3" json:"carbonapi_response_size_bytes,omitempty"`
	Reason                        string   `protobuf:"bytes,19,opt,name=reason,proto3" json:"reason,omitempty"`
	SendGlobs                     bool     `protobuf:"varint,20,opt,name=send_globs,json=sendGlobs,proto3" json:"send_globs,omitempty"`
	From                          int64   `protobuf:"varint,21,opt,name=from,proto3" json:"from,omitempty"`
	Until                         int64   `protobuf:"varint,22,opt,name=until,proto3" json:"until,omitempty"`
	Tz                            string   `protobuf:"bytes,23,opt,name=tz,proto3" json:"tz,omitempty"`
	FromRaw                       string   `protobuf:"bytes,24,opt,name=from_raw,json=fromRaw,proto3" json:"from_raw,omitempty"`
	UntilRaw                      string   `protobuf:"bytes,25,opt,name=until_raw,json=untilRaw,proto3" json:"until_raw,omitempty"`
	Uri                           string   `protobuf:"bytes,26,opt,name=uri,proto3" json:"uri,omitempty"`
	FromCache                     bool     `protobuf:"varint,27,opt,name=from_cache,json=fromCache,proto3" json:"from_cache,omitempty"`
	ZipperRequests                int64    `protobuf:"varint,28,opt,name=zipper_requests,json=zipperRequests,proto3" json:"zipper_requests,omitempty"`
}

func (m *AccessLogDetails) Reset()                    { *m = AccessLogDetails{} }
func (m *AccessLogDetails) String() string            { return proto.CompactTextString(m) }
func (*AccessLogDetails) ProtoMessage()               {}
func (*AccessLogDetails) Descriptor() ([]byte, []int) { return fileDescriptorCarbonapi, []int{0} }

func (m *AccessLogDetails) GetHandler() string {
	if m != nil {
		return m.Handler
	}
	return ""
}

func (m *AccessLogDetails) GetCarbonapiUuid() string {
	if m != nil {
		return m.CarbonapiUuid
	}
	return ""
}

func (m *AccessLogDetails) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccessLogDetails) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AccessLogDetails) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

func (m *AccessLogDetails) GetPeerPort() string {
	if m != nil {
		return m.PeerPort
	}
	return ""
}

func (m *AccessLogDetails) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AccessLogDetails) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

func (m *AccessLogDetails) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *AccessLogDetails) GetUseCache() bool {
	if m != nil {
		return m.UseCache
	}
	return false
}

func (m *AccessLogDetails) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *AccessLogDetails) GetCacheTimeout() int32 {
	if m != nil {
		return m.CacheTimeout
	}
	return 0
}

func (m *AccessLogDetails) GetMetrics() []string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *AccessLogDetails) GetHaveNonFatalErrors() bool {
	if m != nil {
		return m.HaveNonFatalErrors
	}
	return false
}

func (m *AccessLogDetails) GetRuntime() float64 {
	if m != nil {
		return m.Runtime
	}
	return 0
}

func (m *AccessLogDetails) GetHttpCode() int32 {
	if m != nil {
		return m.HttpCode
	}
	return 0
}

func (m *AccessLogDetails) GetCarbonzipperResponseSizeBytes() int64 {
	if m != nil {
		return m.CarbonzipperResponseSizeBytes
	}
	return 0
}

func (m *AccessLogDetails) GetCarbonapiResponseSizeBytes() int64 {
	if m != nil {
		return m.CarbonapiResponseSizeBytes
	}
	return 0
}

func (m *AccessLogDetails) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *AccessLogDetails) GetSendGlobs() bool {
	if m != nil {
		return m.SendGlobs
	}
	return false
}

func (m *AccessLogDetails) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *AccessLogDetails) GetUntil() int64 {
	if m != nil {
		return m.Until
	}
	return 0
}

func (m *AccessLogDetails) GetTz() string {
	if m != nil {
		return m.Tz
	}
	return ""
}

func (m *AccessLogDetails) GetFromRaw() string {
	if m != nil {
		return m.FromRaw
	}
	return ""
}

func (m *AccessLogDetails) GetUntilRaw() string {
	if m != nil {
		return m.UntilRaw
	}
	return ""
}

func (m *AccessLogDetails) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *AccessLogDetails) GetFromCache() bool {
	if m != nil {
		return m.FromCache
	}
	return false
}

func (m *AccessLogDetails) GetZipperRequests() int64 {
	if m != nil {
		return m.ZipperRequests
	}
	return 0
}

func init() {
	proto.RegisterType((*AccessLogDetails)(nil), "carbonapipb.AccessLogDetails")
}
func (m *AccessLogDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessLogDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Handler) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.Handler)))
		i += copy(dAtA[i:], m.Handler)
	}
	if len(m.CarbonapiUuid) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.CarbonapiUuid)))
		i += copy(dAtA[i:], m.CarbonapiUuid)
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Url) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.PeerIp) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.PeerIp)))
		i += copy(dAtA[i:], m.PeerIp)
	}
	if len(m.PeerPort) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.PeerPort)))
		i += copy(dAtA[i:], m.PeerPort)
	}
	if len(m.Host) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.Host)))
		i += copy(dAtA[i:], m.Host)
	}
	if len(m.Referer) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.Referer)))
		i += copy(dAtA[i:], m.Referer)
	}
	if len(m.Format) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.Format)))
		i += copy(dAtA[i:], m.Format)
	}
	if m.UseCache {
		dAtA[i] = 0x50
		i++
		if m.UseCache {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Targets) > 0 {
		for _, s := range m.Targets {
			dAtA[i] = 0x5a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.CacheTimeout != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(m.CacheTimeout))
	}
	if len(m.Metrics) > 0 {
		for _, s := range m.Metrics {
			dAtA[i] = 0x6a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.HaveNonFatalErrors {
		dAtA[i] = 0x70
		i++
		if m.HaveNonFatalErrors {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Runtime != 0 {
		dAtA[i] = 0x79
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Runtime))))
		i += 8
	}
	if m.HttpCode != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(m.HttpCode))
	}
	if m.CarbonzipperResponseSizeBytes != 0 {
		dAtA[i] = 0x88
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(m.CarbonzipperResponseSizeBytes))
	}
	if m.CarbonapiResponseSizeBytes != 0 {
		dAtA[i] = 0x90
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(m.CarbonapiResponseSizeBytes))
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	if m.SendGlobs {
		dAtA[i] = 0xa0
		i++
		dAtA[i] = 0x1
		i++
		if m.SendGlobs {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.From != 0 {
		dAtA[i] = 0xa8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(m.From))
	}
	if m.Until != 0 {
		dAtA[i] = 0xb0
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(m.Until))
	}
	if len(m.Tz) > 0 {
		dAtA[i] = 0xba
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.Tz)))
		i += copy(dAtA[i:], m.Tz)
	}
	if len(m.FromRaw) > 0 {
		dAtA[i] = 0xc2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.FromRaw)))
		i += copy(dAtA[i:], m.FromRaw)
	}
	if len(m.UntilRaw) > 0 {
		dAtA[i] = 0xca
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.UntilRaw)))
		i += copy(dAtA[i:], m.UntilRaw)
	}
	if len(m.Uri) > 0 {
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(len(m.Uri)))
		i += copy(dAtA[i:], m.Uri)
	}
	if m.FromCache {
		dAtA[i] = 0xd8
		i++
		dAtA[i] = 0x1
		i++
		if m.FromCache {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ZipperRequests != 0 {
		dAtA[i] = 0xe0
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCarbonapi(dAtA, i, uint64(m.ZipperRequests))
	}
	return i, nil
}

func encodeVarintCarbonapi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AccessLogDetails) Size() (n int) {
	var l int
	_ = l
	l = len(m.Handler)
	if l > 0 {
		n += 1 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.CarbonapiUuid)
	if l > 0 {
		n += 1 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.PeerIp)
	if l > 0 {
		n += 1 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.PeerPort)
	if l > 0 {
		n += 1 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.Referer)
	if l > 0 {
		n += 1 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.Format)
	if l > 0 {
		n += 1 + l + sovCarbonapi(uint64(l))
	}
	if m.UseCache {
		n += 2
	}
	if len(m.Targets) > 0 {
		for _, s := range m.Targets {
			l = len(s)
			n += 1 + l + sovCarbonapi(uint64(l))
		}
	}
	if m.CacheTimeout != 0 {
		n += 1 + sovCarbonapi(uint64(m.CacheTimeout))
	}
	if len(m.Metrics) > 0 {
		for _, s := range m.Metrics {
			l = len(s)
			n += 1 + l + sovCarbonapi(uint64(l))
		}
	}
	if m.HaveNonFatalErrors {
		n += 2
	}
	if m.Runtime != 0 {
		n += 9
	}
	if m.HttpCode != 0 {
		n += 2 + sovCarbonapi(uint64(m.HttpCode))
	}
	if m.CarbonzipperResponseSizeBytes != 0 {
		n += 2 + sovCarbonapi(uint64(m.CarbonzipperResponseSizeBytes))
	}
	if m.CarbonapiResponseSizeBytes != 0 {
		n += 2 + sovCarbonapi(uint64(m.CarbonapiResponseSizeBytes))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 2 + l + sovCarbonapi(uint64(l))
	}
	if m.SendGlobs {
		n += 3
	}
	if m.From != 0 {
		n += 2 + sovCarbonapi(uint64(m.From))
	}
	if m.Until != 0 {
		n += 2 + sovCarbonapi(uint64(m.Until))
	}
	l = len(m.Tz)
	if l > 0 {
		n += 2 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.FromRaw)
	if l > 0 {
		n += 2 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.UntilRaw)
	if l > 0 {
		n += 2 + l + sovCarbonapi(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 2 + l + sovCarbonapi(uint64(l))
	}
	if m.FromCache {
		n += 3
	}
	if m.ZipperRequests != 0 {
		n += 2 + sovCarbonapi(uint64(m.ZipperRequests))
	}
	return n
}

func sovCarbonapi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCarbonapi(x uint64) (n int) {
	return sovCarbonapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessLogDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCarbonapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AccessLogDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessLogDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CarbonapiUuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CarbonapiUuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerPort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerPort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Referer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Referer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Format = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseCache", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UseCache = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Targets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Targets = append(m.Targets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheTimeout", wireType)
			}
			m.CacheTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CacheTimeout |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metrics = append(m.Metrics, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HaveNonFatalErrors", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HaveNonFatalErrors = bool(v != 0)
		case 15:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Runtime", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Runtime = float64(math.Float64frombits(v))
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpCode", wireType)
			}
			m.HttpCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HttpCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CarbonzipperResponseSizeBytes", wireType)
			}
			m.CarbonzipperResponseSizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CarbonzipperResponseSizeBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CarbonapiResponseSizeBytes", wireType)
			}
			m.CarbonapiResponseSizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CarbonapiResponseSizeBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendGlobs", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SendGlobs = bool(v != 0)
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			m.From = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.From |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Until", wireType)
			}
			m.Until = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Until |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 23:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tz", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tz = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 24:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromRaw", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromRaw = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 25:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UntilRaw", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UntilRaw = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 26:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCarbonapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 27:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromCache", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FromCache = bool(v != 0)
		case 28:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZipperRequests", wireType)
			}
			m.ZipperRequests = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ZipperRequests |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCarbonapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCarbonapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCarbonapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCarbonapi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCarbonapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCarbonapi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCarbonapi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCarbonapi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCarbonapi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCarbonapi   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("carbonapi.proto", fileDescriptorCarbonapi) }

var fileDescriptorCarbonapi = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0x6e, 0x9b, 0xee, 0xba, 0xdd, 0x6d, 0x31, 0xfd, 0x33, 0x6d, 0x69, 0x15, 0x81,
	0x10, 0x39, 0x21, 0x21, 0x9e, 0xa0, 0x2d, 0x50, 0x21, 0x21, 0x84, 0x02, 0x9c, 0x2d, 0x6f, 0x32,
	0xdb, 0xb5, 0x94, 0x8d, 0x83, 0xed, 0x50, 0xb1, 0x4f, 0xc8, 0x91, 0x47, 0x40, 0x3d, 0xf3, 0x10,
	0x68, 0xc6, 0xbb, 0x01, 0x09, 0x6e, 0xf9, 0x7e, 0x33, 0xf3, 0x79, 0xec, 0x99, 0x88, 0xbd, 0x52,
	0xbb, 0xa9, 0x6d, 0x74, 0x6b, 0x9e, 0xb7, 0xce, 0x06, 0x2b, 0x77, 0x7a, 0xd0, 0x4e, 0x1f, 0xff,
	0x4a, 0xc5, 0xfe, 0x65, 0x59, 0xa2, 0xf7, 0xef, 0xec, 0xed, 0x2b, 0x0c, 0xda, 0xd4, 0x5e, 0x82,
	0xd8, 0x9e, 0xeb, 0xa6, 0xaa, 0xd1, 0x41, 0x92, 0x25, 0xf9, 0xa8, 0x58, 0x4b, 0xf9, 0x54, 0x4c,
	0xfa, 0x6a, 0xd5, 0x75, 0xa6, 0x82, 0x0d, 0x4e, 0x18, 0xf7, 0xf4, 0x73, 0x67, 0x2a, 0x79, 0x2a,
	0x86, 0x9d, 0x47, 0xd7, 0xe8, 0x05, 0xc2, 0x80, 0x13, 0x7a, 0x2d, 0xf7, 0xc5, 0xa0, 0x73, 0x35,
	0x6c, 0x32, 0xa6, 0x4f, 0x79, 0x2c, 0xb6, 0x5b, 0x44, 0xa7, 0x4c, 0x0b, 0x5b, 0x4c, 0x53, 0x92,
	0x6f, 0x5b, 0x79, 0x26, 0x46, 0x1c, 0x68, 0xad, 0x0b, 0x90, 0x46, 0x1f, 0x02, 0x1f, 0xac, 0x0b,
	0x52, 0x8a, 0xcd, 0xb9, 0xf5, 0x01, 0xb6, 0x99, 0xf3, 0x37, 0x35, 0xee, 0x70, 0x86, 0x0e, 0x1d,
	0x0c, 0x63, 0xe3, 0x2b, 0x29, 0x8f, 0x44, 0x3a, 0xb3, 0x6e, 0xa1, 0x03, 0x8c, 0xe2, 0x11, 0x51,
	0xd1, 0x11, 0x9d, 0x47, 0x55, 0xea, 0x72, 0x8e, 0x20, 0xb2, 0x24, 0x1f, 0x72, 0xab, 0xd7, 0xa4,
	0xc9, 0x2e, 0x68, 0x77, 0x8b, 0xc1, 0xc3, 0x4e, 0x36, 0x20, 0xbb, 0x95, 0x94, 0x4f, 0xc4, 0x98,
	0x4b, 0x54, 0x30, 0x0b, 0xb4, 0x5d, 0x80, 0xdd, 0x2c, 0xc9, 0xb7, 0x8a, 0x5d, 0x86, 0x9f, 0x22,
	0xa3, 0xf2, 0x05, 0x06, 0x67, 0x4a, 0x0f, 0xe3, 0x58, 0xbe, 0x92, 0xf2, 0x85, 0x38, 0x9c, 0xeb,
	0xaf, 0xa8, 0x1a, 0xdb, 0xa8, 0x99, 0x0e, 0xba, 0x56, 0xe8, 0x9c, 0x75, 0x1e, 0x26, 0xdc, 0x81,
	0xa4, 0xe0, 0x7b, 0xdb, 0xbc, 0xa1, 0xd0, 0x6b, 0x8e, 0xf0, 0xd5, 0xba, 0x86, 0x8e, 0x83, 0xbd,
	0x2c, 0xc9, 0x93, 0x62, 0x2d, 0xe9, 0x0a, 0xf3, 0x10, 0x5a, 0x55, 0xda, 0x0a, 0x61, 0x9f, 0xfb,
	0x18, 0x12, 0xb8, 0xb6, 0x15, 0xca, 0x1b, 0x91, 0xc5, 0xd1, 0x2c, 0x4d, 0xdb, 0xa2, 0x53, 0x0e,
	0x7d, 0x6b, 0x1b, 0x8f, 0xca, 0x9b, 0x25, 0xaa, 0xe9, 0xb7, 0x80, 0x1e, 0x1e, 0x64, 0x49, 0x3e,
	0x28, 0xce, 0xff, 0xce, 0x2b, 0x56, 0x69, 0x1f, 0xcd, 0x12, 0xaf, 0x28, 0x49, 0x5e, 0x8a, 0xf3,
	0x3f, 0x93, 0xff, 0x9f, 0x8b, 0x64, 0x97, 0xd3, 0x3e, 0xe9, 0x5f, 0x8b, 0x23, 0x91, 0x3a, 0xd4,
	0xde, 0x36, 0xf0, 0x30, 0xce, 0x20, 0x2a, 0x79, 0x2e, 0x84, 0xc7, 0xa6, 0x52, 0xb7, 0xb5, 0x9d,
	0x7a, 0x38, 0xe0, 0x27, 0x18, 0x11, 0xb9, 0x21, 0x40, 0x83, 0x9e, 0x39, 0xbb, 0x80, 0xc3, 0x2c,
	0xc9, 0xc7, 0x05, 0x7f, 0xcb, 0x03, 0xb1, 0x45, 0xb7, 0xaf, 0xe1, 0x88, 0x61, 0x14, 0x72, 0x22,
	0x36, 0xc2, 0x12, 0x8e, 0xd9, 0x7c, 0x23, 0x2c, 0xe5, 0x89, 0x18, 0x52, 0xb6, 0x72, 0xfa, 0x0e,
	0x20, 0xee, 0x03, 0xe9, 0x42, 0xdf, 0xf1, 0xdc, 0xa9, 0x86, 0x63, 0x27, 0xab, 0x15, 0x25, 0x40,
	0x41, 0x5e, 0x51, 0x03, 0xa7, 0xeb, 0x15, 0x35, 0xd4, 0x22, 0x3b, 0xc5, 0x3d, 0x39, 0x8b, 0x2d,
	0x12, 0x89, 0x8b, 0xf2, 0x4c, 0xec, 0xf5, 0xef, 0xfb, 0xa5, 0x43, 0x1f, 0x3c, 0x3c, 0xe2, 0xe7,
	0x98, 0xac, 0x9f, 0x33, 0xd2, 0xab, 0xdd, 0xef, 0xf7, 0x17, 0xc9, 0x8f, 0xfb, 0x8b, 0xe4, 0xe7,
	0xfd, 0x45, 0x32, 0x4d, 0xf9, 0x87, 0x7c, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x96, 0x1f,
	0x32, 0xa3, 0x03, 0x00, 0x00,
}
