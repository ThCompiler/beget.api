package result

type UserInfo struct {
	PlanName            string  `json:"plan_name"`
	UserSites           int64   `json:"user_sites"`
	PlanSite            int64   `json:"plan_site"`
	UserDomains         int64   `json:"user_domains"`
	PlanDomain          int64   `json:"plan_domain"`
	UserMysqlSize       int64   `json:"user_mysqlsize"`
	PlanMysql           int64   `json:"plan_mysql"`
	UserQuota           int64   `json:"user_quota"`
	PlanQuota           int64   `json:"plan_quota"`
	UserFtp             int64   `json:"user_ftp"`
	PlanFtp             int64   `json:"plan_ftp"`
	UserMail            int64   `json:"user_mail"`
	PlanMail            int64   `json:"plan_mail"`
	UserBash            string  `json:"user_bash"`
	UserRateCurrent     string  `json:"user_rate_current"`
	UserIsYearPlan      string  `json:"user_is_year_plan"`
	UserRateYear        int64   `json:"user_rate_year"`
	UserRateMonth       int64   `json:"user_rate_month"`
	UserBalance         float64 `json:"user_balance"`
	ServerApacheVersion string  `json:"server_apache_version"`
	ServerMysqlVersion  string  `json:"server_mysql_version"`
	ServerNginxVersion  string  `json:"server_nginx_version"`
	ServerPerlVersion   string  `json:"server_perl_version"`
	ServerPhpVersion    string  `json:"server_php_version"`
	ServerPythonVersion string  `json:"server_python_version"`
	ServerName          string  `json:"server_name"`
	ServerCpuName       string  `json:"server_cpu_name"`
	ServerMemory        string  `json:"server_memory"`
	ServerMemoryCurrent int64   `json:"server_memorycurrent"`
	ServerLoadAverage   string  `json:"server_loadaverage"`
	ServerUptime        string  `json:"server_uptime"`
}
