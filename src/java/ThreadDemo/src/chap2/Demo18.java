package chap2;

public class Demo18 {
    public static void main(String[] args) {
        Thread t1 = new Demo18ThreadA();
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo18ThreadB();
        t2.setName("B");
        t2.start();
    }
}

class Demo18Service{
    synchronized public static void foo1(){
        System.out.println(Thread.currentThread().getName() + "进入方法foo1在" + System.currentTimeMillis());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "结束方法foo1在" + System.currentTimeMillis());
    }
    synchronized public static void foo2(){
        System.out.println(Thread.currentThread().getName() + "进入方法foo2在" + System.currentTimeMillis());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "结束方法foo2在" + System.currentTimeMillis());
    }
}

class Demo18ThreadA extends Thread{
    @Override
    public void run() {
        Demo18Service.foo1();
    }
}

class Demo18ThreadB extends Thread{
    @Override
    public void run() {
        Demo18Service.foo2();
    }
}
