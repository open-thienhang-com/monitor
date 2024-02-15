package framework

import (
	"mono.thienhang.com/pkg/common/logger"
	config "mono.thienhang.com/pkg/config"
	db "mono.thienhang.com/pkg/database"
	plugins "mono.thienhang.com/pkg/plugins"
	service "mono.thienhang.com/pkg/service"
)

type Engine struct {
	PluginList plugins.Plugins
	Adapter    WebFrameWork
	Services   service.List
	config     *config.Config
}

// Default return the default engine instance.
func Default() *Engine {
	return &Engine{
		Adapter:  defaultAdapter,
		Services: service.GetServices(),
	}
}

// Use enable the
func (eng *Engine) Use(router interface{}) error {
	if eng.Adapter == nil {
		panic("adapter is nil, import the default adapter or use AddAdapter method add the adapter")
	}

	eng.initPlugins()

	return eng.Use(router)
}

// ============================
// Config APIs
// ============================

// AddConfig set the global config.
func (eng *Engine) AddConfig(cfg config.Config) *Engine {
	return eng.setConfig(cfg).InitDatabase()
}

// setConfig set the config of engine.
func (eng *Engine) setConfig(cfg config.Config) *Engine {
	eng.config = config.Set(cfg)
	// sysCheck, themeCheck := template.CheckRequirements()
	// if !sysCheck {
	// 	panic(fmt.Sprintf("wrong GoAdmin version, theme %s required GoAdmin version are %s",
	// 		eng.config.Theme, strings.Join(template.Default().GetRequirements(), ",")))
	// }
	// if !themeCheck {
	// 	panic(fmt.Sprintf("wrong Theme version, GoAdmin %s required Theme version are %s",
	// 		system.Version(), strings.Join(system.RequireThemeVersion()[eng.config.Theme], ",")))
	// }
	return eng
}

// AddConfigFromJSON set the global config from json file.
func (eng *Engine) AddConfigFromJSON(path string) *Engine {
	return eng.setConfig(config.ReadFromJson(path)).InitDatabase()
}

// AddConfigFromYAML set the global config from yaml file.
func (eng *Engine) AddConfigFromYAML(path string) *Engine {
	return eng.setConfig(config.ReadFromYaml(path)).InitDatabase()
}

// AddConfigFromINI set the global config from ini file.
func (eng *Engine) AddConfigFromINI(path string) *Engine {
	return eng.setConfig(config.ReadFromINI(path)).InitDatabase()
}

// InitDatabase initialize all database connection.
func (eng *Engine) InitDatabase() *Engine {
	for driver, databaseCfg := range eng.config.Databases.GroupByDriver() {
		eng.Services.Add(driver, db.GetConnectionByDriver(driver).InitDB(databaseCfg))
	}
	if defaultAdapter == nil {
		panic("adapter is nil")
	}
	return eng
}

// AddAdapter add the adapter of engine.
func (eng *Engine) AddAdapter(ada WebFrameWork) *Engine {
	eng.Adapter = ada
	defaultAdapter = ada
	return eng
}

func (eng *Engine) GetTheme() string {
	if eng.config.Theme == "" {
		return ""
	}
	return eng.config.Theme
}

// defaultAdapter is the default adapter of engine.
var defaultAdapter WebFrameWork

// Register set default adapter of engine.
func Register(ada WebFrameWork) {
	if ada == nil {
		panic("adapter is nil")
	}
	defaultAdapter = ada
}

// ============================
// DB Connection APIs
// ============================

// DB return the db connection of given driver.
func (eng *Engine) DB(driver string) db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(driver))
}

// DefaultConnection return the default db connection.
func (eng *Engine) DefaultConnection() db.Connection {
	return eng.DB(eng.config.Databases.GetDefault().Driver)
}

// MysqlConnection return the mysql db connection of given driver.
func (eng *Engine) MysqlConnection() db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(db.DriverMysql))
}

// MssqlConnection return the mssql db connection of given driver.
func (eng *Engine) MssqlConnection() db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(db.DriverMssql))
}

// PostgresqlConnection return the postgresql db connection of given driver.
func (eng *Engine) PostgresqlConnection() db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(db.DriverPostgresql))
}

// SqliteConnection return the sqlite db connection of given driver.
func (eng *Engine) SqliteConnection() db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(db.DriverSqlite))
}

type ConnectionSetter func(db.Connection)

// ResolveConnection resolve the specified driver connection.
func (eng *Engine) ResolveConnection(setter ConnectionSetter, driver string) *Engine {
	setter(eng.DB(driver))
	return eng
}

// ResolveMysqlConnection resolve the mysql connection.
func (eng *Engine) ResolveMysqlConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverMysql)
	return eng
}

// ResolveMssqlConnection resolve the mssql connection.
func (eng *Engine) ResolveMssqlConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverMssql)
	return eng
}

// ResolveSqliteConnection resolve the sqlite connection.
func (eng *Engine) ResolveSqliteConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverSqlite)
	return eng
}

// ResolvePostgresqlConnection resolve the postgres connection.
func (eng *Engine) ResolvePostgresqlConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverPostgresql)
	return eng
}

type Setter func(*Engine)

// Clone copy a new Engine.
func (eng *Engine) Clone(e *Engine) *Engine {
	e = eng
	return eng
}

// ClonedBySetter copy a new Engine by a setter callback function.
func (eng *Engine) ClonedBySetter(setter Setter) *Engine {
	setter(eng)
	return eng
}

func (eng *Engine) initPlugins() {
	logger.Info(len(eng.PluginList))
	// for i := range eng.PluginList {
	// 	if eng.PluginList[i].Name() != "admin" {
	// 		eng.PluginList[i].InitPlugin(eng.Services)
	// 	}
	// }
	logger.Error("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
}

// AddPlugins add the plugins
func (eng *Engine) AddPlugins(plugs ...plugins.Plugin) *Engine {

	if len(plugs) == 0 {
		return eng
	}

	for _, plug := range plugs {
		eng.PluginList = eng.PluginList.Add(plug)
	}

	return eng
}

// AddPluginList add the plugins
func (eng *Engine) AddPluginList(plugs plugins.Plugins) *Engine {

	if len(plugs) == 0 {
		return eng
	}

	for _, plug := range plugs {
		eng.PluginList = eng.PluginList.Add(plug)
	}

	return eng
}

// FindPluginByName find the register plugin by given name.
func (eng *Engine) FindPluginByName(name string) (plugins.Plugin, bool) {
	for _, plug := range eng.PluginList {
		if plug.Name() == name {
			return plug, true
		}
	}
	return nil, false
}
