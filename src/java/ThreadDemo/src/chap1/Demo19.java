package chap1;

public class Demo19 {
    public static void main(String[] args) throws InterruptedException {
        Demo19Thread t = new Demo19Thread();
        t.start();
        Thread.sleep(1000);

        t.suspend();    // 暂停线程
        System.out.println("A=" + System.currentTimeMillis() + ", i=" + t.getI());
        Thread.sleep(1000);
        System.out.println("A=" + System.currentTimeMillis() + ", i=" + t.getI());

        t.resume();     // 恢复暂停线程运行
        Thread.sleep(1000);

        t.suspend();
        System.out.println("B=" + System.currentTimeMillis() + ", i=" + t.getI());
        Thread.sleep(1000);
        System.out.println("B=" + System.currentTimeMillis() + ", i=" + t.getI());
    }
}

class Demo19Thread extends Thread{
    private long i = 0;

    public long getI(){
        return i;
    }

    public void setI(long i){
        this.i = i;
    }

    @Override
    public void run() {
        while (true){
            i++;
        }
    }
}
