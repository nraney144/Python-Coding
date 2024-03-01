foobar = 3.141592
print(f'My number is {foobar:.2f} - look at the nice rounding!')
x=1
sum=0
while x<=100:
    y=(90.9/x)+9.1
    print(f'My number is {y:.2f} for x=',x,'- look at the nice rounding!')
    x+=1
    sum+=y
print(f'The total production is {sum:.2f}.')    