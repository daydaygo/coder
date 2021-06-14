# c++ 基础

> C++Primer.Plus（第6版）中文版
> c++ manual

## mark

- 11/20/23

```cpp
unordered_map
thread/mutex/cond_var/atomic
lambda/std::function
shared_ptr/unique_ptr
regex

coroutines
module
std::networking
std::file_system
std::executor
```

## 基础

常用扩展名: cpp
兼容 c 语法
c++ 使用 c 头文件: `stdlib.h -> cstlib`
变量声明: 有助于编译器检查错误
类型: 有符号, 溢出, bool
引号: 单引号 -> char, 双引号 -> string(需要 '\0')

struct: 多个数据的集合
union: 集合中多个数据中的一个
enum: 符号常量的集合

语法块: 类似函数, 在里面定义的变量外部不可用
循环&分支: 不使用 {} 只对下一行语句生效

运算符优先级: 逻辑(&& ||) < 关系(> = <) < (!)

```cpp
// hello.cpp
#include "iostream"             // cin + cout
#define INT_MAX 327676

int main(){
    using namespace std;        // namespace
    cout<<"hello world"<<endl;  // 运算符重载
    cout.setf(ios::left);       // 感觉比 printf 复杂

    const int n = 10;

    // cin.get();               // keep the window open
    cin.getline(name, 20);      // 面向行的输入
    while(!cin.fail())          // test for EOF
    while(cin)

    // return 0;                // don't need
}
```

```cpp
// namespace
using namespace std;
cout<<"hello"<<endl;

using std::cout;
cout<<"hello"<<std::endl;

// char
cout<<'\a'<<'\032'; // 可以使用 a 的 ASCII 的不同进制表示
cout<<"\u00f6";     // Unicode
wchar_t = L'p';     // 宽字符类型
char16_t / char32_t

// char 0-9 转 int
n = (int)ch - 48;
n = ch - '0';

// string
#include "string"
#include "cstring"      // c-string lib
str.size();             // strlen()

// window 下暂停
#include "cstdlib"
system("pause");

// 读取文件
#include "fstream"
fstream cin("/src/test.txt");
cin>>a>>b>>c;

// math
#include "cmath"
floor() / ceil() / sqrt()

// 数组
int a[SIZE] = {1};                 // 默认初始化为全0; 使用 const 值表示数组元素个数
#include "vector"               // vector
vector<int> vi;
#include "array"                // array
array<int,5> ai;

// struct: 包含多种数据
struct inflatable{
    char name[20];
    float volume;
    double price;
}
inflatable guest = {"czl", 1.88, 29.99};

// 指针
int* p; int * p; int *p;        // 空格在哪都可以
int *p; p = &num;               // 初始化
int *p = &num;
const int* p = &num;            // 可以修改 num, 但是不能修改 p
int *p = new int; *p = 101;     // 分配内存
int *p = arr;                   // 指针 + 数组
int *p = &arr[0];
inflatable *p = &guest;         // 指针 + struct
p->name;

// 基于范围的 for 循环, 主要在 STL 中使用
int arr[5] = {1,2,3,4,5};
for(int x : arr) cout<<x;
for(int &x : arr) x += 10;

if(3==num)                      // 让编译器发现错误
```

## 函数

返回值: void
参数: void, 指针, 可变数量
参数按值传递: 函数被调用时, 为参数创建新的变量, 函数执行完后释放内存
参数按引用传递: 是否加 & 符就看是否值是否为指针
递归调用: 将复杂的工作分解
函数指针
内联函数: 使用 #define 实现, 较普通函数效率高
函数重载: 包含不同参数 但 完成相近功能

```cpp
// 函数 + 数组: 传递过来的是指针
void printArr(int* a, int n)
void printArr(int a[], int n)
void printArr(const int* a, int n)          // 保护数组
void printArr(const int* begin, int* end)   // 使用数组范围
cout<<a[i];
cout<<*(a+i);
// 函数 + 结构体

int & a = b;                                // 引用变量, 别名
int* const a = &b;

int func(int i, int j=1, int k=2)           // 默认参数

// sort
int cmp(const int &a, const int &b){ return a<b;}
sort(arr, arr+n, cmp);

// reverse 函数实现
void reverse2(char* str){
    char* end = str;
    char tmp;
    if(str)
        while(*end) end++;      // 遍历到字符串末尾
    end--;                      // 回退, 最后一个是 null / '\0'

    while(str<end){             // 首尾互换
        tmp = *str;
        *str++ = *end;
        *end-- = tmp;
    }
}
```

## vector(向量 call-by-rank)

数组: 最简单的数据结构, 但是大小固定导致很多场景无法使用

构造: 默认 + 基于复制; `_size + _capacity + _item[]`
load factor = size/capacity, 0.25-0.7 来考虑 expand()/shrink()
常规向量(无序)
有序向量: 判断有序; 去重; 二分查找(查找长度不均衡, 3分支 -> 2分支); fib查找(黄金分割); 有重复元素的查找顺序

排序树: 通过分支数量来衡量算法复杂度(logN 就是这样分析来的)

```cpp
// 二分支 二分查找
int binnarySearch(int a[], int n, int elem){
    int low=0, high=n-1, mid;
    while(1<high-low) {
        mid = (low+high)>>1;
        (elem<a[mid]) ? high=mid : low=mid;
    }
    return (elem==a[low]) ? low : -1; // 这里可以轻松实现 elem 在数组中的序列
}

// 返回 秩 最大者
int binnarySearch(int a[], int n, int elem){
    int low=0, high=n-1, mid;
    while(low<high) {
        mid = (low+high)>>1;
        (elem<a[mid]) ? high=mid : low=mid+1;
    }
    return --low;
}
```

- c++ 简单模板

```cpp
#include <iostream>
#include <cstdio> // cin cout 好用, 但部分题目使用 cin cout 会超时, 比如 1852
using namespace std;
#define MAX_N 1000000 // 如果需要使用数组

// 解题方法
int a,b, arr[MAX_N];
void solve() {
    cout << a+b <<endl;
}

// main() 处理输入
int main()
{
    freopen("in.txt","r",stdin); // 输入重定向到文件, 方便测试
    scanf("%d%d", &a, &b);
    solve();
    return 0;
}
```

- c++ 基础知识/常用结构/算法

```cpp
// stack

// queue

// dfs: 递归
void dfs(Node root){
    if(root == NULL) return;
    root.visited = true;
    for(Node x : root.next)
        if(!x.visited)
            dfs(x);
}

// bfs: 非递归
void bfs(Node root){
    Queue queue = new Queue();
    root.visited = true;
    queue.enqueue(root);

    while(!queue.empty()) {
        Node r = queue.dequeue();
        for(Node x : r.next){
            if(!x.visited){
                x.visited = true;
                queue.enqueue(x);
            }
        }
    }
}

// 算法
#include <algorithm>
int a[] = {1, 3, 2};
sort(a, a+3); // quick sort
bool f = binary_search(a, a+3, 4); // binary search
do {
    // todo
} while(next_permutation(perm, perm + n)); // next_permutation() return false if end
```
