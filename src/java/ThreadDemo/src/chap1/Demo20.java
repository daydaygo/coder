package chap1;

public class Demo20 {
    public static void main(String[] args) throws InterruptedException {
        Demo20Service service = new Demo20Service();
        Thread t1 = new Thread(){
            @Override
            public void run() {
                service.printString();
            }
        };
        t1.setName("A");    // setName修改线程的名称，这里就是把Thread-0修改为A
        t1.start();
        Thread.sleep(10);

        Thread t2 = new Thread(){
            @Override
            public void run() {
                service.printString();
            }
        };
        t2.start();
    }
}

class Demo20Service {
    synchronized public void printString(){
        System.out.println("线程开始");
        if ("A".equals(Thread.currentThread().getName())){
            System.out.println("A线程永远suspend了");
            Thread.currentThread().suspend();
        }
        System.out.println("线程结束");
    }
}
