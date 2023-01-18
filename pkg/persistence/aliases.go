package persistence

import (
	"github.com/Pluto-tv/acyl/pkg/config"
	"github.com/Pluto-tv/acyl/pkg/metrics"
	"github.com/Pluto-tv/acyl/pkg/models"
)

type QAType = models.QAType
type QAEnvironment = models.QAEnvironment
type QAEnvironments = models.QAEnvironments
type EnvironmentStatus = models.EnvironmentStatus
type RepoRevisionData = models.RepoRevisionData
type RefMap = models.RefMap
type QADestroyReason = models.QADestroyReason
type QAEnvironmentEvent = models.QAEnvironmentEvent

type ServerConfig = config.ServerConfig
type MetricsCollector = metrics.Collector
