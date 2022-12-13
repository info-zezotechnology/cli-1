// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package search

import (
	"sync"
)

// Ensure, that SearcherMock does implement Searcher.
// If this is not the case, regenerate this file with moq.
var _ Searcher = &SearcherMock{}

// SearcherMock is a mock implementation of Searcher.
//
//	func TestSomethingThatUsesSearcher(t *testing.T) {
//
//		// make and configure a mocked Searcher
//		mockedSearcher := &SearcherMock{
//			IssuesFunc: func(query Query) (IssuesResult, error) {
//				panic("mock out the Issues method")
//			},
//			RepositoriesFunc: func(query Query) (RepositoriesResult, error) {
//				panic("mock out the Repositories method")
//			},
//			URLFunc: func(query Query) string {
//				panic("mock out the URL method")
//			},
//		}
//
//		// use mockedSearcher in code that requires Searcher
//		// and then make assertions.
//
//	}
type SearcherMock struct {
	// IssuesFunc mocks the Issues method.
	IssuesFunc func(query Query) (IssuesResult, error)

	// RepositoriesFunc mocks the Repositories method.
	RepositoriesFunc func(query Query) (RepositoriesResult, error)

	// URLFunc mocks the URL method.
	URLFunc func(query Query) string

	// calls tracks calls to the methods.
	calls struct {
		// Issues holds details about calls to the Issues method.
		Issues []struct {
			// Query is the query argument value.
			Query Query
		}
		// Repositories holds details about calls to the Repositories method.
		Repositories []struct {
			// Query is the query argument value.
			Query Query
		}
		// URL holds details about calls to the URL method.
		URL []struct {
			// Query is the query argument value.
			Query Query
		}
	}
	lockIssues       sync.RWMutex
	lockRepositories sync.RWMutex
	lockURL          sync.RWMutex
}

// Issues calls IssuesFunc.
func (mock *SearcherMock) Issues(query Query) (IssuesResult, error) {
	if mock.IssuesFunc == nil {
		panic("SearcherMock.IssuesFunc: method is nil but Searcher.Issues was just called")
	}
	callInfo := struct {
		Query Query
	}{
		Query: query,
	}
	mock.lockIssues.Lock()
	mock.calls.Issues = append(mock.calls.Issues, callInfo)
	mock.lockIssues.Unlock()
	return mock.IssuesFunc(query)
}

// IssuesCalls gets all the calls that were made to Issues.
// Check the length with:
//
//	len(mockedSearcher.IssuesCalls())
func (mock *SearcherMock) IssuesCalls() []struct {
	Query Query
} {
	var calls []struct {
		Query Query
	}
	mock.lockIssues.RLock()
	calls = mock.calls.Issues
	mock.lockIssues.RUnlock()
	return calls
}

// Repositories calls RepositoriesFunc.
func (mock *SearcherMock) Repositories(query Query) (RepositoriesResult, error) {
	if mock.RepositoriesFunc == nil {
		panic("SearcherMock.RepositoriesFunc: method is nil but Searcher.Repositories was just called")
	}
	callInfo := struct {
		Query Query
	}{
		Query: query,
	}
	mock.lockRepositories.Lock()
	mock.calls.Repositories = append(mock.calls.Repositories, callInfo)
	mock.lockRepositories.Unlock()
	return mock.RepositoriesFunc(query)
}

// RepositoriesCalls gets all the calls that were made to Repositories.
// Check the length with:
//
//	len(mockedSearcher.RepositoriesCalls())
func (mock *SearcherMock) RepositoriesCalls() []struct {
	Query Query
} {
	var calls []struct {
		Query Query
	}
	mock.lockRepositories.RLock()
	calls = mock.calls.Repositories
	mock.lockRepositories.RUnlock()
	return calls
}

// URL calls URLFunc.
func (mock *SearcherMock) URL(query Query) string {
	if mock.URLFunc == nil {
		panic("SearcherMock.URLFunc: method is nil but Searcher.URL was just called")
	}
	callInfo := struct {
		Query Query
	}{
		Query: query,
	}
	mock.lockURL.Lock()
	mock.calls.URL = append(mock.calls.URL, callInfo)
	mock.lockURL.Unlock()
	return mock.URLFunc(query)
}

// URLCalls gets all the calls that were made to URL.
// Check the length with:
//
//	len(mockedSearcher.URLCalls())
func (mock *SearcherMock) URLCalls() []struct {
	Query Query
} {
	var calls []struct {
		Query Query
	}
	mock.lockURL.RLock()
	calls = mock.calls.URL
	mock.lockURL.RUnlock()
	return calls
}
