package chap1;

public class Demo13 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo13Thread();
        t.start();
        Thread.sleep(2000);
        t.stop();
    }
}

class Demo13Thread extends Thread{
    @Override
    public void run() {
        try {
            while (true) {
                System.out.println("time=" + System.currentTimeMillis());
                Thread.sleep(1000);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }catch (ThreadDeath e){
            System.out.println("进入catch块");
            e.printStackTrace();
        }
    }
}