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

package fake

import (
	v1beta1 "github.com/dichque/strimzi-client/pkg/apis/kafka/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKafkaTopics implements KafkaTopicInterface
type FakeKafkaTopics struct {
	Fake *FakeKafkaV1beta1
	ns   string
}

var kafkatopicsResource = schema.GroupVersionResource{Group: "kafka.strimzi.io", Version: "v1beta1", Resource: "kafkatopics"}

var kafkatopicsKind = schema.GroupVersionKind{Group: "kafka.strimzi.io", Version: "v1beta1", Kind: "KafkaTopic"}

// Get takes name of the kafkaTopic, and returns the corresponding kafkaTopic object, and an error if there is any.
func (c *FakeKafkaTopics) Get(name string, options v1.GetOptions) (result *v1beta1.KafkaTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kafkatopicsResource, c.ns, name), &v1beta1.KafkaTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KafkaTopic), err
}

// List takes label and field selectors, and returns the list of KafkaTopics that match those selectors.
func (c *FakeKafkaTopics) List(opts v1.ListOptions) (result *v1beta1.KafkaTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kafkatopicsResource, kafkatopicsKind, c.ns, opts), &v1beta1.KafkaTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.KafkaTopicList{ListMeta: obj.(*v1beta1.KafkaTopicList).ListMeta}
	for _, item := range obj.(*v1beta1.KafkaTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kafkaTopics.
func (c *FakeKafkaTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kafkatopicsResource, c.ns, opts))

}

// Create takes the representation of a kafkaTopic and creates it.  Returns the server's representation of the kafkaTopic, and an error, if there is any.
func (c *FakeKafkaTopics) Create(kafkaTopic *v1beta1.KafkaTopic) (result *v1beta1.KafkaTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kafkatopicsResource, c.ns, kafkaTopic), &v1beta1.KafkaTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KafkaTopic), err
}

// Update takes the representation of a kafkaTopic and updates it. Returns the server's representation of the kafkaTopic, and an error, if there is any.
func (c *FakeKafkaTopics) Update(kafkaTopic *v1beta1.KafkaTopic) (result *v1beta1.KafkaTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kafkatopicsResource, c.ns, kafkaTopic), &v1beta1.KafkaTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KafkaTopic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKafkaTopics) UpdateStatus(kafkaTopic *v1beta1.KafkaTopic) (*v1beta1.KafkaTopic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kafkatopicsResource, "status", c.ns, kafkaTopic), &v1beta1.KafkaTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KafkaTopic), err
}

// Delete takes name of the kafkaTopic and deletes it. Returns an error if one occurs.
func (c *FakeKafkaTopics) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kafkatopicsResource, c.ns, name), &v1beta1.KafkaTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKafkaTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kafkatopicsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.KafkaTopicList{})
	return err
}

// Patch applies the patch and returns the patched kafkaTopic.
func (c *FakeKafkaTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.KafkaTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kafkatopicsResource, c.ns, name, pt, data, subresources...), &v1beta1.KafkaTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KafkaTopic), err
}
