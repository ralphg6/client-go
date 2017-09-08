package fake

import (
	v1 "github.com/openshift/client-go/project/clientset/typed/project/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeProjectV1 struct {
	*testing.Fake
}

func (c *FakeProjectV1) Projects() v1.ProjectInterface {
	return &FakeProjects{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeProjectV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
