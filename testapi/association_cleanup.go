//
// Test the lbakerchef/chef chef server api /organization/:org/user and /organization/:org/association_requests
// endpoints against a live server
//
package testapi

import (
	"fmt"
	"os"

	"github.com/lbakerchef/chef"
)

// association_cleanup exercise the chef server api
func AssociationCleanup() {
	client := Client()
	deleteUser(client, "usrinvite")
	deleteUser(client, "usr2invite")
	deleteUser(client, "usradd")
}

// deleteUser uses the chef server api to delete a single user
func deleteUser(client *chef.Client, name string) (err error) {
	err = client.Users.Delete(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Issue deleting org:", err)
	}
	return
}
