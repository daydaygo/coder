package chap3;

public class Demo20 {
    public static void main(String[] args) {
        ThreadLocal tl = new ThreadLocal();
        Thread t1 = new Demo20ThreadA(tl);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo20ThreadA(tl);
        t2.setName("B");
        t2.start();
    }
}

class Demo20ThreadA extends Thread{
    private ThreadLocal tl;
    public Demo20ThreadA(ThreadLocal tl){
        this.tl = tl;
    }

    @Override
    public void run() {
        try {
            for (int i = 0; i < 100; i++) {
                tl.set(Thread.currentThread().getName() + (i + 1));
                System.out.println(Thread.currentThread().getName() + " get value = " + tl.get());
                Thread.sleep(100);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
