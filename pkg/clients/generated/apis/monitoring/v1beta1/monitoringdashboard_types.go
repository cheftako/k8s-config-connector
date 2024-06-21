// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DashboardAggregation struct {
	/* The `alignment_period` specifies a time interval, in seconds, that is used
	to divide the data in all the
	[time series][google.monitoring.v3.TimeSeries] into consistent blocks of
	time. This will be done before the per-series aligner can be applied to
	the data.

	The value must be at least 60 seconds. If a per-series aligner other than
	`ALIGN_NONE` is specified, this field is required or an error is returned.
	If no per-series aligner is specified, or the aligner `ALIGN_NONE` is
	specified, then this field is ignored.

	The maximum value of the `alignment_period` is 2 years, or 104 weeks. */
	// +optional
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`

	/* The reduction operation to be used to combine time series into a single
	time series, where the value of each data point in the resulting series is
	a function of all the already aligned values in the input time series.

	Not all reducer operations can be applied to all time series. The valid
	choices depend on the `metric_kind` and the `value_type` of the original
	time series. Reduction can yield a time series with a different
	`metric_kind` or `value_type` than the input time series.

	Time series data must first be aligned (see `per_series_aligner`) in order
	to perform cross-time series reduction. If `cross_series_reducer` is
	specified, then `per_series_aligner` must be specified, and must not be
	`ALIGN_NONE`. An `alignment_period` must also be specified; otherwise, an
	error is returned. */
	// +optional
	CrossSeriesReducer *string `json:"crossSeriesReducer,omitempty"`

	/* The set of fields to preserve when `cross_series_reducer` is specified. The `group_by_fields` determine how the time series are partitioned into subsets prior to applying the aggregation operation. Each subset contains time series that have the same value for each of the grouping fields. Each individual time series is a member of exactly one subset. The `cross_series_reducer` is applied to each subset of time series. It is not possible to reduce across different resource types, so this field implicitly contains `resource.type`.  Fields not specified in `group_by_fields` are aggregated away.  If `group_by_fields` is not specified and all the time series have the same resource type, then the time series are aggregated into a single output time series. If `cross_series_reducer` is not defined, this field is ignored. */
	// +optional
	GroupByFields []string `json:"groupByFields,omitempty"`

	/* An `Aligner` describes how to bring the data points in a single
	time series into temporal alignment. Except for `ALIGN_NONE`, all
	alignments cause all the data points in an `alignment_period` to be
	mathematically grouped together, resulting in a single data point for
	each `alignment_period` with end timestamp at the end of the period.

	Not all alignment operations may be applied to all time series. The valid
	choices depend on the `metric_kind` and `value_type` of the original time
	series. Alignment can change the `metric_kind` or the `value_type` of
	the time series.

	Time series data must be aligned in order to perform cross-time
	series reduction. If `cross_series_reducer` is specified, then
	`per_series_aligner` must be specified and not equal to `ALIGN_NONE`
	and `alignment_period` must be specified; otherwise, an error is
	returned. */
	// +optional
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
}

type DashboardAlertChart struct {
	/* Required. A reference to the MonitoringAlertPolicy. */
	AlertPolicyRef v1alpha1.ResourceRef `json:"alertPolicyRef"`
}

type DashboardBlank struct {
}

type DashboardChartOptions struct {
	/* The chart mode. */
	// +optional
	Mode *string `json:"mode,omitempty"`
}

type DashboardCollapsibleGroup struct {
	/* The collapsed state of the widget on first page load. */
	// +optional
	Collapsed *bool `json:"collapsed,omitempty"`
}

type DashboardColumnLayout struct {
	/* The columns of content to display. */
	// +optional
	Columns []DashboardColumns `json:"columns,omitempty"`
}

type DashboardColumns struct {
	/* The relative weight of this column. The column weight is used to adjust the width of columns on the screen (relative to peers). Greater the weight, greater the width of the column on the screen. If omitted, a value of 1 is used while rendering. */
	// +optional
	Weight *int64 `json:"weight,omitempty"`

	/* The display widgets arranged vertically in this column. */
	// +optional
	Widgets []DashboardWidgets `json:"widgets,omitempty"`
}

type DashboardDataSets struct {
	/* A template string for naming `TimeSeries` in the resulting data set. This should be a string with interpolations of the form `${label_name}`, which will resolve to the label's value. */
	// +optional
	LegendTemplate *string `json:"legendTemplate,omitempty"`

	/* Optional. The lower bound on data point frequency for this data set, implemented by specifying the minimum alignment period to use in a time series query For example, if the data is published once every 10 minutes, the `min_alignment_period` should be at least 10 minutes. It would not make sense to fetch and align data at one minute intervals. */
	// +optional
	MinAlignmentPeriod *string `json:"minAlignmentPeriod,omitempty"`

	/* How this data should be plotted on the chart. */
	// +optional
	PlotType *string `json:"plotType,omitempty"`

	/* Required. Fields for querying time series data from the Stackdriver metrics API. */
	TimeSeriesQuery DashboardTimeSeriesQuery `json:"timeSeriesQuery"`
}

type DashboardDenominator struct {
	/* By default, the raw time series data is returned. Use this field to combine multiple time series for different views of the data. */
	// +optional
	Aggregation *DashboardAggregation `json:"aggregation,omitempty"`

	/* Required. The [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) that identifies the metric types, resources, and projects to query. */
	Filter string `json:"filter"`
}

type DashboardGaugeView struct {
	/* The lower bound for this gauge chart. The value of the chart should always be greater than or equal to this. */
	// +optional
	LowerBound *float64 `json:"lowerBound,omitempty"`

	/* The upper bound for this gauge chart. The value of the chart should always be less than or equal to this. */
	// +optional
	UpperBound *float64 `json:"upperBound,omitempty"`
}

type DashboardGridLayout struct {
	/* The number of columns into which the view's width is divided. If omitted or set to zero, a system default will be used while rendering. */
	// +optional
	Columns *DashboardColumns `json:"columns,omitempty"`

	/* The informational elements that are arranged into the columns row-first. */
	// +optional
	Widgets []DashboardWidgets `json:"widgets,omitempty"`
}

type DashboardLogsPanel struct {
	/* A filter that chooses which log entries to return.  See [Advanced Logs Queries](https://cloud.google.com/logging/docs/view/advanced-queries). Only log entries that match the filter are returned.  An empty filter matches all log entries. */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* The names of logging resources to collect logs for. Currently only projects are supported. If empty, the widget will default to the host project. */
	// +optional
	ResourceNames []DashboardResourceNames `json:"resourceNames,omitempty"`
}

type DashboardMosaicLayout struct {
	/* The number of columns in the mosaic grid. The number of columns must be between 1 and 12, inclusive. */
	// +optional
	Columns *DashboardColumns `json:"columns,omitempty"`

	/* The tiles to display. */
	// +optional
	Tiles []DashboardTiles `json:"tiles,omitempty"`
}

type DashboardNumerator struct {
	/* By default, the raw time series data is returned. Use this field to combine multiple time series for different views of the data. */
	// +optional
	Aggregation *DashboardAggregation `json:"aggregation,omitempty"`

	/* Required. The [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) that identifies the metric types, resources, and projects to query. */
	Filter string `json:"filter"`
}

type DashboardPickTimeSeriesFilter struct {
	/* How to use the ranking to select time series that pass through the filter. */
	// +optional
	Direction *string `json:"direction,omitempty"`

	/* How many time series to allow to pass through the filter. */
	// +optional
	NumTimeSeries *int32 `json:"numTimeSeries,omitempty"`

	/* `ranking_method` is applied to each time series independently to produce the value which will be used to compare the time series to other time series. */
	// +optional
	RankingMethod *string `json:"rankingMethod,omitempty"`
}

type DashboardResourceNames struct {
	/* The external name of the referenced resource */
	// +optional
	External *string `json:"external,omitempty"`

	/* Kind of the referent. */
	// +optional
	Kind *string `json:"kind,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	// +optional
	Name *string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	// +optional
	Namespace *string `json:"namespace,omitempty"`
}

type DashboardRowLayout struct {
	/* The rows of content to display. */
	// +optional
	Rows []DashboardRows `json:"rows,omitempty"`
}

type DashboardRows struct {
	/* The relative weight of this row. The row weight is used to adjust the height of rows on the screen (relative to peers). Greater the weight, greater the height of the row on the screen. If omitted, a value of 1 is used while rendering. */
	// +optional
	Weight *int64 `json:"weight,omitempty"`

	/* The display widgets arranged horizontally in this row. */
	// +optional
	Widgets []DashboardWidgets `json:"widgets,omitempty"`
}

type DashboardScorecard struct {
	/* Will cause the scorecard to show a gauge chart. */
	// +optional
	GaugeView *DashboardGaugeView `json:"gaugeView,omitempty"`

	/* Will cause the scorecard to show a spark chart. */
	// +optional
	SparkChartView *DashboardSparkChartView `json:"sparkChartView,omitempty"`

	/* The thresholds used to determine the state of the scorecard given the
	time series' current value. For an actual value x, the scorecard is in a
	danger state if x is less than or equal to a danger threshold that triggers
	below, or greater than or equal to a danger threshold that triggers above.
	Similarly, if x is above/below a warning threshold that triggers
	above/below, then the scorecard is in a warning state - unless x also puts
	it in a danger state. (Danger trumps warning.)

	As an example, consider a scorecard with the following four thresholds:

	```
	{
	value: 90,
	category: 'DANGER',
	trigger: 'ABOVE',
	},
	{
	value: 70,
	category: 'WARNING',
	trigger: 'ABOVE',
	},
	{
	value: 10,
	category: 'DANGER',
	trigger: 'BELOW',
	},
	{
	value: 20,
	category: 'WARNING',
	trigger: 'BELOW',
	}
	```

	Then: values less than or equal to 10 would put the scorecard in a DANGER
	state, values greater than 10 but less than or equal to 20 a WARNING state,
	values strictly between 20 and 70 an OK state, values greater than or equal
	to 70 but less than 90 a WARNING state, and values greater than or equal to
	90 a DANGER state. */
	// +optional
	Thresholds []DashboardThresholds `json:"thresholds,omitempty"`

	/* Required. Fields for querying time series data from the Stackdriver metrics API. */
	TimeSeriesQuery DashboardTimeSeriesQuery `json:"timeSeriesQuery"`
}

type DashboardSecondaryAggregation struct {
	/* The `alignment_period` specifies a time interval, in seconds, that is used
	to divide the data in all the
	[time series][google.monitoring.v3.TimeSeries] into consistent blocks of
	time. This will be done before the per-series aligner can be applied to
	the data.

	The value must be at least 60 seconds. If a per-series aligner other than
	`ALIGN_NONE` is specified, this field is required or an error is returned.
	If no per-series aligner is specified, or the aligner `ALIGN_NONE` is
	specified, then this field is ignored.

	The maximum value of the `alignment_period` is 2 years, or 104 weeks. */
	// +optional
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`

	/* The reduction operation to be used to combine time series into a single
	time series, where the value of each data point in the resulting series is
	a function of all the already aligned values in the input time series.

	Not all reducer operations can be applied to all time series. The valid
	choices depend on the `metric_kind` and the `value_type` of the original
	time series. Reduction can yield a time series with a different
	`metric_kind` or `value_type` than the input time series.

	Time series data must first be aligned (see `per_series_aligner`) in order
	to perform cross-time series reduction. If `cross_series_reducer` is
	specified, then `per_series_aligner` must be specified, and must not be
	`ALIGN_NONE`. An `alignment_period` must also be specified; otherwise, an
	error is returned. */
	// +optional
	CrossSeriesReducer *string `json:"crossSeriesReducer,omitempty"`

	/* The set of fields to preserve when `cross_series_reducer` is specified. The `group_by_fields` determine how the time series are partitioned into subsets prior to applying the aggregation operation. Each subset contains time series that have the same value for each of the grouping fields. Each individual time series is a member of exactly one subset. The `cross_series_reducer` is applied to each subset of time series. It is not possible to reduce across different resource types, so this field implicitly contains `resource.type`.  Fields not specified in `group_by_fields` are aggregated away.  If `group_by_fields` is not specified and all the time series have the same resource type, then the time series are aggregated into a single output time series. If `cross_series_reducer` is not defined, this field is ignored. */
	// +optional
	GroupByFields []string `json:"groupByFields,omitempty"`

	/* An `Aligner` describes how to bring the data points in a single
	time series into temporal alignment. Except for `ALIGN_NONE`, all
	alignments cause all the data points in an `alignment_period` to be
	mathematically grouped together, resulting in a single data point for
	each `alignment_period` with end timestamp at the end of the period.

	Not all alignment operations may be applied to all time series. The valid
	choices depend on the `metric_kind` and `value_type` of the original time
	series. Alignment can change the `metric_kind` or the `value_type` of
	the time series.

	Time series data must be aligned in order to perform cross-time
	series reduction. If `cross_series_reducer` is specified, then
	`per_series_aligner` must be specified and not equal to `ALIGN_NONE`
	and `alignment_period` must be specified; otherwise, an error is
	returned. */
	// +optional
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
}

type DashboardSectionHeader struct {
	/* Whether to insert a divider below the section in the table of contents */
	// +optional
	DividerBelow *bool `json:"dividerBelow,omitempty"`

	/* The subtitle of the section */
	// +optional
	Subtitle *string `json:"subtitle,omitempty"`
}

type DashboardSparkChartView struct {
	/* The lower bound on data point frequency in the chart implemented by specifying the minimum alignment period to use in a time series query. For example, if the data is published once every 10 minutes it would not make sense to fetch and align data at one minute intervals. This field is optional and exists only as a hint. */
	// +optional
	MinAlignmentPeriod *string `json:"minAlignmentPeriod,omitempty"`

	/* Required. The type of sparkchart to show in this chartView. */
	SparkChartType string `json:"sparkChartType"`
}

type DashboardStyle struct {
	/* The background color as a hex string. "#RRGGBB" or "#RGB" */
	// +optional
	BackgroundColor *string `json:"backgroundColor,omitempty"`

	/* Font sizes for both the title and content. The title will still be larger relative to the content. */
	// +optional
	FontSize *string `json:"fontSize,omitempty"`

	/* The horizontal alignment of both the title and content */
	// +optional
	HorizontalAlignment *string `json:"horizontalAlignment,omitempty"`

	/* The amount of padding around the widget */
	// +optional
	Padding *string `json:"padding,omitempty"`

	/* The pointer location for this widget (also sometimes called a "tail") */
	// +optional
	PointerLocation *string `json:"pointerLocation,omitempty"`

	/* The text color as a hex string. "#RRGGBB" or "#RGB" */
	// +optional
	TextColor *string `json:"textColor,omitempty"`

	/* The vertical alignment of both the title and content */
	// +optional
	VerticalAlignment *string `json:"verticalAlignment,omitempty"`
}

type DashboardText struct {
	/* The text content to be displayed. */
	// +optional
	Content *string `json:"content,omitempty"`

	/* How the text content is formatted. */
	// +optional
	Format *string `json:"format,omitempty"`

	/* How the text is styled */
	// +optional
	Style *DashboardStyle `json:"style,omitempty"`
}

type DashboardThresholds struct {
	/* The state color for this threshold. Color is not allowed in a XyChart. */
	// +optional
	Color *string `json:"color,omitempty"`

	/* The direction for the current threshold. Direction is not allowed in a XyChart. */
	// +optional
	Direction *string `json:"direction,omitempty"`

	/* A label for the threshold. */
	// +optional
	Label *string `json:"label,omitempty"`

	/* The value of the threshold. The value should be defined in the native scale of the metric. */
	// +optional
	Value *float64 `json:"value,omitempty"`
}

type DashboardTiles struct {
	/* The height of the tile, measured in grid blocks. Tiles must have a minimum height of 1. */
	// +optional
	Height *int32 `json:"height,omitempty"`

	/* The informational widget contained in the tile. For example an `XyChart`. */
	// +optional
	Widget *DashboardWidget `json:"widget,omitempty"`

	/* The width of the tile, measured in grid blocks. Tiles must have a minimum width of 1. */
	// +optional
	Width *int32 `json:"width,omitempty"`

	/* The zero-indexed position of the tile in grid blocks relative to the left edge of the grid. Tiles must be contained within the specified number of columns. `x_pos` cannot be negative. */
	// +optional
	XPos *int32 `json:"xPos,omitempty"`

	/* The zero-indexed position of the tile in grid blocks relative to the top edge of the grid. `y_pos` cannot be negative. */
	// +optional
	YPos *int32 `json:"yPos,omitempty"`
}

type DashboardTimeSeriesFilter struct {
	/* By default, the raw time series data is returned. Use this field to combine multiple time series for different views of the data. */
	// +optional
	Aggregation *DashboardAggregation `json:"aggregation,omitempty"`

	/* Required. The [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) that identifies the metric types, resources, and projects to query. */
	Filter string `json:"filter"`

	/* Ranking based time series filter. */
	// +optional
	PickTimeSeriesFilter *DashboardPickTimeSeriesFilter `json:"pickTimeSeriesFilter,omitempty"`

	/* Apply a second aggregation after `aggregation` is applied. */
	// +optional
	SecondaryAggregation *DashboardSecondaryAggregation `json:"secondaryAggregation,omitempty"`
}

type DashboardTimeSeriesFilterRatio struct {
	/* The denominator of the ratio. */
	// +optional
	Denominator *DashboardDenominator `json:"denominator,omitempty"`

	/* The numerator of the ratio. */
	// +optional
	Numerator *DashboardNumerator `json:"numerator,omitempty"`

	/* Ranking based time series filter. */
	// +optional
	PickTimeSeriesFilter *DashboardPickTimeSeriesFilter `json:"pickTimeSeriesFilter,omitempty"`

	/* Apply a second aggregation after the ratio is computed. */
	// +optional
	SecondaryAggregation *DashboardSecondaryAggregation `json:"secondaryAggregation,omitempty"`
}

type DashboardTimeSeriesQuery struct {
	/* Filter parameters to fetch time series. */
	// +optional
	TimeSeriesFilter *DashboardTimeSeriesFilter `json:"timeSeriesFilter,omitempty"`

	/* Parameters to fetch a ratio between two time series filters. */
	// +optional
	TimeSeriesFilterRatio *DashboardTimeSeriesFilterRatio `json:"timeSeriesFilterRatio,omitempty"`

	/* A query used to fetch time series with MQL. */
	// +optional
	TimeSeriesQueryLanguage *string `json:"timeSeriesQueryLanguage,omitempty"`

	/* The unit of data contained in fetched time series. If non-empty, this unit will override any unit that accompanies fetched data. The format is the same as the [`unit`](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors) field in `MetricDescriptor`. */
	// +optional
	UnitOverride *string `json:"unitOverride,omitempty"`
}

type DashboardWidget struct {
	/* A chart of alert policy data. */
	// +optional
	AlertChart *DashboardAlertChart `json:"alertChart,omitempty"`

	/* A blank space. */
	// +optional
	Blank *DashboardBlank `json:"blank,omitempty"`

	/* A widget that groups the other widgets. All widgets that are within the area spanned by the grouping widget are considered member widgets. */
	// +optional
	CollapsibleGroup *DashboardCollapsibleGroup `json:"collapsibleGroup,omitempty"`

	/* A widget that shows a stream of logs. */
	// +optional
	LogsPanel *DashboardLogsPanel `json:"logsPanel,omitempty"`

	/* A scorecard summarizing time series data. */
	// +optional
	Scorecard *DashboardScorecard `json:"scorecard,omitempty"`

	/* A widget that defines a section header for easier navigation of the dashboard. */
	// +optional
	SectionHeader *DashboardSectionHeader `json:"sectionHeader,omitempty"`

	/* A raw string or markdown displaying textual content. */
	// +optional
	Text *DashboardText `json:"text,omitempty"`

	/* Optional. The title of the widget. */
	// +optional
	Title *string `json:"title,omitempty"`

	/* A chart of time series data. */
	// +optional
	XyChart *DashboardXyChart `json:"xyChart,omitempty"`
}

type DashboardWidgets struct {
	/* A chart of alert policy data. */
	// +optional
	AlertChart *DashboardAlertChart `json:"alertChart,omitempty"`

	/* A blank space. */
	// +optional
	Blank *DashboardBlank `json:"blank,omitempty"`

	/* A widget that groups the other widgets. All widgets that are within the area spanned by the grouping widget are considered member widgets. */
	// +optional
	CollapsibleGroup *DashboardCollapsibleGroup `json:"collapsibleGroup,omitempty"`

	/* A widget that shows a stream of logs. */
	// +optional
	LogsPanel *DashboardLogsPanel `json:"logsPanel,omitempty"`

	/* A scorecard summarizing time series data. */
	// +optional
	Scorecard *DashboardScorecard `json:"scorecard,omitempty"`

	/* A widget that defines a section header for easier navigation of the dashboard. */
	// +optional
	SectionHeader *DashboardSectionHeader `json:"sectionHeader,omitempty"`

	/* A raw string or markdown displaying textual content. */
	// +optional
	Text *DashboardText `json:"text,omitempty"`

	/* Optional. The title of the widget. */
	// +optional
	Title *string `json:"title,omitempty"`

	/* A chart of time series data. */
	// +optional
	XyChart *DashboardXyChart `json:"xyChart,omitempty"`
}

type DashboardXAxis struct {
	/* The label of the axis. */
	// +optional
	Label *string `json:"label,omitempty"`

	/* The axis scale. By default, a linear scale is used. */
	// +optional
	Scale *string `json:"scale,omitempty"`
}

type DashboardXyChart struct {
	/* Display options for the chart. */
	// +optional
	ChartOptions *DashboardChartOptions `json:"chartOptions,omitempty"`

	/* Required. The data displayed in this chart. */
	DataSets []DashboardDataSets `json:"dataSets"`

	/* Threshold lines drawn horizontally across the chart. */
	// +optional
	Thresholds []DashboardThresholds `json:"thresholds,omitempty"`

	/* The duration used to display a comparison chart. A comparison chart simultaneously shows values from two similar-length time periods (e.g., week-over-week metrics). The duration must be positive, and it can only be applied to charts with data sets of LINE plot type. */
	// +optional
	TimeshiftDuration *string `json:"timeshiftDuration,omitempty"`

	/* The properties applied to the x-axis. */
	// +optional
	XAxis *DashboardXAxis `json:"xAxis,omitempty"`

	/* The properties applied to the y-axis. */
	// +optional
	YAxis *DashboardYAxis `json:"yAxis,omitempty"`
}

type DashboardYAxis struct {
	/* The label of the axis. */
	// +optional
	Label *string `json:"label,omitempty"`

	/* The axis scale. By default, a linear scale is used. */
	// +optional
	Scale *string `json:"scale,omitempty"`
}

type MonitoringDashboardSpec struct {
	/* The content is divided into equally spaced columns and the widgets are arranged vertically. */
	// +optional
	ColumnLayout *DashboardColumnLayout `json:"columnLayout,omitempty"`

	/* Required. The mutable, human-readable name. */
	DisplayName string `json:"displayName"`

	/* Content is arranged with a basic layout that re-flows a simple list of informational elements like widgets or tiles. */
	// +optional
	GridLayout *DashboardGridLayout `json:"gridLayout,omitempty"`

	/* The content is arranged as a grid of tiles, with each content widget occupying one or more grid blocks. */
	// +optional
	MosaicLayout *DashboardMosaicLayout `json:"mosaicLayout,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The content is divided into equally spaced rows and the widgets are arranged horizontally. */
	// +optional
	RowLayout *DashboardRowLayout `json:"rowLayout,omitempty"`
}

type MonitoringDashboardStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringDashboard's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* \`etag\` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. An \`etag\` is returned in the response to \`GetDashboard\`, and users are expected to put that etag in the request to \`UpdateDashboard\` to ensure that their change will be applied to the same version of the Dashboard configuration. The field should not be passed during dashboard creation. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringdashboard;gcpmonitoringdashboards
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringDashboard is the Schema for the monitoring API
// +k8s:openapi-gen=true
type MonitoringDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringDashboardSpec   `json:"spec,omitempty"`
	Status MonitoringDashboardStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringDashboardList contains a list of MonitoringDashboard
type MonitoringDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringDashboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringDashboard{}, &MonitoringDashboardList{})
}
