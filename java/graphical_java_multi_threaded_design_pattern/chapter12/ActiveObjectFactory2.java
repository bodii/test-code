package chapter12;

public class ActiveObjectFactory2 {
    public static ActiveObject2 createActiveObject() {
        return new ActiveObjectImpl();
    }
}