package chap3;

public class Demo17 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo17Thread();
        t.start();
        t.join(2000);
        System.out.println("主线程结束于" + System.currentTimeMillis());
    }
}

class Demo17Thread extends Thread{
    @Override
    public void run() {
        try {
            System.out.println("子线程开始于" + System.currentTimeMillis());
            Thread.sleep(5000);
            System.out.println("子线程结束于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}