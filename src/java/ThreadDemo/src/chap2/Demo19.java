package chap2;

public class Demo19 {
    public static void main(String[] args) {
        Demo19Service service = new Demo19Service();
        Thread t1 = new Demo19ThreadA(service);
        t1.setName("A");
        t1.start();
        Thread t2 = new Demo19ThreadB(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo19Service{
    synchronized public static void foo1(){
        System.out.println(Thread.currentThread().getName() + "进入foo1方法在" + System.currentTimeMillis());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "结束foo1方法在" + System.currentTimeMillis());
    }

    synchronized public void foo2(){
        System.out.println(Thread.currentThread().getName() + "进入foo2方法在" + System.currentTimeMillis());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "结束foo2方法在" + System.currentTimeMillis());
    }
}

class Demo19ThreadA extends Thread{
    private Demo19Service service;
    public Demo19ThreadA(Demo19Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}

class Demo19ThreadB extends Thread{
    private Demo19Service service;
    public Demo19ThreadB(Demo19Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}
