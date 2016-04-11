package actions

import (
	"kego.io/editor/client/models"
	"kego.io/editor/shared"
)

type NewMessage struct {
	Message string
}

type InitialState struct {
	Info *shared.Info
}

type ToggleBranch struct {
	Branch *models.BranchModel
}

type SelectBranch struct {
	Branch *models.BranchModel
}
type OpenBranch struct {
	Branch *models.BranchModel
}
type CloseBranch struct {
	Branch *models.BranchModel
}

type LoadSource struct {
	Contents models.BranchContentsInterface
	Signal   chan struct{}
}

type KeyboardEvent struct {
	KeyCode int
}

type LoadSourceSent struct {
	Branch *models.BranchModel
}
type LoadSourceCancelled struct {
	Branch *models.BranchModel
}
type LoadSourceSuccess struct {
	Branch *models.BranchModel
}
type LoadSourceError struct {
	Branch *models.BranchModel
}