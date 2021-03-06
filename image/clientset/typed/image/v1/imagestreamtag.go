package v1

import (
	scheme "github.com/openshift/client-go/image/clientset/scheme"
	v1 "github.com/openshift/origin/pkg/image/apis/image/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// ImageStreamTagsGetter has a method to return a ImageStreamTagInterface.
// A group's client should implement this interface.
type ImageStreamTagsGetter interface {
	ImageStreamTags(namespace string) ImageStreamTagInterface
}

// ImageStreamTagInterface has methods to work with ImageStreamTag resources.
type ImageStreamTagInterface interface {
	Create(*v1.ImageStreamTag) (*v1.ImageStreamTag, error)
	Update(*v1.ImageStreamTag) (*v1.ImageStreamTag, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.ImageStreamTag, error)
	ImageStreamTagExpansion
}

// imageStreamTags implements ImageStreamTagInterface
type imageStreamTags struct {
	client rest.Interface
	ns     string
}

// newImageStreamTags returns a ImageStreamTags
func newImageStreamTags(c *ImageV1Client, namespace string) *imageStreamTags {
	return &imageStreamTags{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the imageStreamTag, and returns the corresponding imageStreamTag object, and an error if there is any.
func (c *imageStreamTags) Get(name string, options meta_v1.GetOptions) (result *v1.ImageStreamTag, err error) {
	result = &v1.ImageStreamTag{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagestreamtags").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Create takes the representation of a imageStreamTag and creates it.  Returns the server's representation of the imageStreamTag, and an error, if there is any.
func (c *imageStreamTags) Create(imageStreamTag *v1.ImageStreamTag) (result *v1.ImageStreamTag, err error) {
	result = &v1.ImageStreamTag{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("imagestreamtags").
		Body(imageStreamTag).
		Do().
		Into(result)
	return
}

// Update takes the representation of a imageStreamTag and updates it. Returns the server's representation of the imageStreamTag, and an error, if there is any.
func (c *imageStreamTags) Update(imageStreamTag *v1.ImageStreamTag) (result *v1.ImageStreamTag, err error) {
	result = &v1.ImageStreamTag{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagestreamtags").
		Name(imageStreamTag.Name).
		Body(imageStreamTag).
		Do().
		Into(result)
	return
}

// Delete takes name of the imageStreamTag and deletes it. Returns an error if one occurs.
func (c *imageStreamTags) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagestreamtags").
		Name(name).
		Body(options).
		Do().
		Error()
}
