package ms_time

import (
	"time"
)

var MSTIME_COMMON_FORMAT = "2006-01-02 15:04:05"

type MsTime struct {
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
	Sec   int
	NSec  int
}

func CreateMsTime(inTime time.Time) (otMSTime *MsTime) {
	return &MsTime{
		Year:  inTime.Year(),
		Month: int(inTime.Month()),
		Day:   inTime.Day(),
		Hour:  inTime.Hour(),
		Min:   inTime.Minute(),
		Sec:   inTime.Second(),
		NSec:  inTime.Nanosecond(),
	}
}

func CreateMsTimeByStr(inTimeStr, inFormat string) (otMSTime *MsTime, err error) {
	theTime, err := time.Parse(inFormat, inTimeStr)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}

	return CreateMsTime(theTime), nil
}

func (this *MsTime) ToTime() (otRet time.Time) {
	otRet = time.Date(this.Year, time.Month(this.Month), this.Day, this.Hour, this.Min, this.Sec,
		this.NSec, time.Now().Location())

	return otRet
}

func (this *MsTime) ToUnix() (otRet int64) {
	return this.ToTime().Unix()
}

func (this *MsTime) Compare(inMSTimeComp *MsTime) (otDiff int64) {
	return this.ToUnix() - inMSTimeComp.ToUnix()
}

func (this *MsTime) ToString() (otTimeStr string) {
	timeMid := this.ToTime()
	otTimeStr = timeMid.Format(MSTIME_COMMON_FORMAT)
	return otTimeStr
}
