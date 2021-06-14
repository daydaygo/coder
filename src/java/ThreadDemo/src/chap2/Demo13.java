package chap2;

public class Demo13 {
    public static void main(String[] args) throws InterruptedException {
        Demo13Service service = new Demo13Service();
        Thread t1 = new Demo13ThreadA(service);
        t1.start();
        Thread.sleep(10);
        Thread t2 = new Demo13ThreadB(service);
        t2.start();
    }
}

class Demo13Service{
    public void foo1(){
        try {
            synchronized (this) {
                System.out.println("foo1开始于" + System.currentTimeMillis());
                Thread.sleep(2000);
                System.out.println("foo1结束于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    public void foo2(){
        synchronized (this){
            System.out.println("foo2开始于" + System.currentTimeMillis());
            System.out.println("foo2结束于" + System.currentTimeMillis());
        }
    }
}

class Demo13ThreadA extends Thread{
    private Demo13Service service;
    public Demo13ThreadA(Demo13Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}

class Demo13ThreadB extends Thread{
    private Demo13Service service;
    public Demo13ThreadB(Demo13Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}