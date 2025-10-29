package scheduler

import (
	"fairytale-creator/logger"
	"fairytale-creator/service"
	"fmt"

	"github.com/robfig/cron/v3"
)

var cronInstance *cron.Cron

// StartScheduler 启动定时任务
func StartScheduler() {
	// 创建cron实例，支持到秒级别的定时
	cronInstance = cron.New(cron.WithSeconds())

	// 每天凌晨2点执行一次生成故事任务
	// "0 0 2 * * ?" 表示每天凌晨2点0分0秒执行
	_, err := cronInstance.AddFunc("0 0 2 * * ?", generateDailyStory)
	if err != nil {
		logger.Error("添加定时任务失败: " + err.Error())
		return
	}

	cronInstance.Start()
	logger.Log("定时任务调度器已启动")

	// 可以选择是否立即执行一次（用于测试）
	// generateDailyStory()
}

// StopScheduler 停止定时任务
func StopScheduler() {
	if cronInstance != nil {
		cronInstance.Stop()
		logger.Log("定时任务调度器已停止")
	}
}

// generateDailyStory 生成每日故事
func generateDailyStory() {
	logger.Log("开始执行每日故事生成任务...")

	storyService := service.NewStoryService()
	story := storyService.GenerateStory()

	if story == nil {
		logger.Error("生成故事失败")
		return
	}

	err := storyService.AddStory(story)
	if err != nil {
		logger.Error("添加故事失败: " + err.Error())
		return
	}

	logger.Log(fmt.Sprintf("每日故事生成成功: %s", story.Title))
}
