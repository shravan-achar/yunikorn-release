package framework

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Template struct {
	Kubeconfig []TemplateFields `yaml:"kubeconfig"`
	Node       []TemplateFields `yaml:"node"`
	Job        []TemplateFields `yaml:"job"`
	Scheduler  []TemplateFields `yaml:"scheduler"`
}

type TemplateFields struct {
	Path           *string `yaml:"path"`
	MaxCount       *string `yaml:"maxCount"`
	DesiredCount   *string `yaml:"desiredCount"`
	PodCount       *string `yaml:"podCount"`
	Mode           *string `yaml:"mode"`
	Value          *string `yaml:"value"`
	VcoreRequests  *string `yaml:"vcoreRequests"`
	VcoreLimits    *string `yaml:"vcoreLimits"`
	MemoryRequests *string `yaml:"memoryRequests"`
	MemoryLimits   *string `yaml:"memoryLimits"`
}

type TestCaseParams struct {
	NodeMaxCount      *int `yaml:"nodeMaxCount"`
	NodesDesiredCount *int `yaml:"nodesDesiredCount"`
	NumPods           *int `yaml:"numPods"`
	NumJobs           *int `yaml:"numJobs"`
}

type Prom struct {
	Query      *string `yaml:"query"`
	Expression *string `yaml:"expression"`
	Value      *string `yaml:"value"`
	Op         *string `yaml:"op"`
}

type Metrics struct {
	SchedulerRestarts  *int    `yaml:"schedulerRestarts"`
	MaxAllocationDelay *string `yaml:"maxAllocationDelay"`
	Prom               []Prom  `yaml:"prom"`
}

type Threshold struct {
	MaxRuntime     *string  `yaml:"maxRuntime"`
	PendingPods    *int     `yaml:"pendingPods"`
	DetectDeadlock *bool    `yaml:"detectDeadlock"`
	Metrics        *Metrics `yaml:"metrics"`
}

type TestCase struct {
	Name      *string         `yaml:"name"`
	Params    *TestCaseParams `yaml:"params"`
	Schedule  *string         `yaml:"schedule"`
	Runs      *int            `yaml:"runs"`
	Labels    []string        `yaml:"labels"`
	Threshold *Threshold      `yaml:"threshold"`
}

type Test struct {
	Name      string     `yaml:"name"`
	Template  *Template  `yaml:"template"`
	TestCases []TestCase `yaml:"testCases"`
}

type Config struct {
	Tests []Test `yaml:"tests"`
}

func InitConfig(configFile string) (*Config, error) {
	yamlContent, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file: %s ", err.Error())
	}
	conf := Config{}
	err = yaml.Unmarshal(yamlContent, &conf)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse config file: %s ", err.Error())
	}
	return &conf, nil
}
