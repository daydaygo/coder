package chap2;

public class Demo04 {
    public static void main(String[] args) {
        Demo04Service service = new Demo04Service();
        Thread t1 = new Demo04ThreadA(service);
        Thread t2 = new Demo04ThreadB(service);
        t1.start();
        t2.start();
    }
}

class Demo04Service {
    synchronized public void foo1(){
        System.out.println("开始运行foo1方法，threadName：" + Thread.currentThread().getName());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("foo1方法运行结束");
    }

    public void foo2(){
        System.out.println("开始运行foo2方法，threadName：" + Thread.currentThread().getName());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("foo2方法运行结束");
    }
}

class Demo04ThreadA extends Thread{
    private Demo04Service service;
    public Demo04ThreadA(Demo04Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}

class Demo04ThreadB extends Thread{
    private Demo04Service service;
    public Demo04ThreadB(Demo04Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}
