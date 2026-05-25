package domain

type LibraryPaths map[string]map[string]string

type Library struct {
	LibName string
	LibPath LibraryPaths
}

type LibraryList []Library
