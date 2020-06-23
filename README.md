# InventoryGO
Tech Inventory App using Golang and Console
## Update 
 Add Ability to
```
- Add product
- Delete Product
- Show All Product

- Add Transaction
- Delete Transaction
- Show All Transaction

- Show Daily Report
- Show Monthly Report
- Show All Report
```

## create config file

create [config.env](https://github.com/inact25/InventoryGO), and put it into root of your project folder.Then fill it with :
```env
DbEngine = 'mysql'
DbUser = 'root'
DbPass = 'yourPass'
DbHost = 'localhost'
DbPort = 'yourPort'
DbConnection = '@tcp'
DbSchema = 'yourDbName'
```
## create database

this project using this database design [Download Sql File](https://drive.google.com/file/d/1XJIFJKeI54vIWa7u-Go3NkIGzdxndWUS/view?usp=sharing)
![image of db](https://1.bp.blogspot.com/-K0i0Qhtc0Ts/XvH9C1vf0FI/AAAAAAAAIRA/0gxf3rmxr-8Ez9wBHhI0o25mD9WW6zXvwCK4BGAsYHg/d/cft.PNG)

## Future Update

- Login
- choose Category when add a product
- choose brand when add a product and transaction
- cloudSql
- input validation
- more
