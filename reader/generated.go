// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package reader

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// __getInputInput is used internally by genqlient
type __getInputInput struct {
	InputIndex int `json:"inputIndex"`
}

// GetInputIndex returns __getInputInput.InputIndex, and is useful for accessing the field via an interface.
func (v *__getInputInput) GetInputIndex() int { return v.InputIndex }

// __getLastReportsInput is used internally by genqlient
type __getLastReportsInput struct {
	Last int `json:"last"`
}

// GetLast returns __getLastReportsInput.Last, and is useful for accessing the field via an interface.
func (v *__getLastReportsInput) GetLast() int { return v.Last }

// __getNoticeInput is used internally by genqlient
type __getNoticeInput struct {
	InputIndex  int `json:"inputIndex"`
	NoticeIndex int `json:"noticeIndex"`
}

// GetInputIndex returns __getNoticeInput.InputIndex, and is useful for accessing the field via an interface.
func (v *__getNoticeInput) GetInputIndex() int { return v.InputIndex }

// GetNoticeIndex returns __getNoticeInput.NoticeIndex, and is useful for accessing the field via an interface.
func (v *__getNoticeInput) GetNoticeIndex() int { return v.NoticeIndex }

// __getReportInput is used internally by genqlient
type __getReportInput struct {
	InputIndex  int `json:"inputIndex"`
	ReportIndex int `json:"reportIndex"`
}

// GetInputIndex returns __getReportInput.InputIndex, and is useful for accessing the field via an interface.
func (v *__getReportInput) GetInputIndex() int { return v.InputIndex }

// GetReportIndex returns __getReportInput.ReportIndex, and is useful for accessing the field via an interface.
func (v *__getReportInput) GetReportIndex() int { return v.ReportIndex }

// getInputInput includes the requested fields of the GraphQL type Input.
// The GraphQL type's documentation follows.
//
// Request submitted to the application to advance its state
type getInputInput struct {
	// Number of the base layer block in which the input was recorded
	BlockNumber string `json:"blockNumber"`
	// Get reports from this particular input with support for pagination
	Reports getInputInputReportsReportConnection `json:"reports"`
	// Get notices from this particular input with support for pagination
	Notices getInputInputNoticesNoticeConnection `json:"notices"`
}

// GetBlockNumber returns getInputInput.BlockNumber, and is useful for accessing the field via an interface.
func (v *getInputInput) GetBlockNumber() string { return v.BlockNumber }

// GetReports returns getInputInput.Reports, and is useful for accessing the field via an interface.
func (v *getInputInput) GetReports() getInputInputReportsReportConnection { return v.Reports }

// GetNotices returns getInputInput.Notices, and is useful for accessing the field via an interface.
func (v *getInputInput) GetNotices() getInputInputNoticesNoticeConnection { return v.Notices }

// getInputInputNoticesNoticeConnection includes the requested fields of the GraphQL type NoticeConnection.
// The GraphQL type's documentation follows.
//
// Pagination result
type getInputInputNoticesNoticeConnection struct {
	// Total number of entries that match the query
	TotalCount int `json:"totalCount"`
}

// GetTotalCount returns getInputInputNoticesNoticeConnection.TotalCount, and is useful for accessing the field via an interface.
func (v *getInputInputNoticesNoticeConnection) GetTotalCount() int { return v.TotalCount }

// getInputInputReportsReportConnection includes the requested fields of the GraphQL type ReportConnection.
// The GraphQL type's documentation follows.
//
// Pagination result
type getInputInputReportsReportConnection struct {
	// Total number of entries that match the query
	TotalCount int `json:"totalCount"`
}

// GetTotalCount returns getInputInputReportsReportConnection.TotalCount, and is useful for accessing the field via an interface.
func (v *getInputInputReportsReportConnection) GetTotalCount() int { return v.TotalCount }

// getInputResponse is returned by getInput on success.
type getInputResponse struct {
	// Get input based on its identifier
	Input getInputInput `json:"input"`
}

// GetInput returns getInputResponse.Input, and is useful for accessing the field via an interface.
func (v *getInputResponse) GetInput() getInputInput { return v.Input }

// getLastReportsReportsReportConnection includes the requested fields of the GraphQL type ReportConnection.
// The GraphQL type's documentation follows.
//
// Pagination result
type getLastReportsReportsReportConnection struct {
	// Total number of entries that match the query
	TotalCount int `json:"totalCount"`
	// Pagination metadata
	PageInfo getLastReportsReportsReportConnectionPageInfo `json:"pageInfo"`
	// Pagination entries returned for the current page
	Edges []getLastReportsReportsReportConnectionEdgesReportEdge `json:"edges"`
}

// GetTotalCount returns getLastReportsReportsReportConnection.TotalCount, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnection) GetTotalCount() int { return v.TotalCount }

// GetPageInfo returns getLastReportsReportsReportConnection.PageInfo, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnection) GetPageInfo() getLastReportsReportsReportConnectionPageInfo {
	return v.PageInfo
}

// GetEdges returns getLastReportsReportsReportConnection.Edges, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnection) GetEdges() []getLastReportsReportsReportConnectionEdgesReportEdge {
	return v.Edges
}

// getLastReportsReportsReportConnectionEdgesReportEdge includes the requested fields of the GraphQL type ReportEdge.
// The GraphQL type's documentation follows.
//
// Pagination entry
type getLastReportsReportsReportConnectionEdgesReportEdge struct {
	// Node instance
	Node getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport `json:"node"`
	// Pagination cursor
	Cursor string `json:"cursor"`
}

// GetNode returns getLastReportsReportsReportConnectionEdgesReportEdge.Node, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionEdgesReportEdge) GetNode() getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport {
	return v.Node
}

// GetCursor returns getLastReportsReportsReportConnectionEdgesReportEdge.Cursor, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionEdgesReportEdge) GetCursor() string { return v.Cursor }

// getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport includes the requested fields of the GraphQL type Report.
// The GraphQL type's documentation follows.
//
// Application log or diagnostic information
type getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport struct {
	// Report index within the context of the input that produced it
	Index int `json:"index"`
	// Input whose processing produced the report
	Input getLastReportsReportsReportConnectionEdgesReportEdgeNodeReportInput `json:"input"`
	// Report data as a payload in Ethereum hex binary format, starting with '0x'
	Payload string `json:"payload"`
}

// GetIndex returns getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport.Index, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport) GetIndex() int {
	return v.Index
}

// GetInput returns getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport.Input, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport) GetInput() getLastReportsReportsReportConnectionEdgesReportEdgeNodeReportInput {
	return v.Input
}

// GetPayload returns getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport.Payload, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionEdgesReportEdgeNodeReport) GetPayload() string {
	return v.Payload
}

// getLastReportsReportsReportConnectionEdgesReportEdgeNodeReportInput includes the requested fields of the GraphQL type Input.
// The GraphQL type's documentation follows.
//
// Request submitted to the application to advance its state
type getLastReportsReportsReportConnectionEdgesReportEdgeNodeReportInput struct {
	// Input index starting from genesis
	Index int `json:"index"`
}

// GetIndex returns getLastReportsReportsReportConnectionEdgesReportEdgeNodeReportInput.Index, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionEdgesReportEdgeNodeReportInput) GetIndex() int {
	return v.Index
}

// getLastReportsReportsReportConnectionPageInfo includes the requested fields of the GraphQL type PageInfo.
// The GraphQL type's documentation follows.
//
// Page metadata for the cursor-based Connection pagination pattern
type getLastReportsReportsReportConnectionPageInfo struct {
	// Cursor pointing to the first entry of the page
	StartCursor string `json:"startCursor"`
	// Cursor pointing to the last entry of the page
	EndCursor string `json:"endCursor"`
	// Indicates if there are additional entries after the end curs
	HasNextPage bool `json:"hasNextPage"`
	// Indicates if there are additional entries before the start curs
	HasPreviousPage bool `json:"hasPreviousPage"`
}

// GetStartCursor returns getLastReportsReportsReportConnectionPageInfo.StartCursor, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionPageInfo) GetStartCursor() string { return v.StartCursor }

// GetEndCursor returns getLastReportsReportsReportConnectionPageInfo.EndCursor, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionPageInfo) GetEndCursor() string { return v.EndCursor }

// GetHasNextPage returns getLastReportsReportsReportConnectionPageInfo.HasNextPage, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionPageInfo) GetHasNextPage() bool { return v.HasNextPage }

// GetHasPreviousPage returns getLastReportsReportsReportConnectionPageInfo.HasPreviousPage, and is useful for accessing the field via an interface.
func (v *getLastReportsReportsReportConnectionPageInfo) GetHasPreviousPage() bool {
	return v.HasPreviousPage
}

// getLastReportsResponse is returned by getLastReports on success.
type getLastReportsResponse struct {
	// Get reports with support for pagination
	Reports getLastReportsReportsReportConnection `json:"reports"`
}

// GetReports returns getLastReportsResponse.Reports, and is useful for accessing the field via an interface.
func (v *getLastReportsResponse) GetReports() getLastReportsReportsReportConnection { return v.Reports }

// getNoticeNotice includes the requested fields of the GraphQL type Notice.
// The GraphQL type's documentation follows.
//
// Informational statement that can be validated in the base layer blockchain
type getNoticeNotice struct {
	// Notice data as a payload in Ethereum hex binary format, starting with '0x'
	Payload string `json:"payload"`
}

// GetPayload returns getNoticeNotice.Payload, and is useful for accessing the field via an interface.
func (v *getNoticeNotice) GetPayload() string { return v.Payload }

// getNoticeResponse is returned by getNotice on success.
type getNoticeResponse struct {
	// Get notice based on its index
	Notice getNoticeNotice `json:"notice"`
}

// GetNotice returns getNoticeResponse.Notice, and is useful for accessing the field via an interface.
func (v *getNoticeResponse) GetNotice() getNoticeNotice { return v.Notice }

// getReportReport includes the requested fields of the GraphQL type Report.
// The GraphQL type's documentation follows.
//
// Application log or diagnostic information
type getReportReport struct {
	// Report data as a payload in Ethereum hex binary format, starting with '0x'
	Payload string `json:"payload"`
}

// GetPayload returns getReportReport.Payload, and is useful for accessing the field via an interface.
func (v *getReportReport) GetPayload() string { return v.Payload }

// getReportResponse is returned by getReport on success.
type getReportResponse struct {
	// Get report based on its index
	Report getReportReport `json:"report"`
}

// GetReport returns getReportResponse.Report, and is useful for accessing the field via an interface.
func (v *getReportResponse) GetReport() getReportReport { return v.Report }

// The query or mutation executed by getInput.
const getInput_Operation = `
query getInput ($inputIndex: Int!) {
	input(index: $inputIndex) {
		blockNumber
		reports {
			totalCount
		}
		notices {
			totalCount
		}
	}
}
`

func getInput(
	ctx context.Context,
	client graphql.Client,
	inputIndex int,
) (*getInputResponse, error) {
	req := &graphql.Request{
		OpName: "getInput",
		Query:  getInput_Operation,
		Variables: &__getInputInput{
			InputIndex: inputIndex,
		},
	}
	var err error

	var data getInputResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by getLastReports.
const getLastReports_Operation = `
query getLastReports ($last: Int) {
	reports(last: $last) {
		totalCount
		pageInfo {
			startCursor
			endCursor
			hasNextPage
			hasPreviousPage
		}
		edges {
			node {
				index
				input {
					index
				}
				payload
			}
			cursor
		}
	}
}
`

func getLastReports(
	ctx context.Context,
	client graphql.Client,
	last int,
) (*getLastReportsResponse, error) {
	req := &graphql.Request{
		OpName: "getLastReports",
		Query:  getLastReports_Operation,
		Variables: &__getLastReportsInput{
			Last: last,
		},
	}
	var err error

	var data getLastReportsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by getNotice.
const getNotice_Operation = `
query getNotice ($inputIndex: Int!, $noticeIndex: Int!) {
	notice(noticeIndex: $noticeIndex, inputIndex: $inputIndex) {
		payload
	}
}
`

func getNotice(
	ctx context.Context,
	client graphql.Client,
	inputIndex int,
	noticeIndex int,
) (*getNoticeResponse, error) {
	req := &graphql.Request{
		OpName: "getNotice",
		Query:  getNotice_Operation,
		Variables: &__getNoticeInput{
			InputIndex:  inputIndex,
			NoticeIndex: noticeIndex,
		},
	}
	var err error

	var data getNoticeResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by getReport.
const getReport_Operation = `
query getReport ($inputIndex: Int!, $reportIndex: Int!) {
	report(reportIndex: $reportIndex, inputIndex: $inputIndex) {
		payload
	}
}
`

func getReport(
	ctx context.Context,
	client graphql.Client,
	inputIndex int,
	reportIndex int,
) (*getReportResponse, error) {
	req := &graphql.Request{
		OpName: "getReport",
		Query:  getReport_Operation,
		Variables: &__getReportInput{
			InputIndex:  inputIndex,
			ReportIndex: reportIndex,
		},
	}
	var err error

	var data getReportResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
