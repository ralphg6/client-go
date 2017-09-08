package fake

import (
	v1 "github.com/openshift/client-go/security/clientset/typed/security/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSecurityV1 struct {
	*testing.Fake
}

func (c *FakeSecurityV1) SecurityContextConstraintses() v1.SecurityContextConstraintsInterface {
	return &FakeSecurityContextConstraintses{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSecurityV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
