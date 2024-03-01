#Part 1
#Solving equation for production curve
#y=(A/x)+B
#For x=1, y=100 barrels.
#For x=100, y=10 barrels.
#A production curve equation of y=A/x+B can be solved for A and B from two values each for x and y, respectively, via Wolfram Alpha.
#From Wolfram Alpha, A~90.9 and B~9.1
#Equation for production curve is y=90.9/x+9.1
x=1
s=0
totalprice=0
#Part 2
#Determine daily production for 100 days.
while x<=100:
    y=(90.9/x)+9.1
    print(f'The daily production is {y:.2f} barrels for at day',x,'- look at the nice rounding!')
    x+=1
    s+=y
    if x<=30:
        totalprice+=85*y
    elif 31<=x<=60:
        totalprice+=80*y
    elif 61<=x<=90:
        totalprice+=82*y
    else:
        totalprice+=84*y
#Part 3
print(f'The total production for 100 days is {s:.2f} barrels.')
#Part 4
print(f'The total revenue generated is ${totalprice:.2f}.')