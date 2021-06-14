package chap2;

public class Demo15 {
    public static void main(String[] args) {
        Demo15Service service = new Demo15Service();
        Thread t1 = new Demo15Thread(service);
        t1.setName("A");
        t1.start();
        Thread t2 = new Demo15Thread(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo15Service{
    private Object lockObject = new Object();

    public void foo(){
        try {
            Object objectLock = new Object();
            synchronized (objectLock) {
                System.out.println(Thread.currentThread().getName() + "开始于" + System.currentTimeMillis());
                Thread.sleep(2000);
                System.out.println(Thread.currentThread().getName() + "结束于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo15Thread extends Thread{
    private Demo15Service service;
    public Demo15Thread(Demo15Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}
