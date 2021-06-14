package chap3;

public class Demo18 {
    public static void main(String[] args) {
        Demo18ThreadA t1 = new Demo18ThreadA();
        Demo18ThreadB t2 = new Demo18ThreadB(t1);
        t2.start();
        Demo18ThreadC t3 = new Demo18ThreadC(t1);
        t3.start();
    }
}
class Demo18ThreadA extends Thread{
    @Override
    public void run() {
        try {
            System.out.println("线程A开始于" + System.currentTimeMillis());
            Thread.sleep(5000);
            System.out.println("线程A结束于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    synchronized public void foo(){
        System.out.println("方法执行时间" + System.currentTimeMillis());
    }
}
class Demo18ThreadB extends Thread{
    private Demo18ThreadA t;
    public Demo18ThreadB(Demo18ThreadA t){
        this.t = t;
    }

    @Override
    public void run() {
        synchronized (t) {
            try {
                t.start();
//                Thread.sleep(6000);
                t.join();
                for (int i = 0; i < Integer.MAX_VALUE; i++) {
                    String s = new String();
                    Math.random();
                }
            }catch (InterruptedException e){
                e.printStackTrace();
            }
        }
    }
}
class Demo18ThreadC extends Thread{
    private Demo18ThreadA t;
    public Demo18ThreadC(Demo18ThreadA t){
        this.t = t;
    }

    @Override
    public void run() {
        t.foo();
    }
}
