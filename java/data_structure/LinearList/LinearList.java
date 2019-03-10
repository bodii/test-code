/**
 * 线性表操作接口
 */
public interface LinearList {
	public boolean isEmpty(); // 线性表是否为空

	public int size(); // 返回线性表的大小

	public Object get(int index);

	public void set(int index, Object element);

	public boolean add(Object element); // 向线性表的尾部添加指定的元素

	public boolean add(int index, Object element); // 指定位置插入元素

	public Object remove(int index);

	public void clear(); // 清空线性表

}
