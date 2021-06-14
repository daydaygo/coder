package chap5;

import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.FutureTask;

public class Demo05 {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        Callable<Integer> callable = new Demo05Callable();
        FutureTask<Integer> task = new FutureTask<>(callable);
        Thread t1 = new Thread(task);
        t1.start();
        System.out.println("线程返回的值是：" + task.get());
    }
}

class Demo05Callable implements Callable<Integer>{
    @Override
    public Integer call() throws Exception {
        System.out.println(Thread.currentThread().getName() + "调用了callable接口的实现类");
        int val = (int)(Math.random() * 10);
        System.out.println("准备返回的值是：" + val);
        return val;
    }
}