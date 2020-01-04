package chapter06;

// Program containing marker annotations
/**
 * Sample类有7个静态方法，其中4个被注解为测试。这4个中有2个抛出了异常：
 * m3和m7，另外两个则没有:m1和m5.但是其中一个没有抛出异常的被注解方法：m5，
 * 是一个实例方法，因此不属于注解的有效使用。总之，Sample包含4项测试：一项会
 * 通过，两项会失败，另一项无效。没有用Test注解进行标注的另外4个方法会被测试工
 * 具忽略。
 * Test注解对Sample类的语义没有直接的影响。它们只负责提供信息供相关的程序使用。
 * 更一般地讲，注解永远不会改变被注解代码的语义，但是使它可以通过工具进行特殊的
 * 处理。
 */
public class Sample {
    @Test public static void m1() {  }  // Test should pass
    public static void m2() {  }
    @Test public static void m3() {
        throw new RuntimeException("Boom");
    }

    public static void m4() {  }
    @Test public void m5() {  }  // INVALID USE: nonstatic method
    public static void m6() {  }
    @Test public static void m7() {
        throw new RuntimeException("Crash");
    }
    public static void m8() {  }

}