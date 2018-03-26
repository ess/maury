package accounts

type Entity struct {
	ID string `json:"id,omitempty"`

	// Account Details
	Billable            bool   `json:"billable,omitempty"`
	BillingRequired     bool   `json:"billing_required,omitempty"`
	EmergencyContact    string `json:"emergency_contact,omitempty"`
	Finalized           bool   `json:"finalized,omitempty"`
	Name                string `json:"name,omitempty"`
	Plan                string `json:"plan,omitempty"`
	RdsManagementWebUri string `json:"rds_management_web_uri,omitempty"`
	SignupVia           string `json:"signup_via"`
	SupportPlan         string `json:"support_plan,omitempty"`
	SupportTrialStatus  string `json:"support_trial_status"`
	Type                string `json:"type,omitempty"`

	// Relation URLs
	AccountNotes     string `json:"account_notes,omitempty"`
	AccountTrial     string `json:"account_trial,omitempty"`
	Addresses        string `json:"addresses,omitempty"`
	Applications     string `json:"applications,omitempty"`
	DatabaseServices string `json:"database_services,omitempty"`
	Environments     string `json:"environments,omitempty"`
	Features         string `json:"features,omitempty"`
	LogicalDatabases string `json:"logical_databases,omitempty"`
	Memberships      string `json:"memberships,omitempty"`
	Owners           string `json:"owners,omitempty"`
	Providers        string `json:"providers,omitempty"`
	Requests         string `json:"requests,omitempty"`
	ServerAlerts     string `json:"server_alerts,omitempty"`
	SslCertificates  string `json:"ssl_certificates,omitempty"`
	Users            string `json:"users,omitempty"`

	// Timestamps
	CanceledAt  string `json:"canceled_at,omitempty"`
	CancelledAt string `json:"cancelled_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	Updatedat   string `json:"updated_at,omitempty"`
}
