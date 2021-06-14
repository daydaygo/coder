package chap2;

public class Demo05 {
    public static void main(String[] args) throws InterruptedException {
        Demo05User user = new Demo05User();
        Thread t = new Demo05Thread(user);
        t.start();
        Thread.sleep(200);
        user.getValue();
    }
}

class Demo05User{
    private String username = "a";
    private String password = "aa";

    synchronized public void setUsernameAndPassword(String username, String password){
        this.username = username;
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        this.password = password;
        System.out.println("setUsernameAndPassword方法，线程名称：" + Thread.currentThread().getName() +
                "，username=" +username + ", password=" + password);
    }

    synchronized public void getValue(){
        System.out.println("getValue方法，线程名称" + Thread.currentThread().getName() +
                ", username=" + username + ", password=" + password );
    }
}

class Demo05Thread extends Thread{
    private Demo05User user;
    public Demo05Thread(Demo05User user){
        this.user = user;
    }

    @Override
    public void run() {
        user.setUsernameAndPassword("B", "BB");
    }
}