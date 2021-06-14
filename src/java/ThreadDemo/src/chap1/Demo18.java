package chap1;

public class Demo18 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo18Thread();
        t.start();
        Thread.sleep(10);
        t.interrupt();
    }
}

class Demo18Thread extends Thread{
    @Override
    public void run() {
        try {
            for (int i = 0; i < Integer.MAX_VALUE; i++) {
                String s = new String();
            }
            System.out.println("开始线程");
            Thread.sleep(20000);
            System.out.println("结束线程");
        }catch (InterruptedException e){
            System.out.println("异常进入到catch代码块");
            e.printStackTrace();
        }
    }
}