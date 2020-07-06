package chapter12;

public class ActivObjectFactory2 {
    public static ActiveObject2 createActiveObject() {
        return new ActiveObjectImpl();
    }
}