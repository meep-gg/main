// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: riot/lol/league/v1/league.proto

package v1

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

type RiotLeagueApexReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Queue  string `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *RiotLeagueApexReq) Reset() {
	*x = RiotLeagueApexReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riot_lol_league_v1_league_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiotLeagueApexReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiotLeagueApexReq) ProtoMessage() {}

func (x *RiotLeagueApexReq) ProtoReflect() protoreflect.Message {
	mi := &file_riot_lol_league_v1_league_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiotLeagueApexReq.ProtoReflect.Descriptor instead.
func (*RiotLeagueApexReq) Descriptor() ([]byte, []int) {
	return file_riot_lol_league_v1_league_proto_rawDescGZIP(), []int{0}
}

func (x *RiotLeagueApexReq) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RiotLeagueApexReq) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

type RiotLeagueSummonerIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region     string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	SummonerId string `protobuf:"bytes,2,opt,name=summonerId,proto3" json:"summonerId,omitempty"`
}

func (x *RiotLeagueSummonerIdReq) Reset() {
	*x = RiotLeagueSummonerIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riot_lol_league_v1_league_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiotLeagueSummonerIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiotLeagueSummonerIdReq) ProtoMessage() {}

func (x *RiotLeagueSummonerIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_riot_lol_league_v1_league_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiotLeagueSummonerIdReq.ProtoReflect.Descriptor instead.
func (*RiotLeagueSummonerIdReq) Descriptor() ([]byte, []int) {
	return file_riot_lol_league_v1_league_proto_rawDescGZIP(), []int{1}
}

func (x *RiotLeagueSummonerIdReq) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RiotLeagueSummonerIdReq) GetSummonerId() string {
	if x != nil {
		return x.SummonerId
	}
	return ""
}

type RiotLeagueIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region   string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	LeagueId string `protobuf:"bytes,2,opt,name=leagueId,proto3" json:"leagueId,omitempty"`
}

func (x *RiotLeagueIdReq) Reset() {
	*x = RiotLeagueIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riot_lol_league_v1_league_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiotLeagueIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiotLeagueIdReq) ProtoMessage() {}

func (x *RiotLeagueIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_riot_lol_league_v1_league_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiotLeagueIdReq.ProtoReflect.Descriptor instead.
func (*RiotLeagueIdReq) Descriptor() ([]byte, []int) {
	return file_riot_lol_league_v1_league_proto_rawDescGZIP(), []int{2}
}

func (x *RiotLeagueIdReq) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RiotLeagueIdReq) GetLeagueId() string {
	if x != nil {
		return x.LeagueId
	}
	return ""
}

type RiotLeagueEntryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region   string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Division string `protobuf:"bytes,2,opt,name=division,proto3" json:"division,omitempty"`
	Tier     string `protobuf:"bytes,3,opt,name=tier,proto3" json:"tier,omitempty"`
	Queue    string `protobuf:"bytes,4,opt,name=queue,proto3" json:"queue,omitempty"`
	Page     int32  `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *RiotLeagueEntryReq) Reset() {
	*x = RiotLeagueEntryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riot_lol_league_v1_league_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiotLeagueEntryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiotLeagueEntryReq) ProtoMessage() {}

func (x *RiotLeagueEntryReq) ProtoReflect() protoreflect.Message {
	mi := &file_riot_lol_league_v1_league_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiotLeagueEntryReq.ProtoReflect.Descriptor instead.
func (*RiotLeagueEntryReq) Descriptor() ([]byte, []int) {
	return file_riot_lol_league_v1_league_proto_rawDescGZIP(), []int{3}
}

func (x *RiotLeagueEntryReq) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RiotLeagueEntryReq) GetDivision() string {
	if x != nil {
		return x.Division
	}
	return ""
}

func (x *RiotLeagueEntryReq) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *RiotLeagueEntryReq) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *RiotLeagueEntryReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type RiotLeagueEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leagues []*RiotLeagueEntry `protobuf:"bytes,1,rep,name=leagues,proto3" json:"leagues,omitempty"`
}

func (x *RiotLeagueEntries) Reset() {
	*x = RiotLeagueEntries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riot_lol_league_v1_league_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiotLeagueEntries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiotLeagueEntries) ProtoMessage() {}

func (x *RiotLeagueEntries) ProtoReflect() protoreflect.Message {
	mi := &file_riot_lol_league_v1_league_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiotLeagueEntries.ProtoReflect.Descriptor instead.
func (*RiotLeagueEntries) Descriptor() ([]byte, []int) {
	return file_riot_lol_league_v1_league_proto_rawDescGZIP(), []int{4}
}

func (x *RiotLeagueEntries) GetLeagues() []*RiotLeagueEntry {
	if x != nil {
		return x.Leagues
	}
	return nil
}

type RiotLeagueEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueId     string      `protobuf:"bytes,1,opt,name=leagueId,proto3" json:"leagueId,omitempty"`
	SummonerId   string      `protobuf:"bytes,2,opt,name=summonerId,proto3" json:"summonerId,omitempty"`
	SummonerName string      `protobuf:"bytes,3,opt,name=summonerName,proto3" json:"summonerName,omitempty"`
	QueueType    string      `protobuf:"bytes,4,opt,name=queueType,proto3" json:"queueType,omitempty"`
	Tier         string      `protobuf:"bytes,5,opt,name=tier,proto3" json:"tier,omitempty"`
	Rank         string      `protobuf:"bytes,6,opt,name=rank,proto3" json:"rank,omitempty"`
	LeaguePoints int32       `protobuf:"varint,7,opt,name=leaguePoints,proto3" json:"leaguePoints,omitempty"`
	Wins         int32       `protobuf:"varint,8,opt,name=wins,proto3" json:"wins,omitempty"`
	Losses       int32       `protobuf:"varint,9,opt,name=losses,proto3" json:"losses,omitempty"`
	HotStreak    bool        `protobuf:"varint,10,opt,name=hotStreak,proto3" json:"hotStreak,omitempty"`
	Veteran      bool        `protobuf:"varint,11,opt,name=veteran,proto3" json:"veteran,omitempty"`
	FreshBlood   bool        `protobuf:"varint,12,opt,name=freshBlood,proto3" json:"freshBlood,omitempty"`
	Inactive     bool        `protobuf:"varint,13,opt,name=inactive,proto3" json:"inactive,omitempty"`
	MiniSeries   *MiniSeries `protobuf:"bytes,14,opt,name=miniSeries,proto3" json:"miniSeries,omitempty"`
}

func (x *RiotLeagueEntry) Reset() {
	*x = RiotLeagueEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riot_lol_league_v1_league_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiotLeagueEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiotLeagueEntry) ProtoMessage() {}

func (x *RiotLeagueEntry) ProtoReflect() protoreflect.Message {
	mi := &file_riot_lol_league_v1_league_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiotLeagueEntry.ProtoReflect.Descriptor instead.
func (*RiotLeagueEntry) Descriptor() ([]byte, []int) {
	return file_riot_lol_league_v1_league_proto_rawDescGZIP(), []int{5}
}

func (x *RiotLeagueEntry) GetLeagueId() string {
	if x != nil {
		return x.LeagueId
	}
	return ""
}

func (x *RiotLeagueEntry) GetSummonerId() string {
	if x != nil {
		return x.SummonerId
	}
	return ""
}

func (x *RiotLeagueEntry) GetSummonerName() string {
	if x != nil {
		return x.SummonerName
	}
	return ""
}

func (x *RiotLeagueEntry) GetQueueType() string {
	if x != nil {
		return x.QueueType
	}
	return ""
}

func (x *RiotLeagueEntry) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *RiotLeagueEntry) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

func (x *RiotLeagueEntry) GetLeaguePoints() int32 {
	if x != nil {
		return x.LeaguePoints
	}
	return 0
}

func (x *RiotLeagueEntry) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *RiotLeagueEntry) GetLosses() int32 {
	if x != nil {
		return x.Losses
	}
	return 0
}

func (x *RiotLeagueEntry) GetHotStreak() bool {
	if x != nil {
		return x.HotStreak
	}
	return false
}

func (x *RiotLeagueEntry) GetVeteran() bool {
	if x != nil {
		return x.Veteran
	}
	return false
}

func (x *RiotLeagueEntry) GetFreshBlood() bool {
	if x != nil {
		return x.FreshBlood
	}
	return false
}

func (x *RiotLeagueEntry) GetInactive() bool {
	if x != nil {
		return x.Inactive
	}
	return false
}

func (x *RiotLeagueEntry) GetMiniSeries() *MiniSeries {
	if x != nil {
		return x.MiniSeries
	}
	return nil
}

type RiotLeagueList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueId string        `protobuf:"bytes,1,opt,name=leagueId,proto3" json:"leagueId,omitempty"`
	Entries  []*LeagueItem `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	Tier     string        `protobuf:"bytes,3,opt,name=tier,proto3" json:"tier,omitempty"`
	Name     string        `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Queue    string        `protobuf:"bytes,5,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *RiotLeagueList) Reset() {
	*x = RiotLeagueList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riot_lol_league_v1_league_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiotLeagueList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiotLeagueList) ProtoMessage() {}

func (x *RiotLeagueList) ProtoReflect() protoreflect.Message {
	mi := &file_riot_lol_league_v1_league_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiotLeagueList.ProtoReflect.Descriptor instead.
func (*RiotLeagueList) Descriptor() ([]byte, []int) {
	return file_riot_lol_league_v1_league_proto_rawDescGZIP(), []int{6}
}

func (x *RiotLeagueList) GetLeagueId() string {
	if x != nil {
		return x.LeagueId
	}
	return ""
}

func (x *RiotLeagueList) GetEntries() []*LeagueItem {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *RiotLeagueList) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *RiotLeagueList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RiotLeagueList) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

type LeagueItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FreshBlood   bool        `protobuf:"varint,1,opt,name=freshBlood,proto3" json:"freshBlood,omitempty"`
	Wins         int32       `protobuf:"varint,2,opt,name=wins,proto3" json:"wins,omitempty"`
	SummonerName string      `protobuf:"bytes,3,opt,name=summonerName,proto3" json:"summonerName,omitempty"`
	MiniSeries   *MiniSeries `protobuf:"bytes,4,opt,name=miniSeries,proto3" json:"miniSeries,omitempty"`
	Inactive     bool        `protobuf:"varint,5,opt,name=inactive,proto3" json:"inactive,omitempty"`
	Veteran      bool        `protobuf:"varint,6,opt,name=veteran,proto3" json:"veteran,omitempty"`
	HotStreak    bool        `protobuf:"varint,7,opt,name=hotStreak,proto3" json:"hotStreak,omitempty"`
	Rank         string      `protobuf:"bytes,8,opt,name=rank,proto3" json:"rank,omitempty"`
	LeaguePoints int32       `protobuf:"varint,9,opt,name=leaguePoints,proto3" json:"leaguePoints,omitempty"`
	Losses       int32       `protobuf:"varint,10,opt,name=losses,proto3" json:"losses,omitempty"`
	SummonerId   string      `protobuf:"bytes,11,opt,name=summonerId,proto3" json:"summonerId,omitempty"`
}

func (x *LeagueItem) Reset() {
	*x = LeagueItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riot_lol_league_v1_league_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeagueItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeagueItem) ProtoMessage() {}

func (x *LeagueItem) ProtoReflect() protoreflect.Message {
	mi := &file_riot_lol_league_v1_league_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeagueItem.ProtoReflect.Descriptor instead.
func (*LeagueItem) Descriptor() ([]byte, []int) {
	return file_riot_lol_league_v1_league_proto_rawDescGZIP(), []int{7}
}

func (x *LeagueItem) GetFreshBlood() bool {
	if x != nil {
		return x.FreshBlood
	}
	return false
}

func (x *LeagueItem) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *LeagueItem) GetSummonerName() string {
	if x != nil {
		return x.SummonerName
	}
	return ""
}

func (x *LeagueItem) GetMiniSeries() *MiniSeries {
	if x != nil {
		return x.MiniSeries
	}
	return nil
}

func (x *LeagueItem) GetInactive() bool {
	if x != nil {
		return x.Inactive
	}
	return false
}

func (x *LeagueItem) GetVeteran() bool {
	if x != nil {
		return x.Veteran
	}
	return false
}

func (x *LeagueItem) GetHotStreak() bool {
	if x != nil {
		return x.HotStreak
	}
	return false
}

func (x *LeagueItem) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

func (x *LeagueItem) GetLeaguePoints() int32 {
	if x != nil {
		return x.LeaguePoints
	}
	return 0
}

func (x *LeagueItem) GetLosses() int32 {
	if x != nil {
		return x.Losses
	}
	return 0
}

func (x *LeagueItem) GetSummonerId() string {
	if x != nil {
		return x.SummonerId
	}
	return ""
}

type MiniSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wins     int32  `protobuf:"varint,1,opt,name=wins,proto3" json:"wins,omitempty"`
	Losses   int32  `protobuf:"varint,2,opt,name=losses,proto3" json:"losses,omitempty"`
	Target   int32  `protobuf:"varint,3,opt,name=target,proto3" json:"target,omitempty"`
	Progress string `protobuf:"bytes,4,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *MiniSeries) Reset() {
	*x = MiniSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riot_lol_league_v1_league_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiniSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiniSeries) ProtoMessage() {}

func (x *MiniSeries) ProtoReflect() protoreflect.Message {
	mi := &file_riot_lol_league_v1_league_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiniSeries.ProtoReflect.Descriptor instead.
func (*MiniSeries) Descriptor() ([]byte, []int) {
	return file_riot_lol_league_v1_league_proto_rawDescGZIP(), []int{8}
}

func (x *MiniSeries) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *MiniSeries) GetLosses() int32 {
	if x != nil {
		return x.Losses
	}
	return 0
}

func (x *MiniSeries) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *MiniSeries) GetProgress() string {
	if x != nil {
		return x.Progress
	}
	return ""
}

var File_riot_lol_league_v1_league_proto protoreflect.FileDescriptor

var file_riot_lol_league_v1_league_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x69, 0x6f, 0x74, 0x2f, 0x6c, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67,
	0x75, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x41, 0x0a, 0x11, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x41, 0x70, 0x65, 0x78, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x51, 0x0a, 0x17, 0x52, 0x69, 0x6f, 0x74,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0f, 0x52,
	0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65,
	0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x52, 0x0a, 0x11, 0x52,
	0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x3d, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x22,
	0xbb, 0x03, 0x0a, 0x0f, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x77, 0x69, 0x6e,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x6f,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x74, 0x65, 0x72,
	0x61, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x74, 0x65, 0x72, 0x61,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x6c, 0x6f, 0x6f,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x3e, 0x0a,
	0x0a, 0x6d, 0x69, 0x6e, 0x69, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x69, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0xa4, 0x01,
	0x0a, 0x0e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x22, 0xe8, 0x02, 0x0a, 0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x6c, 0x6f, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x6c,
	0x6f, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x6f,
	0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x6d,
	0x69, 0x6e, 0x69, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x0a, 0x6d, 0x69, 0x6e, 0x69, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x74, 0x65, 0x72,
	0x61, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x74, 0x65, 0x72, 0x61,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x6c, 0x0a, 0x0a, 0x4d, 0x69, 0x6e, 0x69, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x77, 0x69, 0x6e,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x32, 0xb6, 0x04,
	0x0a, 0x11, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x12, 0x25, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x41, 0x70, 0x65, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e,
	0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69,
	0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x5a,
	0x0a, 0x0b, 0x47, 0x72, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e,
	0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x41, 0x70, 0x65,
	0x78, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e,
	0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x06, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e,
	0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x41, 0x70, 0x65, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x72, 0x69,
	0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x00, 0x12, 0x62, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x2b, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x72,
	0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49,
	0x64, 0x12, 0x23, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f,
	0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x05,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c,
	0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c,
	0x65, 0x61, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e,
	0x72, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x69, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x6d, 0x65, 0x65, 0x70, 0x2e, 0x67,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x69, 0x6f, 0x74, 0x2f, 0x6c, 0x6f,
	0x6c, 0x2f, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_riot_lol_league_v1_league_proto_rawDescOnce sync.Once
	file_riot_lol_league_v1_league_proto_rawDescData = file_riot_lol_league_v1_league_proto_rawDesc
)

func file_riot_lol_league_v1_league_proto_rawDescGZIP() []byte {
	file_riot_lol_league_v1_league_proto_rawDescOnce.Do(func() {
		file_riot_lol_league_v1_league_proto_rawDescData = protoimpl.X.CompressGZIP(file_riot_lol_league_v1_league_proto_rawDescData)
	})
	return file_riot_lol_league_v1_league_proto_rawDescData
}

var file_riot_lol_league_v1_league_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_riot_lol_league_v1_league_proto_goTypes = []interface{}{
	(*RiotLeagueApexReq)(nil),       // 0: riot.lol.league.v1.RiotLeagueApexReq
	(*RiotLeagueSummonerIdReq)(nil), // 1: riot.lol.league.v1.RiotLeagueSummonerIdReq
	(*RiotLeagueIdReq)(nil),         // 2: riot.lol.league.v1.RiotLeagueIdReq
	(*RiotLeagueEntryReq)(nil),      // 3: riot.lol.league.v1.RiotLeagueEntryReq
	(*RiotLeagueEntries)(nil),       // 4: riot.lol.league.v1.RiotLeagueEntries
	(*RiotLeagueEntry)(nil),         // 5: riot.lol.league.v1.RiotLeagueEntry
	(*RiotLeagueList)(nil),          // 6: riot.lol.league.v1.RiotLeagueList
	(*LeagueItem)(nil),              // 7: riot.lol.league.v1.LeagueItem
	(*MiniSeries)(nil),              // 8: riot.lol.league.v1.MiniSeries
}
var file_riot_lol_league_v1_league_proto_depIdxs = []int32{
	5,  // 0: riot.lol.league.v1.RiotLeagueEntries.leagues:type_name -> riot.lol.league.v1.RiotLeagueEntry
	8,  // 1: riot.lol.league.v1.RiotLeagueEntry.miniSeries:type_name -> riot.lol.league.v1.MiniSeries
	7,  // 2: riot.lol.league.v1.RiotLeagueList.entries:type_name -> riot.lol.league.v1.LeagueItem
	8,  // 3: riot.lol.league.v1.LeagueItem.miniSeries:type_name -> riot.lol.league.v1.MiniSeries
	0,  // 4: riot.lol.league.v1.RiotLeagueService.Challenger:input_type -> riot.lol.league.v1.RiotLeagueApexReq
	0,  // 5: riot.lol.league.v1.RiotLeagueService.Grandmaster:input_type -> riot.lol.league.v1.RiotLeagueApexReq
	0,  // 6: riot.lol.league.v1.RiotLeagueService.Master:input_type -> riot.lol.league.v1.RiotLeagueApexReq
	1,  // 7: riot.lol.league.v1.RiotLeagueService.SummonerId:input_type -> riot.lol.league.v1.RiotLeagueSummonerIdReq
	2,  // 8: riot.lol.league.v1.RiotLeagueService.LeagueId:input_type -> riot.lol.league.v1.RiotLeagueIdReq
	3,  // 9: riot.lol.league.v1.RiotLeagueService.Entry:input_type -> riot.lol.league.v1.RiotLeagueEntryReq
	6,  // 10: riot.lol.league.v1.RiotLeagueService.Challenger:output_type -> riot.lol.league.v1.RiotLeagueList
	6,  // 11: riot.lol.league.v1.RiotLeagueService.Grandmaster:output_type -> riot.lol.league.v1.RiotLeagueList
	6,  // 12: riot.lol.league.v1.RiotLeagueService.Master:output_type -> riot.lol.league.v1.RiotLeagueList
	4,  // 13: riot.lol.league.v1.RiotLeagueService.SummonerId:output_type -> riot.lol.league.v1.RiotLeagueEntries
	6,  // 14: riot.lol.league.v1.RiotLeagueService.LeagueId:output_type -> riot.lol.league.v1.RiotLeagueList
	4,  // 15: riot.lol.league.v1.RiotLeagueService.Entry:output_type -> riot.lol.league.v1.RiotLeagueEntries
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_riot_lol_league_v1_league_proto_init() }
func file_riot_lol_league_v1_league_proto_init() {
	if File_riot_lol_league_v1_league_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_riot_lol_league_v1_league_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiotLeagueApexReq); i {
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
		file_riot_lol_league_v1_league_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiotLeagueSummonerIdReq); i {
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
		file_riot_lol_league_v1_league_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiotLeagueIdReq); i {
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
		file_riot_lol_league_v1_league_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiotLeagueEntryReq); i {
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
		file_riot_lol_league_v1_league_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiotLeagueEntries); i {
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
		file_riot_lol_league_v1_league_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiotLeagueEntry); i {
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
		file_riot_lol_league_v1_league_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiotLeagueList); i {
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
		file_riot_lol_league_v1_league_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeagueItem); i {
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
		file_riot_lol_league_v1_league_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiniSeries); i {
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
			RawDescriptor: file_riot_lol_league_v1_league_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_riot_lol_league_v1_league_proto_goTypes,
		DependencyIndexes: file_riot_lol_league_v1_league_proto_depIdxs,
		MessageInfos:      file_riot_lol_league_v1_league_proto_msgTypes,
	}.Build()
	File_riot_lol_league_v1_league_proto = out.File
	file_riot_lol_league_v1_league_proto_rawDesc = nil
	file_riot_lol_league_v1_league_proto_goTypes = nil
	file_riot_lol_league_v1_league_proto_depIdxs = nil
}
