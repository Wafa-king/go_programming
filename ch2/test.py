# res1 = foo(["a", "da", "bde", "fdas", "fdakg", "1234567"]) 
# #res1是["a", "da", "bde"]
# res2 = foo(["1", "fdafdas", "fda", "fd", "fewqi", "a"])
# #res2是["1", "fda", "fd", "a"]
# res3 = foo(["42314", "1", "3", "314"])
# # #res3是["1", "3", "314"]
# def foo(x):
#     pass
 
# def func():
#     def mystery(x):
#         return x * 8


# @mystery
# def 
# foo = mystery(8)
# foo(8) # 结果是64
# foo(-1) # 结果是-8



# # 请写出foo的代码。

# multiply_factor = 3
# i = 0.1
# while i != 8.1:
#     # 保留两位小数
#     print("%.2f" % i)
#     i *= multiply_factor
#     if i > 2.7:
#         break

# x = "thisisa测试"
# for i in x:
#     try:
#         if len(i) == len(i.encode("utf-8")):
#             print("%s: %s" %(i, int(ord(i)))) 
#     except Exception:
#         continue


def bar(x):
    if not isinstance(x, list):
        return
    res = 1
    for i in x:
        print(i)
        res *= int(i)
    return res

res1 = bar(["1", "2", "3", "4", "5"])
print(res1)
#res1是120
res2 = bar(["-1", "1", "-1", 2])
print(res2)
#res2是2
res3 = bar(["2", "10", "30", "-1"])
print(res3)
#res3是-600
