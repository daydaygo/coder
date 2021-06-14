package chap1;

public class Demo26 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo26Thread();
        t.setDaemon(true);  // 把指定的线程设置为守护线程
        t.start();
        Thread.sleep(3000);
        System.out.println("主线程结束");
    }
}

class Demo26Thread extends Thread{
    @Override
    public void run() {
        try {
            while (true) {
                System.out.println("time=" + System.currentTimeMillis());
                Thread.sleep(1000);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
