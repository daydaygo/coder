package chap1;

public class Demo21 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo21Thread();
        t.start();
        Thread.sleep(100);
        t.suspend();
        System.out.println("main结束");
    }
}

class Demo21Thread extends Thread{
    private long i = 0;

    @Override
    public void run() {
        while(true){
            i++;
            System.out.println("i=" + i);
        }
    }
}
