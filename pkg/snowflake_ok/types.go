package snowflake_ok

type SnowflakeIF interface {
	GenSnowFlakeID() int64
}
