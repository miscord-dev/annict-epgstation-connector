// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package annict

import (
	"context"

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
	Id         string     `json:"id"`
	Title      string     `json:"title"`
	SeasonName SeasonName `json:"seasonName"`
	SeasonYear int        `json:"seasonYear"`
}

// GetId returns GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork.Id, and is useful for accessing the field via an interface.
func (v *GetOnHoldWorksViewerUserWorksWorkConnectionNodesWork) GetId() string { return v.Id }

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
	Id         string     `json:"id"`
	Title      string     `json:"title"`
	SeasonName SeasonName `json:"seasonName"`
	SeasonYear int        `json:"seasonYear"`
}

// GetId returns GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork.Id, and is useful for accessing the field via an interface.
func (v *GetWannaWatchWorksViewerUserWorksWorkConnectionNodesWork) GetId() string { return v.Id }

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
	Id         string     `json:"id"`
	Title      string     `json:"title"`
	SeasonName SeasonName `json:"seasonName"`
	SeasonYear int        `json:"seasonYear"`
}

// GetId returns GetWatchingWorksViewerUserWorksWorkConnectionNodesWork.Id, and is useful for accessing the field via an interface.
func (v *GetWatchingWorksViewerUserWorksWorkConnectionNodesWork) GetId() string { return v.Id }

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
				id
				title
				seasonName
				seasonYear
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
				id
				title
				seasonName
				seasonYear
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
				id
				title
				seasonName
				seasonYear
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
