package chap2;

public class Demo10 {
    public static void main(String[] args) {
        Demo10Service service = new Demo10Service();
        Thread t1 = new Demo10Thread(service);
        t1.setName("A");
        t1.start();
        Thread t2 = new Demo10Thread(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo10Service{
    public void foo(){
        try {
            synchronized (this) {
                System.out.println(Thread.currentThread().getName() + "开始于" + System.currentTimeMillis());
                Thread.sleep(2000);
                System.out.println(Thread.currentThread().getName() + "结束于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
class Demo10Thread extends Thread{
    private Demo10Service service;
    public Demo10Thread(Demo10Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}
