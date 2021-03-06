package consts

import "time"

const (
	APPLYDISSOLVE_DURATION               = time.Second * 2 * 60 //申请解散房间 等待的时间
	COIN_RAEDY_DURATION                  = time.Second * 60     //金币场，一局结束之后，需要准备，超时不准备需要提出房间
	MJ_TICK_SECOND                 int32 = 15                   //朋友桌麻将打牌定时秒数 给前端显示用
	MJ_TICK_SECOND_COINROOM        int32 = 30                   //金币场麻将打牌定时秒数 给前端显示用
	MJ_TICK_SECOND_COINROOM_HZH    int32 = 20                   //红中金币场麻将打牌定时秒数 给前端显示用
	MJ_TICK_SECOND_COINROOM_ZIGONG int32 = 40                   //自贡金币场麻将打牌定时秒数 给前端显示用
)
