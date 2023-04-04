package toolkit

import (
	"net/http"
	"path/filepath"

	"github.com/dhamanutd/golang-toolkit/logger"
	"github.com/go-openapi/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

// env options
type ServerEnv string

const (
	Dev     ServerEnv = "dev"
	Prod    ServerEnv = "prod"
	Test    ServerEnv = "test"
	Staging ServerEnv = "staging"
)

func NewRuntime() *Runtime {
	r := &Runtime{
		env:  Dev,
		log:  logger.NewLogger(),
		cfg:  viper.New(),
		data: map[string]any{},
	}
	return r
}

// Runtime in application context that handle common functionality that used in microservices
// contain logger, db setup, config and inmemory save for reusable object
type Runtime struct {
	env  ServerEnv
	log  *logger.Logger
	db   *gorm.DB
	cfg  *viper.Viper
	data map[string]any
}

func (r *Runtime) ReadConfig(cfgPath *string) {
	// init environment variable from os.env
	// if the config file is exist, load and replace data wiht config from file

	r.cfg.AutomaticEnv()
	if cfgPath != nil && *cfgPath != "" {
		r.cfg.SetConfigName(filepath.Base(*cfgPath))
		r.cfg.AddConfigPath(filepath.Dir(*cfgPath))

		r.cfg.AddConfigPath("./configs/")
		r.cfg.AddConfigPath("./etc/")
		r.cfg.AddConfigPath("./")

		if err := r.cfg.ReadInConfig(); err != nil {
			r.log.Warning("env file not found. only load environment variabel")
		}
	}
}

func (r *Runtime) SetEnv(env ServerEnv) {
	r.env = env
}

func (r *Runtime) Env() ServerEnv {
	return r.env
}

func (r *Runtime) InitDB(dialector gorm.Dialector, opts ...gorm.Option) error {
	gConfigs := &gorm.Config{}
	if r.env != Prod {
		gConfigs.Logger = gorm_logger.Default.LogMode(gorm_logger.Info)
	}

	opts = append(opts, gConfigs)
	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return err
	}
	r.db = db
	return nil
}

func (r *Runtime) Config() *viper.Viper {
	return r.cfg
}

func (r *Runtime) Get(key string) any {
	if v, ok := r.data[key]; ok {
		return v
	}
	return nil
}

func (r *Runtime) DB() *gorm.DB {
	if r.db == nil {
		r.log.Panic("DB must be iniliazied")
		return nil
	}

	return r.db
}

func (r *Runtime) RunMigration(models ...any) {
	if r.db == nil {
		r.log.Panic("DB must be iniliazied")
	}

	r.log.Info("Migrating DBs")
	r.db.AutoMigrate(models...)
	r.log.Info("Success migrate all DBs")
}

func (r *Runtime) Log() *logger.Logger {
	return r.log
}

func (r *Runtime) SetError(code int, msg string, args ...interface{}) error {
	return errors.New(int32(code), msg)
}

func (r *Runtime) GetError(err error) errors.Error {
	if v, ok := err.(errors.Error); ok {
		return v
	}
	return errors.New(http.StatusInternalServerError, err.Error())
}

// Set is function for set data to map that saven in runtime
func (r *Runtime) Set(key string, value any) {
	r.data[key] = value
}

// check environment is set in os or loaded by config
func (r *Runtime) CheckEnv(environments ...string) {
	for _, env := range environments {
		if r.Config().GetString(env) == "" {
			r.Log().Fatalf("env %s is required to set", env)
			return
		}
	}
}

// Get is function for get data from  dynamic map and return value as generic
func Get[T any](rt *Runtime, key string) (r T) {
	switch t := rt.Get(key).(type) {
	case T:
		return t
	}
	return
}
