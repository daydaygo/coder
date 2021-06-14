package chap2;

public class Demo23 {
    public static void main(String[] args) throws InterruptedException {
        Demo23Thread t = new Demo23Thread();
        t.setFlag("a");

        Thread t1 = new Thread(t);
        t1.start();
        Thread.sleep(10);

        t.setFlag("b");
        Thread t2 = new Thread(t);
        t2.start();
    }
}

class Demo23Thread extends Thread{
    private String flag;    // 标志，控制代码以什么样的方式运行
    private Object lock1 = new Object();
    private Object lock2 = new Object();

    public void setFlag(String flag){
        this.flag = flag;
    }

    @Override
    public void run() {
        try {
            if ("a".equals(flag)) {
                synchronized (lock1) {
                    System.out.println("flag=" + flag);
                    Thread.sleep(3000);
                    synchronized (lock2) {
                        System.out.println("按lock1=>lock2的顺序执行");
                    }
                }
            } else {
                synchronized (lock2){
                    System.out.println("flag=" + flag);
                    Thread.sleep(3000);
                    synchronized (lock1){
                        System.out.println("按lock2->lock1的顺序执行");
                    }
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
