package lambda;

import javax.swing.Timer;

class Greeter {
    public void greet() {
        System.out.println("Hello, world!");
    }
}

class TimedGreeter extends Greeter {
    public void greet() {
        Timer t = new Timer(1000, super::greet);
        t.start();
    }
}