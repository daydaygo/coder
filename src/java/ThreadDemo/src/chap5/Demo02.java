package chap5;

public class Demo02 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo02Thread();
        t.start();
        Thread.sleep(1000);
        System.out.println("线程在main方法中的状态：" + t.getState());
    }
}

class Demo02Thread extends Thread{
    @Override
    public void run() {
        try {
            System.out.println("线程准备进入sleep状态");
            Thread.sleep(10000);
            System.out.println("线程结束sleep状态");
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}