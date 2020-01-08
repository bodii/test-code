package chapter12;

import java.sql.Date;

public class TestMutablePeriod {
    public static void main(String[] args) {
        MutablePeriod mp = new MutablePeriod();
        Period p = mp.Period;
        Date pEnd = mp.end;

        pEnd.setYear(78);
        System.out.println(p);

        pEnd.setYear(69);
        System.out.println(p);
    }
}