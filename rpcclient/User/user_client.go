package client

import (
	"context"

	"banban/app/model"
	"banban/app/pb"
)

// UserProfile 定一个userProfile单例
var UserProfile *userProfile = &userProfile{
	&base{
		name: "User.Profile",
	},
}

type userProfile struct {
	*base
}

// Auth 任何请求
func (serv *userProfile) Auth(ctx context.Context, token string) (*pb.RepUserAuth, error) {
	req := &pb.ReqUserAuth{
		Token: token,
	}
	reply := &pb.RepUserAuth{}
	err := serv.call(ctx, "Auth", req, reply)
	return reply, err
}

// Token 生成token
func (serv *userProfile) Token(ctx context.Context, uid uint32, appId model.AppID, platform string, channel ...string) (*pb.RepUserToken, error) {
	ch := ""
	if len(channel) > 0 {
		ch = channel[0]
	}
	req := &pb.ReqUserToken{
		Uid:      uid,
		AppId:    uint32(appId),
		Platform: platform,
		Channel:  ch,
	}
	reply := &pb.RepUserToken{}
	err := serv.call(ctx, "Token", req, reply)
	return reply, err
}

func (serv *userProfile) GetBadgeAlloc(ctx context.Context, allocId uint32, fields []string) (*pb.EntityBbcBadgeAllocate, error) {
	req := &pb.GetBadgeAllocReq{
		AllocId: allocId,
		Fields:  fields,
	}
	reply := &pb.EntityBbcBadgeAllocate{}
	err := serv.call(ctx, "GetBadgeAlloc", req, reply)
	return reply, err
}

func (serv *userProfile) MgetBadgeAllocs(ctx context.Context, allocIds []uint32, fields []string) ([]*pb.EntityBbcBadgeAllocate, error) {
	req := &pb.MGetBadgeAllocReq{
		AllocIds: allocIds,
		Fields:   fields,
	}
	reply := &pb.MGetBadgesAllocResp{}
	err := serv.call(ctx, "MgetBadgeAllocs", req, reply)
	return reply.Allocs, err
}

func (serv *userProfile) MgetBadgeAllocsFromCache(ctx context.Context, allocIds []uint32, fields []string) ([]*pb.EntityBbcBadgeAllocate, error) {
	req := &pb.MGetBadgeAllocReq{
		AllocIds: allocIds,
		Fields:   fields,
	}
	reply := &pb.MGetBadgesAllocResp{}
	err := serv.call(ctx, "MgetBadgeAllocsFromCache", req, reply)
	return reply.Allocs, err
}

func (serv *userProfile) GetBadge(ctx context.Context, badgeId uint32, fields []string) (*pb.EntityBbcBadge, error) {
	req := &pb.GetBadgeReq{
		BadgeId: badgeId,
		Fields:  fields,
	}
	reply := &pb.EntityBbcBadge{}
	err := serv.call(ctx, "GetBadge", req, reply)
	return reply, err
}

func (serv *userProfile) MgetBadges(ctx context.Context, badgeIds []uint32, fields []string) ([]*pb.EntityBbcBadge, error) {
	req := &pb.MGetBadgeReq{
		BadgeIds: badgeIds,
		Fields:   fields,
	}
	reply := &pb.MGetBadgesResp{}
	err := serv.call(ctx, "MgetBadges", req, reply)
	return reply.Badges, err
}

// Get 获取单个用户信息请求
func (serv *userProfile) Get(ctx context.Context, uid uint32, fields []string) (*pb.EntityXsUserProfile, error) {
	req := &pb.ReqUserProfile{
		Uid:    uid,
		Fields: fields,
	}
	reply := &pb.EntityXsUserProfile{}
	err := serv.call(ctx, "Get", req, reply)
	return reply, err
}

// Mget 批量获取多个用户信息请求
func (serv *userProfile) Mget(ctx context.Context, uids []uint32, fields []string) ([]*pb.EntityXsUserProfile, error) {
	req := &pb.ReqUserProfiles{
		Uids:   uids,
		Fields: fields,
	}
	reply := &pb.RepUserProfiles{}
	err := serv.call(ctx, "Mget", req, reply)
	return reply.Data, err
}

func (serv *userProfile) MgetMap(ctx context.Context, uids []uint32, fields []string) (map[uint32]*pb.EntityXsUserProfile, error) {
	res, err := serv.Mget(ctx, uids, fields)
	data := make(map[uint32]*pb.EntityXsUserProfile)
	if err == nil {
		for _, item := range res {
			data[item.Uid] = item
		}
	}
	return data, err
}

// GetUsersTitleLevelMap 批量获取贵族等级信息
func (serv *userProfile) GetUsersTitleLevelMap(ctx context.Context, uids []uint32) (map[uint32]uint32, error) {
	req := &pb.ReqUserTitleLevels{
		Uids: uids,
	}
	reply := &pb.RepUserTitleLevels{}
	err := serv.call(ctx, "GetUsersTitleLevelMap", req, reply)
	if err != nil {
		return nil, err
	}
	return reply.Data, err
}

// GetUserDecorate 获取用户挂件信息
func (serv *userProfile) GetUserDecorate(ctx context.Context, uid uint32) (*pb.EntityXsGift, error) {
	req := &pb.ReqUid{
		Uid: uid,
	}
	reply := &pb.EntityXsGift{}
	err := serv.call(ctx, "DecorateGet", req, reply)
	if err != nil {
		return nil, err
	}
	return reply, err
}

// GetUserIpInfo 获取ip信息
func (serv *userProfile) GetUserIpInfo(ctx context.Context, req *pb.ReqGetUserIpInfo) (*pb.RepGetUserIpInfo, error) {

	reply := &pb.RepGetUserIpInfo{}
	err := serv.call(ctx, "GetUserIpInfo", req, reply)
	if err != nil {
		return reply, err
	}
	return reply, err
}

func (serv *userProfile) FaceDet(ctx context.Context, req *pb.ReqFaceDetect) (*pb.RepFaceDetect, error) {
	reply := &pb.RepFaceDetect{}
	err := serv.call(ctx, "FaceDetect", req, reply)
	if err != nil {
		return reply, err
	}
	return reply, err
}

// GetGiftPanel 获取面板
func (serv *userProfile) GetGiftPanel(ctx context.Context, req *pb.ReqGiftPanel) (*[]pb.EntityXsGift, error) {

	reply := make([]pb.EntityXsGift, 10)
	err := serv.call(ctx, "GetPanelGift", req, reply)
	if err != nil {
		return &reply, err
	}
	return &reply, err
}

// 装配间礼物刷新
func (serv *userProfile) AssemblyGift(ctx context.Context, req *pb.ReqGiftAssembly) (*pb.RepGiftAssembly, error) {
	reply := &pb.RepGiftAssembly{}
	err := serv.call(ctx, "RefreshAssemblyGift", req, reply)
	return reply, err
}

// GetSetting 获取用户Settings
func (serv *userProfile) GetSetting(ctx context.Context, req *pb.ReqUserProfile) (*pb.EntityXsUserSettings, error) {

	reply := &pb.EntityXsUserSettings{}
	err := serv.call(ctx, "GetSetting", req, reply)
	if err != nil {
		return reply, err
	}
	return reply, err
}

func (serv *userProfile) SearchByName(ctx context.Context, keyword string, appId, searchLevel uint32, limit uint32) (*pb.RepUserSearchDefault, error) {
	req := &pb.ReqUserSearchName{
		Keyword:       keyword,
		AppId:         []uint32{appId},
		Limit:         limit,
		SearcherLevel: searchLevel,
	}
	reply := &pb.RepUserSearchDefault{}
	err := serv.call(ctx, "SearchByName", req, reply)
	if err != nil {
		return nil, err
	}
	return reply, err
}

func (serv *userProfile) SearchFleet(ctx context.Context, keyword string, limit uint32) (*pb.RepSearchFleet, error) {
	req := &pb.ReqFleetSearchName{
		Keyword: keyword,
		Limit:   limit,
	}
	reply := &pb.RepSearchFleet{}
	err := serv.call(ctx, "SearchFleet", req, reply)
	if err != nil {
		return nil, err
	}
	return reply, err
}

func (serv *userProfile) SearchUnion(ctx context.Context, keyword string, batchSize uint32) (*pb.RepSearchUnion, error) {
	req := &pb.ReqFleetSearchName{
		Keyword: keyword,
		Limit:   batchSize,
	}
	reply := &pb.RepSearchUnion{}
	err := serv.call(ctx, "SearchUnion", req, reply)
	if err != nil {
		return nil, err
	}
	return reply, err
}

// NeedVerify 最新版本权限控制
func (serv *userProfile) NeedVerify(ctx context.Context, req *pb.ReqNeedVerify) (*pb.RepNeedVerify, error) {
	reply := &pb.RepNeedVerify{}
	err := serv.call(ctx, "NeedVerify", req, reply)
	return reply, err
}

// UnionRoomFilter 联盟房过滤
func (serv *userProfile) UnionRoomFilter(ctx context.Context, req *pb.ReqUnionRoomFilter) (*pb.RepUnionRoomFilter, error) {
	reply := &pb.RepUnionRoomFilter{}
	err := serv.call(ctx, "UnionRoomFilter", req, reply)
	return reply, err
}

func (serv *userProfile) SearchRoomByName(ctx context.Context, keyword string, appId, limit uint32) (*pb.RepUserSearchDefault, error) {
	req := &pb.ReqUserSearchName{
		Keyword: keyword,
		AppId:   []uint32{appId},
		Limit:   limit,
	}
	reply := &pb.RepUserSearchDefault{}
	err := serv.call(ctx, "SearchRoomByName", req, reply)
	if err != nil {
		return nil, err
	}
	return reply, err
}

func (serv *userProfile) BubbleTailGet(ctx context.Context, uid uint32) (*pb.RepUserBubbleTail, error) {
	req := &pb.ReqUid{Uid: uid}
	reply := &pb.RepUserBubbleTail{}
	err := serv.call(ctx, "BubbleTailGet", req, reply)
	if err != nil {
		return nil, err
	}
	return reply, err
}

// SingerClubRoomFilter 歌友房过滤
func (serv *userProfile) SingerClubRoomFilter(ctx context.Context, req *pb.ReqSingerClubRoomFilter) (*pb.RepSingerClubRoomFilter, error) {
	reply := &pb.RepSingerClubRoomFilter{}
	err := serv.call(ctx, "SingerClubRoomFilter", req, reply)
	return reply, err
}

// GetUserCardDecorate 获取房间资料卡装饰
func (serv *userProfile) GetUserCardDecorate(ctx context.Context, uid uint32) (*pb.RepUserCardDecorate, error) {
	req := &pb.ReqUid{
		Uid: uid,
	}
	reply := &pb.RepUserCardDecorate{}
	err := serv.call(ctx, "CardDecorateGet", req, reply)
	if err != nil {
		return nil, err
	}
	return reply, err
}
