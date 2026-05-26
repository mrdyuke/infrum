package domain

type LibraryFile map[string]string
type LibraryPaths map[string]LibraryFile

type Library struct {
	LibName string
	LibPath LibraryPaths
}

type LibraryList []Library
