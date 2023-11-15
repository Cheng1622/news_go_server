package snowflake

import (
	"fmt"
	"sync"
	"time"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
)

const (
	epoch             = int64(1609459200000) // 设置一个起始时间，例如：2021-01-01 00:00:00
	workerIDBits      = uint(5)              // 工作机器ID所占的位数
	datacenterIDBits  = uint(5)              // 数据中心ID所占的位数
	sequenceBits      = uint(12)             // 序列号所占的位数
	maxWorkerID       = int64(-1) ^ (int64(-1) << workerIDBits)
	maxDatacenterID   = int64(-1) ^ (int64(-1) << datacenterIDBits)
	maxSequence       = int64(-1) ^ (int64(-1) << sequenceBits)
	workerIDShift     = sequenceBits
	datacenterIDShift = sequenceBits + workerIDBits
	timestampShift    = sequenceBits + workerIDBits + datacenterIDBits
)

var SF *Snowflake

type Snowflake struct {
	mu            sync.Mutex
	workerID      int64
	datacenterID  int64
	sequence      int64
	lastTimestamp int64
}

func NewSnowflake(workerID, datacenterID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, fmt.Errorf("worker ID must be between 0 and %d", maxWorkerID)
	}
	if datacenterID < 0 || datacenterID > maxDatacenterID {
		return nil, fmt.Errorf("datacenter ID must be between 0 and %d", maxDatacenterID)
	}
	return &Snowflake{
		workerID:      workerID,
		datacenterID:  datacenterID,
		sequence:      0,
		lastTimestamp: -1,
	}, nil
}

// GenerateID 获得雪花算法ID
func (s *Snowflake) GenerateID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixNano() / 1000000 // 转换为毫秒级时间戳

	if timestamp < s.lastTimestamp {
		return 0, fmt.Errorf("时钟向后移动")
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	id := ((timestamp - epoch) << timestampShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id, nil
}

func InitSnowflake() {
	var err error
	SF, err = NewSnowflake(config.Conf.SnowFlake.WorkerID, config.Conf.SnowFlake.DatacenterID) // 设置workerID和datacenterID
	if err != nil {
		clog.Log.Fatalln("初始化Snowflake失败:", err)
	}
	clog.Log.Infoln("初始化Snowflake成功")
}
