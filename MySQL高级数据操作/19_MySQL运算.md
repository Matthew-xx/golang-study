# MySQL运算

## 算数运算

常见的算术运算+，-，*，/,%（加，减，乘，除，取余）

其中除法可使用 A **DIV** B = A/B

取余可使用 A **MOD** B = A%B

    create table ysf1(
    int_1 int,
    int_2 int,
    int_3 int,
    int_4 int)charset utf8;
    insert into ysf1 values(100,-100,0,default);
    select int_1+int_2,int_1-int_2,int_1*int_2,int_1/int_2,int_2/int_3,int_2%6,int_4/5 from ysf1;

    |           0 |         200 |      -10000 |     -1.0000 |        NULL |      -4|    NULL |
注意：

1. 除法运算结果，为浮点数类型
2. 除数为0的结果为NULL
3. NULL与任何数任何运算结果都为NULL
4. 默认MOD为取余，非取模

## 比较运算

>,>=,<,<=,=,<>(不等于)

用于条件中进行限定结果

1.mysql中没有==，可以用=或者<=>来进行相等判断

    select * from my_student where stu_age >=20;

特殊应用：在字段结果中进行比较运算

    select '1'<=>1,0.02<=>0,0.02<>0; -- select 没有规定只能用于表数据

    | '1'<=>1 | 0.02<=>0 | 0.02<>0 |
    +---------+--------------------+
    |       1 |        0 |       1 |

1.mysql中，数据会先自动转换成同类型，在进行比较
2.mysql中，没有bool值，1代表true，0代表false

### 特殊的比较语句

between and

    select * from my_student where stu_age between 20 and 30;
使用注意：

1. 是闭区间
2. 条件1需要小于条件2

## 逻辑运算

and \or\not  与或非

    select * from my_student where stu_age>=20 and stu_age<=30;

## in运算符

in用来替代=，当结果不是一个值，而是一个结果集的时候

基本语法：in(结果1，结果2，结果3，...)

前面是条件值，in里面是结果值得集合，结果是筛选在符合=运算的数据记录

    select * from my_student where stu_id in('stu0001','stu0002','stu0007');

## is运算符

is用来专门判断字段是否为null的运算符

基本语法：is null/is not null;

    select * from my_int where int_6=null;--查不到数据
    select * from my_int where int_6 is null;--查不到数据

    | int_8 | int_7 | int_6 | int_1 | int_2 | int_3  | int_4    | int_5      |
    +-------+-------+-------+-------+-------+--------+----------+------------+
    |  NULL |  NULL |  NULL |    10 |  1000 | 100000 | 10000000 | 1000000000 |

## like运算符

like用来进行模糊匹配字符串

基本语法：like '匹配模式'

匹配模式中
_表示匹配对应单个字符
%表示匹配对应多个字符

    select * from my_student where stu_name like '小%';