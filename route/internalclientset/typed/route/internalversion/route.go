package internalversion

import (
	scheme "github.com/openshift/client-go/route/internalclientset/scheme"
	route "github.com/openshift/origin/pkg/route/apis/route"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RoutesGetter has a method to return a RouteInterface.
// A group's client should implement this interface.
type RoutesGetter interface {
	Routes(namespace string) RouteInterface
}

// RouteInterface has methods to work with Route resources.
type RouteInterface interface {
	Create(*route.Route) (*route.Route, error)
	Update(*route.Route) (*route.Route, error)
	UpdateStatus(*route.Route) (*route.Route, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*route.Route, error)
	List(opts v1.ListOptions) (*route.RouteList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *route.Route, err error)
	RouteExpansion
}

// routes implements RouteInterface
type routes struct {
	client rest.Interface
	ns     string
}

// newRoutes returns a Routes
func newRoutes(c *RouteClient, namespace string) *routes {
	return &routes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the route, and returns the corresponding route object, and an error if there is any.
func (c *routes) Get(name string, options v1.GetOptions) (result *route.Route, err error) {
	result = &route.Route{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Routes that match those selectors.
func (c *routes) List(opts v1.ListOptions) (result *route.RouteList, err error) {
	result = &route.RouteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested routes.
func (c *routes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("routes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a route and creates it.  Returns the server's representation of the route, and an error, if there is any.
func (c *routes) Create(route *route.Route) (result *route.Route, err error) {
	result = &route.Route{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("routes").
		Body(route).
		Do().
		Into(result)
	return
}

// Update takes the representation of a route and updates it. Returns the server's representation of the route, and an error, if there is any.
func (c *routes) Update(route *route.Route) (result *route.Route, err error) {
	result = &route.Route{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("routes").
		Name(route.Name).
		Body(route).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *routes) UpdateStatus(route *route.Route) (result *route.Route, err error) {
	result = &route.Route{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("routes").
		Name(route.Name).
		SubResource("status").
		Body(route).
		Do().
		Into(result)
	return
}

// Delete takes name of the route and deletes it. Returns an error if one occurs.
func (c *routes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *routes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched route.
func (c *routes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *route.Route, err error) {
	result = &route.Route{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("routes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
