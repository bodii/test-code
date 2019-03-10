/**
 * 队列的基本操作接口
 */

public interface Queue {

	// 向队列未尾添加元素
	public void push(Object e);

	// 从队列头部取出元素
	public Object pop();

	// 返回队列头部的元素
	public Object peek();

	// 判断队列是否为空
	public boolean isEmpty();

	// 返回队列的大小
	public int size();
}
