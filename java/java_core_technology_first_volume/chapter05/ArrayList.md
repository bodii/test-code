#### java.util.ArrayList

* void set(int index, E obj)
	设置数组列表指定位置的元素值，这个操作将覆盖这个位置的原有内容。
	参数: index		位置(必须介于0 - size()-1 之间)
		  obj		新的值

* E get(int index)
	获得指定位置的元素值。
	参数: index		获得的元素位置(必须介于0 ~ size()-1 之间)
	
* void add(int index, E ojb) 
	向后移动元素，以便插入元素。
	参数: index		插入位置(必须介于0 ~ size()-1之间)
		 obj		新元素

* E remove(int index)
	删除一个元素，将将后面的元素向前移动。被删除的元素由返回值返回。
	参数: index			被删除的元素位置(必须介于0 ~ size()-1之间)
