package array;

/**
 * IntArray Class
 * 整型数组类
 *
 */
public class IntArray {
	/* 数组初始值 */
	private int[] data;
	/* 数组的元素个数 */
	private int size = 0;
	/* 数组的容量 */
	private int capacity = 10;
	
	/**
	 * 构造函数
	 * 用于初始化一个int数组，容量默认
	 */
	public IntArray() {
		data = new int[capacity];
	}

	/**
	 * 获取数组的元素个数
	 */
	public int getSize() {
		return size;
	}

	/**
	 * 获取数组的容量
	 */
	public int getCapacity() {
		return data.length;
	}

	/**
	 * 获取数组的元素是否为空
	 */
	public boolean isEmpty() {
		return size <= 0;
	}

	/**
	 * 获取指定索引的元素值
	 * @param index
	 * @return
	 */
	public int get(int index) {
		ensureIndex(index);
		return data[index];
	}

	/**
	 * 设置数组中指定索引的新值
	 * @param index
	 * @param e
	 */
	public void set(int index, int e) {
		ensureIndex(index);
		data[index] = e;
	}

	/**
	 * 检测数组中是否包含指定的元素
	 * @param e
	 * @return
	 */
	public boolean contains(int e) {
		if (isEmpty())
			return false;

		for (int d : data)
			if (d == e) 
				return true;

		return false;
	}

	/**
	 * 查找某个元素在数组中的索引是
	 * @param e
	 * @return 返回值小于0，表示没有
	 */
	public int find(int e) {
		if (isEmpty())
			return -1;
		
		for (int i = 0; i < size; i++)
			if (data[i] == e)
				return i;
		
		return -1;
	}

	/**
	 * 在数组的头部添加元素
	 * @param e
	 */
	public void addFirst(int e) {
		ensureCapacity();
		insert(0, e);
	}

	/**
	 * 在数组的末尾添加元素
	 * @param e int 要添加的元素
	 */
	public void addLast(int e) {
		ensureCapacity();

		data[size++] = e;
	}

	/**
	 * 在数组的第index个位置插入一个新的元素
	 */
	public void insert(int index, int e) {
		ensureIndex(index);
		ensureCapacity();

		if (isEmpty()) {
			addLast(e);
		} else {
			for (int i = size - 1; i >= index; i--)
				data[i+1] = data[i];

			data[index] = e;
			size ++;
		}
	}

	/**
	 * 在数组指定位置删除某个元素
	 * 
	 * @param index int 要删除数组的索引对应的值
	 * @return int 返回被删除的元素
	 */
	public int remove(int index) {
		ensureIndex(index);
		ensureCapacity();

		int result = data[index];
		for (int i = index; i < size-1; i++)
			data[i] = data[i+1];
		data[--size] = 0;
		
		return result;
	}

	/**
	 * 删除数组的第一个元素
	 * 
	 * @return int 返回被删除的元素
	 */
	public int remvoeFirst() {
		return remove(0);
	}

	/**
	 * 删除数组的最后一个元素
	 * 
	 * @return int 返回被删除的元素
	 */
	public int removeLast() {
		return remove(size-1);
	}

	/**
	 * 删除数组中指定的元素
	 * 
	 * @param e int 元素
	 */
	public void removeElement(int e) {
		int findIndex = find(e);
		if (-1 != findIndex) {
			remove(findIndex);
		}
	}

	/**
	 * 确保数组容量的适当大小
	 */
	private void ensureCapacity() {
		int maxCapacity = capacity * 2;
		int minCapacity = capacity / 2 - 1;
		if (size == capacity || (size <= minCapacity && size > 2)) {
			capacity = (size == capacity) ? maxCapacity : minCapacity;
			int[] newData = new int[capacity];

			for (int i = 0; i < size; i++)
				newData[i] = data[i];

			data = newData;
		}
	}

	/**
	 * 确保传入的数组索引是否合法
	 */
	private void ensureIndex(int index) throws IllegalArgumentException {
		if (0 > index || size < index)
			throw new IllegalArgumentException(
				"Failed. Require index >=0 and index <= size"
			);
	}


	/**
	 * 将数组转换成字符串
	 */
	@Override public String toString() {
		StringBuilder result = new StringBuilder();
		result.append(
			String.format(
				"IntArray: size = %d, capacity = %d%n", size, capacity
			)
		);
		result.append("[");
		for (int i : data) 
			result.append(i + ", ");
		result.replace(result.length()-2, result.length(), "]");

		return result.toString();
	}




}
