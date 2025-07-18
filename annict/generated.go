// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package annict

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// GetOnHoldWorksResponse is returned by GetOnHoldWorks on success.
type GetOnHoldWorksResponse struct {
	Viewer GetOnHoldWorksViewerUser `json:"viewer"`
}

// GetViewer returns GetOnHoldWorksResponse.Viewer, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksResponse) GetViewer() GetOnHoldWorksViewerUser { return v.Viewer }

// GetOnHoldWorksViewerUser includes the requested fields of the GraphQL type User.
type GetOnHoldWorksViewerUser struct {
	Works GetOnHoldWorksViewerUserWorksWorkConnection `json:"works"`
}

// GetWorks returns GetOnHoldWorksViewerUser.Works, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUser) GetWorks() GetOnHoldWorksViewerUserWorksWorkConnection {
	return v.Works
}

// GetOnHoldWorksViewerUserWorksWorkConnection includes the requested fields of the GraphQL type WorkConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Work.
type GetOnHoldWorksViewerUserWorksWorkConnection struct {
	// A list of nodes.
	Nodes []GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork `json:"nodes"`
}

// GetNodes returns GetOnHoldWorksViewerUserWorksWorkConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUserWorksWorkConnection) GetNodes() []GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork {
	return v.Nodes
}

// GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork includes the requested fields of the GraphQL type Work.
// The GraphQL type's documentation follows.
//
// An anime title
type GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork struct {
	AnnictId   int                                                                           `json:"annictId"`
	Title      string                                                                        `json:"title"`
	SeasonName SeasonName                                                                    `json:"seasonName"`
	SeasonYear int                                                                           `json:"seasonYear"`
	Programs   GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection `json:"programs"`
}

// GetAnnictId returns GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork.AnnictId, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork) GetAnnictId() int { return v.AnnictId }

// GetTitle returns GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork.Title, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork) GetTitle() string { return v.Title }

// GetSeasonName returns GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork.SeasonName, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonName() SeasonName {
	return v.SeasonName
}

// GetSeasonYear returns GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork.SeasonYear, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonYear() int {
	return v.SeasonYear
}

// GetPrograms returns GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork.Programs, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork) GetPrograms() GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection {
	return v.Programs
}

// GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection includes the requested fields of the GraphQL type ProgramConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Program.
type GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection struct {
	// A list of nodes.
	Nodes []GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram `json:"nodes"`
}

// GetNodes returns GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection) GetNodes() []GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram {
	return v.Nodes
}

// GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram includes the requested fields of the GraphQL type Program.
type GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram struct {
	StartedAt time.Time `json:"startedAt"`
}

// GetStartedAt returns GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram.StartedAt, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram) GetStartedAt() time.Time {
	return v.StartedAt
}

// GetStopWatchingWorksResponse is returned by GetStopWatchingWorks on success.
type GetStopWatchingWorksResponse struct {
	Viewer GetStopWatchingWorksViewerUser `json:"viewer"`
}

// GetViewer returns GetStopWatchingWorksResponse.Viewer, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksResponse) GetViewer() GetStopWatchingWorksViewerUser { return v.Viewer }

// GetStopWatchingWorksViewerUser includes the requested fields of the GraphQL type User.
type GetStopWatchingWorksViewerUser struct {
	Works GetStopWatchingWorksViewerUserWorksWorkConnection `json:"works"`
}

// GetWorks returns GetStopWatchingWorksViewerUser.Works, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksViewerUser) GetWorks() GetStopWatchingWorksViewerUserWorksWorkConnection {
	return v.Works
}

// GetStopWatchingWorksViewerUserWorksWorkConnection includes the requested fields of the GraphQL type WorkConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Work.
type GetStopWatchingWorksViewerUserWorksWorkConnection struct {
	// A list of nodes.
	Nodes []GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork `json:"nodes"`
}

// GetNodes returns GetStopWatchingWorksViewerUserWorksWorkConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksViewerUserWorksWorkConnection) GetNodes() []GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork {
	return v.Nodes
}

// GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork includes the requested fields of the GraphQL type Work.
// The GraphQL type's documentation follows.
//
// An anime title
type GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork struct {
	AnnictId   int                                                                                 `json:"annictId"`
	Title      string                                                                              `json:"title"`
	SeasonName SeasonName                                                                          `json:"seasonName"`
	SeasonYear int                                                                                 `json:"seasonYear"`
	Programs   GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection `json:"programs"`
}

// GetAnnictId returns GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork.AnnictId, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetAnnictId() int {
	return v.AnnictId
}

// GetTitle returns GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork.Title, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetTitle() string {
	return v.Title
}

// GetSeasonName returns GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork.SeasonName, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonName() SeasonName {
	return v.SeasonName
}

// GetSeasonYear returns GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork.SeasonYear, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonYear() int {
	return v.SeasonYear
}

// GetPrograms returns GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork.Programs, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetPrograms() GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection {
	return v.Programs
}

// GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection includes the requested fields of the GraphQL type ProgramConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Program.
type GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection struct {
	// A list of nodes.
	Nodes []GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram `json:"nodes"`
}

// GetNodes returns GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection) GetNodes() []GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram {
	return v.Nodes
}

// GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram includes the requested fields of the GraphQL type Program.
type GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram struct {
	StartedAt time.Time `json:"startedAt"`
}

// GetStartedAt returns GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram.StartedAt, and is useful for accessing the field via an interface.
func (v *GetStopWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram) GetStartedAt() time.Time {
	return v.StartedAt
}

// GetWannaWatchWorksResponse is returned by GetWannaWatchWorks on success.
type GetWannaWatchWorksResponse struct {
	Viewer GetWannaWatchWorksViewerUser `json:"viewer"`
}

// GetViewer returns GetWannaWatchWorksResponse.Viewer, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksResponse) GetViewer() GetWannaWatchWorksViewerUser { return v.Viewer }

// GetWannaWatchWorksViewerUser includes the requested fields of the GraphQL type User.
type GetWannaWatchWorksViewerUser struct {
	Works GetWannaWatchWorksViewerUserWorksWorkConnection `json:"works"`
}

// GetWorks returns GetWannaWatchWorksViewerUser.Works, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUser) GetWorks() GetWannaWatchWorksViewerUserWorksWorkConnection {
	return v.Works
}

// GetWannaWatchWorksViewerUserWorksWorkConnection includes the requested fields of the GraphQL type WorkConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Work.
type GetWannaWatchWorksViewerUserWorksWorkConnection struct {
	// A list of nodes.
	Nodes []GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork `json:"nodes"`
}

// GetNodes returns GetWannaWatchWorksViewerUserWorksWorkConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUserWorksWorkConnection) GetNodes() []GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork {
	return v.Nodes
}

// GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork includes the requested fields of the GraphQL type Work.
// The GraphQL type's documentation follows.
//
// An anime title
type GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork struct {
	AnnictId   int                                                                               `json:"annictId"`
	Title      string                                                                            `json:"title"`
	SeasonName SeasonName                                                                        `json:"seasonName"`
	SeasonYear int                                                                               `json:"seasonYear"`
	Programs   GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection `json:"programs"`
}

// GetAnnictId returns GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork.AnnictId, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork) GetAnnictId() int {
	return v.AnnictId
}

// GetTitle returns GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork.Title, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork) GetTitle() string { return v.Title }

// GetSeasonName returns GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork.SeasonName, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonName() SeasonName {
	return v.SeasonName
}

// GetSeasonYear returns GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork.SeasonYear, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonYear() int {
	return v.SeasonYear
}

// GetPrograms returns GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork.Programs, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork) GetPrograms() GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection {
	return v.Programs
}

// GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection includes the requested fields of the GraphQL type ProgramConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Program.
type GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection struct {
	// A list of nodes.
	Nodes []GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram `json:"nodes"`
}

// GetNodes returns GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection) GetNodes() []GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram {
	return v.Nodes
}

// GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram includes the requested fields of the GraphQL type Program.
type GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram struct {
	StartedAt time.Time `json:"startedAt"`
}

// GetStartedAt returns GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram.StartedAt, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram) GetStartedAt() time.Time {
	return v.StartedAt
}

// GetWatchedWorksResponse is returned by GetWatchedWorks on success.
type GetWatchedWorksResponse struct {
	Viewer GetWatchedWorksViewerUser `json:"viewer"`
}

// GetViewer returns GetWatchedWorksResponse.Viewer, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksResponse) GetViewer() GetWatchedWorksViewerUser { return v.Viewer }

// GetWatchedWorksViewerUser includes the requested fields of the GraphQL type User.
type GetWatchedWorksViewerUser struct {
	Works GetWatchedWorksViewerUserWorksWorkConnection `json:"works"`
}

// GetWorks returns GetWatchedWorksViewerUser.Works, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksViewerUser) GetWorks() GetWatchedWorksViewerUserWorksWorkConnection {
	return v.Works
}

// GetWatchedWorksViewerUserWorksWorkConnection includes the requested fields of the GraphQL type WorkConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Work.
type GetWatchedWorksViewerUserWorksWorkConnection struct {
	// A list of nodes.
	Nodes []GetWatchedWorksViewerUserWorksWorkConnectionNodesWork `json:"nodes"`
}

// GetNodes returns GetWatchedWorksViewerUserWorksWorkConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksViewerUserWorksWorkConnection) GetNodes() []GetWatchedWorksViewerUserWorksWorkConnectionNodesWork {
	return v.Nodes
}

// GetWatchedWorksViewerUserWorksWorkConnectionNodesWork includes the requested fields of the GraphQL type Work.
// The GraphQL type's documentation follows.
//
// An anime title
type GetWatchedWorksViewerUserWorksWorkConnectionNodesWork struct {
	AnnictId   int                                                                            `json:"annictId"`
	Title      string                                                                         `json:"title"`
	SeasonName SeasonName                                                                     `json:"seasonName"`
	SeasonYear int                                                                            `json:"seasonYear"`
	Programs   GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection `json:"programs"`
}

// GetAnnictId returns GetWatchedWorksViewerUserWorksWorkConnectionNodesWork.AnnictId, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksViewerUserWorksWorkConnectionNodesWork) GetAnnictId() int { return v.AnnictId }

// GetTitle returns GetWatchedWorksViewerUserWorksWorkConnectionNodesWork.Title, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksViewerUserWorksWorkConnectionNodesWork) GetTitle() string { return v.Title }

// GetSeasonName returns GetWatchedWorksViewerUserWorksWorkConnectionNodesWork.SeasonName, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonName() SeasonName {
	return v.SeasonName
}

// GetSeasonYear returns GetWatchedWorksViewerUserWorksWorkConnectionNodesWork.SeasonYear, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonYear() int {
	return v.SeasonYear
}

// GetPrograms returns GetWatchedWorksViewerUserWorksWorkConnectionNodesWork.Programs, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksViewerUserWorksWorkConnectionNodesWork) GetPrograms() GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection {
	return v.Programs
}

// GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection includes the requested fields of the GraphQL type ProgramConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Program.
type GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection struct {
	// A list of nodes.
	Nodes []GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram `json:"nodes"`
}

// GetNodes returns GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection) GetNodes() []GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram {
	return v.Nodes
}

// GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram includes the requested fields of the GraphQL type Program.
type GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram struct {
	StartedAt time.Time `json:"startedAt"`
}

// GetStartedAt returns GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram.StartedAt, and is useful for accessing the field via an interface.
func (v *GetWatchedWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram) GetStartedAt() time.Time {
	return v.StartedAt
}

// GetWatchingWorksResponse is returned by GetWatchingWorks on success.
type GetWatchingWorksResponse struct {
	Viewer GetWatchingWorksViewerUser `json:"viewer"`
}

// GetViewer returns GetWatchingWorksResponse.Viewer, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksResponse) GetViewer() GetWatchingWorksViewerUser { return v.Viewer }

// GetWatchingWorksViewerUser includes the requested fields of the GraphQL type User.
type GetWatchingWorksViewerUser struct {
	Works GetWatchingWorksViewerUserWorksWorkConnection `json:"works"`
}

// GetWorks returns GetWatchingWorksViewerUser.Works, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUser) GetWorks() GetWatchingWorksViewerUserWorksWorkConnection {
	return v.Works
}

// GetWatchingWorksViewerUserWorksWorkConnection includes the requested fields of the GraphQL type WorkConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Work.
type GetWatchingWorksViewerUserWorksWorkConnection struct {
	// A list of nodes.
	Nodes []GetWatchingWorksViewerUserWorksWorkConnectionNodesWork `json:"nodes"`
}

// GetNodes returns GetWatchingWorksViewerUserWorksWorkConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUserWorksWorkConnection) GetNodes() []GetWatchingWorksViewerUserWorksWorkConnectionNodesWork {
	return v.Nodes
}

// GetWatchingWorksViewerUserWorksWorkConnectionNodesWork includes the requested fields of the GraphQL type Work.
// The GraphQL type's documentation follows.
//
// An anime title
type GetWatchingWorksViewerUserWorksWorkConnectionNodesWork struct {
	AnnictId   int                                                                             `json:"annictId"`
	Title      string                                                                          `json:"title"`
	SeasonName SeasonName                                                                      `json:"seasonName"`
	SeasonYear int                                                                             `json:"seasonYear"`
	Programs   GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection `json:"programs"`
}

// GetAnnictId returns GetWatchingWorksViewerUserWorksWorkConnectionNodesWork.AnnictId, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetAnnictId() int { return v.AnnictId }

// GetTitle returns GetWatchingWorksViewerUserWorksWorkConnectionNodesWork.Title, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetTitle() string { return v.Title }

// GetSeasonName returns GetWatchingWorksViewerUserWorksWorkConnectionNodesWork.SeasonName, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonName() SeasonName {
	return v.SeasonName
}

// GetSeasonYear returns GetWatchingWorksViewerUserWorksWorkConnectionNodesWork.SeasonYear, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetSeasonYear() int {
	return v.SeasonYear
}

// GetPrograms returns GetWatchingWorksViewerUserWorksWorkConnectionNodesWork.Programs, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetPrograms() GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection {
	return v.Programs
}

// GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection includes the requested fields of the GraphQL type ProgramConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Program.
type GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection struct {
	// A list of nodes.
	Nodes []GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram `json:"nodes"`
}

// GetNodes returns GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnection) GetNodes() []GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram {
	return v.Nodes
}

// GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram includes the requested fields of the GraphQL type Program.
type GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram struct {
	StartedAt time.Time `json:"startedAt"`
}

// GetStartedAt returns GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram.StartedAt, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUserWorksWorkConnectionNodesWorkProgramsProgramConnectionNodesProgram) GetStartedAt() time.Time {
	return v.StartedAt
}

// Season name
type SeasonName string

const (
	SeasonNameWinter SeasonName = "WINTER"
	SeasonNameSpring SeasonName = "SPRING"
	SeasonNameSummer SeasonName = "SUMMER"
	SeasonNameAutumn SeasonName = "AUTUMN"
)

var AllSeasonName = []SeasonName{
	SeasonNameWinter,
	SeasonNameSpring,
	SeasonNameSummer,
	SeasonNameAutumn,
}

// The query executed by GetOnHoldWorks.
const GetOnHoldWorks_Operation = `
query GetOnHoldWorks {
	viewer {
		works(state: ON_HOLD) {
			nodes {
				annictId
				title
				seasonName
				seasonYear
				programs(first: 1) {
					nodes {
						startedAt
					}
				}
			}
		}
	}
}
`

func GetOnHoldWorks(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *GetOnHoldWorksResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "GetOnHoldWorks",
		Query:  GetOnHoldWorks_Operation,
	}

	data_ = &GetOnHoldWorksResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

// The query executed by GetStopWatchingWorks.
const GetStopWatchingWorks_Operation = `
query GetStopWatchingWorks {
	viewer {
		works(state: STOP_WATCHING) {
			nodes {
				annictId
				title
				seasonName
				seasonYear
				programs(first: 1) {
					nodes {
						startedAt
					}
				}
			}
		}
	}
}
`

func GetStopWatchingWorks(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *GetStopWatchingWorksResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "GetStopWatchingWorks",
		Query:  GetStopWatchingWorks_Operation,
	}

	data_ = &GetStopWatchingWorksResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

// The query executed by GetWannaWatchWorks.
const GetWannaWatchWorks_Operation = `
query GetWannaWatchWorks {
	viewer {
		works(state: WANNA_WATCH) {
			nodes {
				annictId
				title
				seasonName
				seasonYear
				programs(first: 1) {
					nodes {
						startedAt
					}
				}
			}
		}
	}
}
`

func GetWannaWatchWorks(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *GetWannaWatchWorksResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "GetWannaWatchWorks",
		Query:  GetWannaWatchWorks_Operation,
	}

	data_ = &GetWannaWatchWorksResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

// The query executed by GetWatchedWorks.
const GetWatchedWorks_Operation = `
query GetWatchedWorks {
	viewer {
		works(state: WATCHED) {
			nodes {
				annictId
				title
				seasonName
				seasonYear
				programs(first: 1) {
					nodes {
						startedAt
					}
				}
			}
		}
	}
}
`

func GetWatchedWorks(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *GetWatchedWorksResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "GetWatchedWorks",
		Query:  GetWatchedWorks_Operation,
	}

	data_ = &GetWatchedWorksResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

// The query executed by GetWatchingWorks.
const GetWatchingWorks_Operation = `
query GetWatchingWorks {
	viewer {
		works(state: WATCHING) {
			nodes {
				annictId
				title
				seasonName
				seasonYear
				programs(first: 1) {
					nodes {
						startedAt
					}
				}
			}
		}
	}
}
`

func GetWatchingWorks(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *GetWatchingWorksResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "GetWatchingWorks",
		Query:  GetWatchingWorks_Operation,
	}

	data_ = &GetWatchingWorksResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}
