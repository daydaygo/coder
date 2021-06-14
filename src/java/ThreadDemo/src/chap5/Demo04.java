package chap5;

public class Demo04 {
    public static void main(String[] args) throws InterruptedException {
        Demo04Service service = new Demo04Service();
        Thread t1 = new Demo04Thread(service);
        t1.setName("A");
        Thread t2 = new Demo04Thread(service);
        t2.setName("B");
        t1.start();

        Thread.sleep(10);
        t2.start();
        System.out.println("线程B在main方法中的状态：" + t2.getState());
    }
}

class Demo04Service{
    synchronized static public void foo(){
        try {
            System.out.println(Thread.currentThread().getName() + "调用了foo方法");
            Thread.sleep(100000);
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo04Thread extends Thread{
    private Demo04Service service;
    public Demo04Thread(Demo04Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}