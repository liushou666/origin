package fake

import (
	unversioned "github.com/openshift/origin/pkg/authorization/client/clientset_generated/internalclientset/typed/core/unversioned"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeCore struct {
	*core.Fake
}

func (c *FakeCore) Policies(namespace string) unversioned.PolicyInterface {
	return &FakePolicies{c, namespace}
}

// GetRESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCore) GetRESTClient() *restclient.RESTClient {
	return nil
}
