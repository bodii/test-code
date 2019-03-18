/**
 * 直接插入排序算法
 */

public class InsertSort {

	public static void main(String[] args) {
		int[] data = { 23, 45, 16, 7, 42 };
		int len = data.length;

		for (int i = 1; i < len; i++) {
			int currentElement = data[i];
			int temp = i;
			// 如果前一位大于当前元素，就将前一位向右移
			while (temp > 0 && data[temp - 1] > currentElement) {
				data[temp] = data[temp - 1];
				temp--;
			}
			data[temp] = currentElement; // 数据交换
		}

		System.out.println("sorted:");
		for (int j : data) {
			System.out.println(j);
		}
	}
}
