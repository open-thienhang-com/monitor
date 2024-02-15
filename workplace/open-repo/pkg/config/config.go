package config

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"

	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
	"mono.thienhang.com/pkg/common/logger"
	"mono.thienhang.com/pkg/utils"
)

// Database is a type of database connection config.
//
// Because a little difference of different database driver.
// The Config has multiple options but may not be used.
// Such as the sqlite driver only use the File option which
// can be ignored when the driver is mysql.
//
// If the Dsn is configured, when driver is mysql/postgresql/
// mssql, the other configurations will be ignored, except for
// MaxIdleCon and MaxOpenCon.
type Database struct {
	BaseConfig
	Host       string `json:"host,omitempty",yaml:"host",ini:"host"`
	Port       string `json:"port,omitempty",yaml:"port",ini:"port"`
	User       string `json:"user,omitempty",yaml:"user",ini:"user"`
	Pwd        string `json:"pwd,omitempty",yaml:"pwd",ini:"pwd"`
	Name       string `json:"name,omitempty",yaml:"name",ini:"name"`
	MaxIdleCon int    `json:"max_idle_con,omitempty",yaml:"max_idle_con",ini:"max_idle_con"`
	MaxOpenCon int    `json:"max_open_con,omitempty",yaml:"max_open_con",ini:"max_open_con"`
	Driver     string `json:"driver,omitempty",yaml:"driver",ini:"driver"`
	File       string `json:"file,omitempty",yaml:"file",ini:"file"`
	Dsn        string `json:"dsn,omitempty",yaml:"dsn",ini:"dsn"`
}

// DatabaseList is a map of Database.
type DatabaseList map[string]Database

// GetDefault get the default Database.
func (d DatabaseList) GetDefault() Database {
	return d["default"]
}

// Add add a Database to the DatabaseList.
func (d DatabaseList) Add(key string, db Database) {
	d[key] = db
}

// GroupByDriver group the Databases with the drivers.
func (d DatabaseList) GroupByDriver() map[string]DatabaseList {
	drivers := make(map[string]DatabaseList)
	for key, item := range d {
		if driverList, ok := drivers[item.Driver]; ok {
			driverList.Add(key, item)
		} else {
			drivers[item.Driver] = make(DatabaseList)
			drivers[item.Driver].Add(key, item)
		}
	}
	return drivers
}

func (d DatabaseList) JSON() string {
	return utils.JSON(d)
}

func (d DatabaseList) Copy() DatabaseList {
	var c = make(DatabaseList)
	for k, v := range d {
		c[k] = v
	}
	return c
}

func (d DatabaseList) Connections() []string {
	conns := make([]string, len(d))
	count := 0
	for key := range d {
		conns[count] = key
		count++
	}
	return conns
}

func GetDatabaseListFromJSON(m string) DatabaseList {
	var d = make(DatabaseList, 0)
	if m == "" {
		panic("wrong config")
	}
	_ = json.Unmarshal([]byte(m), &d)
	return d
}

const (
	// EnvTest is a const value of test environment.
	EnvTest = "test"
	// EnvLocal is a const value of local environment.
	EnvLocal = "local"
	// EnvProd is a const value of production environment.
	EnvProd = "prod"

	// DriverMysql is a const value of mysql driver.
	DriverMysql = "mysql"
	// DriverSqlite is a const value of sqlite driver.
	DriverSqlite = "sqlite"
	// DriverPostgresql is a const value of postgresql driver.
	DriverPostgresql = "postgresql"
	// DriverMssql is a const value of mssql driver.
	DriverMssql = "mssql"
)

// Store is the file store config. Path is the local store path.
// and prefix is the url prefix used to visit it.
type Store struct {
	Path   string `json:"path,omitempty",yaml:"path",ini:"path"`
	Prefix string `json:"prefix,omitempty",yaml:"prefix",ini:"prefix"`
}

func (s Store) URL(suffix string) string {
	if len(suffix) > 4 && suffix[:4] == "http" {
		return suffix
	}
	if s.Prefix == "" {
		if suffix[0] == '/' {
			return suffix
		}
		return "/" + suffix
	}
	if s.Prefix[0] == '/' {
		if suffix[0] == '/' {
			return s.Prefix + suffix
		}
		return s.Prefix + "/" + suffix
	}
	if suffix[0] == '/' {
		if len(s.Prefix) > 4 && s.Prefix[:4] == "http" {
			return s.Prefix + suffix
		}
		return "/" + s.Prefix + suffix
	}
	if len(s.Prefix) > 4 && s.Prefix[:4] == "http" {
		return s.Prefix + "/" + suffix
	}
	return "/" + s.Prefix + "/" + suffix
}

func (s Store) JSON() string {
	if s.Path == "" && s.Prefix == "" {
		return ""
	}
	return utils.JSON(s)
}

func GetStoreFromJSON(m string) Store {
	var s Store
	if m == "" {
		return s
	}
	_ = json.Unmarshal([]byte(m), &s)
	return s
}

// Config type is the global config of goAdmin. It will be
// initialized in the engine.
type Config struct {
	BaseConfig
	// An map supports multi database connection. The first
	// element of Databases is the default connection. See the
	// file connection.go.
	Databases DatabaseList `json:"database,omitempty",yaml:"database",ini:"database"`

	// The cookie domain used in the auth modules. see
	// the session.go.
	Domain string `json:"domain,omitempty",yaml:"domain",ini:"domain"`

	// Used to set as the localize language which show in the
	// interface.
	Language string `json:"language,omitempty",yaml:"language",ini:"language"`

	// The global url prefix.
	UrlPrefix string `json:"prefix,omitempty",yaml:"prefix",ini:"prefix"`

	// The theme name of template.
	Theme string `json:"theme,omitempty",yaml:"theme",ini:"theme"`

	// The path where files will be stored into.
	Store Store `json:"store,omitempty",yaml:"store",ini:"store"`

	// The title of web page.
	Title string `json:"title,omitempty",yaml:"title",ini:"title"`

	// Logo is the top text in the sidebar.
	Logo template.HTML `json:"logo,omitempty",yaml:"logo",ini:"logo"`

	// Mini-logo is the top text in the sidebar when folding.
	MiniLogo template.HTML `json:"mini_logo,omitempty",yaml:"mini_logo",ini:"mini_logo"`

	// The url redirect to after login.
	IndexUrl string `json:"index,omitempty",yaml:"index",ini:"index"`

	// Login page URL
	LoginUrl string `json:"login_url,omitempty",yaml:"login_url",ini:"login_url"`

	// Debug mode
	Debug bool `json:"debug,omitempty",yaml:"debug",ini:"debug"`

	// Env is the environment,which maybe local,test,prod.
	Env string `json:"env,omitempty",yaml:"env",ini:"env"`

	// Info log path.
	InfoLogPath string `json:"info_log,omitempty",yaml:"info_log",ini:"info_log"`

	// Error log path.
	ErrorLogPath string `json:"error_log,omitempty",yaml:"error_log",ini:"error_log"`

	// Access log path.
	AccessLogPath string `json:"access_log,omitempty",yaml:"access_log",ini:"access_log"`

	// Access assets log off
	AccessAssetsLogOff bool `json:"access_assets_log_off,omitempty",yaml:"access_assets_log_off",ini:"access_assets_log_off"`

	// Sql operator record log switch.
	SqlLog bool `json:"sql_log,omitempty",yaml:"sql_log",ini:"sql_log"`

	AccessLogOff bool `json:"access_log_off,omitempty",yaml:"access_log_off",ini:"access_log_off"`
	InfoLogOff   bool `json:"info_log_off,omitempty",yaml:"info_log_off",ini:"info_log_off"`
	ErrorLogOff  bool `json:"error_log_off,omitempty",yaml:"error_log_off",ini:"error_log_off"`

	Logger Logger `json:"logger,omitempty",yaml:"logger",ini:"logger"`

	// Color scheme.
	ColorScheme string `json:"color_scheme,omitempty",yaml:"color_scheme",ini:"color_scheme"`

	// Session valid time duration,units are seconds. Default 7200.
	SessionLifeTime int `json:"session_life_time,omitempty",yaml:"session_life_time",ini:"session_life_time"`

	// Assets visit link.
	AssetUrl string `json:"asset_url,omitempty",yaml:"asset_url",ini:"asset_url"`

	// File upload engine,default "local"
	FileUploadEngine FileUploadEngine `json:"file_upload_engine,omitempty",yaml:"file_upload_engine",ini:"file_upload_engine"`

	// Custom html in the tag head.
	CustomHeadHtml template.HTML `json:"custom_head_html,omitempty",yaml:"custom_head_html",ini:"custom_head_html"`

	// Custom html after body.
	CustomFootHtml template.HTML `json:"custom_foot_html,omitempty",yaml:"custom_foot_html",ini:"custom_foot_html"`

	// Footer Info html
	FooterInfo template.HTML `json:"footer_info,omitempty",yaml:"footer_info",ini:"footer_info"`

	// Login page title
	LoginTitle string `json:"login_title,omitempty",yaml:"login_title",ini:"login_title"`

	// Login page logo
	LoginLogo template.HTML `json:"login_logo,omitempty",yaml:"login_logo",ini:"login_logo"`

	// Auth user table
	AuthUserTable string `json:"auth_user_table,omitempty",yaml:"auth_user_table",ini:"auth_user_table"`

	// Extra config info
	Extra ExtraInfo `json:"extra,omitempty",yaml:"extra",ini:"extra"`

	// Page animation
	Animation PageAnimation `json:"animation,omitempty",yaml:"animation",ini:"animation"`

	// Limit login with different IPs
	NoLimitLoginIP bool `json:"no_limit_login_ip,omitempty",yaml:"no_limit_login_ip",ini:"no_limit_login_ip"`

	// When site off is true, website will be closed
	SiteOff bool `json:"site_off,omitempty",yaml:"site_off",ini:"site_off"`

	// Hide config center entrance flag
	HideConfigCenterEntrance bool `json:"hide_config_center_entrance,omitempty",yaml:"hide_config_center_entrance",ini:"hide_config_center_entrance"`

	// Hide app info entrance flag
	HideAppInfoEntrance bool `json:"hide_app_info_entrance,omitempty",yaml:"hide_app_info_entrance",ini:"hide_app_info_entrance"`

	// Hide tool entrance flag
	HideToolEntrance bool `json:"hide_tool_entrance,omitempty",yaml:"hide_tool_entrance",ini:"hide_tool_entrance"`

	Custom404HTML template.HTML `json:"custom_404_html,omitempty",yaml:"custom_404_html",ini:"custom_404_html"`

	Custom403HTML template.HTML `json:"custom_403_html,omitempty",yaml:"custom_403_html",ini:"custom_403_html"`

	Custom500HTML template.HTML `json:"custom_500_html,omitempty",yaml:"custom_500_html",ini:"custom_500_html"`

	// Update Process Function
	// UpdateProcessFn UpdateConfigProcessFn `json:"-",,omitemptyyaml:"-",ini:"-"`

	Favicon string `json:"favicon,omitempty",yaml:"favicon",ini:"favicon"`

	// Is open admin plugin json api
	OpenAdminApi bool `json:"open_admin_api,omitempty",yaml:"open_admin_api",ini:"open_admin_api"`

	HideVisitorUserCenterEntrance bool `json:"hide_visitor_user_center_entrance,omitempty",yaml:"hide_visitor_user_center_entrance",ini:"hide_visitor_user_center_entrance"`

	ExcludeThemeComponents []string `json:"exclude_theme_components,omitempty",yaml:"exclude_theme_components",ini:"exclude_theme_components"`

	prefix string
}

type Logger struct {
	Encoder EncoderCfg `json:"encoder,omitempty",yaml:"encoder",ini:"encoder"`
	Rotate  RotateCfg  `json:"rotate,omitempty",yaml:"rotate",ini:"rotate"`
	Level   int8       `json:"level,omitempty",yaml:"level",ini:"level"`
}

type EncoderCfg struct {
	TimeKey       string `json:"time_key,omitempty",yaml:"time_key",ini:"time_key"`
	LevelKey      string `json:"level_key,omitempty",yaml:"level_key",ini:"level_key"`
	NameKey       string `json:"name_key,omitempty",yaml:"name_key",ini:"name_key"`
	CallerKey     string `json:"caller_key,omitempty",yaml:"caller_key",ini:"caller_key"`
	MessageKey    string `json:"message_key,omitempty",yaml:"message_key",ini:"message_key"`
	StacktraceKey string `json:"stacktrace_key,omitempty",yaml:"stacktrace_key",ini:"stacktrace_key"`
	Level         string `json:"level,omitempty",yaml:"level",ini:"level"`
	Time          string `json:"time,omitempty",yaml:"time",ini:"time"`
	Duration      string `json:"duration,omitempty",yaml:"duration",ini:"duration"`
	Caller        string `json:"caller,omitempty",yaml:"caller",ini:"caller"`
	Encoding      string `json:"encoding,omitempty",yaml:"encoding",ini:"encoding"`
}

type RotateCfg struct {
	MaxSize    int  `json:"max_size,omitempty",yaml:"max_size",ini:"max_size"`
	MaxBackups int  `json:"max_backups,omitempty",yaml:"max_backups",ini:"max_backups"`
	MaxAge     int  `json:"max_age,omitempty",yaml:"max_age",ini:"max_age"`
	Compress   bool `json:"compress,omitempty",yaml:"compress",ini:"compress"`
}

type ExtraInfo map[string]interface{}

// type UpdateConfigProcessFn func(values form.Values) (form.Values, error)

// see more: https://daneden.github.io/animate.css/
type PageAnimation struct {
	Type     string  `json:"type,omitempty",yaml:"type",ini:"type"`
	Duration float32 `json:"duration,omitempty",yaml:"duration",ini:"duration"`
	Delay    float32 `json:"delay,omitempty",yaml:"delay",ini:"delay"`
}

func (p PageAnimation) JSON() string {
	if p.Type == "" {
		return ""
	}
	return utils.JSON(p)
}

func GetPageAnimationFromJSON(m string) PageAnimation {
	var p PageAnimation
	if m == "" {
		return p
	}
	_ = json.Unmarshal([]byte(m), &p)
	return p
}

// FileUploadEngine is a file upload engine.
type FileUploadEngine struct {
	Name   string                 `json:"name,omitempty",yaml:"name",ini:"name"`
	Config map[string]interface{} `json:"config,omitempty",yaml:"config",ini:"config"`
}

func (f FileUploadEngine) JSON() string {
	if f.Name == "" {
		return ""
	}
	if len(f.Config) == 0 {
		f.Config = nil
	}
	return utils.JSON(f)
}

func GetFileUploadEngineFromJSON(m string) FileUploadEngine {
	var f FileUploadEngine
	if m == "" {
		return f
	}
	_ = json.Unmarshal([]byte(m), &f)
	return f
}

// GetIndexURL get the index url with prefix.
func (c *Config) GetIndexURL() string {
	index := c.Index()
	if index == "/" {
		return c.Prefix()
	}

	return c.Prefix() + index
}

// Url get url with the given suffix.
func (c *Config) Url(suffix string) string {
	if c.prefix == "/" {
		return suffix
	}
	if suffix == "/" {
		return c.prefix
	}
	return c.prefix + suffix
}

// IsTestEnvironment check the environment if it is test.
func (c *Config) IsTestEnvironment() bool {
	return c.Env == EnvTest
}

// IsLocalEnvironment check the environment if it is local.
func (c *Config) IsLocalEnvironment() bool {
	return c.Env == EnvLocal
}

// IsProductionEnvironment check the environment if it is production.
func (c *Config) IsProductionEnvironment() bool {
	return c.Env == EnvProd
}

// URLRemovePrefix remove prefix from the given url.
func (c *Config) URLRemovePrefix(url string) string {
	if url == c.prefix {
		return "/"
	}
	if c.prefix == "/" {
		return url
	}
	return strings.Replace(url, c.prefix, "", 1)
}

// Index return the index url without prefix.
func (c *Config) Index() string {
	if c.IndexUrl == "" {
		return "/"
	}
	if c.IndexUrl[0] != '/' {
		return "/" + c.IndexUrl
	}
	return c.IndexUrl
}

// Prefix return the prefix.
func (c *Config) Prefix() string {
	return c.prefix
}

// AssertPrefix return the prefix of assert.
func (c *Config) AssertPrefix() string {
	if c.prefix == "/" {
		return ""
	}
	return c.prefix
}

// func (c *Config) AddUpdateProcessFn(fn UpdateConfigProcessFn) *Config {
// 	c.UpdateProcessFn = fn
// 	return c
// }

// PrefixFixSlash return the prefix fix the slash error.
func (c *Config) PrefixFixSlash() string {
	if c.UrlPrefix == "/" {
		return ""
	}
	if c.UrlPrefix != "" && c.UrlPrefix[0] != '/' {
		return "/" + c.UrlPrefix
	}
	return c.UrlPrefix
}

func (c *Config) Copy() *Config {
	return &Config{
		Databases:                c.Databases.Copy(),
		Domain:                   c.Domain,
		Language:                 c.Language,
		UrlPrefix:                c.UrlPrefix,
		Theme:                    c.Theme,
		Store:                    c.Store,
		Title:                    c.Title,
		Logo:                     c.Logo,
		MiniLogo:                 c.MiniLogo,
		IndexUrl:                 c.IndexUrl,
		LoginUrl:                 c.LoginUrl,
		Debug:                    c.Debug,
		Env:                      c.Env,
		InfoLogPath:              c.InfoLogPath,
		ErrorLogPath:             c.ErrorLogPath,
		AccessLogPath:            c.AccessLogPath,
		SqlLog:                   c.SqlLog,
		AccessLogOff:             c.AccessLogOff,
		InfoLogOff:               c.InfoLogOff,
		ErrorLogOff:              c.ErrorLogOff,
		ColorScheme:              c.ColorScheme,
		SessionLifeTime:          c.SessionLifeTime,
		AssetUrl:                 c.AssetUrl,
		FileUploadEngine:         c.FileUploadEngine,
		CustomHeadHtml:           c.CustomHeadHtml,
		CustomFootHtml:           c.CustomFootHtml,
		FooterInfo:               c.FooterInfo,
		LoginTitle:               c.LoginTitle,
		LoginLogo:                c.LoginLogo,
		AuthUserTable:            c.AuthUserTable,
		Extra:                    c.Extra,
		Animation:                c.Animation,
		NoLimitLoginIP:           c.NoLimitLoginIP,
		Logger:                   c.Logger,
		SiteOff:                  c.SiteOff,
		HideConfigCenterEntrance: c.HideConfigCenterEntrance,
		HideAppInfoEntrance:      c.HideAppInfoEntrance,
		HideToolEntrance:         c.HideToolEntrance,
		Custom404HTML:            c.Custom404HTML,
		Custom500HTML:            c.Custom500HTML,
		// UpdateProcessFn:               c.UpdateProcessFn,
		OpenAdminApi:                  c.OpenAdminApi,
		HideVisitorUserCenterEntrance: c.HideVisitorUserCenterEntrance,
		ExcludeThemeComponents:        c.ExcludeThemeComponents,
		prefix:                        c.prefix,
	}
}

func (c *Config) ToMap() map[string]string {
	var m = make(map[string]string, 0)
	m["language"] = c.Language
	m["databases"] = c.Databases.JSON()
	m["domain"] = c.Domain
	m["url_prefix"] = c.UrlPrefix
	m["theme"] = c.Theme
	m["store"] = c.Store.JSON()
	m["title"] = c.Title
	m["logo"] = string(c.Logo)
	m["mini_logo"] = string(c.MiniLogo)
	m["index_url"] = c.IndexUrl
	m["site_off"] = strconv.FormatBool(c.SiteOff)
	m["login_url"] = c.LoginUrl
	m["debug"] = strconv.FormatBool(c.Debug)
	m["env"] = c.Env

	// Logger config
	// ========================

	m["info_log_path"] = c.InfoLogPath
	m["error_log_path"] = c.ErrorLogPath
	m["access_log_path"] = c.AccessLogPath
	m["sql_log"] = strconv.FormatBool(c.SqlLog)
	m["access_log_off"] = strconv.FormatBool(c.AccessLogOff)
	m["info_log_off"] = strconv.FormatBool(c.InfoLogOff)
	m["error_log_off"] = strconv.FormatBool(c.ErrorLogOff)
	m["access_assets_log_off"] = strconv.FormatBool(c.AccessAssetsLogOff)

	m["logger_rotate_max_size"] = strconv.Itoa(c.Logger.Rotate.MaxSize)
	m["logger_rotate_max_backups"] = strconv.Itoa(c.Logger.Rotate.MaxBackups)
	m["logger_rotate_max_age"] = strconv.Itoa(c.Logger.Rotate.MaxAge)
	m["logger_rotate_compress"] = strconv.FormatBool(c.Logger.Rotate.Compress)

	m["logger_encoder_time_key"] = c.Logger.Encoder.TimeKey
	m["logger_encoder_level_key"] = c.Logger.Encoder.LevelKey
	m["logger_encoder_name_key"] = c.Logger.Encoder.NameKey
	m["logger_encoder_caller_key"] = c.Logger.Encoder.CallerKey
	m["logger_encoder_message_key"] = c.Logger.Encoder.MessageKey
	m["logger_encoder_stacktrace_key"] = c.Logger.Encoder.StacktraceKey
	m["logger_encoder_level"] = c.Logger.Encoder.Level
	m["logger_encoder_time"] = c.Logger.Encoder.Time
	m["logger_encoder_duration"] = c.Logger.Encoder.Duration
	m["logger_encoder_caller"] = c.Logger.Encoder.Caller
	m["logger_encoder_encoding"] = c.Logger.Encoder.Encoding
	m["logger_level"] = strconv.Itoa(int(c.Logger.Level))

	m["color_scheme"] = c.ColorScheme
	m["session_life_time"] = strconv.Itoa(c.SessionLifeTime)
	m["asset_url"] = c.AssetUrl
	m["file_upload_engine"] = c.FileUploadEngine.JSON()
	m["custom_head_html"] = string(c.CustomHeadHtml)
	m["custom_foot_html"] = string(c.CustomFootHtml)
	m["custom_404_html"] = string(c.Custom404HTML)
	m["custom_403_html"] = string(c.Custom403HTML)
	m["custom_500_html"] = string(c.Custom500HTML)
	m["footer_info"] = string(c.FooterInfo)
	m["login_title"] = c.LoginTitle
	m["login_logo"] = string(c.LoginLogo)
	m["auth_user_table"] = c.AuthUserTable
	if len(c.Extra) == 0 {
		m["extra"] = ""
	} else {
		m["extra"] = utils.JSON(c.Extra)
	}

	m["animation_type"] = c.Animation.Type
	m["animation_duration"] = fmt.Sprintf("%.2f", c.Animation.Duration)
	m["animation_delay"] = fmt.Sprintf("%.2f", c.Animation.Delay)

	m["no_limit_login_ip"] = strconv.FormatBool(c.NoLimitLoginIP)

	m["hide_config_center_entrance"] = strconv.FormatBool(c.HideConfigCenterEntrance)
	m["hide_app_info_entrance"] = strconv.FormatBool(c.HideAppInfoEntrance)
	m["hide_tool_entrance"] = strconv.FormatBool(c.HideToolEntrance)

	return m
}

func (c *Config) Update(m map[string]string) error {
	updateLock.Lock()
	defer updateLock.Unlock()
	c.Language = m["language"]
	c.Domain = m["domain"]
	c.Theme = m["theme"]
	c.Title = m["title"]
	c.Logo = template.HTML(m["logo"])
	c.MiniLogo = template.HTML(m["mini_logo"])
	c.Debug = utils.ParseBool(m["debug"])
	c.Env = m["env"]
	c.SiteOff = utils.ParseBool(m["site_off"])

	c.AccessLogOff = utils.ParseBool(m["access_log_off"])
	c.InfoLogOff = utils.ParseBool(m["info_log_off"])
	c.ErrorLogOff = utils.ParseBool(m["error_log_off"])
	c.AccessAssetsLogOff = utils.ParseBool(m["access_assets_log_off"])

	if c.InfoLogPath != m["info_log_path"] {
		c.InfoLogPath = m["info_log_path"]
	}
	if c.ErrorLogPath != m["error_log_path"] {
		c.ErrorLogPath = m["error_log_path"]
	}
	if c.AccessLogPath != m["access_log_path"] {
		c.AccessLogPath = m["access_log_path"]
	}
	c.SqlLog = utils.ParseBool(m["sql_log"])

	c.Logger.Rotate.MaxSize, _ = strconv.Atoi(m["logger_rotate_max_size"])
	c.Logger.Rotate.MaxBackups, _ = strconv.Atoi(m["logger_rotate_max_backups"])
	c.Logger.Rotate.MaxAge, _ = strconv.Atoi(m["logger_rotate_max_age"])
	c.Logger.Rotate.Compress = utils.ParseBool(m["logger_rotate_compress"])

	c.Logger.Encoder.Encoding = m["logger_encoder_encoding"]
	loggerLevel, _ := strconv.Atoi(m["logger_level"])
	c.Logger.Level = int8(loggerLevel)

	if c.Logger.Encoder.Encoding == "json" {
		c.Logger.Encoder.TimeKey = m["logger_encoder_time_key"]
		c.Logger.Encoder.LevelKey = m["logger_encoder_level_key"]
		c.Logger.Encoder.NameKey = m["logger_encoder_name_key"]
		c.Logger.Encoder.CallerKey = m["logger_encoder_caller_key"]
		c.Logger.Encoder.MessageKey = m["logger_encoder_message_key"]
		c.Logger.Encoder.StacktraceKey = m["logger_encoder_stacktrace_key"]
		c.Logger.Encoder.Level = m["logger_encoder_level"]
		c.Logger.Encoder.Time = m["logger_encoder_time"]
		c.Logger.Encoder.Duration = m["logger_encoder_duration"]
		c.Logger.Encoder.Caller = m["logger_encoder_caller"]
	}

	initLogger(*c)

	if c.Theme == "adminlte" {
		c.ColorScheme = m["color_scheme"]
	}
	ses, _ := strconv.Atoi(m["session_life_time"])
	if ses != 0 {
		c.SessionLifeTime = ses
	}
	c.CustomHeadHtml = template.HTML(m["custom_head_html"])
	c.CustomFootHtml = template.HTML(m["custom_foot_html"])
	c.Custom404HTML = template.HTML(m["custom_404_html"])
	c.Custom403HTML = template.HTML(m["custom_403_html"])
	c.Custom500HTML = template.HTML(m["custom_500_html"])
	c.FooterInfo = template.HTML(m["footer_info"])
	c.LoginTitle = m["login_title"]
	c.AssetUrl = m["asset_url"]
	c.LoginLogo = template.HTML(m["login_logo"])
	c.NoLimitLoginIP = utils.ParseBool(m["no_limit_login_ip"])

	c.HideConfigCenterEntrance = utils.ParseBool(m["hide_config_center_entrance"])
	c.HideAppInfoEntrance = utils.ParseBool(m["hide_app_info_entrance"])
	c.HideToolEntrance = utils.ParseBool(m["hide_tool_entrance"])

	c.FileUploadEngine = GetFileUploadEngineFromJSON(m["file_upload_engine"])

	c.Animation.Type = m["animation_type"]
	c.Animation.Duration = utils.ParseFloat32(m["animation_duration"])
	c.Animation.Delay = utils.ParseFloat32(m["animation_delay"])

	if m["extra"] != "" {
		var extra = make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(m["extra"]), &extra)
		c.Extra = extra
	}

	return nil
}

// eraseSens erase sensitive info.
func (c *Config) EraseSens() *Config {
	for key := range c.Databases {
		c.Databases[key] = Database{
			Driver: c.Databases[key].Driver,
		}
	}
	return c
}

var (
	globalCfg  = new(Config)
	declare    sync.Once
	updateLock sync.Mutex
)

// ReadFromJson read the Config from a JSON file.
func ReadFromJson(path string) Config {
	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var cfg Config

	err = json.Unmarshal(jsonByte, &cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}

// ReadFromYaml read the Config from a YAML file.
func ReadFromYaml(path string) Config {
	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var cfg Config

	err = yaml.Unmarshal(jsonByte, &cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}

// ReadFromINI read the Config from a INI file.
func ReadFromINI(path string) Config {
	iniCfg, err := ini.Load(path)

	if err != nil {
		panic(err)
	}

	var cfg Config

	err = iniCfg.MapTo(&cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}

var (
	count uint32
	lock  sync.Mutex
)

func SetDefault(cfg Config) Config {
	cfg.Title = utils.SetDefault(cfg.Title, "", "GoAdmin")
	cfg.LoginTitle = utils.SetDefault(cfg.LoginTitle, "", "GoAdmin")
	cfg.Logo = template.HTML(utils.SetDefault(string(cfg.Logo), "", "<b>Go</b>Admin"))
	cfg.MiniLogo = template.HTML(utils.SetDefault(string(cfg.MiniLogo), "", "<b>G</b>A"))
	cfg.Theme = utils.SetDefault(cfg.Theme, "", "adminlte")
	cfg.IndexUrl = utils.SetDefault(cfg.IndexUrl, "", "/info/manager")
	cfg.LoginUrl = utils.SetDefault(cfg.LoginUrl, "", "/login")
	cfg.AuthUserTable = utils.SetDefault(cfg.AuthUserTable, "", "goadmin_users")
	if cfg.Theme == "adminlte" {
		cfg.ColorScheme = utils.SetDefault(cfg.ColorScheme, "", "skin-black")
	}
	cfg.FileUploadEngine.Name = utils.SetDefault(cfg.FileUploadEngine.Name, "", "local")
	cfg.Env = utils.SetDefault(cfg.Env, "", EnvProd)
	if cfg.SessionLifeTime == 0 {
		// default two hours
		cfg.SessionLifeTime = 7200
	}
	return cfg
}

// Set sets the config.
func Set(cfg Config) *Config {

	lock.Lock()
	defer lock.Unlock()

	if atomic.LoadUint32(&count) != 0 {
		panic("can not set config twice")
	}
	atomic.StoreUint32(&count, 1)

	cfg = SetDefault(cfg)

	if cfg.UrlPrefix == "" {
		cfg.prefix = "/"
	} else if cfg.UrlPrefix[0] != '/' {
		cfg.prefix = "/" + cfg.UrlPrefix
	} else {
		cfg.prefix = cfg.UrlPrefix
	}

	initLogger(cfg)

	if cfg.SqlLog {
		logger.OpenSQLLog()
	}

	if cfg.Debug {
		declare.Do(func() {
			fmt.Println(`Server is now running.
Running in "debug" mode. Switch to "release" mode in production.`)
			fmt.Println()
		})
	}

	globalCfg = &cfg

	return globalCfg
}

func initLogger(cfg Config) {
	logger.InitWithConfig(logger.Config{
		InfoLogOff:         cfg.InfoLogOff,
		ErrorLogOff:        cfg.ErrorLogOff,
		AccessLogOff:       cfg.AccessLogOff,
		SqlLogOpen:         cfg.SqlLog,
		InfoLogPath:        cfg.InfoLogPath,
		ErrorLogPath:       cfg.ErrorLogPath,
		AccessLogPath:      cfg.AccessLogPath,
		AccessAssetsLogOff: cfg.AccessAssetsLogOff,
		Rotate: logger.RotateCfg{
			MaxSize:    cfg.Logger.Rotate.MaxSize,
			MaxBackups: cfg.Logger.Rotate.MaxBackups,
			MaxAge:     cfg.Logger.Rotate.MaxAge,
			Compress:   cfg.Logger.Rotate.Compress,
		},
		Encode: logger.EncoderCfg{
			TimeKey:       cfg.Logger.Encoder.TimeKey,
			LevelKey:      cfg.Logger.Encoder.LevelKey,
			NameKey:       cfg.Logger.Encoder.NameKey,
			CallerKey:     cfg.Logger.Encoder.CallerKey,
			MessageKey:    cfg.Logger.Encoder.MessageKey,
			StacktraceKey: cfg.Logger.Encoder.StacktraceKey,
			Level:         cfg.Logger.Encoder.Level,
			Time:          cfg.Logger.Encoder.Time,
			Duration:      cfg.Logger.Encoder.Duration,
			Caller:        cfg.Logger.Encoder.Caller,
			Encoding:      cfg.Logger.Encoder.Encoding,
		},
		Debug: cfg.Debug,
		Level: cfg.Logger.Level,
	})
}

// AssertPrefix return the prefix of assert.
func AssertPrefix() string {
	return globalCfg.AssertPrefix()
}

// GetIndexURL get the index url with prefix.
func GetIndexURL() string {
	return globalCfg.GetIndexURL()
}

// IsProductionEnvironment check the environment if it is production.
func IsProductionEnvironment() bool {
	return globalCfg.IsProductionEnvironment()
}

// URLRemovePrefix remove prefix from the given url.
func URLRemovePrefix(url string) string {
	return globalCfg.URLRemovePrefix(url)
}

func Url(suffix string) string {
	return globalCfg.Url(suffix)
}

// Prefix return the prefix.
func Prefix() string {
	return globalCfg.prefix
}

// PrefixFixSlash return the prefix fix the slash error.
func PrefixFixSlash() string {
	return globalCfg.PrefixFixSlash()
}

// Get gets the config.
func Get() *Config {
	return globalCfg.Copy().EraseSens()
}

// Getter methods
// ============================

func GetDatabases() DatabaseList {
	var list = make(DatabaseList, len(globalCfg.Databases))
	for key := range globalCfg.Databases {
		list[key] = Database{
			Driver: globalCfg.Databases[key].Driver,
		}
	}
	return list
}

func GetDomain() string {
	return globalCfg.Domain
}

func GetLanguage() string {
	return globalCfg.Language
}

func GetUrlPrefix() string {
	return globalCfg.UrlPrefix
}

func GetOpenAdminApi() bool {
	return globalCfg.OpenAdminApi
}

func GetCustom500HTML() template.HTML {
	return globalCfg.Custom500HTML
}

func GetCustom404HTML() template.HTML {
	return globalCfg.Custom404HTML
}

func GetCustom403HTML() template.HTML {
	return globalCfg.Custom403HTML
}

func GetTheme() string {
	return globalCfg.Theme
}

func GetStore() Store {
	return globalCfg.Store
}

func GetTitle() string {
	return globalCfg.Title
}

func GetLogo() template.HTML {
	return globalCfg.Logo
}

func GetSiteOff() bool {
	return globalCfg.SiteOff
}

func GetMiniLogo() template.HTML {
	return globalCfg.MiniLogo
}

func GetIndexUrl() string {
	return globalCfg.IndexUrl
}

func GetLoginUrl() string {
	return globalCfg.LoginUrl
}

func GetDebug() bool {
	return globalCfg.Debug
}

func GetEnv() string {
	return globalCfg.Env
}

func GetInfoLogPath() string {
	return globalCfg.InfoLogPath
}

func GetErrorLogPath() string {
	return globalCfg.ErrorLogPath
}

func GetAccessLogPath() string {
	return globalCfg.AccessLogPath
}

func GetSqlLog() bool {
	return globalCfg.SqlLog
}

func GetAccessLogOff() bool {
	return globalCfg.AccessLogOff
}
func GetInfoLogOff() bool {
	return globalCfg.InfoLogOff
}
func GetErrorLogOff() bool {
	return globalCfg.ErrorLogOff
}

func GetColorScheme() string {
	return globalCfg.ColorScheme
}

func GetSessionLifeTime() int {
	return globalCfg.SessionLifeTime
}

func GetAssetUrl() string {
	return globalCfg.AssetUrl
}

func GetFileUploadEngine() FileUploadEngine {
	return globalCfg.FileUploadEngine
}

func GetCustomHeadHtml() template.HTML {
	return globalCfg.CustomHeadHtml
}

func GetCustomFootHtml() template.HTML {
	return globalCfg.CustomFootHtml
}

func GetFooterInfo() template.HTML {
	return globalCfg.FooterInfo
}

func GetLoginTitle() string {
	return globalCfg.LoginTitle
}

func GetLoginLogo() template.HTML {
	return globalCfg.LoginLogo
}

func GetAuthUserTable() string {
	return globalCfg.AuthUserTable
}

func GetExtra() map[string]interface{} {
	return globalCfg.Extra
}

func GetAnimation() PageAnimation {
	return globalCfg.Animation
}

func GetNoLimitLoginIP() bool {
	return globalCfg.NoLimitLoginIP
}

func GetHideVisitorUserCenterEntrance() bool {
	return globalCfg.HideVisitorUserCenterEntrance
}

func GetExcludeThemeComponents() []string {
	return globalCfg.ExcludeThemeComponents
}

type Service struct {
	C *Config
}

func (s *Service) Name() string {
	return "config"
}

func SrvWithConfig(c *Config) *Service {
	return &Service{c}
}

func GetService(s interface{}) *Config {
	if srv, ok := s.(*Service); ok {
		return srv.C
	}
	panic("wrong service")
}
