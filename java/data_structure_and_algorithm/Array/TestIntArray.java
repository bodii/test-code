package array;

public class TestIntArray {
	public static void main(String[] args) {
		IntArray it = new IntArray();
		System.out.println("添加0－32到数组中");
		for (int i = 0; i < 33; i++)
			it.addLast(i);
		System.out.println(it + "\n");
		System.out.println("在数组头部添加55: ");
		it.addFirst(55);
		System.out.println(it + "\n");
		System.out.println("在数组尾部添加99: ");
		it.addLast(99);
		System.out.println(it + "\n");


		System.out.println("删除索引为2的元素： ");
		it.remove(2);
		System.out.println(it + "\n");
		System.out.println("删除索引为7的元素,且删除13次，每次索引加1： ");
		for (int i = 7; i < 20; i++)
			it.remove(i);
		System.out.println(it + "\n");
		System.out.println("删除第一个元素： " + it.remvoeFirst());
		System.out.println(it + "\n");
		System.out.println("删除最后一个元素： " + it.removeLast());
		System.out.println(it + "\n");
		System.out.println("删除指定元素: 3");
		it.removeElement(3);
		System.out.println(it + "\n");

		System.out.println("查询20是否存在：" + it.contains(20));
		System.out.println();

		System.out.println("设置索引6的值为41:");
		it.set(6, 41);
		System.out.println(it + "\n");

	}
}
