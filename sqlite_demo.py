import sqlite3
from employee import Employee
conn=sqlite3.connect(:memory)
c=conn.cursor()
#c.execute("""CREATE TABLE employees (
#    first text,
#    last text,
#    pay integer
 #   """)
emp_1=employee('John','Doe',80000)
emp_2=employee('Jane','Doe',90000)
c.execute(INSERT INTO employees VALUES('?','?','?',(emp_1.first,emp_1.last,emp_1.pay))
conn.commit()
c.execute(INSERT INTO employees VALUES(':first',':last',':pay',('first':emp_1.first,':last':emp_1.last,':pay'emp_1.pay))
conn.commit()
c.execute(SELECT * FROM EMPLOYEES WHERE LAST = '?'.('Raney'))
print(c.fetchall())
c.execute(SELECT * FROM EMPLOYEES WHERE LAST = ':last'.('Doe'))
print(c.fetchall())
conn.commit()
conn.close()