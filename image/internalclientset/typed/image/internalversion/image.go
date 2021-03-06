package internalversion

import (
	scheme "github.com/openshift/client-go/image/internalclientset/scheme"
	image "github.com/openshift/origin/pkg/image/apis/image"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ImagesGetter has a method to return a ImageInterface.
// A group's client should implement this interface.
type ImagesGetter interface {
	Images() ImageInterface
}

// ImageInterface has methods to work with Image resources.
type ImageInterface interface {
	Create(*image.Image) (*image.Image, error)
	Update(*image.Image) (*image.Image, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*image.Image, error)
	List(opts v1.ListOptions) (*image.ImageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *image.Image, err error)
	ImageExpansion
}

// images implements ImageInterface
type images struct {
	client rest.Interface
}

// newImages returns a Images
func newImages(c *ImageClient) *images {
	return &images{
		client: c.RESTClient(),
	}
}

// Get takes name of the image, and returns the corresponding image object, and an error if there is any.
func (c *images) Get(name string, options v1.GetOptions) (result *image.Image, err error) {
	result = &image.Image{}
	err = c.client.Get().
		Resource("images").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Images that match those selectors.
func (c *images) List(opts v1.ListOptions) (result *image.ImageList, err error) {
	result = &image.ImageList{}
	err = c.client.Get().
		Resource("images").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested images.
func (c *images) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("images").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a image and creates it.  Returns the server's representation of the image, and an error, if there is any.
func (c *images) Create(image *image.Image) (result *image.Image, err error) {
	result = &image.Image{}
	err = c.client.Post().
		Resource("images").
		Body(image).
		Do().
		Into(result)
	return
}

// Update takes the representation of a image and updates it. Returns the server's representation of the image, and an error, if there is any.
func (c *images) Update(image *image.Image) (result *image.Image, err error) {
	result = &image.Image{}
	err = c.client.Put().
		Resource("images").
		Name(image.Name).
		Body(image).
		Do().
		Into(result)
	return
}

// Delete takes name of the image and deletes it. Returns an error if one occurs.
func (c *images) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("images").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *images) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("images").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched image.
func (c *images) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *image.Image, err error) {
	result = &image.Image{}
	err = c.client.Patch(pt).
		Resource("images").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
