package fake

import (
	internalversion "github.com/openshift/client-go/project/internalclientset/typed/project/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeProject struct {
	*testing.Fake
}

func (c *FakeProject) Projects() internalversion.ProjectInterface {
	return &FakeProjects{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeProject) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
