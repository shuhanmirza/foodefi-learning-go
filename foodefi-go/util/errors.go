package util

type InvalidEventFieldType struct{}

func (m *InvalidEventFieldType) Error() string {
	return "invalid event field type"
}

type UserRoleNotPermitted struct{}

func (m *UserRoleNotPermitted) Error() string {
	return "user role not permitted to do this action"
}

type InvalidLoginRequest struct{}

func (m *InvalidLoginRequest) Error() string {
	return "invalid login request"
}

type InvalidBlockchainId struct{}

func (m *InvalidBlockchainId) Error() string {
	return "invalid blockchain id"
}

type InvalidEvent struct{}

func (m *InvalidEvent) Error() string {
	return "invalid blockchain event"
}
