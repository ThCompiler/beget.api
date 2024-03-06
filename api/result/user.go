package result

type UserInfo struct {
	PlanName            string  `json:"plan_name"`
	UserSites           uint64  `json:"user_sites"`
	PlanSite            uint64  `json:"plan_site"`
	UserDomains         uint64  `json:"user_domains"`
	PlanDomain          uint64  `json:"plan_domain"`
	UserMysqlSize       uint64  `json:"user_mysqlsize"`
	PlanMysql           uint64  `json:"plan_mysql"`
	UserQuota           uint64  `json:"user_quota"`
	PlanQuota           uint64  `json:"plan_quota"`
	UserFtp             uint64  `json:"user_ftp"`
	PlanFtp             uint64  `json:"plan_ftp"`
	UserMail            uint64  `json:"user_mail"`
	PlanMail            uint64  `json:"plan_mail"`
	UserBash            string  `json:"user_bash"`
	PlanCp              int64   `json:"plan_cp,omitempty"`
	UserRateCurrent     float64 `json:"user_rate_current"`
	UserIsYearPlan      string  `json:"user_is_year_plan"`
	UserRateYear        int64   `json:"user_rate_year"`
	UserRateMonth       int64   `json:"user_rate_month"`
	UserBalance         float64 `json:"user_balance"`
	UserDaysToBlock     uint64  `json:"user_days_to_block,omitempty"`
	ServerApacheVersion string  `json:"server_apache_version"`
	ServerMysqlVersion  string  `json:"server_mysql_version"`
	ServerNginxVersion  string  `json:"server_nginx_version"`
	ServerPerlVersion   string  `json:"server_perl_version"`
	ServerPhpVersion    string  `json:"server_php_version"`
	ServerPythonVersion string  `json:"server_python_version"`
	ServerName          string  `json:"server_name"`
	ServerCpuName       string  `json:"server_cpu_name"`
	ServerMemory        uint64  `json:"server_memory"`
	ServerMemoryCurrent uint64  `json:"server_memorycurrent"`
	ServerLoadAverage   float64 `json:"server_loadaverage"`
	ServerUptime        uint64  `json:"server_uptime"`
}

type SSHToggle []any
