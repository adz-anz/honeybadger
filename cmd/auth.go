package cmd

import (
	"net/http"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

type authAPIKeyAccess struct {
	Events         bool `json:"events,omitempty"`
	Markers        bool `json:"markers,omitempty"`
	Triggers       bool `json:"triggers,omitempty"`
	Boards         bool `json:"boards,omitempty"`
	Queries        bool `json:"queries,omitempty"`
	Columns        bool `json:"columns,omitempty"`
	CreateDatasets bool `json:"createDatasets,omitempty"`
	SLOs           bool `json:"slos,omitempty"`
	Recipients     bool `json:"recipients,omitempty"`
	PrivateBoards  bool `json:"privateBoards,omitempty"`
}

type authEnvironment struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type authTeam struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type auth struct {
	APIKeyAccess authAPIKeyAccess `json:"api_key_access,omitempty"`
	Environment  authEnvironment  `json:"environment,omitempty"`
	Team         authTeam         `json:"team,omitempty"`
}

// Authorizations
// https://docs.honeycomb.io/api/tag/Auth
func newAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Aliases: []string{"a"},
		Short: "Manage API keys",
		Long: "API keys can have various permissions and belong to a specific Environment.\n" +
			"\n" +
			"This command can be used to validate authentication for a key, to determine\n" +
			"what authorizations have been granted to a key, and to determine the Team and\n" +
			"Environment that a key belongs to.",
	}

	cmd.AddCommand(newAuthListCmd())

	return cmd
}

// List Authorizations
// https://docs.honeycomb.io/api/tag/Auth#operation/getAuth
func newAuthListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "get"},
		Short:   "List authorizations",
		Long: "Lists all authorizations that have been granted for an API key within a team\n" +
			"and environment.\n" +
			"\n" +
			"Note: a Honeycomb Classic API key will return an empty string for both of the\n" +
			"environment values.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = Payload{
				Method: http.MethodGet,
				Path:   "/1/auth",
			}

			var err = p.Execute()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newAuthListCmd",
					"err": err,
					"payload": p,
				}).Fatal("Error received when attempting to list authorizations.")
			}
		},
	}

	return cmd
}
