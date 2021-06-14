package chap2;

public class Demo24 {
    public static void main(String[] args) throws InterruptedException {
        Demo24Service service = new Demo24Service();
        Thread t1 = new Demo24Thread(service);
        t1.setName("A");
        t1.start();

        Thread.sleep(50);

        Thread t2 = new Demo24Thread(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo24Service{
    private String lockObject = "123";

    public void foo(){
        try {
            synchronized (lockObject) {
                System.out.println(Thread.currentThread().getName() + "开始于" + System.currentTimeMillis());
                lockObject = "456";
                Thread.sleep(2000);
                System.out.println(Thread.currentThread().getName() + "结束于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo24Thread extends Thread{
    private Demo24Service service;
    public Demo24Thread(Demo24Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}