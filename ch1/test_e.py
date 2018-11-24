# coding=utf-8

# i = 0
# fibos = [0, 1]
# while i < 15:
# 	# a , b = b, a+b
# 	i += 1
# 	# fibos.append(b)
# 	fibos.append(fibos[-2]+fibos[-1])
# print(fibos)

# x = 1
# y = 2
# c = 3
# n  = 0
# print(x and y and c)
# print(x or y or c)
# print(n and x and y and c)
# print(n or x or y or c )

class MyClass(object):
	class_name = "xiaobai"
	def __init__(self, *args):
		self.class_name = args[0]
	def print_name(self):
		print("this class name: {}".format(self.class_name))
	@staticmethod
	def print_name_stc():
		print('This class name: {}'.format(MyClass.class_name))
	@classmethod
	def print_name_cls(cls):
		print('This class name: {}'.format(cls.class_name))

my_class = MyClass("xiaohei")	
my_class.print_name()
my_class.print_name_cls()
my_class.print_name_stc()
MyClass.print_name_cls()
MyClass.print_name_stc()
# 注：print_name是实例方法，类对象Myclass没有这个方法，不能调用
# MyClass.print_name()








# class A:
# 	x = 5
# 	s = "aaa"
# 	def __init__(self):
# 		self.x = 8
	
# class B(A):
# 	def __init__(self):
# 		self.s = "bbb"
# # 注意：这里没有拍全，看看是不是这个
# a = A()
# a.x = 9
# b = B()
# print(a.x, a.s)
# print(b.x, b.s)

# A0 = dict(zip(("a", "b", "c", "d", "e"), (1, 2, 3, 4, 5)))
# A1 = range(10)
# A2 = [i for i in A1 if i in A0]
# A3 = [A0[s] for s in A0]
# A4 = [i for i in A1 if i in A3]
# A5 = {i: i*i for i in A1}
# A6 = [[i, i*i] for i in A1]
# print(A0)
# print(A1)
# print(A2)
# print(A3)
# print(A4)
# print(A5)
# print(A6)

# t = tuple('asd')
# print(t)

# def f(x, I=[]):
# 	for i in range(x):
# 		I.append(i*i)
# 	print(I)

# f(2)
# f(3, [3, 2, 1])
# f(3)