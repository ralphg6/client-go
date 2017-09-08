package fake

import (
	internalversion "github.com/openshift/client-go/route/internalclientset/typed/route/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRoute struct {
	*testing.Fake
}

func (c *FakeRoute) Routes(namespace string) internalversion.RouteInterface {
	return &FakeRoutes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRoute) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
