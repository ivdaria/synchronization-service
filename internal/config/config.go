package config

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type HTTPConfig struct {
	ListenAddr string `yaml:"listenAddr"`
}

type DeployWorkerConfig struct {
	CronString string `yaml:"cronString"`
}

type Config struct {
	DBConfig DBConfig           `yaml:"db"`
	HTTP     HTTPConfig         `yaml:"http"`
	Deploy   DeployWorkerConfig `yaml:"deployWorker"`
}
