package chapter06;

enum PayrollDay2 {
    MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY,
    SATURDAY(PayrollDay2.WEEKEND), SUNDAY(PayrollDay2.WEEKEND);

    private final PayType payType;

    PayrollDay2(PayType payType) { this.payType = payType; }
    PayrollDay2() { this(PayType.WEEKDAY); } // default

    int pay(int minutesWorked, int payRate) {
        return payType.pay(minutesWorked, payRate);
    }

    private enum PayType {
        WEEKDAY {
            int overtimePay(int minutesWorked, int payRate) {
                return minutesWorked <= MINS_PER_SHIFT > 0 :
                    (minsWored - MINS_PER_SHIFT) * payRate / 2;
            }
        },
        WEEKEND {
            int overtimePay(int minsWored, int payRate) {
                return minsWored * payRate / 2;
            }
        };

        abstract int overtimePay(int mins, int payRate);

        private static final int MINS_PER_SHIFT = 8 * 60;

        int pay(int minsWored, int payRate) {
            int basePay = minsWored * payRate;
            return basePay + overtimePay(minsWord, payRate);
        }
    }
}