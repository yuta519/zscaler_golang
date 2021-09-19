package pkg

type AdminRole struct {
	Id       int    `json:"id"`
	Rank     int    `json:"rank"`
	Name     string `json:"name"`
	RoleType string `json:"roleType"`
}

type ApiCredential struct {
	APIKey    string `json:"apiKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Timestamp string `json:"timestamp"`
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

type UrlCategory struct {
	Id                               string   `json:"id"`
	ConfiguredName                   string   `json:"configuredName"`
	Urls                             []string `json:"urls"`
	DbCategorizedUrls                []string `json:"dbCategorizedUrls"`
	CustomCateogry                   bool     `json:"customCateogry"`
	Editable                         bool     `json:"editable"`
	Description                      string   `json:"description"`
	Type                             string   `json:"Type"`
	Val                              int      `json:"val"`
	CustomUrlsCount                  int      `json:"customUrlsCount"`
	UrlsRetainingParentCategoryCount int      `json:"urlsRetainingParentCategoryCount"`
}

type CreateUrlFilterRuleParameter struct {
	AccessControl       string        `json:"accessControl"`
	Name                string        `json:"name"`
	Order               int           `json:"order"`
	Protocols           []string      `json:"protocols"`
	Users               []interface{} `json:"users"`
	UrlCategories       []string      `json:"urlCategories"`
	State               string        `json:"state"`
	Rank                int           `json:"rank"`
	RequestMethods      []string      `json:"requestMethods"`
	BlockOverride       bool          `json:"blockOverride"`
	EnforceTimeValidity bool          `json:"enforceTimeValidity"`
	CbiProfileId        int           `json:"cbiProfileId"`
	Action              string        `json:"action"`
}
