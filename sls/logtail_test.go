package sls
import (
	"testing"
	"fmt"
)

func TestLogtailConfigs(t *testing.T) {
	p := DefaultProject()
	list, err := p.LogtailConfigs(0, 100)
	if err != nil {
		t.Fatalf("error list logtail configs: %v", err)
	}
	fmt.Println(list)
}

func TestDelete(t *testing.T) {
	p := DefaultProject()
	c, e := p.LogtailConfig("logtail-test")
	if e != nil {
		t.Fatalf("error %v", e)
	}
	c.Delete()
}

func TestCreateLogtailConfig(t *testing.T) {
	p := DefaultProject()
	logtailConfig := &LogtailConfig{
		Name: "logtail-test",
		InputType: "file",
		InputDetail: LogtailInput{
			LogType: "common_reg_log",
			LogPath: "/abc",
			FilePattern: "*.log",
			LocalStorage: false,
			TimeFormat: "",
			LogBeginRegex: ".*",
			Regex: "(.*)",
			Key: []string{"content"},
			FilterKey:[]string{"content"},
			FilterRegex: []string{".*"},
			TopicFormat: "none",
		},
		OutputType:"LogService",
		Sample: "sample",
		OutputDetail:LogtailOutput{
			LogstoreName: "test-jjz",
		},
	}

	if err := p.CreateLogtailConfig(logtailConfig); err != nil {
		t.Fatalf("error create logtail config: %v", err)
	}
}