#Program to check whether or not a number is prime
#Take input from the user.
num=int(input("Enter a number: "))
#define a flag variable
flag=False
if num==1:
    print(num,' is not a prime number.')
elif num>1:
    #check for factors
    for i in range(2,num):
        if (num%i==0):
            #if factor is found, set flag to true
            flag=True
            #break out of loop
            break

#check if flag is true
if flag==True:
    print(num,' is not a prime number.')
else:
    print(num,'is a prime number.')