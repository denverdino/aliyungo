package sls

import (
	"testing"
)

func TestLogtailConfigs(t *testing.T) {
	p := DefaultProject(t)
	_, err := p.ListConfig(0, 100)
	if err != nil {
		t.Fatalf("error list logtail configs: %v", err)
	}
}

func TestDelete(t *testing.T) {
	p := DefaultProject(t)
	_, e := p.GetConfig("logtail-test")
	if e != nil {
		t.Fatalf("error %v", e)
	}
	p.DeleteConfig("logtail-test")
}

func TestCreateLogtailConfig(t *testing.T) {
	p := DefaultProject(t)
	logtailConfig := &LogtailConfig{
		Name:      "logtail-test",
		InputType: "file",
		InputDetail: LogtailInput{
			LogType:       "common_reg_log",
			LogPath:       "/abc",
			FilePattern:   "*.log",
			LocalStorage:  false,
			TimeFormat:    "",
			LogBeginRegex: ".*",
			Regex:         "(.*)",
			Key:           []string{"content"},
			FilterKey:     []string{"content"},
			FilterRegex:   []string{".*"},
			TopicFormat:   "none",
		},
		OutputType: "LogService",
		Sample:     "sample",
		OutputDetail: LogtailOutput{
			LogstoreName: TestLogstoreName,
		},
	}

	if err := p.CreateConfig(logtailConfig); err != nil {
		t.Fatalf("error create logtail config: %v", err)
	}
}
