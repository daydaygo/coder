package chap2;

public class Demo22 {
    public static void main(String[] args) {
        Demo22Service service = new Demo22Service();
        Thread t1 = new Demo22ThreadA(service);
        t1.start();
        Thread t2 = new Demo22ThreadB(service);
        t2.start();
    }
}

class Demo22Service{
    private Object lockObject1 = new Object();
    /*synchronized */public void foo1(){
        synchronized (lockObject1) {
            System.out.println("foo1方法开始执行");
            boolean isContinue = true;
            while (isContinue) {

            }
            System.out.println("foo2方法执行结束");
        }
    }
    private Object lockObject2 = new Object();
    /*synchronized */public void foo2(){
        synchronized (lockObject2) {
            System.out.println("foo2方法开始执行");
            System.out.println("foo2方法执行结束");
        }
    }
}

class Demo22ThreadA extends Thread{
    private Demo22Service service;
    public Demo22ThreadA(Demo22Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}

class Demo22ThreadB extends Thread{
    private Demo22Service service;
    public Demo22ThreadB(Demo22Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}
