package result

// Unlimited is the value of the limiting fields in [UserInfo], which indicates the absence of limits.
const Unlimited = 2147483647

// UserInfo is the result of a successful call to the [getAccountInfo] method.
//
// [getAccountInfo]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#getaccountinfo
type UserInfo struct {
	PlanName            string  `json:"plan_name"`                    // name of the rate
	UserSites           uint64  `json:"user_sites"`                   // the actual number of user sites
	PlanSite            uint64  `json:"plan_site"`                    // maximum number of sites
	UserDomains         uint64  `json:"user_domains"`                 // the actual number of user domains
	PlanDomain          uint64  `json:"plan_domain"`                  // maximum number of domains
	UserMysqlSize       uint64  `json:"user_mysqlsize"`               // the actual volume of the MySQL user database
	PlanMysql           uint64  `json:"plan_mysql"`                   // maximum volume of the MySQL database
	UserQuota           uint64  `json:"user_quota"`                   // the size of the disk quota used by the user
	PlanQuota           uint64  `json:"plan_quota"`                   // maximum disk quota size
	UserFtp             uint64  `json:"user_ftp"`                     // the actual number of user FTP accounts
	PlanFtp             uint64  `json:"plan_ftp"`                     // maximum number of FTP accounts
	UserMail            uint64  `json:"user_mail"`                    // the actual number of user mailboxes
	PlanMail            uint64  `json:"plan_mail"`                    // maximum number of mailboxes
	UserBash            string  `json:"user_bash"`                    // the command shell used
	PlanCp              int64   `json:"plan_cp,omitempty"`            // not documented by Beget API
	UserRateCurrent     float64 `json:"user_rate_current"`            // the current cost of the rate plan per day
	UserIsYearPlan      string  `json:"user_is_year_plan"`            // reports whether the annual discount is used
	UserRateYear        int64   `json:"user_rate_year"`               // the current cost of the rate plan per year
	UserRateMonth       int64   `json:"user_rate_month"`              // the current cost of the rate plan per month
	UserBalance         float64 `json:"user_balance"`                 // User's current balance
	UserDaysToBlock     uint64  `json:"user_days_to_block,omitempty"` // not documented by Beget API
	ServerApacheVersion string  `json:"server_apache_version"`        // apache version on the user's server
	ServerMysqlVersion  string  `json:"server_mysql_version"`         // mysql version on the user's server
	ServerNginxVersion  string  `json:"server_nginx_version"`         // nginx version on the user's server
	ServerPerlVersion   string  `json:"server_perl_version"`          // perl version on the user's server
	ServerPhpVersion    string  `json:"server_php_version"`           // php version on the user's server
	ServerPythonVersion string  `json:"server_python_version"`        // python version on the user's server
	ServerName          string  `json:"server_name"`                  // user's server name
	ServerCPUName       string  `json:"server_cpu_name"`              // user's server cpu name
	ServerMemory        uint64  `json:"server_memory"`                // user's server RAM count
	ServerMemoryCurrent uint64  `json:"server_memorycurrent"`         // the amount of RAM used by the user's server
	ServerLoadAverage   float64 `json:"server_loadaverage"`           // load Average of user's server
	ServerUptime        uint64  `json:"server_uptime"`                // uptime of user's server
}

// SSHToggle is the result of a successful call to the [toggleSsh] method.
// At the moment, in any case, the result will be an empty array.
//
// [toggleSsh]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#togglessh
type SSHToggle []any
