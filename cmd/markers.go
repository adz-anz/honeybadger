package cmd

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

type marker struct {
	// Indicates the time the Marker should be placed. If missing, defaults to
	// the time the request arrives. Expressed in Unix Time.
	StartTime int64 `json:"start_time,omitempty"`

	// Specifies end time, and allows a Marker to be recorded as representing a
	// time range, such as a 5 minute deploy. Expressed in Unix Time.
	EndTime int64 `json:"end_time,omitempty"`

	// A message to describe this specific Marker.
	Message string `json:"message,omitempty"`

	// Groups similar Markers. For example, 'deploys'. All Markers of the same
	// type appear with the same color on the graph.
	Type string `json:"type,omitempty"`

	// A target for the marker. Clicking the marker text will take you to this
	// URL.
	URL string `json:"url,omitempty"`

	// The unique identifier (ID) of a Marker.
	ID string `json:"id,omitempty"`

	// Color can be assigned to Markers using the Marker Settings endpoint. This
	// field will be populated when List All Markers is called.
	Color string `json:"color,omitempty"`

	// The ISO8601-formatted time when the Marker was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The ISO8601-formatted time when the Marker was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func newMarkersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "markers",
		Aliases: []string{"m"},
		Short:   "Manage Markers",
		Long: "Markers indicate points in time on graphs where interesting things happen,\n" +
			"such as deploys or outages.\n" +
			"\n" +
			"These commands allow you to list, create, update, and delete Markers.",
	}

	cmd.PersistentFlags().StringVarP(&targetDataset, "dataset", "d", "__all__",
		"The dataset slug or use __all__ (or omit) for endpoints that support environment-wide operations.")

	cmd.AddCommand(
		newMarkersCreateCmd(),
		newMarkersListCmd(),
		newMarkersUpdateCmd(),
		newMarkersDeleteCmd(),
	)

	return cmd
}

// Create a Marker
// https://docs.honeycomb.io/api/tag/Markers#operation/createMarker
func newMarkersCreateCmd() *cobra.Command {
	var (
		mStartTime int64
		mEndTime   int64
		mMsg       string
		mType      string
		mUrl       string
	)
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"add", "new"},
		Short:   "Create a Marker in the specified dataset.",
		Long: "Create a Marker in the specified dataset. To create an environment marker,\n" +
			"use the __all__ dataset (or omit the dataset) and an API key associated with\n" +
			"the desired environment.",
		Run: func(cmd *cobra.Command, args []string) {
			var m = marker{
				StartTime: mStartTime,
				EndTime:   mEndTime,
				Message:   mMsg,
				Type:      mType,
				URL:       mUrl,
			}
			var bodyMarshal, err = json.Marshal(m)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersCreateCmd",
					"err":       err,
					"marker":    m,
				}).Fatal("Error received when attempting to marshal a marker.")
			}
			var p = Payload{
				Method:   http.MethodPost,
				Path:     "/1/markers/" + targetDataset,
				Body:     bodyMarshal,
				Response: &marker{},
			}

			err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersCreateCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to create a new marker.")
			}
		},
	}

	cmd.Flags().Int64VarP(&mStartTime, "start_time", "s", 0,
		"Indicates the time the Marker should be placed. If missing, defaults to the time the request arrives. Expressed in Unix Time.")
	cmd.Flags().Int64VarP(&mEndTime, "end_time", "e", 0,
		"Specifies end time, and allows a Marker to be recorded as representing a time range, such as a 5 minute deploy. Expressed in Unix Time.")
	cmd.Flags().StringVarP(&mMsg, "msg", "m", "", "A message to describe this specific Marker.")
	cmd.Flags().StringVarP(&mType, "type", "t", "",
		"Groups similar Markers. For example, 'deploys'. All Markers of the same type appear with the same color on the graph.")
	cmd.Flags().StringVarP(&mUrl, "url", "u", "",
		"A target for the marker. Clicking the marker text will take you to this URL.")

	return cmd
}

// List All Markers
// https://docs.honeycomb.io/api/tag/Markers#operation/getMarker
func newMarkersListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "get"},
		Short:   "Lists all Markers for a dataset.",
		Long: "Lists all Markers for a dataset. To list environment markers, use the __all__\n" +
			"dataset (or omit the dataset) and an API key associated with the desired environment.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = Payload{
				Method:   http.MethodGet,
				Path:     "/1/markers/" + targetDataset,
				Response: &[]marker{},
			}

			var err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersListCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to list all markers.")
			}
		},
	}

	return cmd
}

// Update a Marker
// https://docs.honeycomb.io/api/tag/Markers#operation/updateMarker
func newMarkersUpdateCmd() *cobra.Command {
	var (
		mId        string
		mStartTime int64
		mEndTime   int64
		mMsg       string
		mType      string
		mUrl       string
	)
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"up", "edit", "modify", "change", "set"},
		Short:   "Update a Marker in the specified dataset.",
		Long: "Update a Marker in the specified dataset. To update an environment marker, use the\n" +
			"_all__ dataset (or omit the dataset) and an API key associated with the desired environment.",
		Run: func(cmd *cobra.Command, args []string) {
			var m = marker{
				ID:        mId,
				StartTime: mStartTime,
				EndTime:   mEndTime,
				Message:   mMsg,
				Type:      mType,
				URL:       mUrl,
			}
			var bodyMarshal, err = json.Marshal(m)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersUpdateCmd",
					"err":       err,
					"marker":    m,
				}).Fatal("Error received when attempting to marshal a marker.")
			}
			var p = Payload{
				Method:   http.MethodPut,
				Path:     "/1/markers/" + targetDataset + "/" + m.ID,
				Body:     bodyMarshal,
				Response: &marker{},
			}

			err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersUpdateCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to update an existing marker.")
			}
		},
	}

	cmd.Flags().StringVarP(&mId, "id", "i", "", "The unique identifier (ID) of a Marker.")
	cmd.MarkFlagRequired("id")
	cmd.Flags().Int64VarP(&mStartTime, "start_time", "s", 0,
		"Indicates the time the Marker should be placed. If missing, defaults to the time the request arrives. Expressed in Unix Time.")
	cmd.Flags().Int64VarP(&mEndTime, "end_time", "e", 0,
		"Specifies end time, and allows a Marker to be recorded as representing a time range, such as a 5 minute deploy. Expressed in Unix Time.")
	cmd.Flags().StringVarP(&mMsg, "msg", "m", "", "A message to describe this specific Marker.")
	cmd.Flags().StringVarP(&mType, "type", "t", "",
		"Groups similar Markers. For example, 'deploys'. All Markers of the same type appear with the same color on the graph.")
	cmd.Flags().StringVarP(&mUrl, "url", "u", "",
		"A target for the marker. Clicking the marker text will take you to this URL.")

	return cmd
}

// Delete a Marker
// https://docs.honeycomb.io/api/tag/Markers#operation/deleteMarker
func newMarkersDeleteCmd() *cobra.Command {
	var (
		mId string
	)

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm", "remove", "del"},
		Short:   "Delete a Marker in the specified dataset.",
		Long: "Delete a Marker in the specified dataset. To delete an environment marker, use the __all__\n" +
			"dataset (or omit the dataset) and an API key associated with the desired environment.",
		Run: func(cmd *cobra.Command, args []string) {
			var m = marker{
				ID: mId,
			}
			var p = Payload{
				Method:   http.MethodDelete,
				Path:     "/1/markers/" + targetDataset + "/" + m.ID,
				Response: &marker{},
			}

			var err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newMarkersDeleteCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to delete an existing marker.")
			}
		},
	}

	cmd.Flags().StringVarP(&mId, "id", "i", "", "The unique identifier (ID) of a Marker.")
	cmd.MarkFlagRequired("id")

	return cmd
}
