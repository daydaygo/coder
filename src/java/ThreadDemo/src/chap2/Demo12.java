package chap2;

public class Demo12 {
    public static void main(String[] args) {
        Demo12Service service = new Demo12Service();
        Thread t1 = new Demo12Thread(service);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo12Thread(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo12Service {
    public void foo(){
        try{
            for (int i = 0; i < 100; i++) {
                System.out.println("非同步线程" + Thread.currentThread().getName() + ", i=" + i);
                Thread.sleep(10);
            }
            System.out.println();
            synchronized (this){
                for (int i = 0; i < 100; i++) {
                    System.out.println("同步线程" + Thread.currentThread().getName() + ", i=" + i);
                    Thread.sleep(10);
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo12Thread extends Thread{
    private Demo12Service service;
    public Demo12Thread(Demo12Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}