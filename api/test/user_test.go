package test

import (
	"net/url"

	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/api/user"
	"github.com/ThCompiler/go.beget.api/core"
	"github.com/stretchr/testify/require"
)

func (ap *APISuite) TestGetAccountInfo() {
	expectedUserInfo := result.UserInfo{
		PlanName:            "Tog",
		UserSites:           1,
		PlanSite:            3,
		UserDomains:         2,
		PlanDomain:          2147483647,
		UserMysqlSize:       1,
		PlanMysql:           2147483647,
		UserQuota:           1,
		PlanQuota:           10000,
		UserFtp:             1,
		PlanFtp:             2147483647,
		UserMail:            1,
		PlanMail:            2147483647,
		UserBash:            "/bin/",
		PlanCp:              80,
		UserRateCurrent:     12.82,
		UserIsYearPlan:      "0",
		UserRateYear:        5321,
		UserRateMonth:       200,
		UserBalance:         2500,
		UserDaysToBlock:     572,
		ServerApacheVersion: "2.4.55",
		ServerMysqlVersion:  "5.7.21-20",
		ServerNginxVersion:  "2.4.55",
		ServerPerlVersion:   "5.34.0",
		ServerPythonVersion: "3.10.12",
		ServerPhpVersion:    "8.2.13",
		ServerName:          "burut.beget.com",
		ServerCPUName:       "62 *  AMD EPYC 7763 64-Core Processor",
		ServerMemory:        772905,
		ServerMemoryCurrent: 30460,
		ServerLoadAverage:   20.920000076294,
		ServerUptime:        71,
	}

	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.UserGetAccountInfo(getAccountResponse, func(values url.Values) {
		RequireEqualValues(ap.T(), url.Values{}, values, client)
	})

	req, err := core.PrepareRequest[result.UserInfo](
		client,
		user.CallGetAccountInfo(),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.EqualValues(ap.T(), expectedUserInfo, *res)
}

func (ap *APISuite) TestToggleSSHForUser() {
	expectedResponse := result.SSHToggle{}

	toggleStatus := user.ENABLE
	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.UserToggleSSH(toggleSSHResponse, func(values url.Values) {
		RequireEqualValues(ap.T(), url.Values{
			"status": []string{string(toggleStatus)},
		}, values, client)
	})

	req, err := core.PrepareRequest[result.SSHToggle](
		client,
		user.CallToggleSSH(toggleStatus),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.EqualValues(ap.T(), expectedResponse, *res)
}

func (ap *APISuite) TestToggleSSHForFTPUser() {
	expectedResponse := result.SSHToggle{}

	toggleStatus := user.ENABLE
	ftpUserName := "ftp"
	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.UserToggleSSH(toggleSSHResponse, func(values url.Values) {
		RequireEqualValues(ap.T(), url.Values{
			"status":   []string{string(toggleStatus)},
			"ftplogin": []string{ftpUserName},
		}, values, client)
	})

	req, err := core.PrepareRequest[result.SSHToggle](
		client,
		user.CallToggleSSHFTP(toggleStatus, ftpUserName),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.EqualValues(ap.T(), expectedResponse, *res)
}

const getAccountResponse = `
{
    "status": "success",
    "answer": {
        "status": "success",
        "result": {
            "plan_name": "Tog",
            "user_sites": 1,
            "plan_site": 3,
            "user_domains": 2,
            "plan_domain": 2147483647,
            "user_mysqlsize": 1,
            "plan_mysql": 2147483647,
            "user_quota": 1,
            "plan_quota": 10000,
            "user_ftp": 1,
            "plan_ftp": 2147483647,
            "user_mail": 1,
            "plan_mail": 2147483647,
            "user_bash": "/bin/",
            "plan_cp": 80,
            "user_rate_current": 12.82,
            "user_is_year_plan": "0",
            "user_rate_year": 5321,
            "user_rate_month": 200,
            "user_balance": 2500,
            "user_days_to_block": 572,
            "server_apache_version": "2.4.55",
            "server_mysql_version": "5.7.21-20",
            "server_nginx_version": "2.4.55",
            "server_perl_version": "5.34.0",
            "server_python_version": "3.10.12",
            "server_php_version": "8.2.13",
            "server_name": "burut.beget.com",
            "server_cpu_name": "62 *  AMD EPYC 7763 64-Core Processor",
            "server_memory": 772905,
            "server_memorycurrent": 30460,
            "server_loadaverage": 20.920000076294,
            "server_uptime": 71
        }
    }
}
`

const toggleSSHResponse = `
{
    "status": "success",
    "answer": {
        "status": "success",
        "result": []
    }
}
`
