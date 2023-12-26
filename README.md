<div align="center">
  <h1>python 基础知识点</h1>
  <p>对python3及部分实践</p>
</div>

## 目录

- [目录](#目录)
- [数据类型转换](#数据类型转换)
- [Example](#example)
- [MD 数学公式](#md-数学公式)
- [提交规范](#提交规范)
- [参考链接](#参考链接)

## 数据类型转换

| 函数                   | 描述                                                  |
| ---------------------- | ----------------------------------------------------- |
| int(x [,base])         | 将 x 转换为一个整数                                   |
| float(x)               | 将 x 转换到一个浮点数                                 |
| complex(real [,image]) | 创建一个复数                                          |
| str(x)                 | 将对象 x 转换为字符串                                 |
| repr(x)                | 将对象 x 转换为表达式字符串                           |
| eval(str)              | 用来计算在字符串中的有效 Python 表达式,并返回一个对象 |
| tuple(s)               | 将序列 s 转换为一个元组                               |
| list(s)                | 将序列 s 转换为一个列表                               |
| set(s)                 | 转换为可变集合                                        |
| dict(d)                | 创建一个字典。d 必须是一个 (key, value)元组序列。     |
| frozenset(s)           | 转换为不可变集合                                      |
| chr(x)                 | 将一个整数转换为一个字符                              |
| ord(x)                 | 将一个字符转换为它的整数值                            |
| hex(x)                 | 将一个整数转换为一个十六进制字符串                    |
| oct(x)                 | 将一个整数转换为一个八进制字符串                      |

## Example

- [01.Hello World 实例](./examples/hello_world.py)
- [02.数字求和](./examples/sum.py)
- [03.平方根](./examples/square_root.py)
- [04.一元二次方程](./examples/quadratic_equation.py)
- [05.三角形面积](./examples/triangle_area.py)
- [06.圆形面积](./examples/circular_area.py)
- [07.随机生成数字](./examples/random_number.py)
- [08.交换变量](./examples/swap_variables.py)
- [09.if 语句](./examples/if_example.py)
- [10.判断奇数偶数](./examples/odd_even.py)
- [11.判断闰年](./examples/leap_year.py)
- [12.质数判断](./examples/prime_number.py)
- [13.输出指定范围内的素数](./examples/prime_number_intervals.py)
- [14.阶乘计算](./examples/factorial.py)
- [15.九九乘法表](./examples/99_table.py)
- [16.斐波那契数列](./examples/fibonacci_sequence.py)
- [17.阿姆斯特朗数](./examples/armstrong_number.py)
- [18.十进制转二进制、八进制、十六进制](./examples/base_conversion.py)
- [19.ASCII 码与字符相互转换](./examples/ascii_character.py)
- [20.最大公约数算法](./examples/hcf.py)
- [21.简单计算器实现](./examples/calculator.py)
- [22.生成日历](./examples/calendar_sample.py)
- [23.使用递归斐波那契数列](./examples/fibonacci_recursion.py)
- [24.文件 IO](./examples/file_io.py)
- [25.字符串判断](./examples/check_string.py)
- [26.字符串大小写转换](./examples/upper_lower.py)
- [27.计算每个月天数](./examples/month_days.py)
- [28.获取昨天日期](./examples/get_yesterday.py)
- [29.约瑟夫生者死者小游戏](./examples/joseph_life_dead_game.py)
- [30.五人分鱼](./examples/five_fish.py)
- [31.秒表功能](./examples/simple_stop_watch.py)
- [32.计算 n 个自然数的立方和](./examples/cube_sum.py)
- [33.计算数组元素之和](./examples/sum_array.py)
- [34.数组翻转指定个数的元素](./examples/array_rotation.py)
- [35.列表中的头尾两个元素对调](./examples/list_interchange.py)
- [36.列表中的指定位置的两个元素对调](./examples/list_swap_two_elements.py)
- [37.翻转列表](./examples/reversing_list.py)
- [38.判断元素是否在列表中存在](./examples/check_element_exists_in_list.py)
- [39.清空列表](./examples/clear_list.py)
- [40.移除列表中重复的元素](./examples/remove_duplicate_from_list.py)
- [41.计算元素在列表中出现的次数](./examples/count_occurrences_element_list.py)
- [42.计算列表元素之和、积](./examples/sum_list.py)

## MD 数学公式

1. markdown 实现公式内实现换行（简单版）`\\`
2. md Latex 中的空格。（a 符号 b,ab 为两个字符）
   | 符号 | 宽度 |
   | ---- | ---- |
   |a \qquad b|两个 m 的宽度|
   |a \quad b|一个 m 的宽度|
   |a\ b|1/3m 宽度|
   |a\;b|2/7m 宽度|
   |a\,b|1/6m 宽度|
   |ab|没有空格|
   |a\!b|缩进 1/6m 宽度|

## 提交规范

- `feat` 增加新功能
- `fix` 修复问题/BUG
- `style` 代码风格相关无影响运行结果的
- `perf` 优化/性能提升
- `refactor` 重构
- `revert` 撤销修改
- `test` 测试相关
- `docs` 文档/注释
- `build` 对构建系统或者外部依赖项进行了修改
- `chore` 依赖更新/脚手架配置修改等
- `workflow` 工作流改进
- `ci` 持续集成
- `types` 类型定义文件更改
- `wip` 开发中

## 参考链接

- [Markdown/LaTeX 数学公式和符号表](https://zhuanlan.zhihu.com/p/450465546)
- [markdown 中插入数学公式](https://blog.csdn.net/LeonSUST/article/details/84204723)
- [数学公式速查](https://blog.csdn.net/jyfu2_12/article/details/79207643)
