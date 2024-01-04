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

// The query or mutation executed by GetOnHoldWorks.
const GetOnHoldWorks_Operation = `
query GetOnHoldWorks {
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

func GetOnHoldWorks(
	ctx context.Context,
	client graphql.Client,
) (*GetOnHoldWorksResponse, error) {
	req := &graphql.Request{
		OpName: "GetOnHoldWorks",
		Query:  GetOnHoldWorks_Operation,
	}
	var err error

	var data GetOnHoldWorksResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by GetWannaWatchWorks.
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
	ctx context.Context,
	client graphql.Client,
) (*GetWannaWatchWorksResponse, error) {
	req := &graphql.Request{
		OpName: "GetWannaWatchWorks",
		Query:  GetWannaWatchWorks_Operation,
	}
	var err error

	var data GetWannaWatchWorksResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by GetWatchingWorks.
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
	ctx context.Context,
	client graphql.Client,
) (*GetWatchingWorksResponse, error) {
	req := &graphql.Request{
		OpName: "GetWatchingWorks",
		Query:  GetWatchingWorks_Operation,
	}
	var err error

	var data GetWatchingWorksResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
