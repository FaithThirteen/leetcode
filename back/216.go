package main


// 回溯模板
//	void backtracking(参数) {
//		if (终止条件) {
//			存放结果;
//			return;
//		}
//
//		for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
//			处理节点;
//			backtracking(路径，选择列表); // 递归
//			回溯，撤销处理结果
//		}
//	}

// 本题和77题套路一样
func combinationSum3(k int, n int) [][]int {
	backtracking216(k,n,1,[]int{})
	return res
}

func backtracking216(k, n, start int, track []int) {
	if len(track) == k {
		sum := 0
		for _, v := range track {
			sum += v
		}
		if sum == n {
			temp := make([]int, k)
			copy(temp, track)
			res = append(res, temp)
		}
	}
	// 剪枝
	if len(track)+(9-start+1) < k {
		return
	}

	for i := start; i <= 9; i++ {
		track = append(track, i)
		backtracking216(k, n, i+1, track)
		track = track[:len(track)-1]
	}
}

func combinationSum3_3(k int, n int) [][]int {
	var track []int// 遍历路径
	var result [][]int// 存放结果集
	backTree(n,k,1,&track,&result)
	return result
}
func backTree(n,k,startIndex int,track *[]int,result *[][]int){
	if len(*track)==k{
		var sum int
		tmp:=make([]int,k)
		for k,v:=range *track{
			sum+=v
			tmp[k]=v
		}
		if sum==n{
			*result=append(*result,tmp)
		}
		return
	}
	for i:=startIndex;i<=9-(k-len(*track))+1;i++{//减枝（k-len(*track)表示还剩多少个可填充的元素）
		*track=append(*track,i)//记录路径
		backTree(n,k,i+1,track,result)//递归
		*track=(*track)[:len(*track)-1]//回溯
	}
}
