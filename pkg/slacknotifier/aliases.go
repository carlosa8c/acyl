package slacknotifier

import (
	"github.com/Pluto-tv/acyl/pkg/ghclient"
	"github.com/Pluto-tv/acyl/pkg/models"
)

type RepoRevisionData = models.RepoRevisionData
type QADestroyReason = models.QADestroyReason

type RepoClient = ghclient.RepoClient
type BranchInfo = ghclient.BranchInfo
type CommitStatus = ghclient.CommitStatus

const (
	DestroyApiRequest = models.DestroyApiRequest
)
