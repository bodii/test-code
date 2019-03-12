public class Dome {

    public static void main(String[] args) {
        SequenceList s1 = new SequenceList();
        s1.add("Jack");
        s1.add("John");
        System.out.println(s1.get(0));
        System.out.println(s1.isEmpty());
        System.out.println(s1.size());
        // s1.clear();
        // System.out.println(s1.size());
        s1.set(0, "Lile");
        System.out.println(s1.get(0));
        System.out.println(s1.get(1));
        System.out.println(s1.size());
    }
}