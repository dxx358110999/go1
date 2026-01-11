package app_config

type MysqlConfig struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DB           string `yaml:"dbname"`
	Port         int    `yaml:"port"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `yaml:"host"`
	Password     string `yaml:"password"`
	Port         int    `yaml:"port"`
	DB           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

type AliSms struct {
	AKId     string `yaml:"ak_id"`
	AKSecret string `yaml:"ak_secret"`
	Host     string `yaml:"host"`
}

type JwtUser struct {
	AccessSecret  string `yaml:"access_secret"`
	AccessExpire  int64  `yaml:"access_expire"`
	RefreshSecret string `yaml:"refresh_secret"`
	RefreshExpire int64  `yaml:"refresh_expire"`
}

type EmailSmtp struct {
	SmtpHost    string `yaml:"host"`         // QQ邮箱SMTP服务器地址
	SmtpPort    int    `yaml:"port"`         // SMTP端口：465（SSL）或 587（TLS）
	SenderEmail string `yaml:"sender_email"` // 发件人邮箱（你的QQ邮箱）
	AuthCode    string `yaml:"auth_code"`    // QQ邮箱授权码（不是登录密码）
}

type Config struct {
	Name      string `yaml:"name"`
	Mode      string `yaml:"mode"`
	Version   string `yaml:"version"`
	StartTime string `yaml:"start_time"`
	Port      uint64 `yaml:"port"`

	*LogConfig   `yaml:"log"`
	*MysqlConfig `yaml:"mysql"`
	*RedisConfig `yaml:"redis"`
	*AliSms      `yaml:"ali_sms"`
	*JwtUser     `yaml:"jwt_user"`
	*EmailSmtp   `yaml:"email_smtp"`
}
