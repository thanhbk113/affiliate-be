package config

// ENV ...
type ENV struct {
	Env   string `env:"ENV"`
	Admin struct {
		Server string `env:"ADMIN_SERVER"`
		Port   string `env:"ADMIN_PORT"`
	}
	FileHost      string      `env:"FILE_HOST"`
	AppName       string      `env:"APP_NAME"`
	MongoDB       MongoConfig `env:",prefix=MONGO_"`
	Port          string      `env:"PORT"`
	AuthSecret    AuthSecret  `env:",prefix=AUTH_SECRET_"`
	HostFile      string      `env:"HOST_FILE"`
	Redis         RedisCfg    `env:",prefix=REDIS_"`
	SecretKeyTOTP string      `env:"SecretKeyTOTP"`
	CloudDinary   CLOUDINARY  `env:",prefix=CLOUDINARY_"`
	CRON_URL      string      `env:"CRON_URL"`
	DOMAIN_ADMIN  string      `env:"DOMAIN_ADMIN"`
	PASS_AUTHEN   string      `env:"PASS_AUTHEN"`
}

type CLOUDINARY struct {
	Name       string `env:"NAME"`
	ApiKey     string `env:"API_KEY"`
	ApiSecret  string `env:"API_SECRET"`
	FolderName string `env:"FOLDER_NAME"`
}

type AuthSecret struct {
	Admin  string `env:"ADMIN"`
	Public string `env:"PUBLIC"`
}

// RedisCfg ...
type RedisCfg struct {
	URI      string `env:"URI"`
	Password string `env:"PASSWORD"`
}

// MongoConfig ...
type MongoConfig struct {
	Host   string `env:"HOST"`
	DBName string `env:"DB_NAME"`

	// Standalone
	Mechanism string `env:"MECHANISM"`
	Source    string `env:"SOURCE"`
	Username  string `env:"USERNAME"`
	Password  string `env:"PASSWORD"`
}
