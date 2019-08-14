package inheritance;

public class Manager extends Employee {
    private double bonus;

    /**
     * @param name the employee's name
     * @param salary the salary
     * @param year the hire year
     * @param month the hire month
     * @param day the hire day
     */

    public Manager(String name, double salary, int year, int month, int day) {
        super(name, salary, year, month, day);
        bonus = 0;
    }

    public double getSalary() {
        return super.getSalary() + bonus;
    }

    public void setBonus(double bonus) {
        this.bonus = bonus;
    }
}