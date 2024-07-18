package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// ----- QUERY STRUCTS -----

type queryCalculation struct {
	// Enum:
	//		"COUNT" "CONCURRENCY" "SUM" "AVG" "COUNT_DISTINCT" "HEATMAP" "MAX"
	//		"MIN" "P001" "P01" "P05" "P10" "P25" "P50" "P75" "P90" "P95" "P99"
	//		"P999" "RATE_AVG" "RATE_SUM" "RATE_MAX"
	Op string `json:"op,omitempty"`

	// The name of the column.
	Column string `json:"column,omitempty"`
}

type queryFilter struct {
	// Enum:
	//		"=" "!=" ">" ">=" "<" "<=" "starts-with" "does-not-start-with"
	//		"exists" "does-not-exist" "contains" "does-not-contain" "in"
	//		"not-in"
	Op string `json:"op,omitempty"`

	// The name of the column.
	Column string `json:"column,omitempty"`

	// The value to compare against.
	Value interface{} `json:"value,omitempty"`
}

type queryOrder struct {
	// The name of the column.
	Column string `json:"column,omitempty"`

	// Enum:
	//		"COUNT" "CONCURRENCY" "SUM" "AVG" "COUNT_DISTINCT" "HEATMAP" "MAX"
	//		"MIN" "P001" "P01" "P05" "P10" "P25" "P50" "P75" "P90" "P95" "P99"
	//		"P999" "RATE_AVG" "RATE_SUM" "RATE_MAX"
	Op string `json:"op,omitempty"`

	Order string `json:"order,omitempty"`
}

type queryHaving struct {
	// Enum:
	//		"COUNT" "CONCURRENCY" "SUM" "AVG" "COUNT_DISTINCT" "HEATMAP" "MAX"
	//		"MIN" "P001" "P01" "P05" "P10" "P25" "P50" "P75" "P90" "P95" "P99"
	//		"P999" "RATE_AVG" "RATE_SUM" "RATE_MAX"
	CalculateOp string `json:"calculate_op,omitempty"`

	// The name of the column to filter against.
	Column string `json:"column,omitempty"`

	// Enum:
	//		 "=" "!=" ">" ">=" "<" "<="
	Op string `json:"op,omitempty"`

	Value int `json:"value,omitempty"`
}

type query struct {
	// The ID of a query returned from the Queries endpoint.
	ID string `json:"id,omitempty"`

	// The columns by which to break events down into groups.
	Breakdowns []string `json:"breakdowns,omitempty"`

	// The calculations to return as a time series and summary table.
	Calculations []queryCalculation `json:"calculations,omitempty"`

	// The filters with which to restrict the considered events.
	Filters []queryFilter `json:"filters,omitempty"`

	// set to "OR" to match ANY filter in the filter list.
	FilterCombination string `json:"filter_combination,omitempty"`

	// The time resolution of the query's graph, in seconds. Given a query time
	// range T, valid values (T/1000...T/10).
	Granularity int `json:"granularity,omitempty"`

	// The terms on which to order the query results. Each term must appear in
	// either the breakdowns field or the calculations field.
	Orders []queryOrder `json:"orders,omitempty"`

	// The maximum number of unique groups returned in 'results'. Aggregating
	// many unique groups across a large time range is computationally
	// expensive, and too high a limit with too many unique groups may cause
	// queries to fail completey. Limiting the results to only the needed
	// values can significantly speed up queries. The normal allowed maximum
	// value when creating a query is 1_000. When running 'disable_series'
	// queries, this can be overridden to be up to 10_000, so the maximum value
	// returned from the API when fetching a query may be up to 10_000.
	Limit int `json:"limit,omitempty"`

	// Absolute start time of query, in seconds since UNIX epoch.
	// Must be <= end_time.
	StartTime int `json:"start_time,omitempty"`

	// Absolute end time of query, in seconds since UNIX epoch.
	EndTime int `json:"end_time,omitempty"`

	// Time range of query in seconds. Can be used with either start_time
	// (seconds after start_time), end_time (seconds before end_time), or
	// without either (seconds before now).
	TimeRange int `json:"time_range,omitempty"`

	// The Having clause allows you to filter on the results table. This
	// operation is distinct from the Where clause, which filters the
	// underlying events. Order By allows you to order the results, and Having
	// filters them.
	Havings []queryHaving `json:"havings,omitempty"`
}

// ----- QUERY RESULT STRUCTS -----

type queryResultCreateRequest struct {
	// The ID of a query returned from the Queries endpoint.
	QueryID string `json:"query_id,omitempty"`

	// If true, will disable calculation and return of the full time-series
	// data, usually included in the 'series' response field, instead only
	// returning the summarized 'results'.
	DisableSeries bool `json:"disable_series,omitempty"`

	// If 'disable_series' is true, then a limit may optionally be given, which
	// will override the limit normally associated with the query. Unlike normal
	// query results which are limited to 1_000 results, 'disable_series'
	// results may have a limit of up to 10_000. If 'disable_series' is false,
	// then this field will be ignored.
	Limit int `json:"limit,omitempty"`
}

type queryResultLinks struct {
	QueryURL string `json:"query_url,omitempty"`

	GraphImageURL string `json:"graph_image_url,omitempty"`
}

type queryResultSeries struct {
	Time string `json:"time,omitempty"`

	Data map[string]interface{} `json:"data,omitempty"`
}

type queryResultResults struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

type queryResultData struct {
	Series []queryResultSeries `json:"series,omitempty"`

	Results []queryResultResults `json:"results,omitempty"`
}

type queryResult struct {
	// The query that this result is for.
	Query query `json:"query,omitempty"`

	// The unique identifier (ID) of a Query Result
	ID string `json:"id,omitempty"`

	// Indicates if the query results are available yet or not. For example, is
	// the query still being processed or complete?
	Complete bool `json:"complete"`

	// Data
	Data queryResultData `json:"data,omitempty"`

	// Links
	Links queryResultLinks `json:"links,omitempty"`
}

func newQueryResultCreateCmd() *cobra.Command {
	var (
		queryID string
	)

	cmd := &cobra.Command{
		Use:     "create-query-result",
		Aliases: []string{"cqr"},
		Short:   "Kick off processing of a Query to then get back the Query Results.",
		Long: "Kick off processing of a Query to then get back the Query Results.\n" +
			"\n" +
			"Once the Query Result has been created, the query will be run asynchronously, allowing the result data to be fetched from the GET query result endpoint.\n" +
			"\n" +
			"Only the last 7 days of data can be queried. Any queries with a `start_time`, `end_time`, or `time_range` older than last 7 days will result in a `400` error response.",
		Run: func(cmd *cobra.Command, args []string) {
			var qr = queryResultCreateRequest{
				QueryID: queryID,
			}

			var bodyMarshal, err = json.Marshal(qr)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newQueryResultCreateCmd",
					"err":       err,
					"query":     qr,
				}).Fatal("Error received when attempting to marshal a query result create request.")
			}

			var p = payload{
				Method:   http.MethodPost,
				Path:     "/1/query_results/" + targetDataset,
				Body:     bodyMarshal,
				Response: &queryResult{},
			}

			err = p.GetResponse(true)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newQueryResultCreateCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to create a query result.")
			}
		},
	}

	cmd.Flags().StringVarP(&queryID, "query-id", "q", "",
		"The ID of a query returned from the Queries endpoint.")
	cmd.MarkFlagRequired("query-id")

	return cmd
}

func newQueryResultGetCmd() *cobra.Command {
	var (
		queryResultID string
	)

	cmd := &cobra.Command{
		Use:     "get-query-result",
		Aliases: []string{"gqr"},
		Short:   "Get the Query Result details for a specific Query Result ID.",
		Long: "Get the Query Result details for a specific Query Result ID.\n" +
			"\n" +
			"This endpoint is used to fetch the results of a query that had previously been created. It is recommended to follow the Location header included in the Create Query Result output, but the URL can also be constructed manually with the <query-result-id>.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = payload{
				Method:   http.MethodGet,
				Path:     "/1/query_results/" + targetDataset + "/" + queryResultID,
				Response: &queryResult{},
			}

			err := p.GetResponse(true)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newQueryResultGetCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to get a query result.")
			}
		},
	}

	cmd.Flags().StringVarP(&queryResultID, "query-result-id", "q", "",
		"The unique identifier (ID) of the query result.")
	cmd.MarkFlagRequired("query-result-id")

	return cmd
}
