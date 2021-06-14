package chap1;

public class Demo17 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo17Thread();
        t.start();
//        Thread.sleep(10);
        t.interrupt();
    }
}

class Demo17Thread extends Thread{
    @Override
    public void run() {
        try {
            for (int i = 0; i < 10000; i++) {
                if (this.isInterrupted()) {
                    System.out.println("已经是停止状态了！我有退出");
//                break;
                    throw new InterruptedException();
                }
                System.out.println("i=" + i);
            }
            System.out.println("这里是结束循环后的代码");
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
