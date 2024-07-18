package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

type dataset struct {
	// The name of the dataset.
	Name string `json:"name,omitempty"`

	// A description for the dataset.
	Description string `json:"description,omitempty"`

	// The maximum unpacking depth of nested JSON fields.
	ExpandJSONDepth int `json:"expand_json_depth,omitempty"`

	// The 'slug' of the dataset to be used in URLs.
	Slug string `json:"slug,omitempty"`

	// The total number of unique fields for this Dataset. The value will be
	// null if the dataset does not contain any fields yet.
	RegularColumnsCount int `json:"regular_columns_count,omitempty"`

	// The ISO8601-formatted time when the dataset last received event data.
	// The value will be null if no data has been received yet.
	LastWrittenAt string `json:"last_written_at,omitempty"`

	// The ISO8601-formatted time when the dataset was created.
	CreatedAt string `json:"created_at,omitempty"`
}

func newDatasetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "datasets",
		Aliases: []string{"d"},
		Short:   "Manage Datasets",
		Long: "A Dataset represents a collection of related events that come from the\n" +
			"same source, or are related to the same source. \n" +
			"\n" +
			"These commands allow you to list, create, update, and delete Datasets.",
	}

	cmd.AddCommand(
		newDatasetsCreateCmd(),
		newDatasetsListCmd(),
		newDatasetsGetCmd(),
		newDatasetsDeleteCmd(),
		newDatasetsUpdateCmd(),
	)

	return cmd
}

// Create a Dataset
// https://docs.honeycomb.io/api/tag/Datasets#operation/createDataset
func newDatasetsCreateCmd() *cobra.Command {
	var (
		dName            string
		dDescription     string
		dExpandJSONDepth int
	)
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"add", "new"},
		Short:   "Create a Dataset.",
		Long: "Create a Dataset.  If a Dataset already exists by that name (or slug), then\n" +
			"the existing dataset will be returned.",
		Run: func(cmd *cobra.Command, args []string) {
			var d = dataset{
				Name:            dName,
				Description:     dDescription,
				ExpandJSONDepth: dExpandJSONDepth,
			}
			var bodyMarshal, err = json.Marshal(d)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newDatasetsCreateCmd",
					"err":       err,
					"dataset":   d,
				}).Fatal("Error received when attempting to marshal a dataset.")
			}
			var p = payload{
				Method:   http.MethodPost,
				Path:     "/1/datasets",
				Body:     bodyMarshal,
				Response: &dataset{},
			}

			err = p.GetResponse(true)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newDatasetsCreateCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to create a new dataset.")
			}
		},
	}

	cmd.Flags().StringVarP(&dName, "name", "n", "",
		"The name of the dataset.")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&dDescription, "description", "d", "",
		"A description for the dataset.")
	cmd.Flags().IntVarP(&dExpandJSONDepth, "expand_json_depth", "e", 0,
		"The maximum unpacking depth of nested JSON fields.")

	return cmd
}

// List All Datasets
// https://docs.honeycomb.io/api/tag/Datasets#operation/listDatasets
func newDatasetsListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all Datasets.",
		Long:    "Lists all Datasets for an environment.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = payload{
				Method:   http.MethodGet,
				Path:     "/1/datasets",
				Response: &[]dataset{},
			}

			var err = p.GetResponse(true)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newDatasetsListCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to list all datasets.")
			}
		},
	}

	return cmd
}

// Get a Dataset
// https://docs.honeycomb.io/api/tag/Datasets#operation/getDataset
func newDatasetsGetCmd() *cobra.Command {
	var (
		dSlug string
	)
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{},
		Short:   "Get a Dataset.",
		Long:    "Get a single Dataset by slug.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = payload{
				Method:   http.MethodGet,
				Path:     "/1/datasets/" + dSlug,
				Response: &dataset{},
			}

			var err = p.GetResponse(true)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newDatasetsGetCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to get a dataset.")
			}
		},
	}

	cmd.Flags().StringVarP(&dSlug, "slug", "s", "",
		"The dataset slug.")
	cmd.MarkFlagRequired("slug")

	return cmd
}

// Delete a Dataset
// https://docs.honeycomb.io/api/tag/Datasets#operation/deleteDataset
func newDatasetsDeleteCmd() *cobra.Command {
	var (
		dSlug string
	)
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm", "remove", "del"},
		Short:   "Delete a Dataset.",
		Long: "Asynchronously delete a dataset.\n" +
			"\n" +
			"WARNING: This endpoint will allow anyone with an API key that has the manage dataset\n" +
			"permission to delete any dataset in the environment (or any dataset in the whole team\n" +
			"for Classic customers).\n" +
			"This might make you sad.\n" +
			"\n" +
			"This endpoint is not enabled by default and will return a 403 Forbidden if you try to\n" +
			"use it. If you would like access to this endpoint despite the above-listed risks, please\n" +
			"have your Honeycomb team owner contact Honeycomb Support or email Support.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = payload{
				Method:   http.MethodDelete,
				Path:     "/1/datasets/" + dSlug,
				Response: nil,
			}

			var err = p.GetResponse(true)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newDatasetsDeleteCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to delete a dataset.")
			}
		},
	}

	cmd.Flags().StringVarP(&dSlug, "slug", "s", "",
		"The dataset slug.")
	cmd.MarkFlagRequired("slug")

	return cmd
}

// Update a Dataset
// https://docs.honeycomb.io/api/tag/Datasets#operation/updateDataset
func newDatasetsUpdateCmd() *cobra.Command {
	var (
		dSlug            string
		dDescription     string
		dExpandJSONDepth int
	)
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"up", "edit", "modify", "change", "set"},
		Short:   "Update a Dataset.",
		Long: "Update a dataset's description or expand_json_depth setting.\n" +
			"\n" +
			"Both fields must be specified, as omitting one will have the effect of reverting the\n" +
			"setting to the default.",
		Run: func(cmd *cobra.Command, args []string) {
			var d = dataset{
				Description:     dDescription,
				ExpandJSONDepth: dExpandJSONDepth,
			}
			var bodyMarshal, err = json.Marshal(d)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newDatasetsUpdateCmd",
					"err":       err,
					"dataset":   d,
				}).Fatal("Error received when attempting to marshal a dataset.")
			}
			var p = payload{
				Method:   http.MethodPut,
				Path:     "/1/datasets/" + dSlug,
				Body:     bodyMarshal,
				Response: &dataset{},
			}

			err = p.GetResponse(true)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newDatasetsUpdateCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to update a dataset.")
			}
		},
	}

	cmd.Flags().StringVarP(&dSlug, "slug", "s", "",
		"The dataset slug.")
	cmd.MarkFlagRequired("slug")
	cmd.Flags().StringVarP(&dDescription, "description", "d", "",
		"A description for the dataset.")
	cmd.MarkFlagRequired("description")
	cmd.Flags().IntVarP(&dExpandJSONDepth, "expand_json_depth", "e", 0,
		"The maximum unpacking depth of nested JSON fields.")
	cmd.MarkFlagRequired("expand_json_depth")

	return cmd
}
