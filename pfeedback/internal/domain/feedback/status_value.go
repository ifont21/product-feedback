package domain

import "errors"

type StatusValue int64

const (
	Pending StatusValue = iota
	InProgress
	InReview
	Done
)

func (s StatusValue) String() string {
	switch s {
	case Pending:
		return "pending"
	case InProgress:
		return "in_progress"
	case InReview:
		return "in_review"
	case Done:
		return "done"
	}

	return "unknown"
}

func (s StatusValue) Label() string {
	switch s {
	case Pending:
		return "Pending"
	case InProgress:
		return "In Progress"
	case InReview:
		return "In Review"
	case Done:
		return "Done"
	}

	return "unknown"
}

func ToStatusValue(s string) (StatusValue, error) {
	switch s {
	case "pending":
		return 1, nil
	case "in_progress":
		return 2, nil
	case "in_review":
		return 3, nil
	case "done":
		return 4, nil
	}

	return 0, errors.New("invalid status")
}
