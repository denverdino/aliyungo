package sls
import (
	"encoding/json"
	"strconv"
)

type LogtailInput struct {
	LogType       string `json:"logType,omitempty"`
	LogPath       string `json:"logPath,omitempty"`
	FilePattern   string `json:"filePattern,omitempty"`
	LocalStorage  bool `json:"localStorage,omitempty"`
	TimeFormat    string `json:"timeFormat,omitempty"`
	logBeginRegex string `json:"logBeginRegex,omitempty"`
	regex         string `json:"regex,omitempty"`
	key           []string `json:"key,omitempty"`
	filterKey     []string `json:"filterKey,omitempty"`
	filterRegex   []string `json:"filterRegex,omitempty"`
	topicFormat   string `json:"topicFormat,omitempty"`
}

type LogtailOutput struct {
	LogstoreName string `json:"logstoreName,omitempty"`
}

type LogtailConfig struct {
	client       *Client
	Name         string `json:"configName,omitempty"`
	InputType    string `json:"inputType,omitempty"`
	InputDetail  LogtailInput `json:inputDetail,omitempty`
	OutputType   LogtailInput `json:outputType,omitempty`
	OutputDetail LogtailOutput `json:outputDetail,omitempty`
}

func (proj *Project) CreateLogtailConfig(config *LogtailConfig) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	req := &request{
		method: METHOD_POST,
		path: "/configs",
		payload: data,
		contentType: "application/json",
	}

	return proj.client.requestWithClose(req)
}

type LogtailConfigList struct {
	count   int `json:"count,omitempty"`
	total   int `json:"total,omitempty"`
	configs int `json:"configs,omitempty"`
}

func (proj *Project) LogtailConfigs(offset, size int) (*LogtailConfigList, error) {
	req := &request{
		method: METHOD_GET,
		path: "/configs",
		params: map[string]string{
			"size": strconv.Itoa(size),
			"offset": strconv.Itoa(offset),
		},
	}

	list := &LogtailConfigList{}
	if err := proj.client.requestWithJsonResponse(req, list); err != nil {
		return nil, err
	}
	return list, nil
}

func (proj *Project) LogtailConfig(name string) (*LogtailConfig, error) {
	req := &request{
		method: METHOD_GET,
		path: "/configs/" + name,
	}

	config := &LogtailConfig{}
	if err := proj.client.requestWithJsonResponse(req, config); err != nil {
		return nil, err
	}

	config.client = proj.client
	return config, nil

}

func (lc *LogtailConfig) AppliedMachineGroups() ([]string, error) {
	type appliedMachineGroups struct {
		machinegroups []string `json:"machinegroups,omitempty"`
	}

	req := &request{
		method: METHOD_GET,
		path:"/configs/" + lc.Name + "/machinegroups",
	}

	group := &appliedMachineGroups{}

	if err := lc.client.requestWithJsonResponse(req, group); err != nil {
		return nil, err
	}

	return group.machinegroups, nil
}

func (lc *LogtailConfig) Delete() error {
	req := request{
		method: METHOD_DELETE,
		path:"/configs/" + lc.Name,
	}

	return lc.client.requestWithClose(req)
}


func (lc *LogtailConfig) Update() error {
	data, err := json.Marshal(lc)
	if err != nil {
		return err
	}

	req := &request{
		method: METHOD_POST,
		path: "/configs/" + lc.Name,
		payload: data,
		contentType: "application/json",
	}

	return lc.client.requestWithClose(req)
}
