package onepassword

var (
	// Known categories from https://support.1password.com/opvault-design/
	CatLogin           = Category{"001", "Login"}
	CatCreditCard      = Category{"002", "Credit Card"}
	CatSecureNote      = Category{"003", "Secure Note"}
	CatIdentity        = Category{"004", "Identity"}
	CatPassword        = Category{"005", "Password"}
	CatTombstone       = Category{"099", "Tombstone"}
	CatSoftwareLicense = Category{"100", "Software License"}
	CatBankAccount     = Category{"101", "Bank Account"}
	CatDatabase        = Category{"102", "Database"}
	CatDriverLicense   = Category{"103", "Driver License"}
	CatOutdoorLicense  = Category{"104", "Outdoor License"}
	CatMembership      = Category{"105", "Membership"}
	CatPassport        = Category{"106", "Passport"}
	CatRewards         = Category{"107", "Rewards"}
	CatSSN             = Category{"108", "SSN"}
	CatRouter          = Category{"109", "Router"}
	CatServer          = Category{"110", "Server"}
	CatEmail           = Category{"111", "Email"}
)

type Category struct {
	Uuid string `json: "uuid"`
	Name string `json: "name"`
}

type Field struct {
	Value string `json:"v"`
	Name  string `json:"t"`
}

type Section struct {
	Fields []Field `json:"fields"`
}

type Note struct {
	Sections    []Section `json:"sections"`
	Description string    `json:"notesPlain"`
}

type Item struct {
	Title     string   `json:"title"`
	Url       string   `json:"url"`
	Tags      []string `json:"tags"`
	Category  Category `json:"cat"`
	Details   []byte // JSON encoded object. Structure is based on category.
}
