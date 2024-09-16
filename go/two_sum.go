func twoSum(nums []int, target int) []int {
	//create a map
	tmp := make(map[int]int);
	res:= []int{};

	// assign 
	tmp[nums[0]] = 1;

	for i:= 1; i < len(nums); i++{
		diff := target - nums[i];

		if(tmp[diff]!=0){
			res = append(res,tmp[diff],i-1);
			break;
		}
		else{
			tmp[nums[i]] = i+1;
		}
	}

	return res;
}