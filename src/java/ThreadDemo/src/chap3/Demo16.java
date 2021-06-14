package chap3;

public class Demo16 {
    public static void main(String[] args) throws InterruptedException {
        Demo16ThreadB t = new Demo16ThreadB();
        t.start();
        Thread.sleep(10);
        Thread t2 = new Demo16ThreadC(t);
        t2.start();
    }
}

class Demo16ThreadA extends Thread{
    @Override
    public void run() {
        for (int i = 0; i < Integer.MAX_VALUE; i++) {
            // 耗时操作
            String s = new String();
            Math.random();
        }
    }
}
class Demo16ThreadB extends Thread{
    @Override
    public void run() {
        try {
            Thread t = new Demo16ThreadA();
            t.start();
            t.join();
            System.out.println("B线程正常结束");
        }catch (InterruptedException e){
            System.out.println("B线程异常结束");
            e.printStackTrace();
        }
    }
}
class Demo16ThreadC extends Thread{
    private Demo16ThreadB t;
    public Demo16ThreadC(Demo16ThreadB t){
        this.t = t;
    }

    @Override
    public void run() {
        t.interrupt();
    }
}