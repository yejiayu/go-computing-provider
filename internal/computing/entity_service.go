package computing

import (
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/db"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/gorm"
	"time"
)

type TaskService struct {
	*gorm.DB
}

func (taskServ TaskService) GetAllTask() (list []*models.TaskEntity, err error) {
	err = taskServ.Model(&models.TaskEntity{}).Order("create_time").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

func (taskServ TaskService) GetTaskList(taskStatus int) (list []*models.TaskEntity, err error) {
	err = taskServ.Where(&models.TaskEntity{Status: taskStatus}).Order("create_time").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

func (taskServ TaskService) SaveTaskEntity(task *models.TaskEntity) (err error) {
	if task.Status == models.TASK_FAILED_STATUS || task.Status == models.TASK_SUCCESS_STATUS {
		task.EndTime = time.Now().Unix()
	}
	return taskServ.Save(task).Error
}

func (taskServ TaskService) GetTaskEntity(taskId int64) (task *models.TaskEntity, err error) {
	err = taskServ.Where(&models.TaskEntity{Id: taskId}).First(task).Error
	if err != nil {
		return nil, err
	}
	return
}

func (taskServ TaskService) GetTaskListNoReward() (list []*models.TaskEntity, err error) {
	err = taskServ.Where("reward ==0").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

type JobService struct {
	*gorm.DB
}

var taskSet = wire.NewSet(db.NewDbService, wire.Struct(new(TaskService), "*"))
var jobSet = wire.NewSet(db.NewDbService, wire.Struct(new(JobService), "*"))
