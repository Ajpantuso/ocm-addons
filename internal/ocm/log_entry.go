package ocm

import (
	"fmt"
	"strings"

	"github.com/mt-sre/ocm-addons/internal/output"
	slv1 "github.com/openshift-online/ocm-sdk-go/servicelogs/v1"
)

func NewLogEntry(cluster *Cluster, opts ...LogEntryOption) (LogEntry, error) {
	var cfg LogEntryConfig

	for _, opt := range opts {
		opt(&cfg)
	}

	entry, err := slv1.NewLogEntry().
		ClusterID(cluster.ID()).
		ClusterUUID(cluster.ExternalID()).
		Description(cfg.Description).
		InternalOnly(cfg.InternalOnly).
		ServiceName(cfg.ServiceName).
		Severity(slv1.Severity(cfg.Severity)).
		Summary(cfg.Summary).
		Build()
	if err != nil {
		return LogEntry{}, fmt.Errorf("building log entry: %w", err)
	}

	return LogEntry{
		Entry: entry,
	}, nil
}

type LogEntry struct {
	Entry *slv1.LogEntry
}

func (l *LogEntry) ToRow() output.Row {
	severity := strings.ToUpper(string(l.Entry.Severity()))

	return output.Row{
		{Name: "timestamp", Value: l.Entry.Timestamp()},
		{Name: "cluster_uuid", Value: l.Entry.ClusterUUID()},
		{Name: "description", Value: l.Entry.Description()},
		{Name: "id", Value: l.Entry.ID()},
		{Name: "service_name", Value: l.Entry.ServiceName()},
		{Name: "severity", Value: severity},
		{Name: "summary", Value: l.Entry.Summary()},
		{Name: "username", Value: l.Entry.Username()},
	}
}

type LogLevel string

const (
	LogLevelNone    = ""
	LogLevelDebug   = "Debug"
	LogLevelInfo    = "Info"
	LogLevelWarning = "Warning"
	LogLevelError   = "Error"
	LogLevelFatal   = "Fatal"
)

func NewLogEntrySorter(size int, sortFunc LogEntrySortFunc) *LogEntrySorter {
	return &LogEntrySorter{
		entries:  make([]LogEntry, 0, size),
		sortFunc: sortFunc,
	}
}

type LogEntrySorter struct {
	entries  []LogEntry
	sortFunc LogEntrySortFunc
}

func (s *LogEntrySorter) Len() int           { return len(s.entries) }
func (s *LogEntrySorter) Swap(i, j int)      { s.entries[i], s.entries[j] = s.entries[j], s.entries[i] }
func (s *LogEntrySorter) Less(i, j int) bool { return s.sortFunc(s.entries[i], s.entries[j]) }

func (s *LogEntrySorter) Append(e LogEntry) {
	s.entries = append(s.entries, e)
}

func (s *LogEntrySorter) Entries() []LogEntry {
	result := make([]LogEntry, len(s.entries))

	copy(result, s.entries)

	return result
}

type LogEntrySortFunc func(LogEntry, LogEntry) bool

func LogEntryByTime(ord Order) LogEntrySortFunc {
	return func(e1, e2 LogEntry) bool {
		if ord == OrderAsc {
			return e1.Entry.Timestamp().Before(e2.Entry.Timestamp())
		}

		return e1.Entry.Timestamp().After(e2.Entry.Timestamp())
	}
}

type Order string

const (
	OrderNone = ""
	OrderAsc  = "ascending"
	OrderDesc = "descending"
)

type LogEntryOption func(*LogEntryConfig)

func LogEntryDescription(desc string) LogEntryOption {
	return func(c *LogEntryConfig) {
		c.Description = desc
	}
}

func LogEntryInternalOnly(c *LogEntryConfig) {
	c.InternalOnly = true
}

func LogEntryServiceName(name string) LogEntryOption {
	return func(c *LogEntryConfig) {
		c.ServiceName = name
	}
}

func LogEntrySeverity(sev string) LogEntryOption {
	return func(c *LogEntryConfig) {
		c.Severity = sev
	}
}

func LogEntrySummary(sum string) LogEntryOption {
	return func(c *LogEntryConfig) {
		c.Summary = sum
	}
}

type LogEntryConfig struct {
	Description  string
	InternalOnly bool
	ServiceName  string
	Severity     string
	Summary      string
}
