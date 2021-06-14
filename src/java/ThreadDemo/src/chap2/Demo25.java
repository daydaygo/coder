package chap2;

public class Demo25 {
    public static void main(String[] args) throws InterruptedException {
        Demo25User user = new Demo25User();
        Demo25Service service = new Demo25Service();

        Thread t1 = new Demo25Thread(service, user);
        t1.setName("A");
        t1.start();

        Thread.sleep(50);

        Thread t2 = new Demo25Thread(service, user);
        t2.setName("B");
        t2.start();
    }
}

class Demo25User{
    public String username;
    public String password;
}

class Demo25Service{
    public void foo(Demo25User user){
        try {
            synchronized (user) {
                System.out.println(Thread.currentThread().getName() + "开始" + System.currentTimeMillis());
                user.username = "b";
                user.password = "bb";
                Thread.sleep(2000);
                System.out.println(Thread.currentThread().getName() + "结束" + System.currentTimeMillis());
            }

        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo25Thread extends Thread{
    private Demo25Service service;
    private Demo25User user;

    public Demo25Thread(Demo25Service service, Demo25User user){
        this.service = service;
        this.user = user;
    }

    @Override
    public void run() {
        service.foo(user);
    }
}