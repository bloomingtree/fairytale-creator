// cmd/test_d1/main.go
package main

import (
	"bufio"
	"fairytale-creator/modelapi"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			val = strings.Trim(val, " \t\"'")
			if os.Getenv(key) == "" {
				_ = os.Setenv(key, val)
			}
		}
	}
	return nil
}

func main() {
	// 加载 .env 文件（从项目根目录）
	_ = loadEnvFile(".env")
	accountID := os.Getenv("CF_ACCOUNT_ID")
	dbID := os.Getenv("D1_DATABASE_ID")
	apiKey := os.Getenv("D1_API_KEY")
	if accountID == "" || dbID == "" || apiKey == "" {
		fmt.Println("缺少环境变量：CF_ACCOUNT_ID / D1_DATABASE_ID / D1_API_KEY")
		os.Exit(1)
	}
	client := modelapi.NewD1Client(accountID, dbID, apiKey)
	resp, err := client.ExecuteQuery("SELECT 1;", nil)
	if err != nil {
		fmt.Println("请求失败:", err)
		os.Exit(2)
	}
	if !resp.Success {
		fmt.Println("D1 不可用或不存在:", resp.Errors)
		os.Exit(3)
	}
	fmt.Println("D1 可用 ✅")
	if len(resp.Result) > 0 {
		fmt.Printf("查询结果: %+v\n", resp.Result[0].Results)
	}
}
