package chapter12;

import java.io.Serializable;

/**
 * 即使你确定了默认的序列化形式是合适的，通常还必须提供一个readObject方法以
 * 保证约束关系和安全性。
 * 对于这个Name类而言，readObject方法必须确保lastName和firstName是非null的。
 */
public class Name implements Serializable {
    /**
     * Last name. Must be non-null.
     * @serial
     */
     private final String lastName;

     /**
      * First name. Must be non-null.
      * @serial
      */
     private final String firstName;

     /**
      * Middle name, or null if there is none.
      * @serial
      */
     private final String middleName;
}