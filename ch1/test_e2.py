# coding=utf-8
# class ClassTest(object):
#     def __init__(self, a):
#         self.a = a
#     def printk(self):
#         print "a is ", self.a 

#     # 注：此处有错误
#     @classmethod
#     def class_method(cls, *args):
#         print "args is ", args
#     @classmethod
#     def class_method2(cls):
#         print "cls is ", cls
#     @staticmethod
#     def static_method(*args):
#         print "args is ", args

# ct = ClassTest(100)
# ct.printk()
# ct.class_method()
# ct.static_method()
# ct.class_method2()

# a = 1
# def func(a):
#     a = 2
# func(a)
# print a

# a = []
# def func1(a):
#     a.append(1)
# func1(a)
# print a


# # 第四题
# result = []
# for x in range(10):
#     result.append(x ** 2)
#     print(result)
# # 法一：
# result = [(i**2) for i in range(10)]
# print(result)

# print("-"*10)
# # 法二：
# def func2(x):
#     return x ** 2
# print(map(func2, [i for i in range(10)]))

# 10题
# a = [12, 3 , 4 , 4, 5 , 10]
# # 先将列表逆序 
# b = reversed(a)
# print(sorted(b))