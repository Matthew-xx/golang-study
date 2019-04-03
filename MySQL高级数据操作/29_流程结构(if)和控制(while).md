# 流程结构if和流程控制while

代码的执行顺序

## if分支

基本语法：

1.用在select，当做条件进行判断
if(条件，为真结果，为假结果)

    select *,if(stu_age>20,'符合','不符合')as judge from my_student;
结果

    | stu_id  | stu_name | class_id | stu_age | stu_height | gender | judge     |
    +---------+----------+----------+---------+------------+--------+-----------+
    | stu0001 | 夏洛     |        1 |      18 |        185 | 男     | 不符合    |
    | stu0002 | 张四     |        1 |      28 |        165 | 女     | 符合      |
    | stu0003 | 张五     |        2 |      22 |        187 | 男     | 符合      |
    | stu0004 | 小婷     |        2 |      32 |        187 | 女     | 符合      |
    | stu0005 | 小猪     |        1 |      30 |        173 | 女     | 符合      |
    | stu0006 | 小狗     |        2 |      18 |        170 | 男     | 不符合    |
    | stu0007 | 小江     |        6 |      25 |        178 | 女     | 符合      |

2.用在复杂的语句块中（函数、存储过程、触发器）

if 条件表达式 then
    满足条件要执行的语句；
end if;

if 条件表达式 then
    满足条件要执行的语句；
else
    满足条件要执行的语句；
    if 条件表达式 then
        满足条件要执行的语句；
    end if;
end if;

## while循环

while 条件 do
     要循环执行的代码;
end while;

结构标识符：为某些特定的结构命名，以便在某些地方使用它

[标识名字:]while 条件 do
     要循环执行的代码;
end while[标识名字];

一般标识符的存在，是为了在循环体重使用循环控制
iterate相当于其他编程语言中的continue；
leave 相当于break；

[标识名字:]while 条件 do
     if 条件判断 then
         iterate [标识名字]
      end if;
     循环体
end while[标识名字];