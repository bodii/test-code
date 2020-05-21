package chapter18;

import java.io.Serializable;

public class SerializableNameTest implements Serializable {

    /**
     * Last name. Must be non-null
     * @serial
     */
    private final String lastName;

    /**
     * First name. Must be non-null
     * @serial
     */
    private final String firstName;

    /**
     * Middle name, or null if three is nono.
     * @serial
     */
    private final String middleName;
}