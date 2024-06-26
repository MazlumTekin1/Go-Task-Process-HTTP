package service

import (
	"sort"
	"task-process-service/internal/domain"
)

func DistributeTasks(developers []domain.UserGetDataList, tasks []domain.TaskGetDataList) (map[string][]domain.TaskGetDataList, int) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Difficulty > tasks[j].Difficulty || (tasks[i].Difficulty == tasks[j].Difficulty && tasks[i].Duration > tasks[j].Duration)
	})

	developerWorkloads := make(map[string]float64)
	for _, dev := range developers {
		developerWorkloads[dev.FirstName] = 0
	}

	taskDistribution := make(map[string][]domain.TaskGetDataList)

	// Distribute tasks
	for _, task := range tasks {
		// Find the best developer for the task
		var bestDev string
		minExtraTime := float64(999999)
		for _, dev := range developers {
			if float64(task.Difficulty) <= float64(dev.DeveloperWorkHourDifficulty) {
				extraTime := developerWorkloads[dev.FirstName] + float64(task.Duration)
				if extraTime < minExtraTime {
					minExtraTime = extraTime
					bestDev = dev.FirstName
				}
			}
		}
		// Assign to the fit developer
		taskDistribution[bestDev] = append(taskDistribution[bestDev], task)
		developerWorkloads[bestDev] += float64(task.Duration)
	}

	maxHours := 0.0
	for _, workload := range developerWorkloads {
		if workload > maxHours {
			maxHours = workload
		}
	}
	maxWeeks := int(maxHours / 45)
	if maxHours > float64(maxWeeks*45) {
		maxWeeks++
	}

	return taskDistribution, maxWeeks
}
