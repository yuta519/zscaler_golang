package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"zscaler_golang/pkg/infra"
	"zscaler_golang/pkg/zia/auth"
	"zscaler_golang/pkg/zia/config"
)

type AdminUser struct {
	Id                          int                    `json:"id"`
	LoginName                   string                 `json:"loginName"`
	UserName                    string                 `json:"userName"`
	Email                       string                 `json:"email"`
	Role                        AdminUserRole          `json:"role"`
	Comments                    string                 `json:"comments"`
	AdminScope                  map[string]interface{} `json:"adminScop"`
	IsNoneEditable              bool                   `json:"isNonEditable"`
	Disabled                    bool                   `json:"disabled"`
	IsAuditor                   bool                   `json:"isAuditor"`
	Password                    string                 `json:"password"`
	IsPasswordLoginAllowed      bool                   `json:"isPasswordLoginAllowed"`
	IsSecurityReportCommEnabled bool                   `json:"isSecurityReportCommEnabled"`
	IsPasswordExpired           bool                   `json:"isPasswordExpired"`
	IsExecMobileAppEnabled      bool                   `json:"isExecMobileAppEnabled"`
	ExecMobileAppTokens         map[string]interface{} `json:"ExecMobileAppTokens"`
}

type AdminUserRole struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Extensions map[string]string `json:"extensions"`
}

type CreateAdminUserParameter struct {
	LoginName              string         `json:"loginName"`
	UserName               string         `json:"userName"`
	Email                  string         `json:"email"`
	Password               string         `json:"password"`
	Role                   map[string]int `json:"role"`
	AdminScopeType         string         `json:"adminScopeType"`
	IsPasswordLoginAllowed bool           `json:"isPasswordLoginAllowed"`
	Name                   string         `json:"name"`
}

type SuccessResultOfCreatedAdminUser struct {
	Id                                 int                    `json:"id"`
	LoginName                          string                 `json:"loginName"`
	UserName                           string                 `json:"userName"`
	Email                              string                 `json:"email"`
	Password                           string                 `json:"password"`
	Role                               map[string]interface{} `json:"role"`
	AdminScopescopeGroupMemberEntities []string               `json:"adminScopescopeGroupMemberEntities"`
	AdminScopeType                     string                 `json:"adminScopeType"`
	PwdLastModifiedTime                int                    `json:"pwdLastModifiedTime"`
	Name                               string                 `json:"name"`
}

func FetchAllAdminUsers() []AdminUser {
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	adminuserRef, _ := url.Parse("/api/v1/adminUsers")
	adminuserEdpoint := baseUrl.ResolveReference(adminuserRef).String()

	session_id := auth.Login()
	response := infra.GetApi(adminuserEdpoint, session_id)
	auth.Logout()

	var adminUsers []AdminUser
	json.Unmarshal(response, &adminUsers)
	return adminUsers
}

func CreateAdminUsers(
	role_name string,
	login_name string,
	user_name string,
	email string,
	isPassword bool,
	password string,
	admin_scope string,
	name string,
) string {
	session_id := auth.Login()

	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	adminrole_reference, _ := url.Parse("/api/v1/adminRoles/lite")
	adminuser_reference, _ := url.Parse("/api/v1/adminUsers")
	adminrole_endpoint := url_base.ResolveReference(adminrole_reference).String()
	adminuser_endpoint := url_base.ResolveReference(adminuser_reference).String()

	adminrole_response := infra.GetApi(adminrole_endpoint, session_id)
	var admin_roles []AdminRole
	json.Unmarshal(adminrole_response, &admin_roles)

	var target_role_id int
	for _, admin_role := range admin_roles {
		if admin_role.Name == role_name {
			target_role_id = admin_role.Id
		}
	}

	var payload CreateAdminUserParameter
	payload.LoginName = login_name
	payload.UserName = user_name
	payload.Email = email
	payload.IsPasswordLoginAllowed = isPassword
	payload.Password = password
	payload.Role = map[string]int{"id": target_role_id}
	payload.AdminScopeType = admin_scope
	payload.Name = name
	payload_json, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		adminuser_endpoint,
		bytes.NewBuffer(payload_json),
	)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", session_id)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	auth.Logout()

	resp_byte, _ := ioutil.ReadAll(resp.Body)
	var message string
	if resp.StatusCode == 200 {
		var adminuser AdminUser
		json.Unmarshal(resp_byte, &adminuser)
		message = "Success: " + adminuser.LoginName + " is created."
	} else {
		var failed_message map[string]string
		json.Unmarshal(resp_byte, &failed_message)
		message = "Failed: " + strconv.Itoa(resp.StatusCode) + failed_message["code"] + failed_message["message"]
	}

	return message
}
