package h264

const eAVEncH264VProfile_unknown = 0
const eAVEncH264VProfile_Simple = 66
const eAVEncH264VProfile_Base = 66
const eAVEncH264VProfile_Main = 77
const eAVEncH264VProfile_High = 100
const eAVEncH264VProfile_422 = 122
const eAVEncH264VProfile_High10 = 110
const eAVEncH264VProfile_444 = 244
const eAVEncH264VProfile_Extended = 88
const eAVEncH264VProfile_ScalableBase = 83
const eAVEncH264VProfile_ScalableHigh = 86
const eAVEncH264VProfile_MultiviewHigh = 118
const eAVEncH264VProfile_StereoHigh = 128
const eAVEncH264VProfile_ConstrainedBase = 256
const eAVEncH264VProfile_UCConstrainedHigh = 257
const eAVEncH264VProfile_UCScalableConstrainedBase = 258
const eAVEncH264VProfile_UCScalableConstrainedHigh = 259

type H264ProfileInfo struct {
	profileIndex int
	fps          float64
	frame_size   float64
	bitrate      int
}

var profiles [5]H264ProfileInfo = [...]H264ProfileInfo{
	{eAVEncH264VProfile_Base, 15, 176 / 144, 128000},
	{},
	{},
	{},
	{},
}

