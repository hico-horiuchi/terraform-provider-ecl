/*
Package listeners contains functionality for working with ECL Managed Load Balancer resources.

Example to list listeners

	listOpts := listeners.ListOpts{}

	allPages, err := listeners.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allListeners, err := listeners.ExtractListeners(allPages)
	if err != nil {
		panic(err)
	}

	for _, listener := range allListeners {
		fmt.Printf("%+v\n", listener)
	}

Example to create a listener


	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	createOpts := listeners.CreateOpts{
		Name: "listener",
		Description: "description",
		Tags: tags,
		IPAddress: "10.0.0.1",
		Port: 443,
		Protocol: "https",
		LoadBalancerID: "67fea379-cff0-4191-9175-de7d6941a040",
	}

	listener, err := listeners.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", listener)

Example to show a listener

	showOpts := listeners.ShowOpts{}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	listener, err := listeners.Show(managedLoadBalancerClient, id, showOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", listener)

Example to update a listener

	name := "listener"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	updateOpts := listeners.UpdateOpts{
		Name: &name,
		Description: &description,
		Tags: &tags,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	listener, err := listeners.Update(managedLoadBalancerClient, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", listener)

Example to delete a listener

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := listeners.Delete(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to create staged listener configurations

	createStagedOpts := listeners.CreateStagedOpts{
		IPAddress: "10.0.0.1",
		Port: 443,
		Protocol: "https",
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	listenerConfigurations, err := listeners.CreateStaged(managedLoadBalancerClient, id, createStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", listenerConfigurations)

Example to show staged listener configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	listenerConfigurations, err := listeners.ShowStaged(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", listenerConfigurations)

Example to update staged listener configurations

	ipAddress := "10.0.0.1"
	port := 443
	protocol := "https"
	updateStagedOpts := listeners.UpdateStagedOpts{
		IPAddress: &ipAddress,
		Port: &port,
		Protocol: &protocol,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	listenerConfigurations, err := listeners.UpdateStaged(managedLoadBalancerClient, updateStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", listenerConfigurations)

Example to cancel staged listener configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := listeners.CancelStaged(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package listeners
