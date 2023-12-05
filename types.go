package main

type Task struct {
	ID         int    `json:"id"`
	TaskName   string `json:"taskName"`
	TaskDetail string `json:"taskDetail"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type CreateTaskReq struct {
	TaskName   string `json:"taskName" validate:"required"`
	TaskDetail string `json:"taskDetail" validate:"required"`
}

func NewTask(taskName, taskDetail string) (*Task, error) {
	brazilTime, err := GetBrazilCurrentTimeHelper()
	if err != nil {
		return nil, err
	}

	return &Task{
		TaskName:   taskName,
		TaskDetail: taskDetail,
		CreatedAt:  brazilTime,
		UpdatedAt:  brazilTime,
	}, nil
}
