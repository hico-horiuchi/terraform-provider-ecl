/*
Package system_updates contains functionality for working with ECL Managed Load Balancer resources.

Example to list system updates

	listOpts := system_updates.ListOpts{}

	allPages, err := system_updates.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allSystemUpdates, err := system_updates.ExtractSystemUpdates(allPages)
	if err != nil {
		panic(err)
	}

	for _, systemUpdate := range allSystemUpdates {
		fmt.Printf("%+v\n", systemUpdate)
	}

Example to show a system update

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	systemUpdate, err := system_updates.Show(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", systemUpdate)
*/
package system_updates
