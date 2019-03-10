/**
 * 栈的基本操作接口
 */

public interface Stack {
	// 将一个元素添加到栈顶
	public void push(Object e);

	// 将栈顶的元素移出
	public Object pop();

	// 查询栈顶的值 
	public Object peek();

	// 判断是否是空栈
	public boolean isEmpty();

	// 获取栈的元素个数
	public int size();
}
