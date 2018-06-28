package lookout

import (
	"github.com/src-d/lookout/pb"
)

// Event represents a repository event.
type Event interface {
	// ID returns the EventID.
	ID() EventID
	// Type returns the EventType, in order to identify the concreate type of
	// the event.
	Type() EventType
	// Revision returns a commit revision, containing the head and the base of
	// the changes.
	Revision() *CommitRevision
}

type EventID = pb.EventID
type EventType = pb.EventType

type CommitRevision = pb.CommitRevision
type RepositoryInfo = pb.RepositoryInfo
type ReferencePointer = pb.ReferencePointer

type PushEvent = pb.PushEvent
type PullRequestEvent = pb.PullRequestEvent
