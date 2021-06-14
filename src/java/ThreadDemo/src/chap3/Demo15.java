package chap3;

public class Demo15 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo15Thread();
        t.start();
        t.join();
        System.out.println("子线程执行完成以后再执行");
    }
}

class Demo15Thread extends Thread{
    @Override
    public void run() {
        try {
            int value = (int) (Math.random() * 10000);
            System.out.println("需要等待" + value + "毫秒");
            Thread.sleep(value);
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}