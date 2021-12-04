package sub_timeline_fiexer

type SubTimelineFixerConfig struct {
	// V1 的设置
	MaxCompareDialogue int     // 最大需要匹配的连续对白，默认3
	MaxStartTimeDiffSD float64 // 对白开始时间的统计 SD 最大误差，超过则不进行修正
	MinMatchedPercent  float64 // 两个文件的匹配百分比（src/base），高于这个才比例进行修正
	MinOffset          float64 // 超过这个(+-)偏移的时间轴才校正，否则跳过，单位秒
	// V2 的设置
	SubOneUnitProcessTimeOut int     // 字幕时间轴校正一个单元的超时时间，单位秒
	FrontAndEndPerBase       float64 // 前百分之 15 和后百分之 15 都不进行识别
	FrontAndEndPerSrc        float64 // 前百分之 20 和后百分之 20 都不进行识别
	WindowMatchPer           float64 // SrcSub 滑动窗体的占比
	CompareParts             int     // 滑动窗体分段次数
	FixThreads               int     // 字幕校正的并发线程
}

// CheckDefault 检测默认值（比如某些之默认不能为0），不对就重置到默认值上
func (s *SubTimelineFixerConfig) CheckDefault() {
	// V1
	if s.MaxCompareDialogue <= 0 {
		s.MaxCompareDialogue = 3
	}
	if s.MaxStartTimeDiffSD <= 0 {
		s.MaxStartTimeDiffSD = 0.1
	}
	if s.MinMatchedPercent <= 0 {
		s.MinMatchedPercent = 0.1
	}
	if s.MinOffset <= 0 {
		s.MinOffset = 0.1
	}
	// V2
	if s.SubOneUnitProcessTimeOut <= 0 {
		s.SubOneUnitProcessTimeOut = 30
	}
	if s.FrontAndEndPerBase <= 0 || s.FrontAndEndPerBase >= 1.0 {
		s.FrontAndEndPerBase = 0.15
	}
	if s.FrontAndEndPerSrc <= 0 || s.FrontAndEndPerSrc >= 1.0 {
		s.FrontAndEndPerSrc = 0.0
	}
	if s.WindowMatchPer <= 0 || s.WindowMatchPer >= 1.0 {
		s.WindowMatchPer = 0.7
	}
	if s.CompareParts <= 0 {
		s.CompareParts = 5
	}
	if s.CompareParts <= 0 {
		s.CompareParts = 3
	}
}
