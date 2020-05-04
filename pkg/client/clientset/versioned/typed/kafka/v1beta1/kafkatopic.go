/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/dichque/strimzi-client/pkg/apis/kafka/v1beta1"
	scheme "github.com/dichque/strimzi-client/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KafkaTopicsGetter has a method to return a KafkaTopicInterface.
// A group's client should implement this interface.
type KafkaTopicsGetter interface {
	KafkaTopics(namespace string) KafkaTopicInterface
}

// KafkaTopicInterface has methods to work with KafkaTopic resources.
type KafkaTopicInterface interface {
	Create(*v1beta1.KafkaTopic) (*v1beta1.KafkaTopic, error)
	Update(*v1beta1.KafkaTopic) (*v1beta1.KafkaTopic, error)
	UpdateStatus(*v1beta1.KafkaTopic) (*v1beta1.KafkaTopic, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.KafkaTopic, error)
	List(opts v1.ListOptions) (*v1beta1.KafkaTopicList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.KafkaTopic, err error)
	KafkaTopicExpansion
}

// kafkaTopics implements KafkaTopicInterface
type kafkaTopics struct {
	client rest.Interface
	ns     string
}

// newKafkaTopics returns a KafkaTopics
func newKafkaTopics(c *KafkaV1beta1Client, namespace string) *kafkaTopics {
	return &kafkaTopics{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kafkaTopic, and returns the corresponding kafkaTopic object, and an error if there is any.
func (c *kafkaTopics) Get(name string, options v1.GetOptions) (result *v1beta1.KafkaTopic, err error) {
	result = &v1beta1.KafkaTopic{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkatopics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KafkaTopics that match those selectors.
func (c *kafkaTopics) List(opts v1.ListOptions) (result *v1beta1.KafkaTopicList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.KafkaTopicList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkatopics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kafkaTopics.
func (c *kafkaTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kafkatopics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kafkaTopic and creates it.  Returns the server's representation of the kafkaTopic, and an error, if there is any.
func (c *kafkaTopics) Create(kafkaTopic *v1beta1.KafkaTopic) (result *v1beta1.KafkaTopic, err error) {
	result = &v1beta1.KafkaTopic{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kafkatopics").
		Body(kafkaTopic).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kafkaTopic and updates it. Returns the server's representation of the kafkaTopic, and an error, if there is any.
func (c *kafkaTopics) Update(kafkaTopic *v1beta1.KafkaTopic) (result *v1beta1.KafkaTopic, err error) {
	result = &v1beta1.KafkaTopic{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkatopics").
		Name(kafkaTopic.Name).
		Body(kafkaTopic).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kafkaTopics) UpdateStatus(kafkaTopic *v1beta1.KafkaTopic) (result *v1beta1.KafkaTopic, err error) {
	result = &v1beta1.KafkaTopic{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkatopics").
		Name(kafkaTopic.Name).
		SubResource("status").
		Body(kafkaTopic).
		Do().
		Into(result)
	return
}

// Delete takes name of the kafkaTopic and deletes it. Returns an error if one occurs.
func (c *kafkaTopics) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkatopics").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kafkaTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkatopics").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kafkaTopic.
func (c *kafkaTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.KafkaTopic, err error) {
	result = &v1beta1.KafkaTopic{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kafkatopics").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
