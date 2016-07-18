package structs

type Configuration struct {
	DBConfig       *DBConf `json:"database,omitempty"`
	JwtKey         string  `json:"jwtkey,omitempty"`
	WebDir         string  `json:"webdir,omitempty"`
	RulebaseDir    string  `json:"rulebasedir,omitempty"`
	Gopathrelative bool    `json:"gopathrelative,omitempty"`
	Loadtestdata   bool    `json:"loadtestdata,omitempty"`
}

type DBConf struct {
	Location string `json:"location"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty`
}

type User struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Firstname  string     `json:"firstname"`
	Lastname   string     `json:"lastname"`
	Locale     string     `json:"locale"`
	Revision   string     `json:"revision"`
	Dictionary Dictionary `json:"dictionary"`
	//Documents []string `json:"documents"`
}

//Response represents a JSON response from the ducklib server
type Response struct {
	Ok        bool        `json:"ok"`
	Reason    *string     `json:"reason,omitempty"`
	ID        *string     `json:"id,omitempty"`
	Documents *[]Document `json:"documents,omitempty"`
}

type ComplianceResponse struct {
	Ok        bool        `json:"ok"`
	Compliant string      `json:"compliant"`
	Documents []*Document `json:"documents,omitempty"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) FromValueMap(mp map[string]interface{}) {

	if id, ok := mp["_id"]; ok {
		u.ID = id.(string)
	}
	if rev, ok := mp["_rev"]; ok {
		u.Revision = rev.(string)
	}
	if name, ok := mp["firstname"]; ok {
		u.Firstname = name.(string)
	}
	if owner, ok := mp["lastname"]; ok {
		u.Lastname = owner.(string)
	}
	if owner, ok := mp["password"]; ok {
		u.Password = owner.(string)
	}
	if owner, ok := mp["email"]; ok {
		u.Email = owner.(string)
	}
	if locale, ok := mp["locale"]; ok {
		u.Locale = locale.(string)
	}

	/*	if docs, prs := mp["documents"].([]interface{}); prs {
		u.Documents = make([]string, len(docs))
		for i, v := range docs {
			u.Documents[i] = v.(string)
		}
	}*/

}

type Rulebase struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Revision string `json:"_rev"`
}

//can this be a map?
type Taxonomy struct {
	Scope []struct {
		Value    string `json:"value"`
		Code     string `json:"code"`
		Category string `json:"category"`
		Fixed    bool   `json:"fixed"`
	} `json:"scope"`
	Qualifier []struct {
		Value    string `json:"value"`
		Code     string `json:"code"`
		Category string `json:"category"`
		Fixed    bool   `json:"fixed"`
	} `json:"qualifier"`
	DataCategory []struct {
		Value    string `json:"value"`
		Code     string `json:"code"`
		Category string `json:"category"`
		Fixed    bool   `json:"fixed"`
	} `json:"dataCategory"`
	Action []struct {
		Value    string `json:"value"`
		Code     string `json:"code"`
		Category string `json:"category"`
		Fixed    bool   `json:"fixed"`
	} `json:"action"`
}

type DictionaryEntry struct {
	Value          string
	Type           string
	Code           string
	Category       string
	DictionaryType string
}

type Dictionary []DictionaryEntry
