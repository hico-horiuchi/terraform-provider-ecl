/*
Package target_groups contains functionality for working with ECL Managed Load Balancer resources.

Example to list target groups

	listOpts := target_groups.ListOpts{}

	allPages, err := target_groups.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allTargetGroups, err := target_groups.ExtractTargetGroups(allPages)
	if err != nil {
		panic(err)
	}

	for _, targetGroup := range allTargetGroups {
		fmt.Printf("%+v\n", targetGroup)
	}

Example to create a target group

	member1 := target_groups.CreateOptsMember{
		IPAddress: "192.168.0.7",
		Port: 80,
		Weight: 1,
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	createOpts := target_groups.CreateOpts{
		Name: "target_group",
		Description: "description",
		Tags: tags,
		LoadBalancerID: "67fea379-cff0-4191-9175-de7d6941a040",
		Members: &[]target_groups.CreateOptsMember{member1},
	}

	targetGroup, err := target_groups.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", targetGroup)

Example to show a target group

	showOpts := target_groups.ShowOpts{}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	targetGroup, err := target_groups.Show(managedLoadBalancerClient, id, showOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", targetGroup)

Example to update a target group

	name := "target_group"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	updateOpts := target_groups.UpdateOpts{
		Name: &name,
		Description: &description,
		Tags: &tags,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	targetGroup, err := target_groups.Update(managedLoadBalancerClient, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", targetGroup)

Example to delete a target group

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := target_groups.Delete(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to create staged target group configurations

	member1 := target_groups.CreateStagedOptsMember{
		IPAddress: "192.168.0.7",
		Port: 80,
		Weight: 1,
	}
	createStagedOpts := target_groups.CreateStagedOpts{
		Members: &[]target_groups.CreateStagedOptsMember{member1},
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	targetGroupConfigurations, err := target_groups.CreateStaged(managedLoadBalancerClient, id, createStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", targetGroupConfigurations)

Example to show staged target group configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	targetGroupConfigurations, err := target_groups.ShowStaged(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", targetGroupConfigurations)

Example to update staged target group configurations

	member1IPAddress := "192.168.0.7"
	member1Port := 80
	member1Weight := 1
	member1 := target_groups.UpdateStagedOptsMember{
		IPAddress: &member1IPAddress,
		Port: &member1Port,
		Weight: &member1Weight,
	}

	updateStagedOpts := target_groups.UpdateStagedOpts{
		Members: &[]target_groups.UpdateStagedOptsMember{member1},
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	targetGroupConfigurations, err := target_groups.UpdateStaged(managedLoadBalancerClient, updateStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", targetGroupConfigurations)

Example to cancel staged target group configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := target_groups.CancelStaged(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package target_groups
