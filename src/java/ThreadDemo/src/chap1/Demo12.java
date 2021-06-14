package chap1;

public class Demo12 {
    public static void main(String[] args) throws InterruptedException {
        Demo12Thread t = new Demo12Thread();
        t.start();
        Thread.sleep(2000);
        t.stopThread();
    }
}

class Demo12Thread extends Thread{
    private boolean flag = true;

    @Override
    public void run() {
        try {
            while (flag) {
                System.out.println("time=" + System.currentTimeMillis());
                Thread.sleep(1000);
            }
            System.out.println("线程执行结束");
        }catch(InterruptedException e){
            e.printStackTrace();
        }
    }

    public void stopThread(){
        flag = false;
    }
}