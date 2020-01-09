package chapter12;

import java.io.Serializable;
import java.util.Date;

private static class SerializationProxy implements Serializable {
    private final Date start;
    private final Date end;

    SerializationProxy(Period p) {
        this.start = p.start;
        this.end = p.end;
    }

    private static final long serialVersionUID = 2340982438485285L;
}