package chapter16;

import java.sql.Date;

public final class Period {
    private final Date start;
    private final Date end;

    public Period(Date start, Date end) {
        // 保护性拷贝
        // 是在检查参数的有效性之前进行的，并且有效性检查是针对拷贝之后的对象，
        // 而不是针对原始的对象
        this.start = new Date(start.getTime());
        this.end = new Date(end.getTime());
        if (this.start.compareTo(this.end) > 0)
        throw new IllegalArgumentException(this.start + " after " + this.end);
    }

    public Date start() {
        return start;
    }

    public Date end() {
        return end;
    }
}