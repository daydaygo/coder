package chap4;

import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Demo07 {
    private static Lock lock = new ReentrantLock();
    private static Condition conditionA = lock.newCondition();
    private static Condition conditionB = lock.newCondition();
    private static Condition conditionC = lock.newCondition();
    volatile private static int nextPrintWho = 1;

    public static void main(String[] args) {
        Thread t1 = new Thread(){
            @Override
            public void run() {
                try{
                    lock.lock();
                    while (nextPrintWho != 1){
                        conditionA.await();
                    }
                    for (int i = 0; i < 3; i++) {
                        System.out.println("ThreadA打印第" + (i + 1) + "次");
                    }
                    nextPrintWho = 2;
                    conditionB.signalAll();
                }catch (InterruptedException e){
                    e.printStackTrace();
                }finally {
                    lock.unlock();
                }

            }
        };
        Thread t2 = new Thread(){
            @Override
            public void run() {
                try{
                    lock.lock();
                    while (nextPrintWho != 2){
                        conditionB.await();
                    }
                    for (int i = 0; i < 3; i++) {
                        System.out.println("ThreadB打印第" + (i + 1) + "次");
                    }
                    nextPrintWho = 3;
                    conditionC.signalAll();
                }catch (InterruptedException e){
                    e.printStackTrace();
                }finally {
                    lock.unlock();
                }

            }
        };
        Thread t3 = new Thread(){
            @Override
            public void run() {
                try{
                    lock.lock();
                    while (nextPrintWho != 3){
                        conditionC.await();
                    }
                    for (int i = 0; i < 3; i++) {
                        System.out.println("ThreadC打印第" + (i + 1) + "次");
                    }
                    nextPrintWho = 1;
                    conditionA.signalAll();
                }catch (InterruptedException e){
                    e.printStackTrace();
                }finally {
                    lock.unlock();
                }

            }
        };

        Thread[] threadsA = new Thread[5];
        Thread[] threadsB = new Thread[5];
        Thread[] threadsC = new Thread[5];
        for (int i = 0; i < threadsA.length; i++) {
            threadsA[i] = new Thread(t1);
            threadsB[i] = new Thread(t2);
            threadsC[i] = new Thread(t3);

            threadsA[i].start();
            threadsB[i].start();
            threadsC[i].start();
        }
    }
}
