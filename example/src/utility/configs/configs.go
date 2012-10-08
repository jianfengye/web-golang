package configs

import (
    "errors"
    "io/ioutil"
    "path"
    "strconv"
    "strings"
)

type Config struct {
    data map[string]string
}

func NewConfig() *Config {
    return &Config{data: make(map[string]string)}
}

const emptyChars = " \r\t\v"

func (this *Config) Load(configFile string) error {
    stream, err := ioutil.ReadFile(configFile)
    if err != nil {
        return errors.New("cannot load config file")
    }
    content := string(stream)
    lines := strings.Split(content, "\n")
    for _, line := range lines {
        line = strings.Trim(line, emptyChars)
        if line == "" || line[0] == '#' {
            continue
        }
        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            for i, part := range parts {
                parts[i] = strings.Trim(part, emptyChars)
            }
            this.data[parts[0]] = parts[1]
        } else {
            // 判断并处理include条目，load相应的config文件
            includes := strings.SplitN(parts[0], " ", 2)
            if len(includes) == 2 && strings.EqualFold(includes[0], "include") {
                // 拼解新包含config文件的path
                confDir := path.Dir(configFile)
                newConfName := strings.Trim(includes[1], emptyChars)
                newConfPath := path.Join(confDir, newConfName)
                // 载入include的config文件，调用Load自身
                err := this.Load(newConfPath)
                if err != nil {
                    return errors.New("load include config file failed")
                }
                continue
            } else {
                return errors.New("invalid config file syntax")
            }
        }
    }
    return nil
}

func (this *Config) Get(key string) string {
    if value, ok := this.data[key]; ok {
        return value
    }
    return ""
}

func (this *Config) GetInt(key string) int {
    value := this.Get(key)
    if value == "" {
        return 0
    }
    result, err := strconv.Atoi(value)
    if err != nil {
        return 0
    }
    return result
}

func (this *Config) GetInt64(key string) int64 {
    value := this.Get(key)
    if value == "" {
        return 0
    }
    result, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
        return 0
    }
    return result
}

func (this *Config) GetSlice(key string, separator string) []string {
    slice := []string{}
    value := this.Get(key)
    if value != "" {
        for _, part := range strings.Split(value, separator) {
            slice = append(slice, strings.Trim(part, emptyChars))
        }
    }
    return slice
}

func (this *Config) GetSliceInt(key string, separator string) []int {
    slice := this.GetSlice(key, separator)
    results := []int{}
    for _, part := range slice {
        result, err := strconv.Atoi(part)
        if err != nil {
            continue
        }
        results = append(results, result)
    }
    return results
}
