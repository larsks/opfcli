package cmd

import "fmt"

type GroupExistsError struct {
	Name string
}

type NamespaceExistsError struct {
	Name string
}

type NamespaceMissingError struct {
	Name string
}

type GroupMissingError struct {
	Name string
}

// PathError is a common set of attributes for errors involving
// operations on filesystem paths.
type PathError struct {
	Err  error
	Path string
}

type DirectoryCreateFailed struct {
	PathError
}

type WriteFailed struct {
	PathError
}

type CommandFailedError struct {
	Err     error
	Command string
}

func (err GroupExistsError) Error() string {
	return fmt.Sprintf("Group %s already exists", err.Name)
}

func (err NamespaceExistsError) Error() string {
	return fmt.Sprintf("Namespace %s already exists", err.Name)
}

func (err DirectoryCreateFailed) Error() string {
	return fmt.Sprintf(
		"Failed to create directory %s: %v",
		err.Path, err.Err,
	)
}

func (err WriteFailed) Error() string {
	return fmt.Sprintf(
		"Failed to write to %s: %v",
		err.Path, err.Err,
	)
}

func (err NamespaceMissingError) Error() string {
	return fmt.Sprintf("Namespace %s does not exist", err.Name)
}

func (err GroupMissingError) Error() string {
	return fmt.Sprintf("Group %s does not exist", err.Name)
}
