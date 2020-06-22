# InventoryGO
Tech Inventory App using Golang and Console

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

this project using this database design
```sql
create database goInventory;
use goInventory;

create table category (
categoryID varchar (5) primary key,
categoryDesc varchar(50) not null
);

create table discount (
discountID varchar(5) primary key,
discountDesc varchar(50) not null
);

create table product (
productID varchar(5) primary key,
productDesc varchar(50) not null,
categoryID varchar (5),
productPrice int,
discountID varchar(5),
foreign key product_PtoC (categoryID) references category(categoryID),
foreign key product_PtoD (discountID) references discount(discountID)
);

create table user (
userID int primary key auto_increment,
userName varchar (15) not null,
userPass varchar (50) not null
);

create table userDetail (
user int,
user_fname varchar (25),
user_lname varchar (25),
userAddress varchar(50)
);

alter table userDetail add foreign key userDetail_UDtoU(userID) references user(userID);
alter table discount modify column discountDesc int not null;

select p.productID,p.productDesc, c.categoryDesc,p.productPrice, d.discountDesc from product p 
inner join category c on p.categoryID = c.categoryID
inner join discount d on p.discountID = d.discountID order by p.productID;
```
