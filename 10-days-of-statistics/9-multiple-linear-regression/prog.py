from sklearn import linear_model

m, n = map(int, input().split())
x = []
y = []
for i in range(n):
    data = list(map(float, input().split()))
    x.append(data[:m])
    y.append(data[-1])
lm = linear_model.LinearRegression()
lm.fit(x, y)
a = lm.intercept_
b = lm.coef_

q = int(input())
for i in range(q) :
    f = list(map(float,input().split()))
    y = a
    for j in range(m):
        y += f[j] * b[j] 
    print('{:.2f}'.format(y))

