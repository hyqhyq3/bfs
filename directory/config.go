package main

import (
	"github.com/Terry-Mao/goconf"
	"time"
)

const (
	// snowflake zookeeper
	configSnowflakeZkTimeout = time.Second * 15 // 1s
	configSnowflakeZkPath    = "/bfs_sf"
	configSnowflakeWorkId    = 0
	// zookeeper
	configZkTimeout    = time.Second * 15 // 1s
	configZkStoreRoot  = "/rack"
	configZkVolumeRoot = "/volume"
	configZkGroupRoot  = "/group"
	configPullInterval = time.Second * 10
	// hbase
	configHbaseAddr      = "localhost:9090"
	configHbaseMaxActive = 50
	configHbaseMaxIdle   = 10
	configHbaseTimeout   = time.Second * 5
	// http
	configApiListen = "localhost:6065"
	configMaxNum    = 16
	// pprof
	configPprofListen = "localhost:6066"
)

var (
	configSnowflakeZkAddrs = []string{"localhost:2181"}
	configZkAddrs          = []string{"localhost:2181"}
)

type Config struct {
	// snowflake zookeeper
	SnowflakeZkAddrs   []string      `goconf:"snowflake:sfzkaddrs:,"`
	SnowflakeZkTimeout time.Duration `goconf:"snowflake:sfzktimeout"`
	SnowflakeZkPath    string        `goconf:"snowflake:sfzkpath"`
	SnowflakeWorkId    int64         `goconf:"snowflake:workid"`
	// bfs zookeeper
	ZkAddrs      []string      `goconf:"zookeeper:addrs:,"`
	ZkTimeout    time.Duration `goconf:"zookeeper:timeout:,"`
	ZkStoreRoot  string        `goconf:"zookeeper:storeroot"`
	ZkVolumeRoot string        `goconf:"zookeeper:volumeroot"`
	ZkGroupRoot  string        `goconf:"zookeeper:grouproot"`
	PullInterval time.Duration `goconf:"zookeeper:pullinterval"`
	// hbase
	HbaseAddr      string        `goconf:"hbase:addr"`
	HbaseMaxActive int           `goconf:"hbase:max.active"`
	HbaseMaxIdle   int           `goconf:"hbase:max.idle"`
	HbaseTimeout   time.Duration `goconf:"hbase:timeout:time"`
	// http
	MaxNum    int    `goconf:"http:maxnum"`
	ApiListen string `goconf:"http:apilisten"`
	// pprof
	PprofEnable bool   `goconf:"pprof:enable"`
	PprofListen string `goconf:"pprof:listen"`
}

// NewConfig new a config.
func NewConfig(file string) (c *Config, err error) {
	var gconf = goconf.New()
	c = &Config{}
	if err = gconf.Parse(file); err != nil {
		return
	}
	if err = gconf.Unmarshal(c); err != nil {
		return
	}
	c.setDefault()
	return
}

// setDefault set the config default value.
func (c *Config) setDefault() {
	if len(c.SnowflakeZkAddrs) == 0 {
		c.SnowflakeZkAddrs = configSnowflakeZkAddrs
	}
	if c.SnowflakeZkTimeout < 1*time.Second {
		c.SnowflakeZkTimeout = configSnowflakeZkTimeout
	}
	if c.SnowflakeZkPath == "" {
		c.SnowflakeZkPath = configSnowflakeZkPath
	}
	if c.SnowflakeWorkId == 0 {
		c.SnowflakeWorkId = configSnowflakeWorkId
	}
	if len(c.ZkAddrs) == 0 {
		c.ZkAddrs = configZkAddrs
	}
	if c.ZkTimeout < 1*time.Second {
		c.ZkTimeout = configZkTimeout
	}
	if c.ZkStoreRoot == "" {
		c.ZkStoreRoot = configZkStoreRoot
	}
	if c.ZkVolumeRoot == "" {
		c.ZkVolumeRoot = configZkVolumeRoot
	}
	if c.ZkGroupRoot == "" {
		c.ZkGroupRoot = configZkGroupRoot
	}
	if c.PullInterval < 1*time.Second {
		c.PullInterval = configPullInterval
	}
	if c.HbaseAddr == "" {
		c.HbaseAddr = configHbaseAddr
	}
	if c.HbaseMaxActive == 0 {
		c.HbaseMaxActive = configHbaseMaxActive
	}
	if c.HbaseMaxIdle == 0 {
		c.HbaseMaxIdle = configHbaseMaxIdle
	}
	if c.HbaseTimeout < 1*time.Second {
		c.HbaseTimeout = configHbaseTimeout
	}
	if c.MaxNum == 0 || c.MaxNum > configMaxNum {
		c.MaxNum = configMaxNum
	}
	if c.ApiListen == "" {
		c.ApiListen = configApiListen
	}
	if len(c.PprofListen) == 0 {
		c.PprofListen = configPprofListen
	}
}
