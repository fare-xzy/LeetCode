counter = 100  # 整型变量
miles = 1000.0  # 浮点型变量
name = "runoob"  # 字符串

print(counter)
print(miles)
print(name)

a = b = c = 1

a, b, c = 1, 2, "runoob"

a, b, c, d = 20, 5.5, True, 4 + 3j

print(type(a), type(b), type(c), type(d))
print(isinstance(a, int))


class A:
    pass


class B(A):
    pass


# type 会认为B就是B  isinstance则会认为B的类型时他的父类A
print(type(B()), isinstance(B(), A))

var1 = 1
var2 = 10
# 删除变量
del var1, var2

print(5 + 4)  # 加法
print(4.3 - 2)  # 减法
print(3 * 7)  # 乘法
print(2 / 4)  # 除法，得到一个浮点数
print(2 // 4)  # 除法，得到一个整数
print(17 % 3)  # 取余
print(2 ** 5)  # 乘方

word = 'Python'
print(word[0], word[5])
print(word[-1], word[-6])
# 注意python字符串不可以被改变例如：word[0] = 'm'会导致错误。

# 列表
myList = ['abcd', 786, 2.23, 'runoob', 70.2]
tinyList = [123, 'runoob']
allList = [123, '123', (1, 1, 1)]
print(myList)  # 输出完整列表
print(myList[0])  # 输出列表第一个元素
print(myList[1:3])  # 从第二个开始输出到第三个元素
print(myList[2:])  # 输出从第三个元素开始的所有元素
print(tinyList * 2)  # 输出两次列表
print(myList + tinyList)  # 连接列表

# 列表元素可以被改变
a = [1, 2, 3, 4, 5, 6]
a[0] = 9
a[2:5] = [13, 14, 15]
print(a)
a[2:5] = []
print(a)

# b列表中可以通过设置步长来进行输出如下
# print(b[start:end:step])  start开始索引， end结束索引，  step步长。如果步长为2 则会在规定范围内隔一个元素一输出
b = [1, 2, 3, 4, 5, 6]
print(b[0:5:2])

# 上述内容可以实现翻转字符串
strC = '1234567890'
inputWords = list(strC)
outputWords = inputWords[-1::-1]
output = ''.join(outputWords)
print(output)

# Tuple元组, 元组不可修改
tuple = ('abcd', 786, 2.23, 'runoob', 70.2)
tinytuple = (123, 'runoob')
allTuple = (123, '123', ['123', '123'])
print(tuple)  # 输出完整元组
print(tuple[0])  # 输出元组的第一个元素
print(tuple[1:3])  # 输出从第二个元素开始到第三个元素
print(tuple[2:])  # 输出从第三个元素开始的所有元素
print(tinytuple * 2)  # 输出两次元组
print(tuple + tinytuple)  # 连接元组

tup1 = ()  # 空元组
tup2 = (20,)  # 一个元素，需要在元素后添加逗号

# set 集合（唯一） TODO：输出内容是否随机，看现象是随机顺序输出
sites = {'Google', 'Taobao', 'Runoob', 'Facebook', 'Zhihu', 'Baidu'}
print(sites)  # 输出集合，重复的元素被自动去掉
# 成员测试
if 'Runoob' in sites:
    print('Runoob 在集合中')
else:
    print('Runoob 不在集合中')
# set可以进行集合运算
a = set('abracadabra')
b = set('alacazam')
print(a)

# 字典
dict = {}
dict['one'] = "1 - 菜鸟教程"
dict[2] = "2 - 菜鸟工具"

print(dict)

tinydict = {'name': 'runoob', 'code': 1, 'site': 'www.runoob.com'}
print(dict['one'])  # 输出键为 'one' 的值
print(dict[2])  # 输出键为 2 的值
print(tinydict)  # 输出完整的字典
print(tinydict.keys())  # 输出所有键
print(tinydict.values())  # 输出所有值

alldict = {'1': 1, '2': '2', '3': (1, 2, 3,), '4': [1, 2, 3, 4], '5': {1, 2, 3, 4, 5}}
print(alldict)
print(type(alldict.get('3')))

# 数据类型转换
# int(x [,base]) 将x转换为一个整数
# float(x) 将x转换到一个浮点数
# complex(real [,imag]) 创建一个复数
# str(x) 将对象 x 转换为字符串
# repr(x) 将对象 x 转换为表达式字符串
# eval(str) 用来计算在字符串中的有效Python表达式,并返回一个对象
# tuple(s) 将序列 s 转换为一个元组
# list(s) 将序列 s 转换为一个列表
# set(s) 转换为可变集合
# dict(d) 创建一个字典。d 必须是一个 (key, value)元组序列。
# frozenset(s) 转换为不可变集合
# chr(x) 将一个整数转换为一个字符
# ord(x) 将一个字符转换为它的整数值
# hex(x) 将一个整数转换为一个十六进制字符串
# oct(x) 将一个整数转换为一个八进制字符串


