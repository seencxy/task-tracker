package main

import (
	"encoding/json"
	"sort"
	"time"
)

const (
	Todo       = "todo"
	Done       = "done"        // 1
	InProgress = "in-progress" // 2
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func NewTask(nextId int, description string) *Task {
	return &Task{
		ID:          nextId,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (t *Task) update(opts ...Option) {
	t.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	options := &Options{
		Description: t.Description,
		Status:      t.Status,
	}
	for _, opt := range opts {
		opt(options)
	}

	t.Description = options.Description
	t.Status = options.Status
}

type Options struct {
	Description string
	Status      string
}

type Option func(opts *Options)

func WithDescOpts(description string) Option {
	return func(opts *Options) {
		opts.Description = description
	}
}

func WithStatusOpts(status string) Option {
	return func(opts *Options) {
		opts.Status = status
	}
}

func unmarshalTasks(fileBody []byte) ([]*Task, error) {
	var tasks []*Task
	err := json.Unmarshal(fileBody, &tasks)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})
	return tasks, err
}
