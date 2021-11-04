# 分治和回溯

大问题 可以分解为 子问题 有重复行 就可以通过分治来解决

一般泛型递归代码模板

```go
package main

func recursion(level int) {
	//终止条件
	if level > max {
		print(level)
		return 
	}
	//处理当前层逻辑
	process(level, data)
	
	//下探到下一层
	recursion(level)
}
```
分治比递归多一步在最后就是需要将子结果再组装到一起 最后返回

```
def divide_conquer(problem, param1, param2, ...):
	# recursion terminator  
    if problem is None:
        return  
    #prepare data
    data = prepare_data(problem)
    subproblems = split_problem(problem, data)
    # conquer subproblems
    subresult1 = self.divide_conquer(subproblems[0], p1, ...)
    subresult2 = self.divide_conquer(subproblems[1], p1, ...)
    subresult3 = self.divide_conquer(subproblems[2], p1, ...)
    # process and generate the final result
    result = process_result(subresult1, subresult2, subresult3, …)
    # revert the current level states
```

