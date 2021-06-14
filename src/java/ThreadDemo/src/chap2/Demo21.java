package chap2;

public class Demo21 {
    public static void main(String[] args) {
        Thread t1 = new Demo21ThreadA();
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo21ThreadB();
        t2.setName("B");
        t2.start();
    }
}

class Demo21Service{
    public static void foo1(String lockObject){
        try {
            synchronized (lockObject) {
                while (true) {
                    System.out.println("线程" + Thread.currentThread().getName());
                    Thread.sleep(1000);
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
    public static void foo2(Object lockObject){
        try {
            synchronized (lockObject) {
                while (true) {
                    System.out.println("线程" + Thread.currentThread().getName());
                    Thread.sleep(1000);
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo21ThreadA extends Thread{
    @Override
    public void run() {
        Demo21Service.foo1("AA");
    }
}

class Demo21ThreadB extends Thread{
    @Override
    public void run() {
        Demo21Service.foo2("AA");
    }
}