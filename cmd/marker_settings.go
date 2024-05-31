package cmd

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

type markerSettings struct {
	// Groups similar Markers. For example, 'deploys'. All Markers of the same
	// type appear with the same color on the graph.
	Type string `json:"type,omitempty"`

	// Color to use for display of this marker type. Specified as hexadecimal
	// RGB. For example, "#F96E11".
	Color string `json:"color,omitempty"`

	// The unique identifier (ID) of a Marker Setting.
	ID string `json:"id,omitempty"`

	// The ISO8601-formatted time when the Marker Setting was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The ISO8601-formatted time when the Marker Setting was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func newMarkerSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "marker_settings",
		Aliases: []string{"ms"},
		Short:   "Manage Marker Settings",
		Long: "Marker Settings apply to groups of similar Markers. For example, `deploys` markers\n" +
			"appear with the same color on a graph.\n" +
			"\n" +
			"These commands allow you to list, create, update, and delete Marker Settings.",
	}

	cmd.PersistentFlags().StringVarP(&targetDataset, "dataset", "d", "__all__",
		"The dataset slug or use __all__ (or omit) for endpoints that support environment-wide operations.")

	cmd.AddCommand(
		newMarkersSettingsCreateCmd(),
		newMarkersSettingsGetCmd(),
		newMarkersSettingsUpdateCmd(),
		newMarkersSettingsDeleteCmd(),
	)

	return cmd
}

// Create a Marker Setting
// https://docs.honeycomb.io/api/tag/Marker-Settings#operation/createMarkerSetting
func newMarkersSettingsCreateCmd() *cobra.Command {
	var (
		msType  string
		msColor string
	)

	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"add", "new", "insert", "put"},
		Short:   "Create a Marker Setting in the specified dataset.",
		Long:    "TODO: Update this with the actual description.",
		Run: func(cmd *cobra.Command, args []string) {
			var ms = markerSettings{
				Type:  msType,
				Color: msColor,
			}

			var bodyMarshal, err = json.Marshal(ms)
			if err != nil {
				log.WithFields(log.Fields{
					"_function":      "newMarkersSettingsCreateCmd",
					"err":            err,
					"marker_setting": ms,
				}).Fatal("Error received when attempting to marshal a marker setting.")
			}
			var p = Payload{
				Method:   http.MethodPost,
				Path:     "/1/marker_settings/" + targetDataset,
				Body:     bodyMarshal,
				Response: &markerSettings{},
			}

			err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersSettingsCreateCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to create a new marker setting.")
			}
		},
	}

	cmd.Flags().StringVarP(&msType, "type", "t", "",
		"Groups similar Markers. For example, 'deploys'. All Markers of the same type appears with the same color on the graph.")
	cmd.MarkFlagRequired("type")
	cmd.Flags().StringVarP(&msColor, "color", "c", "",
		"Color to use for display of this marker type. Specified as hexadecimal RGB. For example, \"#F96E11\".")
	cmd.MarkFlagRequired("color")

	return cmd
}

// Get a Marker Setting
// https://docs.honeycomb.io/api/tag/Marker-Settings#operation/listMarkerSettings
func newMarkersSettingsGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"ls", "list"},
		Short:   `List all Marker Settings in the specified dataset.`,
		Long:    `TODO: Update this with the actual description.`,
		Example: `Example`,
		Run: func(cmd *cobra.Command, args []string) {
			var p = Payload{
				Method:   http.MethodGet,
				Path:     "/1/marker_settings/" + targetDataset,
				Response: &[]markerSettings{},
			}

			var err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersSettingsGetCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to list all marker settings.")
			}
		},
	}

	return cmd
}

// Update a Marker Setting
// https://docs.honeycomb.io/api/tag/Marker-Settings#operation/updateMarkerSettings
func newMarkersSettingsUpdateCmd() *cobra.Command {
	var (
		msId    string
		msType  string
		msColor string
	)

	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"up", "edit", "modify", "change", "set"},
		Short:   "Update a Marker Setting in the specified dataset.",
		Long:    `TODO: Update this with the actual description.`,
		Example: `Example`,
		Run: func(cmd *cobra.Command, args []string) {
			var ms = markerSettings{
				ID:    msId,
				Type:  msType,
				Color: msColor,
			}

			var bodyMarshal, err = json.Marshal(ms)
			if err != nil {
				log.WithFields(log.Fields{
					"_function":      "newMarkersSettingsUpdateCmd",
					"err":            err,
					"marker_setting": ms,
				}).Fatal("Error received when attempting to marshal a marker setting.")
			}
			var p = Payload{
				Method:   http.MethodPut,
				Path:     "/1/markers/" + targetDataset + "/" + ms.ID,
				Body:     bodyMarshal,
				Response: &markerSettings{},
			}

			err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersSettingsUpdateCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to update a marker setting.")
			}
		},
	}

	cmd.Flags().StringVarP(&msId, "id", "i", "", "The ID of the marker setting to update.")
	cmd.MarkFlagRequired("id")
	cmd.Flags().StringVarP(&msType, "type", "t", "",
		"Groups similar Markers. For example, 'deploys'. All Markers of the same type appears with the same color on the graph.")
	cmd.MarkFlagRequired("type")
	cmd.Flags().StringVarP(&msColor, "color", "c", "",
		"Color to use for display of this marker type. Specified as hexadecimal RGB. For example, \"#F96E11\".")
	cmd.MarkFlagRequired("color")

	return cmd
}

// Delete a Marker Setting
// https://docs.honeycomb.io/api/tag/Marker-Settings#operation/deleteMarkerSettings
func newMarkersSettingsDeleteCmd() *cobra.Command {
	var (
		msId string
	)

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm", "remove", "del"},
		Short:   "Delete a Marker Setting in the specified dataset.",
		Long:    `TODO: Update this with the actual description.`,
		Run: func(cmd *cobra.Command, args []string) {
			var ms = markerSettings{
				ID: msId,
			}
			var p = Payload{
				Method:   http.MethodDelete,
				Path:     "/1/marker_settings/" + targetDataset + "/" + ms.ID,
				Response: nil,
			}

			var err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersSettingsDeleteCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to delete a marker setting.")
			}
		},
	}

	cmd.Flags().StringVarP(&msId, "id", "i", "", "The ID of the marker to delete")
	cmd.MarkFlagRequired("id")

	return cmd
}
