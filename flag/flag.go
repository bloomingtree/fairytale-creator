package flag

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	Username              string
	Password              string
	VideoRoot             string
	DeepSeekAPIKey        string
	DeepSeekUrl           string
	JimengAccessKeyID     string
	JimengSecretAccessKey string
	StoryRoot             string
	ImageRoot             string
	CosyVoiceAPIKey       string
	VoiceRoot             string
	DoubaoSeedreamAPIKey  string
	MysqlUsername         string
	MysqlPassword         string
	MysqlHost             string
	MysqlPort             string
	MysqlDatabase         string
	CfAccountID           string
	D1DatabaseID          string
	D1Email               string
	D1APIKey              string
	R2AccessKeyID         string
	R2AccessKeySecret     string
)

var once sync.Once

func init() {
	once.Do(func() {
		_ = loadEnvFile(".env")
		Username = getEnvOrDefault("USERNAME", "admin")
		Password = getEnvOrDefault("PASSWORD", "20240316")
		VideoRoot = getEnvOrDefault("VIDEO_ROOT", "")
		DeepSeekAPIKey = getEnvOrDefault("DEEPSEEK_API_KEY", "")
		DeepSeekUrl = getEnvOrDefault("DEEPSEEK_URL", "https://api.deepseek.com")
		JimengAccessKeyID = getEnvOrDefault("JIMENG_ACCESS_KEY_ID", "")
		JimengSecretAccessKey = getEnvOrDefault("JIMENG_SECRET_ACCESS_KEY", "")
		StoryRoot = getEnvOrDefault("STORY_ROOT", "stories")
		ImageRoot = getEnvOrDefault("IMAGE_ROOT", "images")
		CosyVoiceAPIKey = getEnvOrDefault("COSY_VOICE_API_KEY", "")
		VoiceRoot = getEnvOrDefault("VOICE_ROOT", "voices")
		DoubaoSeedreamAPIKey = getEnvOrDefault("DOUBAO_SEEDREAM_API_KEY", "")
		MysqlUsername = getEnvOrDefault("MYSQL_USERNAME", "admin")
		MysqlPassword = getEnvOrDefault("MYSQL_PASSWORD", "20240316")
		MysqlHost = getEnvOrDefault("MYSQL_HOST", "localhost")
		MysqlPort = getEnvOrDefault("MYSQL_PORT", "3306")
		MysqlDatabase = getEnvOrDefault("MYSQL_DATABASE", "fairytale")
		CfAccountID = getEnvOrDefault("CF_ACCOUNT_ID", "")
		D1DatabaseID = getEnvOrDefault("D1_DATABASE_ID", "")
		D1Email = getEnvOrDefault("D1_EMAIL", "")
		D1APIKey = getEnvOrDefault("D1_API_KEY", "")
		R2AccessKeyID = getEnvOrDefault("R2_ACCESS_KEY_ID", "")
		R2AccessKeySecret = getEnvOrDefault("R2_ACCESS_KEY_SECRET", "")
	})
}

func getEnvOrDefault(key, def string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return def
	}
	return v
}

func loadEnvFile(filename string) error {
	abs, err := filepath.Abs(filename)
	if err == nil {
		filename = abs
	}
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// 使用 SplitN 只分割第一个 `=`，支持值中包含等号
		// 例如: KEY=value=with=equals=sign
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			// 去掉首尾的引号（支持 "value" 或 'value'）
			val = strings.Trim(val, " \t\"'")
			if os.Getenv(key) == "" { // do not override existing env
				_ = os.Setenv(key, val)
			}
		}
	}
	return nil
}
