// Package accounts provides the data structures and functions for modeling
// the Accounts endpoint ont he Engine Yard API
package accounts

// Entity is a flat data structure that maps to an upstream Account
type Entity struct {
	ID string `json:"id,omitempty"`

	// Account Details
	Billable            bool   `json:"billable,omitempty"`
	BillingRequired     bool   `json:"billing_required,omitempty"`
	EmergencyContact    string `json:"emergency_contact,omitempty"`
	Finalized           bool   `json:"finalized,omitempty"`
	Name                string `json:"name,omitempty"`
	Plan                string `json:"plan,omitempty"`
	RdsManagementWebURI string `json:"rds_management_web_uri,omitempty"`
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

// Copyright 2018 Dennis Walters
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
