package chap2;

public class Demo08 {
    public static void main(String[] args) {
        Demo08ServiceB serviceB = new Demo08ServiceB();
        Thread t1 = new Demo08Thread(serviceB);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo08Thread(serviceB);
        t2.setName("B");
        t2.start();
    }
}

class Demo08ServiceA{
    synchronized public void foo(){
        try {
            System.out.println("父类：" + Thread.currentThread().getName() + "，开始于" + System.currentTimeMillis());
            Thread.sleep(5000);
            System.out.println("父类：" + Thread.currentThread().getName() + "，结束于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo08ServiceB extends Demo08ServiceA{
    public void foo(){
        try {
            System.out.println("子类：" + Thread.currentThread().getName() + "，开始于" + System.currentTimeMillis());
            Thread.sleep(5000);
            System.out.println("子类：" + Thread.currentThread().getName() + "，结束于" + System.currentTimeMillis());
            super.foo();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo08Thread extends Thread{
    Demo08ServiceB service;
    public Demo08Thread(Demo08ServiceB service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}