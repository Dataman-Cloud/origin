package fake

import (
	user_v1 "github.com/openshift/origin/pkg/user/apis/user/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUsers implements UserResourceInterface
type FakeUsers struct {
	Fake *FakeUserV1
	ns   string
}

var usersResource = schema.GroupVersionResource{Group: "user.openshift.io", Version: "v1", Resource: "users"}

var usersKind = schema.GroupVersionKind{Group: "user.openshift.io", Version: "v1", Kind: "User"}

// Get takes name of the userResource, and returns the corresponding userResource object, and an error if there is any.
func (c *FakeUsers) Get(name string, options v1.GetOptions) (result *user_v1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(usersResource, c.ns, name), &user_v1.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*user_v1.User), err
}

// List takes label and field selectors, and returns the list of Users that match those selectors.
func (c *FakeUsers) List(opts v1.ListOptions) (result *user_v1.UserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(usersResource, usersKind, c.ns, opts), &user_v1.UserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &user_v1.UserList{}
	for _, item := range obj.(*user_v1.UserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested users.
func (c *FakeUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(usersResource, c.ns, opts))

}

// Create takes the representation of a userResource and creates it.  Returns the server's representation of the userResource, and an error, if there is any.
func (c *FakeUsers) Create(userResource *user_v1.User) (result *user_v1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(usersResource, c.ns, userResource), &user_v1.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*user_v1.User), err
}

// Update takes the representation of a userResource and updates it. Returns the server's representation of the userResource, and an error, if there is any.
func (c *FakeUsers) Update(userResource *user_v1.User) (result *user_v1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(usersResource, c.ns, userResource), &user_v1.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*user_v1.User), err
}

// Delete takes name of the userResource and deletes it. Returns an error if one occurs.
func (c *FakeUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(usersResource, c.ns, name), &user_v1.User{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(usersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &user_v1.UserList{})
	return err
}

// Patch applies the patch and returns the patched userResource.
func (c *FakeUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *user_v1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(usersResource, c.ns, name, data, subresources...), &user_v1.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*user_v1.User), err
}
