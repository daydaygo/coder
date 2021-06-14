package chap1;

public class Demo11 {
    public static void main(String[] args) {
        Thread t = Thread.currentThread();
        System.out.println(t.getName() + ", " + t.getId());
    }
}
