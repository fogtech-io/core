// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OrderType int32

const (
	OrderType_ANY OrderType = 0
	OrderType_BID OrderType = 1
	OrderType_ASK OrderType = 2
)

var OrderType_name = map[int32]string{
	0: "ANY",
	1: "BID",
	2: "ASK",
}
var OrderType_value = map[string]int32{
	"ANY": 0,
	"BID": 1,
	"ASK": 2,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

type OrderStatus int32

const (
	OrderStatus_ORDER_UNKNOWN  OrderStatus = 0
	OrderStatus_ORDER_INACTIVE OrderStatus = 1
	OrderStatus_ORDER_ACTIVE   OrderStatus = 2
)

var OrderStatus_name = map[int32]string{
	0: "ORDER_UNKNOWN",
	1: "ORDER_INACTIVE",
	2: "ORDER_ACTIVE",
}
var OrderStatus_value = map[string]int32{
	"ORDER_UNKNOWN":  0,
	"ORDER_INACTIVE": 1,
	"ORDER_ACTIVE":   2,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}
func (OrderStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

type IdentityLevel int32

const (
	IdentityLevel_UNKNOWN      IdentityLevel = 0
	IdentityLevel_ANONYMOUS    IdentityLevel = 1
	IdentityLevel_REGISTERED   IdentityLevel = 2
	IdentityLevel_IDENTIFIED   IdentityLevel = 3
	IdentityLevel_PROFESSIONAL IdentityLevel = 4
)

var IdentityLevel_name = map[int32]string{
	0: "UNKNOWN",
	1: "ANONYMOUS",
	2: "REGISTERED",
	3: "IDENTIFIED",
	4: "PROFESSIONAL",
}
var IdentityLevel_value = map[string]int32{
	"UNKNOWN":      0,
	"ANONYMOUS":    1,
	"REGISTERED":   2,
	"IDENTIFIED":   3,
	"PROFESSIONAL": 4,
}

func (x IdentityLevel) String() string {
	return proto.EnumName(IdentityLevel_name, int32(x))
}
func (IdentityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

type DealStatus int32

const (
	DealStatus_DEAL_UNKNOWN  DealStatus = 0
	DealStatus_DEAL_ACCEPTED DealStatus = 1
	DealStatus_DEAL_CLOSED   DealStatus = 2
)

var DealStatus_name = map[int32]string{
	0: "DEAL_UNKNOWN",
	1: "DEAL_ACCEPTED",
	2: "DEAL_CLOSED",
}
var DealStatus_value = map[string]int32{
	"DEAL_UNKNOWN":  0,
	"DEAL_ACCEPTED": 1,
	"DEAL_CLOSED":   2,
}

func (x DealStatus) String() string {
	return proto.EnumName(DealStatus_name, int32(x))
}
func (DealStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

type ChangeRequestStatus int32

const (
	ChangeRequestStatus_REQUEST_UNKNOWN  ChangeRequestStatus = 0
	ChangeRequestStatus_REQUEST_CREATED  ChangeRequestStatus = 1
	ChangeRequestStatus_REQUEST_CANCELED ChangeRequestStatus = 2
	ChangeRequestStatus_REQUEST_REJECTED ChangeRequestStatus = 3
	ChangeRequestStatus_REQUEST_ACCEPTED ChangeRequestStatus = 4
)

var ChangeRequestStatus_name = map[int32]string{
	0: "REQUEST_UNKNOWN",
	1: "REQUEST_CREATED",
	2: "REQUEST_CANCELED",
	3: "REQUEST_REJECTED",
	4: "REQUEST_ACCEPTED",
}
var ChangeRequestStatus_value = map[string]int32{
	"REQUEST_UNKNOWN":  0,
	"REQUEST_CREATED":  1,
	"REQUEST_CANCELED": 2,
	"REQUEST_REJECTED": 3,
	"REQUEST_ACCEPTED": 4,
}

func (x ChangeRequestStatus) String() string {
	return proto.EnumName(ChangeRequestStatus_name, int32(x))
}
func (ChangeRequestStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

type BlacklistType int32

const (
	BlacklistType_BLACKLIST_NOBODY BlacklistType = 0
	BlacklistType_BLACKLIST_WORKER BlacklistType = 1
	BlacklistType_BLACKLIST_MASTER BlacklistType = 2
)

var BlacklistType_name = map[int32]string{
	0: "BLACKLIST_NOBODY",
	1: "BLACKLIST_WORKER",
	2: "BLACKLIST_MASTER",
}
var BlacklistType_value = map[string]int32{
	"BLACKLIST_NOBODY": 0,
	"BLACKLIST_WORKER": 1,
	"BLACKLIST_MASTER": 2,
}

func (x BlacklistType) String() string {
	return proto.EnumName(BlacklistType_name, int32(x))
}
func (BlacklistType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

type ErrorByID struct {
	Response []*ErrorByID_Item `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *ErrorByID) Reset()                    { *m = ErrorByID{} }
func (m *ErrorByID) String() string            { return proto.CompactTextString(m) }
func (*ErrorByID) ProtoMessage()               {}
func (*ErrorByID) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *ErrorByID) GetResponse() []*ErrorByID_Item {
	if m != nil {
		return m.Response
	}
	return nil
}

type ErrorByID_Item struct {
	Id    *BigInt `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Error string  `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *ErrorByID_Item) Reset()                    { *m = ErrorByID_Item{} }
func (m *ErrorByID_Item) String() string            { return proto.CompactTextString(m) }
func (*ErrorByID_Item) ProtoMessage()               {}
func (*ErrorByID_Item) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0, 0} }

func (m *ErrorByID_Item) GetId() *BigInt {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ErrorByID_Item) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ErrorByStringID struct {
	Response []*ErrorByStringID_Item `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *ErrorByStringID) Reset()                    { *m = ErrorByStringID{} }
func (m *ErrorByStringID) String() string            { return proto.CompactTextString(m) }
func (*ErrorByStringID) ProtoMessage()               {}
func (*ErrorByStringID) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *ErrorByStringID) GetResponse() []*ErrorByStringID_Item {
	if m != nil {
		return m.Response
	}
	return nil
}

type ErrorByStringID_Item struct {
	ID    string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ErrorByStringID_Item) Reset()                    { *m = ErrorByStringID_Item{} }
func (m *ErrorByStringID_Item) String() string            { return proto.CompactTextString(m) }
func (*ErrorByStringID_Item) ProtoMessage()               {}
func (*ErrorByStringID_Item) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1, 0} }

func (m *ErrorByStringID_Item) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ErrorByStringID_Item) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type OrderIDs struct {
	Ids []*BigInt `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *OrderIDs) Reset()                    { *m = OrderIDs{} }
func (m *OrderIDs) String() string            { return proto.CompactTextString(m) }
func (*OrderIDs) ProtoMessage()               {}
func (*OrderIDs) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *OrderIDs) GetIds() []*BigInt {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetOrdersReply struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *GetOrdersReply) Reset()                    { *m = GetOrdersReply{} }
func (m *GetOrdersReply) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersReply) ProtoMessage()               {}
func (*GetOrdersReply) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *GetOrdersReply) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type Benchmarks struct {
	Values []uint64 `protobuf:"varint,1,rep,packed,name=values" json:"values,omitempty"`
}

func (m *Benchmarks) Reset()                    { *m = Benchmarks{} }
func (m *Benchmarks) String() string            { return proto.CompactTextString(m) }
func (*Benchmarks) ProtoMessage()               {}
func (*Benchmarks) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *Benchmarks) GetValues() []uint64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Deal struct {
	Id             *BigInt     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Benchmarks     *Benchmarks `protobuf:"bytes,2,opt,name=benchmarks" json:"benchmarks,omitempty"`
	SupplierID     *EthAddress `protobuf:"bytes,3,opt,name=supplierID" json:"supplierID,omitempty"`
	ConsumerID     *EthAddress `protobuf:"bytes,4,opt,name=consumerID" json:"consumerID,omitempty"`
	MasterID       *EthAddress `protobuf:"bytes,5,opt,name=masterID" json:"masterID,omitempty"`
	AskID          *BigInt     `protobuf:"bytes,6,opt,name=askID" json:"askID,omitempty"`
	BidID          *BigInt     `protobuf:"bytes,7,opt,name=bidID" json:"bidID,omitempty"`
	Duration       uint64      `protobuf:"varint,8,opt,name=duration" json:"duration,omitempty"`
	Price          *BigInt     `protobuf:"bytes,9,opt,name=price" json:"price,omitempty"`
	StartTime      *Timestamp  `protobuf:"bytes,10,opt,name=startTime" json:"startTime,omitempty"`
	EndTime        *Timestamp  `protobuf:"bytes,11,opt,name=endTime" json:"endTime,omitempty"`
	Status         DealStatus  `protobuf:"varint,12,opt,name=status,enum=sonm.DealStatus" json:"status,omitempty"`
	BlockedBalance *BigInt     `protobuf:"bytes,13,opt,name=blockedBalance" json:"blockedBalance,omitempty"`
	TotalPayout    *BigInt     `protobuf:"bytes,14,opt,name=totalPayout" json:"totalPayout,omitempty"`
	LastBillTS     *Timestamp  `protobuf:"bytes,15,opt,name=lastBillTS" json:"lastBillTS,omitempty"`
}

func (m *Deal) Reset()                    { *m = Deal{} }
func (m *Deal) String() string            { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()               {}
func (*Deal) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *Deal) GetId() *BigInt {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Deal) GetBenchmarks() *Benchmarks {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

func (m *Deal) GetSupplierID() *EthAddress {
	if m != nil {
		return m.SupplierID
	}
	return nil
}

func (m *Deal) GetConsumerID() *EthAddress {
	if m != nil {
		return m.ConsumerID
	}
	return nil
}

func (m *Deal) GetMasterID() *EthAddress {
	if m != nil {
		return m.MasterID
	}
	return nil
}

func (m *Deal) GetAskID() *BigInt {
	if m != nil {
		return m.AskID
	}
	return nil
}

func (m *Deal) GetBidID() *BigInt {
	if m != nil {
		return m.BidID
	}
	return nil
}

func (m *Deal) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Deal) GetPrice() *BigInt {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Deal) GetStartTime() *Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Deal) GetEndTime() *Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Deal) GetStatus() DealStatus {
	if m != nil {
		return m.Status
	}
	return DealStatus_DEAL_UNKNOWN
}

func (m *Deal) GetBlockedBalance() *BigInt {
	if m != nil {
		return m.BlockedBalance
	}
	return nil
}

func (m *Deal) GetTotalPayout() *BigInt {
	if m != nil {
		return m.TotalPayout
	}
	return nil
}

func (m *Deal) GetLastBillTS() *Timestamp {
	if m != nil {
		return m.LastBillTS
	}
	return nil
}

type Order struct {
	Id             *BigInt       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DealID         *BigInt       `protobuf:"bytes,2,opt,name=dealID" json:"dealID,omitempty"`
	OrderType      OrderType     `protobuf:"varint,3,opt,name=orderType,enum=sonm.OrderType" json:"orderType,omitempty"`
	OrderStatus    OrderStatus   `protobuf:"varint,4,opt,name=orderStatus,enum=sonm.OrderStatus" json:"orderStatus,omitempty"`
	AuthorID       *EthAddress   `protobuf:"bytes,5,opt,name=authorID" json:"authorID,omitempty"`
	CounterpartyID *EthAddress   `protobuf:"bytes,6,opt,name=counterpartyID" json:"counterpartyID,omitempty"`
	Duration       uint64        `protobuf:"varint,7,opt,name=duration" json:"duration,omitempty"`
	Price          *BigInt       `protobuf:"bytes,8,opt,name=price" json:"price,omitempty"`
	Netflags       *NetFlags     `protobuf:"bytes,9,opt,name=netflags" json:"netflags,omitempty"`
	IdentityLevel  IdentityLevel `protobuf:"varint,10,opt,name=identityLevel,enum=sonm.IdentityLevel" json:"identityLevel,omitempty"`
	Blacklist      string        `protobuf:"bytes,11,opt,name=blacklist" json:"blacklist,omitempty"`
	Tag            []byte        `protobuf:"bytes,12,opt,name=tag,proto3" json:"tag,omitempty"`
	Benchmarks     *Benchmarks   `protobuf:"bytes,13,opt,name=benchmarks" json:"benchmarks,omitempty"`
	FrozenSum      *BigInt       `protobuf:"bytes,14,opt,name=frozenSum" json:"frozenSum,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *Order) GetId() *BigInt {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Order) GetDealID() *BigInt {
	if m != nil {
		return m.DealID
	}
	return nil
}

func (m *Order) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_ANY
}

func (m *Order) GetOrderStatus() OrderStatus {
	if m != nil {
		return m.OrderStatus
	}
	return OrderStatus_ORDER_UNKNOWN
}

func (m *Order) GetAuthorID() *EthAddress {
	if m != nil {
		return m.AuthorID
	}
	return nil
}

func (m *Order) GetCounterpartyID() *EthAddress {
	if m != nil {
		return m.CounterpartyID
	}
	return nil
}

func (m *Order) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Order) GetPrice() *BigInt {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Order) GetNetflags() *NetFlags {
	if m != nil {
		return m.Netflags
	}
	return nil
}

func (m *Order) GetIdentityLevel() IdentityLevel {
	if m != nil {
		return m.IdentityLevel
	}
	return IdentityLevel_UNKNOWN
}

func (m *Order) GetBlacklist() string {
	if m != nil {
		return m.Blacklist
	}
	return ""
}

func (m *Order) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *Order) GetBenchmarks() *Benchmarks {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

func (m *Order) GetFrozenSum() *BigInt {
	if m != nil {
		return m.FrozenSum
	}
	return nil
}

type BidNetwork struct {
	Overlay  bool `protobuf:"varint,1,opt,name=overlay" json:"overlay,omitempty"`
	Outbound bool `protobuf:"varint,2,opt,name=outbound" json:"outbound,omitempty"`
	Incoming bool `protobuf:"varint,3,opt,name=incoming" json:"incoming,omitempty"`
}

func (m *BidNetwork) Reset()                    { *m = BidNetwork{} }
func (m *BidNetwork) String() string            { return proto.CompactTextString(m) }
func (*BidNetwork) ProtoMessage()               {}
func (*BidNetwork) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *BidNetwork) GetOverlay() bool {
	if m != nil {
		return m.Overlay
	}
	return false
}

func (m *BidNetwork) GetOutbound() bool {
	if m != nil {
		return m.Outbound
	}
	return false
}

func (m *BidNetwork) GetIncoming() bool {
	if m != nil {
		return m.Incoming
	}
	return false
}

type BidResources struct {
	Network    *BidNetwork       `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	Benchmarks map[string]uint64 `protobuf:"bytes,2,rep,name=benchmarks" json:"benchmarks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *BidResources) Reset()                    { *m = BidResources{} }
func (m *BidResources) String() string            { return proto.CompactTextString(m) }
func (*BidResources) ProtoMessage()               {}
func (*BidResources) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *BidResources) GetNetwork() *BidNetwork {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *BidResources) GetBenchmarks() map[string]uint64 {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

type BidOrder struct {
	ID           string        `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Duration     *Duration     `protobuf:"bytes,2,opt,name=duration" json:"duration,omitempty"`
	Price        *Price        `protobuf:"bytes,3,opt,name=price" json:"price,omitempty"`
	Blacklist    *EthAddress   `protobuf:"bytes,4,opt,name=blacklist" json:"blacklist,omitempty"`
	Identity     IdentityLevel `protobuf:"varint,5,opt,name=identity,enum=sonm.IdentityLevel" json:"identity,omitempty"`
	Tag          string        `protobuf:"bytes,6,opt,name=tag" json:"tag,omitempty"`
	Counterparty *EthAddress   `protobuf:"bytes,7,opt,name=Counterparty" json:"Counterparty,omitempty"`
	Resources    *BidResources `protobuf:"bytes,8,opt,name=resources" json:"resources,omitempty"`
}

func (m *BidOrder) Reset()                    { *m = BidOrder{} }
func (m *BidOrder) String() string            { return proto.CompactTextString(m) }
func (*BidOrder) ProtoMessage()               {}
func (*BidOrder) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *BidOrder) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *BidOrder) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *BidOrder) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *BidOrder) GetBlacklist() *EthAddress {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *BidOrder) GetIdentity() IdentityLevel {
	if m != nil {
		return m.Identity
	}
	return IdentityLevel_UNKNOWN
}

func (m *BidOrder) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *BidOrder) GetCounterparty() *EthAddress {
	if m != nil {
		return m.Counterparty
	}
	return nil
}

func (m *BidOrder) GetResources() *BidResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorByID)(nil), "sonm.ErrorByID")
	proto.RegisterType((*ErrorByID_Item)(nil), "sonm.ErrorByID.Item")
	proto.RegisterType((*ErrorByStringID)(nil), "sonm.ErrorByStringID")
	proto.RegisterType((*ErrorByStringID_Item)(nil), "sonm.ErrorByStringID.Item")
	proto.RegisterType((*OrderIDs)(nil), "sonm.OrderIDs")
	proto.RegisterType((*GetOrdersReply)(nil), "sonm.GetOrdersReply")
	proto.RegisterType((*Benchmarks)(nil), "sonm.Benchmarks")
	proto.RegisterType((*Deal)(nil), "sonm.Deal")
	proto.RegisterType((*Order)(nil), "sonm.Order")
	proto.RegisterType((*BidNetwork)(nil), "sonm.BidNetwork")
	proto.RegisterType((*BidResources)(nil), "sonm.BidResources")
	proto.RegisterType((*BidOrder)(nil), "sonm.BidOrder")
	proto.RegisterEnum("sonm.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("sonm.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterEnum("sonm.IdentityLevel", IdentityLevel_name, IdentityLevel_value)
	proto.RegisterEnum("sonm.DealStatus", DealStatus_name, DealStatus_value)
	proto.RegisterEnum("sonm.ChangeRequestStatus", ChangeRequestStatus_name, ChangeRequestStatus_value)
	proto.RegisterEnum("sonm.BlacklistType", BlacklistType_name, BlacklistType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Market service

type MarketClient interface {
	// GetOrders returns orders by given filter parameters.
	// Note that set of filters may be changed in the closest future.
	GetOrders(ctx context.Context, in *Count, opts ...grpc.CallOption) (*GetOrdersReply, error)
	// CreateOrder places new order on the Marketplace.
	// Note that current impl of Node API prevents you from
	// creating ASKs orders.
	CreateOrder(ctx context.Context, in *BidOrder, opts ...grpc.CallOption) (*Order, error)
	// GetOrderByID returns order by given ID.
	// If order save an `inactive` status returns error instead.
	// TODO: get rid of string ID #1237
	GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error)
	// CancelOrder removes active order from the Marketplace.
	// TODO: get rid of string ID #1237
	// Deprecated: use CancelOrders instead
	CancelOrder(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	// CancelOrders removes specified orders from the Marketplace.
	CancelOrders(ctx context.Context, in *OrderIDs, opts ...grpc.CallOption) (*ErrorByID, error)
	// Purge remove all active orders from marketplace
	// Deprecated: use PurgeVerbose
	Purge(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// PurgeVerbose remove all active orders from marketplace and return detailed status on each order
	PurgeVerbose(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ErrorByID, error)
}

type marketClient struct {
	cc *grpc.ClientConn
}

func NewMarketClient(cc *grpc.ClientConn) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) GetOrders(ctx context.Context, in *Count, opts ...grpc.CallOption) (*GetOrdersReply, error) {
	out := new(GetOrdersReply)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CreateOrder(ctx context.Context, in *BidOrder, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/CreateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrderByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CancelOrder(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Market/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CancelOrders(ctx context.Context, in *OrderIDs, opts ...grpc.CallOption) (*ErrorByID, error) {
	out := new(ErrorByID)
	err := grpc.Invoke(ctx, "/sonm.Market/CancelOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) Purge(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Market/Purge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) PurgeVerbose(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ErrorByID, error) {
	out := new(ErrorByID)
	err := grpc.Invoke(ctx, "/sonm.Market/PurgeVerbose", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Market service

type MarketServer interface {
	// GetOrders returns orders by given filter parameters.
	// Note that set of filters may be changed in the closest future.
	GetOrders(context.Context, *Count) (*GetOrdersReply, error)
	// CreateOrder places new order on the Marketplace.
	// Note that current impl of Node API prevents you from
	// creating ASKs orders.
	CreateOrder(context.Context, *BidOrder) (*Order, error)
	// GetOrderByID returns order by given ID.
	// If order save an `inactive` status returns error instead.
	// TODO: get rid of string ID #1237
	GetOrderByID(context.Context, *ID) (*Order, error)
	// CancelOrder removes active order from the Marketplace.
	// TODO: get rid of string ID #1237
	// Deprecated: use CancelOrders instead
	CancelOrder(context.Context, *ID) (*Empty, error)
	// CancelOrders removes specified orders from the Marketplace.
	CancelOrders(context.Context, *OrderIDs) (*ErrorByID, error)
	// Purge remove all active orders from marketplace
	// Deprecated: use PurgeVerbose
	Purge(context.Context, *Empty) (*Empty, error)
	// PurgeVerbose remove all active orders from marketplace and return detailed status on each order
	PurgeVerbose(context.Context, *Empty) (*ErrorByID, error)
}

func RegisterMarketServer(s *grpc.Server, srv MarketServer) {
	s.RegisterService(&_Market_serviceDesc, srv)
}

func _Market_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Count)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrders(ctx, req.(*Count))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CreateOrder(ctx, req.(*BidOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrderByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CancelOrder(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CancelOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CancelOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CancelOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CancelOrders(ctx, req.(*OrderIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/Purge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).Purge(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_PurgeVerbose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).PurgeVerbose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/PurgeVerbose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).PurgeVerbose(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Market_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrders",
			Handler:    _Market_GetOrders_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Market_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _Market_GetOrderByID_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Market_CancelOrder_Handler,
		},
		{
			MethodName: "CancelOrders",
			Handler:    _Market_CancelOrders_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _Market_Purge_Handler,
		},
		{
			MethodName: "PurgeVerbose",
			Handler:    _Market_PurgeVerbose_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketplace.proto",
}

func init() { proto.RegisterFile("marketplace.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 1335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x13, 0xb6, 0x64, 0x59, 0xa6, 0x46, 0x5f, 0xf4, 0xc6, 0x78, 0x41, 0x08, 0xc1, 0x0b, 0x97, 0x09,
	0x52, 0x57, 0x48, 0x95, 0xd4, 0x49, 0x8b, 0x34, 0x40, 0x0f, 0xa2, 0x48, 0x07, 0xac, 0x15, 0xca,
	0x59, 0x29, 0x09, 0x72, 0x68, 0x83, 0x95, 0xb8, 0x91, 0x17, 0xa6, 0x48, 0x95, 0x5c, 0x26, 0x50,
	0x6f, 0x3d, 0xf7, 0x37, 0xf5, 0x50, 0xa0, 0xc7, 0xfe, 0xa8, 0x62, 0x97, 0x1f, 0xa2, 0x64, 0xb9,
	0xe9, 0x8d, 0xf3, 0xcc, 0x33, 0xb3, 0xb3, 0xc3, 0x67, 0x67, 0x17, 0x8e, 0x16, 0x24, 0xbc, 0xa6,
	0x7c, 0xe9, 0x91, 0x19, 0xed, 0x2d, 0xc3, 0x80, 0x07, 0xa8, 0x12, 0x05, 0xfe, 0xa2, 0xd3, 0x98,
	0xb2, 0x39, 0xf3, 0x79, 0x82, 0x75, 0xd0, 0x8c, 0x2c, 0xc9, 0x94, 0x79, 0x8c, 0x33, 0x1a, 0xa5,
	0x58, 0x9b, 0xf9, 0x82, 0xe9, 0x33, 0x92, 0x01, 0x9c, 0x2d, 0x68, 0xc4, 0xc9, 0x62, 0x99, 0x00,
	0xfa, 0x0a, 0x6a, 0x56, 0x18, 0x06, 0xa1, 0xb1, 0xb2, 0x4d, 0xf4, 0x18, 0x94, 0x90, 0x46, 0xcb,
	0xc0, 0x8f, 0xa8, 0x56, 0x3a, 0xd9, 0x3f, 0xad, 0x9f, 0x1d, 0xf7, 0x44, 0x7c, 0x2f, 0xa7, 0xf4,
	0x6c, 0x4e, 0x17, 0x38, 0x67, 0x75, 0x9e, 0x43, 0x45, 0x20, 0xe8, 0x2e, 0x94, 0x99, 0xab, 0x95,
	0x4e, 0x4a, 0xa7, 0xf5, 0xb3, 0x46, 0x12, 0x63, 0xb0, 0xb9, 0xed, 0x73, 0x5c, 0x66, 0x2e, 0x3a,
	0x86, 0x03, 0x2a, 0x32, 0x68, 0xfb, 0x27, 0xa5, 0xd3, 0x1a, 0x4e, 0x0c, 0xfd, 0x13, 0xb4, 0xd3,
	0xbc, 0x63, 0x1e, 0x32, 0x7f, 0x6e, 0x9b, 0xe8, 0xbb, 0x1b, 0x05, 0x74, 0x36, 0x0a, 0xc8, 0x88,
	0xdb, 0x65, 0x3c, 0x4c, 0xcb, 0x68, 0x41, 0xd9, 0x36, 0x65, 0x19, 0x35, 0x5c, 0xb6, 0xcd, 0xf5,
	0xc2, 0xe5, 0xe2, 0xc2, 0x5d, 0x50, 0x46, 0xa1, 0x4b, 0x43, 0xdb, 0x8c, 0xd0, 0xff, 0x61, 0x9f,
	0xb9, 0x51, 0xba, 0xd8, 0x66, 0xe5, 0xc2, 0xa1, 0x7f, 0x0b, 0xad, 0x17, 0x94, 0x4b, 0x7a, 0x84,
	0xe9, 0xd2, 0x5b, 0xa1, 0x7b, 0x50, 0x0d, 0xa4, 0x99, 0x06, 0xd5, 0x93, 0x20, 0x49, 0xc1, 0xa9,
	0x4b, 0xbf, 0x0f, 0x60, 0x50, 0x7f, 0x76, 0x25, 0x7e, 0x5d, 0x84, 0xfe, 0x07, 0xd5, 0x8f, 0xc4,
	0x8b, 0x69, 0x12, 0x52, 0xc1, 0xa9, 0xa5, 0xff, 0x7e, 0x00, 0x15, 0x93, 0x12, 0xef, 0x33, 0xed,
	0x7b, 0x0c, 0x30, 0xcd, 0x93, 0xc9, 0xad, 0xd4, 0xcf, 0xd4, 0x94, 0x95, 0xe3, 0xb8, 0xc0, 0x11,
	0x11, 0x51, 0xbc, 0x5c, 0x7a, 0x4c, 0x6c, 0x52, 0x76, 0x3d, 0x8f, 0xb0, 0xf8, 0x55, 0xdf, 0x75,
	0x43, 0x1a, 0x45, 0xb8, 0xc0, 0x11, 0x11, 0xb3, 0xc0, 0x8f, 0xe2, 0x85, 0x8c, 0xa8, 0xdc, 0x16,
	0xb1, 0xe6, 0xa0, 0x87, 0xa0, 0x2c, 0x48, 0xc4, 0x25, 0xff, 0xe0, 0x16, 0x7e, 0xce, 0x40, 0x3a,
	0x1c, 0x90, 0xe8, 0xda, 0x36, 0xb5, 0xea, 0x8e, 0x4d, 0x26, 0x2e, 0xc1, 0x99, 0x32, 0xd7, 0x36,
	0xb5, 0xc3, 0x5d, 0x1c, 0xe9, 0x42, 0x1d, 0x50, 0xdc, 0x38, 0x24, 0x9c, 0x05, 0xbe, 0xa6, 0x9c,
	0x94, 0x4e, 0x2b, 0x38, 0xb7, 0x45, 0xfc, 0x32, 0x64, 0x33, 0xaa, 0xd5, 0x76, 0xc5, 0x4b, 0x17,
	0xfa, 0x1a, 0x6a, 0x11, 0x27, 0x21, 0x9f, 0xb0, 0x05, 0xd5, 0x40, 0xf2, 0xda, 0x09, 0x6f, 0x92,
	0x9d, 0x0c, 0xbc, 0x66, 0xa0, 0xaf, 0xe0, 0x90, 0xfa, 0xae, 0x24, 0xd7, 0x77, 0x93, 0x33, 0x3f,
	0x3a, 0x85, 0x6a, 0xc4, 0x09, 0x8f, 0x23, 0xad, 0x71, 0x52, 0x3a, 0x6d, 0x65, 0xdd, 0x10, 0xff,
	0x77, 0x2c, 0x71, 0x9c, 0xfa, 0xd1, 0x53, 0x68, 0x4d, 0xbd, 0x60, 0x76, 0x4d, 0x5d, 0x83, 0x78,
	0xc4, 0x9f, 0x51, 0xad, 0xb9, 0xa3, 0xe0, 0x2d, 0x0e, 0xea, 0x41, 0x9d, 0x07, 0x9c, 0x78, 0x97,
	0x64, 0x15, 0xc4, 0x5c, 0x6b, 0xed, 0x08, 0x29, 0x12, 0xd0, 0x23, 0x00, 0x8f, 0x44, 0xdc, 0x60,
	0x9e, 0x37, 0x19, 0x6b, 0xed, 0xdd, 0xd5, 0x17, 0x28, 0xfa, 0xdf, 0x15, 0x38, 0x90, 0x2a, 0xfe,
	0x8c, 0x1c, 0xef, 0x43, 0xd5, 0xa5, 0xc4, 0xb3, 0xcd, 0x54, 0x8a, 0x9b, 0x8c, 0xd4, 0x27, 0x1a,
	0x2d, 0xcf, 0xc2, 0x64, 0xb5, 0xa4, 0x52, 0x81, 0xad, 0x6c, 0xf5, 0x51, 0x06, 0xe3, 0x35, 0x03,
	0x3d, 0x81, 0xba, 0x34, 0x92, 0x56, 0x49, 0x01, 0xb6, 0xce, 0x8e, 0x0a, 0x01, 0x69, 0x0f, 0x8b,
	0x2c, 0x21, 0x41, 0x12, 0xf3, 0xab, 0xe0, 0x5f, 0x25, 0x98, 0x31, 0xd0, 0x33, 0x68, 0xcd, 0x82,
	0xd8, 0xe7, 0x34, 0x5c, 0x92, 0x90, 0xaf, 0x72, 0x2d, 0xde, 0x8c, 0xd9, 0xe2, 0x6d, 0x88, 0xee,
	0xf0, 0x36, 0xd1, 0x29, 0xb7, 0x8b, 0xae, 0x0b, 0x8a, 0x4f, 0xf9, 0x07, 0x8f, 0xcc, 0xa3, 0x54,
	0x9b, 0xad, 0x84, 0xe6, 0x50, 0x7e, 0x2e, 0x50, 0x9c, 0xfb, 0xd1, 0xf7, 0xd0, 0x64, 0x2e, 0xf5,
	0x39, 0xe3, 0xab, 0x21, 0xfd, 0x48, 0x3d, 0x29, 0xd2, 0xd6, 0xd9, 0x9d, 0x24, 0xc0, 0x2e, 0xba,
	0xf0, 0x26, 0x13, 0xdd, 0x85, 0xda, 0xd4, 0x23, 0xb3, 0x6b, 0x8f, 0x45, 0x5c, 0xca, 0xb5, 0x86,
	0xd7, 0x00, 0x52, 0x61, 0x9f, 0x93, 0xb9, 0x14, 0x67, 0x03, 0x8b, 0xcf, 0xad, 0xb9, 0xd2, 0xfc,
	0x0f, 0x73, 0xa5, 0x0b, 0xb5, 0x0f, 0x61, 0xf0, 0x2b, 0xf5, 0xc7, 0xf1, 0x62, 0xa7, 0x02, 0xd7,
	0x6e, 0xfd, 0x67, 0x00, 0x83, 0xb9, 0x0e, 0xe5, 0x9f, 0x82, 0xf0, 0x1a, 0x69, 0x70, 0x18, 0x7c,
	0xa4, 0xa1, 0x47, 0x56, 0x52, 0x57, 0x0a, 0xce, 0x4c, 0xd1, 0xdc, 0x20, 0xe6, 0xd3, 0x20, 0xf6,
	0x5d, 0x29, 0x28, 0x05, 0xe7, 0xb6, 0xf0, 0x31, 0x7f, 0x16, 0x2c, 0x98, 0x3f, 0x97, 0x1a, 0x52,
	0x70, 0x6e, 0xeb, 0x7f, 0x94, 0xa0, 0x61, 0x30, 0x17, 0xd3, 0x28, 0x88, 0xc3, 0x19, 0x15, 0xc5,
	0x1d, 0xfa, 0xc9, 0x6a, 0xa9, 0x74, 0xb3, 0xbd, 0xe4, 0x55, 0xe0, 0x8c, 0x80, 0x8c, 0xad, 0x91,
	0x2a, 0x06, 0xb9, 0x9e, 0xd3, 0xf3, 0x9c, 0x85, 0x3e, 0x58, 0x3e, 0x0f, 0x57, 0xc5, 0x66, 0x74,
	0x7e, 0x80, 0xf6, 0x96, 0x5b, 0xf4, 0xf8, 0x9a, 0xae, 0xd2, 0x0b, 0x48, 0x7c, 0x8a, 0x1b, 0x48,
	0x0e, 0x7b, 0xb9, 0xb5, 0x0a, 0x4e, 0x8c, 0xe7, 0xe5, 0x67, 0x25, 0xfd, 0xaf, 0x32, 0x28, 0x06,
	0x73, 0x93, 0x13, 0xb7, 0x7d, 0x71, 0x75, 0x0b, 0x8a, 0x2b, 0x17, 0x15, 0x63, 0xa6, 0x68, 0x41,
	0x81, 0x5f, 0x64, 0x0a, 0x4c, 0xe6, 0x7c, 0x7a, 0x1f, 0x5d, 0x0a, 0x28, 0x13, 0x60, 0xaf, 0xa8,
	0x8c, 0xdb, 0x86, 0x7b, 0x41, 0x2b, 0x8f, 0x40, 0xc9, 0xa4, 0x25, 0x0f, 0xd6, 0x2d, 0xfa, 0xcb,
	0x49, 0x99, 0xb8, 0xaa, 0xc9, 0xc6, 0x85, 0xb8, 0x9e, 0x42, 0x63, 0x50, 0x38, 0x45, 0xe9, 0x4c,
	0xbf, 0xb9, 0xea, 0x06, 0x0b, 0x3d, 0x86, 0x5a, 0x98, 0x35, 0x3f, 0x3d, 0x51, 0xe8, 0xe6, 0x6f,
	0xc1, 0x6b, 0x52, 0xf7, 0x01, 0xd4, 0xf2, 0x81, 0x82, 0x0e, 0x61, 0xbf, 0xef, 0xbc, 0x53, 0xf7,
	0xc4, 0x87, 0x61, 0x9b, 0x6a, 0x49, 0x22, 0xe3, 0x0b, 0xb5, 0xdc, 0x3d, 0x87, 0x7a, 0x61, 0x8e,
	0xa0, 0x23, 0x68, 0x8e, 0xb0, 0x69, 0xe1, 0xf7, 0xaf, 0x9d, 0x0b, 0x67, 0xf4, 0xd6, 0x51, 0xf7,
	0x10, 0x82, 0x56, 0x02, 0xd9, 0x4e, 0x7f, 0x30, 0xb1, 0xdf, 0x58, 0x6a, 0x09, 0xa9, 0xd0, 0x48,
	0xb0, 0x14, 0x29, 0x77, 0x7f, 0x82, 0xe6, 0x46, 0x13, 0x50, 0x1d, 0x0e, 0xd7, 0x39, 0x9a, 0x50,
	0xeb, 0x3b, 0x23, 0xe7, 0xdd, 0xcb, 0xd1, 0xeb, 0xb1, 0x5a, 0x42, 0x2d, 0x00, 0x6c, 0xbd, 0xb0,
	0xc7, 0x13, 0x0b, 0x5b, 0xa6, 0x5a, 0x16, 0xb6, 0x6d, 0x5a, 0xce, 0xc4, 0x3e, 0xb7, 0x2d, 0x53,
	0xdd, 0x17, 0xe9, 0x2f, 0xf1, 0xe8, 0xdc, 0x1a, 0x8f, 0xed, 0x91, 0xd3, 0x1f, 0xaa, 0x95, 0xae,
	0x01, 0xb0, 0xbe, 0x31, 0x84, 0xdf, 0xb4, 0xfa, 0xc3, 0x42, 0x91, 0x47, 0xd0, 0x94, 0x48, 0x7f,
	0x30, 0xb0, 0x2e, 0x27, 0x96, 0xd8, 0x62, 0x1b, 0xea, 0x12, 0x1a, 0x0c, 0x47, 0x63, 0xb1, 0x4a,
	0xf7, 0xb7, 0x12, 0xdc, 0x19, 0x5c, 0x11, 0x7f, 0x4e, 0x31, 0xfd, 0x25, 0xa6, 0x11, 0x4f, 0xb3,
	0xdd, 0x81, 0x36, 0xb6, 0x5e, 0xbd, 0xb6, 0xc6, 0x93, 0x42, 0xc2, 0x02, 0x38, 0xc0, 0x56, 0x3f,
	0x49, 0x79, 0x0c, 0x6a, 0x0e, 0xf6, 0x9d, 0x81, 0x35, 0x94, 0xd5, 0x17, 0x50, 0x6c, 0xfd, 0x68,
	0x0d, 0x26, 0x72, 0x0f, 0x05, 0x34, 0x2f, 0xaa, 0xd2, 0x7d, 0x05, 0x4d, 0x23, 0x93, 0x93, 0xfc,
	0x35, 0xc7, 0xa0, 0x1a, 0xc3, 0xfe, 0xe0, 0x62, 0x68, 0x8f, 0x27, 0xef, 0x9d, 0x91, 0x31, 0x32,
	0xc5, 0x7f, 0xda, 0x40, 0xdf, 0x8e, 0xf0, 0x85, 0x85, 0x93, 0xe5, 0xd7, 0xe8, 0xcb, 0xbe, 0x68,
	0x9f, 0x5a, 0x3e, 0xfb, 0xb3, 0x0c, 0xd5, 0x97, 0xf2, 0x29, 0x2c, 0x64, 0x92, 0xbf, 0xca, 0x50,
	0x2a, 0x78, 0xa9, 0xa3, 0x4e, 0xfa, 0x60, 0xdd, 0x7c, 0xb3, 0xe9, 0x7b, 0xe8, 0x21, 0xd4, 0x07,
	0x21, 0x25, 0x9c, 0xa6, 0xe7, 0x2d, 0x17, 0x95, 0xb4, 0x3b, 0xc5, 0x47, 0x9c, 0xbe, 0x87, 0xbe,
	0x84, 0x46, 0x96, 0x41, 0x3e, 0x8c, 0x95, 0x54, 0xfd, 0xe6, 0x36, 0xf1, 0x01, 0xd4, 0x07, 0xe2,
	0x76, 0xf6, 0x92, 0xb4, 0x37, 0x78, 0xd6, 0x62, 0xc9, 0xc5, 0xf2, 0xdf, 0x40, 0xa3, 0xc0, 0x8b,
	0xb2, 0xf5, 0xb3, 0x67, 0x68, 0xa7, 0xbd, 0xf5, 0xce, 0xd6, 0xf7, 0xd0, 0x3d, 0x38, 0xb8, 0x8c,
	0xc3, 0x39, 0x45, 0xc5, 0x54, 0xdb, 0x79, 0x7b, 0xd0, 0x90, 0xa4, 0x37, 0x34, 0x9c, 0x06, 0xd1,
	0x16, 0xf7, 0x66, 0xd2, 0x69, 0x55, 0xbe, 0xfa, 0x9f, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x80,
	0x53, 0x81, 0x2f, 0x54, 0x0c, 0x00, 0x00,
}
