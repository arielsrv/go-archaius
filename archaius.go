// Package archaius provides you APIs which helps to manage files,
// remote config center configurations
package archaius

import (
	"errors"
	"fmt"
	"io"

	"github.com/arielsrv/go-archaius/event"
	"github.com/arielsrv/go-archaius/pkg/cast"
	"github.com/arielsrv/go-archaius/source"
	"github.com/arielsrv/go-archaius/source/cli"
	"github.com/arielsrv/go-archaius/source/env"
	"github.com/arielsrv/go-archaius/source/mem"
	"github.com/sirupsen/logrus"

	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	filesource "github.com/arielsrv/go-archaius/source/file"
)

var (
	manager             *source.Manager
	fs                  filesource.FileSource
	running             = false
	configServerRunning = false
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&nested.Formatter{
		FieldsOrder:      []string{"component", "category"},
		TimestampFormat:  "2006-01-02 15:04:05",
		HideKeys:         true,
		TrimMessages:     true,
		NoUppercaseLevel: true,
	})
}
func initFileSource(o *Options) (filesource.FileSource, error) {
	files := make([]string, 0)
	// created file source object
	fs = filesource.NewFileSource()
	// adding all files with file source
	for _, v := range o.RequiredFiles {
		if err := fs.AddFile(v, filesource.DefaultFilePriority, o.FileHandler); err != nil {
			logrus.Error(fmt.Sprintf("add file source error [%s].", err.Error()))
			return nil, err
		}
		files = append(files, v)
	}
	for _, v := range o.OptionalFiles {
		_, err := os.Stat(v)
		if os.IsNotExist(err) {
			logrus.Info(fmt.Sprintf("[%s] not exist", v))
			continue
		}
		if err := fs.AddFile(v, filesource.DefaultFilePriority, o.FileHandler); err != nil {
			logrus.Info(err.Error())
			return nil, err
		}
		files = append(files, v)
	}
	// logrus.Info(fmt.Sprintf("Configuration files: %s", strings.Join(files, ", ")))
	for _, file := range files {
		logrus.Info(fmt.Sprintf("loaded configuration file: %s", file))
	}
	return fs, nil
}

// Init create a Archaius config singleton.
func Init(opts ...Option) error {
	if running {
		logrus.Warn("can not init archaius again, call Clean first")
		return nil
	}
	var err error
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}

	manager = source.NewManager()

	fs, err := initFileSource(o)
	if err != nil {
		return err
	}
	err = manager.AddSource(fs)
	if err != nil {
		return err
	}

	if o.RemoteSource != "" {
		if err = EnableRemoteSource(o.RemoteSource, o.RemoteInfo); err != nil {
			return err
		}
	}

	// build-in config sources
	if o.UseMemSource {
		ms := mem.NewMemoryConfigurationSource()
		if err = manager.AddSource(ms); err != nil {
			return err
		}
	}
	if o.UseCLISource {
		cmdSource := cli.NewCommandlineConfigSource()
		if err = manager.AddSource(cmdSource); err != nil {
			return err
		}
	}
	if o.UseENVSource {
		envSource := env.NewEnvConfigurationSource()
		if err = manager.AddSource(envSource); err != nil {
			return err
		}
	}

	logrus.Info("archaius init success")
	running = true
	return nil
}

// CustomInit accept a list of config source, add it into archaius runtime.
// it almost like Init(), but you can fully control config sources you inject to archaius.

// EnableRemoteSource create a remote source singleton
// A config center source pull remote config server key values into local memory
// so that you can use GetXXX to get value easily.
func EnableRemoteSource(remoteSource string, ci *RemoteInfo) error {
	if ci == nil {
		return errors.New("RemoteInfo can not be empty")
	}
	if configServerRunning {
		logrus.Warn("can not init config server again, call Clean first")
		return nil
	}

	f, ok := newFuncMap[remoteSource]
	if !ok {
		return errors.New("don not support remote source: " + remoteSource)
	}
	s, err := f(ci)
	if err != nil {
		return err
	}
	err = manager.AddSource(s)
	if err != nil {
		return err
	}
	configServerRunning = true
	return nil
}

// Get is for to get the value of configuration key.
func Get(key string) interface{} {
	return manager.GetConfig(key)
}

// GetValue return interface.
func GetValue(key string) cast.Value {
	var confValue cast.Value
	val := manager.GetConfig(key)
	if val == nil {
		confValue = cast.NewValue(nil, source.ErrKeyNotExist)
	} else {
		confValue = cast.NewValue(val, nil)
	}
	return confValue
}

// Exist check the configuration key existence.
func Exist(key string) bool {
	return manager.IsKeyExist(key)
}

// UnmarshalConfig unmarshal the config of receiving object.
func UnmarshalConfig(obj interface{}) error {
	return manager.Unmarshal(obj)
}

// WriteTo write the config to writer by yaml.
func WriteTo(w io.Writer) error {
	return manager.Marshal(w)
}

// GetBool is gives the key value in the form of bool.
func GetBool(key string, defaultValue bool) bool {
	b, err := GetValue(key).ToBool()
	if err != nil {
		return defaultValue
	}
	return b
}

// GetFloat64 gives the key value in the form of float64.

// GetInt gives the key value in the form of GetInt.
func GetInt(key string, defaultValue int) int {
	result, err := GetValue(key).ToInt()
	if err != nil {
		return defaultValue
	}
	return result
}

// GetInt64 gives the key value in the form of int64.
func GetInt64(key string, defaultValue int64) int64 {
	result, err := GetValue(key).ToInt64()
	if err != nil {
		return defaultValue
	}
	return result
}

// GetString gives the key value in the form of GetString.
func GetString(key string, defaultValue string) string {
	result, err := GetValue(key).ToString()
	if err != nil {
		return defaultValue
	}
	return result
}

// GetConfigs gives the information about all configurations.
func GetConfigs() map[string]interface{} {
	return manager.Configs()
}

// GetConfigsWithSourceNames gives the information about all configurations
// each config key, along with its source will be returned
// the returned map will be like:
//
//	map[string]interface{}{
//			key string: map[string]interface{"value": value, "sourceName": sourceName}
//	}
func GetConfigsWithSourceNames() map[string]interface{} {
	return manager.ConfigsWithSourceNames()
}

// AddDimensionInfo adds a NewDimensionInfo of which configurations needs to be taken.

// RegisterListener to Register all listener for different key changes, each key could be a regular expression.
func RegisterListener(listenerObj event.Listener, key ...string) error {
	return manager.RegisterListener(listenerObj, key...)
}

// UnRegisterListener is to remove the listener.
func UnRegisterListener(listenerObj event.Listener, key ...string) error {
	return manager.UnRegisterListener(listenerObj, key...)
}

// RegisterModuleListener to Register all moduleListener for different key(prefix) changes.
func RegisterModuleListener(listenerObj event.ModuleListener, prefix ...string) error {
	return manager.RegisterModuleListener(listenerObj, prefix...)
}

// UnRegisterModuleListener is to remove the moduleListener.

// AddFile is for to add the configuration files at runtime.
func AddFile(file string, opts ...FileOption) error {
	o := &FileOptions{}
	for _, f := range opts {
		f(o)
	}
	if err := fs.AddFile(file, filesource.DefaultFilePriority, o.Handler); err != nil {
		return err
	}
	return manager.Refresh(fs.GetSourceName())
}

// Set add the configuration key, value pairs into memory source at runtime
// it is just affect the local configs.
func Set(key string, value interface{}) error {
	return manager.Set(key, value)
}

// Delete delete the configuration key, value pairs in memory source.
func Delete(key string) error {
	return manager.Delete(key)
}

// AddSource add source implementation.

// Clean will call config manager CleanUp Method,
// it deletes all sources which means all of key value is deleted.
// after you call Clean, you can init archaius again.
func Clean() error {
	manager.Cleanup()
	running = false
	return nil
}
